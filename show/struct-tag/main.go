package main

import (
	"reflect"
	"strconv"
	"strings"
)

const testFmptStr string = "a=fmtp:111 minptime=10;useinbandfec=1"

type OpusRtpFormatParams struct {
	Minptime     uint8 `json:"minptime" sdp:"minptime"`
	Stereo       uint8 `json:"stereo" sdp:"stereo"`
	SpropStereo  uint8 `json:"sprop-stereo" sdp:"sprop-stereo"`
	Useinbandfec uint8 `json:"useinbandfec" sdp:"useinbandfec"`
}

func ExtractFmtp(fmtp string) (res OpusRtpFormatParams) {
	resVal := reflect.ValueOf(&res).Elem()
	ss := strings.Split(fmtp, ";")
	mapFmtp := map[string]string{}
	for _, s := range ss {
		kv := strings.Split(s, "=")
		if len(kv) == 2 {
			mapFmtp[kv[0]] = kv[1]
		}
	}
	for k, v := range mapFmtp {
		for i := 0; i < resVal.NumField(); i++ {
			sdpTag := resVal.Type().Field(i).Tag.Get("sdp")
			if sdpTag == k {
				switch resVal.Type().Field(i).Type.Kind() {
				case reflect.Uint8:
					if u, err := strconv.ParseUint(v, 10, 8); err == nil {
						resVal.Field(i).Set(reflect.ValueOf(uint8(u)))
					}
				case reflect.String:
					resVal.Field(i).Set(reflect.ValueOf(v))
				}
			}
		}
	}
	return res
}

func main() {
	ExtractFmtp(testFmptStr)
}
