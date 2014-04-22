package pkgb

import (
	"fmt"
	"github.com/Centny/Cny4go/routing"
	"github.com/Centny/igtest_ex/pkga"
)

var Vars map[string]interface{} = map[string]interface{}{}

func AddB(hs *routing.HTTPSession) routing.HResult {
	var key string
	var val string
	err := hs.ValidCheckVal(`
		key,R|S,L:0;
		val,R|S,L:0;
		`, &key, &val)
	if err != nil {
		hs.W.Write([]byte(`
			{
				"code":1,
				"msg":"error"
			}
			`))
		return routing.HRES_RETURN
	}
	if _, ok := pkga.Vars[key]; ok {
		Vars[key] = val
		hs.W.Write([]byte(`
			{
				"code":0,
				"data":""
			}
			`))
	} else {
		hs.W.Write([]byte(`
			{
				"code":1,
				"msg":"key not found"
			}
			`))
	}
	return routing.HRES_RETURN
}

func GetB(hs *routing.HTTPSession) routing.HResult {
	var key string
	err := hs.ValidCheckVal(`
		key,R|S,L:0;
		`, &key)
	if err != nil {
		hs.W.Write([]byte(`
			{
				"code":1,
				"msg":"error"
			}
			`))
		return routing.HRES_RETURN
	}
	if v, ok := Vars[key]; ok {
		hs.W.Write([]byte(fmt.Sprintf(`
			{
				"code":0,
				"data":"%v"
			}
			`, v)))
	} else {
		hs.W.Write([]byte(`
			{
				"code":1,
				"msg":"key not found"
			}
			`))
	}
	return routing.HRES_RETURN
}
