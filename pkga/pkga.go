package pkga

import (
	"github.com/Centny/Cny4go/routing"
)

var Vars map[string]interface{} = map[string]interface{}{}

func AddA(hs *routing.HTTPSession) routing.HResult {
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
	Vars[key] = val
	hs.W.Write([]byte(`
			{
				"code":0,
				"data":""
			}
			`))
	return routing.HRES_RETURN
}

func DelA(hs *routing.HTTPSession) routing.HResult {
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
	if _, ok := Vars[key]; ok {
		delete(Vars, key)
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
