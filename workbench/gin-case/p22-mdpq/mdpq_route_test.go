package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFormatMethod(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected string
	}{
		"test01": {
			input:    "",
			expected: "",
		},
		"test02": {
			input:    "POST",
			expected: "Post",
		},
		"test03": {
			input:    "post",
			expected: "Post",
		},
		"test04": {
			input:    "OpTIONs",
			expected: "Options",
		},
	}

	for n, c := range testCases {
		t.Run(n, func(t *testing.T) {
			infos := formatMethod(c.input)
			assert.Equal(t, c.expected, infos)
		})
	}
}

func TestParseRawQuery(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected string
	}{
		"test01": {input: "foo=bar&zoo", expected: ""},
		"test02": {input: "fOo", expected: "FOo"},
		"test03": {input: "", expected: ""},
		"test04": {input: "foo&bar=zoo", expected: "Foo"},
		"test05": {input: "&", expected: ""},
		"test06": {input: "=", expected: ""},
		"test07": {input: "foo=", expected: "Foo"},
		"test08": {input: "=bar", expected: ""},
		"test09": {input: "q", expected: "Q"},
	}

	for n, c := range testCases {
		t.Run(n, func(t *testing.T) {
			infos := parseRawQuery(c.input)
			assert.Equal(t, c.expected, infos)
		})
	}
}

func TestParsePathSlice(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected []string
	}{
		"test01": {
			input:    "/",
			expected: []string{},
		},
		"test02": {
			input:    "",
			expected: []string{},
		},
		"test03": {
			input:    "rho/sigma/tou",
			expected: []string{"rho", "sigma", "tou"},
		},
		"test04": {
			input:    "/zeta/eta/theta/",
			expected: []string{"zeta", "eta", "theta"},
		},
	}

	for n, c := range testCases {
		t.Run(n, func(t *testing.T) {
			infos, _ := parsePathSlice(c.input)
			assert.Equal(t, c.expected, infos)
		})
	}
}

func TestParseDomainSlice(t *testing.T) {
	testCases := map[string]struct {
		input01    string
		input02    []string
		expected01 []string
		expected02 bool
	}{
		"test01": {
			input01:    "zoo.bar.foo.com",
			input02:    []string{"foo.org", "foo.com"},
			expected01: []string{"bar", "zoo"},
			expected02: false,
		},
		"test02": {
			input01:    "zoo.bar.foo.com",
			input02:    []string{"zoo.com"},
			expected01: nil,
			expected02: true,
		},
		"test03": {
			input01:    "",
			input02:    []string{"foo.com"},
			expected01: nil,
			expected02: true,
		},
		"test04": {
			input01:    "zoobarfoo.com",
			input02:    []string{"foo.com"},
			expected01: nil,
			expected02: true,
		},
		"test05": {
			input01:    "foo.com",
			input02:    []string{"foo.com"},
			expected01: []string{},
			expected02: false,
		},
	}

	for n, c := range testCases {
		t.Run(n, func(t *testing.T) {
			infos, err := parseDomainSlice(c.input01, c.input02)
			assert.Equal(t, c.expected01, infos)
			if !c.expected02 {
				assert.Nil(t, err)
			}
		})
	}
}

type mockSrv struct{}

func (ms *mockSrv) GetIotaLambdaSigmaTau(ctx *gin.Context) {
	rho, exist := ctx.Get("Rho")
	if !exist {
		ctx.String(499, ctx.Request.Host)
		return
	}
	lambda, exist := ctx.Get("Lambda")
	if !exist {
		ctx.String(499, ctx.Request.Host)
		return
	}
	sigma, exist := ctx.Get("Sigma")
	if !exist {
		ctx.String(499, ctx.Request.Host)
		return
	}
	ctx.String(200, rho.(string)+"."+lambda.(string)+"."+sigma.(string))
}

func (ms *mockSrv) PostIota(ctx *gin.Context) {
	rho, exist := ctx.Get("Rho")
	if !exist {
		ctx.String(499, ctx.Request.Host)
		return
	}
	ctx.String(200, rho.(string))
}

func (ms *mockSrv) AllIota(ctx *gin.Context) {
	rho, exist := ctx.Get("Rho")
	if !exist {
		ctx.String(499, ctx.Request.Host)
		return
	}
	ctx.String(200, "all."+rho.(string))
}

func TestMdpqRoute(t *testing.T) {
	r := gin.Default()
	r.Use(service(&MdpqConf{
		Router:         &mockSrv{},
		DomainSuffixes: []string{"psi.omega", "phi.chi"},
		DomainCasts: []MdpqFd{
			{
				Kind:      MdpqFdCtxOnly,
				Mandatory: true,
				Key:       "Rho",
			},
			{
				Kind:      MdpqFdRoute,
				Mandatory: true,
				Maps: map[string]string{
					"k-a-p-p-a": "Kappa",
					"i-o-t-a":   "Iota",
				},
			},
			{
				Kind:      MdpqFdRoute,
				Mandatory: false,
				Key:       "Lambda",
			},
		},
		PathCasts: []MdpqFd{
			{
				Kind:      MdpqFdRoute,
				Mandatory: false,
				Key:       "Sigma",
			},
		},
		UseQuery: true,
	}))

	testCases := map[string]struct {
		input01    string
		input02    string
		expected01 int
		expected02 string
	}{
		"test01": {
			input01:    "GET",
			input02:    "http://alpha.i-o-t-a.zeta.psi.omega/beta?tau",
			expected01: 200,
			expected02: "zeta.alpha.beta",
		},
		"test02": {
			input01:    "POST",
			input02:    "http://i-o-t-a.zeta.psi.omega/",
			expected01: 200,
			expected02: "zeta",
		},
		"test03": {
			input01:    "DELETE",
			input02:    "http://i-o-t-a.zeta.psi.omega/",
			expected01: 200,
			expected02: "all.zeta",
		},
		"test04": {
			input01:    "GET",
			input02:    "http://alpha.beta.org/gamma?deta",
			expected01: 400,
			expected02: "{\"error\":\"Invalid Host\"}",
		},
		"test05": {
			input01:    "POST",
			input02:    "http://alpha.i-o-t-a.zeta.psi.omega/beta?tau",
			expected01: 400,
			expected02: "{\"error\":\"Invalid Route\"}",
		},
	}

	for n, c := range testCases {
		t.Run(n, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(c.input01, c.input02, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, c.expected01, w.Code)
			assert.Equal(t, c.expected02, w.Body.String())
		})
	}
}
