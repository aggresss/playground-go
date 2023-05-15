/*
Method-Domain-Path-Query Access (MDPQ)

MDPQ 中间件支持从 HTTP 请求的 Method/Host/Path/Query 中提取字段(field)，并作用于：

1. 将字段通过 Method-Domain-Path-Query 顺序进行组合并映射为注册组件的方法调用，从而实现路径分发；
2. 将字段中的变量信息加载到上下文中，为下游方法提供所需信息；

字段提取规则：

- Method
	将方法名称映射为首字母大写其余字母小写的格式作为路径分发的前缀，当注册组件中没有对应方法时会尝试使用 "All" 作为前缀；
- Host
	Host 首先需要满足初始化时提供的固定后缀，字段提取在剔除匹配的固定后缀后进行；
	剩余部分以 "." 为分隔符逆序匹配；
- Path
	以 "/" 为分隔符正序匹配；
- Query
	仅支持 Primary Non-Key-Value Query ，即 Raw Query 中以 "&" 为分隔符的第一字段，并且只能是 "key" 的格式，而非 "key=value" 格式；
*/

package middleware

import (
	"errors"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type FdKind uint

const (
	MdpqFdInvalid FdKind = iota
	MdpqFdRoute
	MdpqFdCtxOnly
)

// Method-Domain-Path-Query Field Description
type MdpqFd struct {
	Kind      FdKind
	Mandatory bool
	Key       string            // Primary
	Maps      map[string]string // Secondary
}

// Method-Domain-Path-Query Configuration
type MdpqConf struct {
	Router         interface{}
	DomainSuffixes []string
	DomainCasts    []MdpqFd
	PathCasts      []MdpqFd
	UseQuery       bool
}

func service(conf *MdpqConf) gin.HandlerFunc {
	if conf.Router == nil {
		log.Panic("Registed Nil Router")
	}
	domainMandNum := 0
	for _, v := range conf.DomainCasts {
		if v.Mandatory {
			domainMandNum++
		} else {
			break
		}
	}
	pathMandNum := 0
	for _, v := range conf.PathCasts {
		if v.Mandatory {
			pathMandNum++
		} else {
			break
		}
	}
	routerType := reflect.TypeOf(conf.Router)
	routerValue := reflect.ValueOf(conf.Router)
	// use cache instead of MethodByName()
	methodCache := make(map[string]int)
	for i := 0; i < routerType.NumMethod(); i++ {
		if routerType.Method(i).IsExported() {
			methodCache[routerType.Method(i).Name] = i
		}
	}
	return func(ctx *gin.Context) {
		var route string
		domains, err := parseDomainSlice(ctx.Request.Host, conf.DomainSuffixes)
		if err != nil || len(domains) < domainMandNum || len(domains) > len(conf.DomainCasts) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Host",
			})
			return
		}
		for i, v := range domains {
			cast := conf.DomainCasts[i]
			switch cast.Kind {
			case MdpqFdRoute:
				if cast.Key != "" {
					route += cast.Key
					ctx.Set(cast.Key, v)
				} else if cast.Maps != nil {
					if r, ok := cast.Maps[v]; ok {
						route += r
					} else if cast.Mandatory {
						ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
							"error": "Invalid Host",
						})
						return
					}
				}
			case MdpqFdCtxOnly:
				if cast.Key != "" {
					ctx.Set(cast.Key, v)
				}
			}
		}
		paths, _ := parsePathSlice(ctx.Request.URL.Path)
		if len(paths) < pathMandNum || len(paths) > len(conf.PathCasts) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Path",
			})
			return
		}
		for i, v := range paths {
			cast := conf.PathCasts[i]
			switch cast.Kind {
			case MdpqFdRoute:
				if cast.Key != "" {
					route += cast.Key
					ctx.Set(cast.Key, v)
				} else if cast.Maps != nil {
					if r, ok := cast.Maps[v]; ok {
						route += r
					} else if cast.Mandatory {
						ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
							"error": "Invalid Path",
						})
						return
					}
				}
			case MdpqFdCtxOnly:
				if cast.Key != "" {
					ctx.Set(cast.Key, v)
				}
			}
		}
		if conf.UseQuery {
			route += parseRawQuery(ctx.Request.URL.RawQuery)
		}
		i, ok := methodCache[formatMethod(ctx.Request.Method)+route]
		if !ok {
			i, ok = methodCache["All"+route]
		}
		if ok {
			routerValue.Method(i).Call([]reflect.Value{reflect.ValueOf(ctx)})
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Route",
			})
		}
	}
}

func parseDomainSlice(host string, suffix []string) (s []string, err error) {
	isMatched := false
	for _, sf := range suffix {
		if index := strings.Index(host, sf); (index > 1 && host[index-1] == '.') || (index == 0) {
			isMatched = true
			s = []string{}
			if index > 1 {
				s = strings.Split(strings.TrimSuffix(host[:index], "."), ".")
				for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
					s[i], s[j] = s[j], s[i]
				}
			}
			break
		}
	}
	if !isMatched {
		return nil, errors.New("mismathed host")
	}
	return s, err
}

func parsePathSlice(path string) (s []string, err error) {
	s = []string{}
	p := strings.TrimRight(strings.TrimLeft(path, "/"), "/")
	if len(p) > 0 {
		s = strings.Split(p, "/")
	}
	return s, nil
}

func parseRawQuery(rawQuery string) (fq string) {
	firstQuery := strings.Split(rawQuery, "&")[0]
	if len(firstQuery) == 0 {
		return ""
	}
	i := strings.Index(firstQuery, "=")
	if i == len(firstQuery)-1 {
		firstQuery = firstQuery[:len(firstQuery)-1]
	} else if i >= 0 {
		return ""
	}
	if len(firstQuery) > 0 {
		fq = strings.ToUpper(string(firstQuery[0]))
	}
	if len(firstQuery) > 1 {
		fq += firstQuery[1:]
	}
	return fq
}

func formatMethod(i string) (o string) {
	o = strings.ToLower(i)
	if len(o) > 0 {
		o = strings.ToUpper(string(o[0])) + o[1:]
	}
	return o
}
