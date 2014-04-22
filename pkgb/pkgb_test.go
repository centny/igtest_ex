package pkgb

import (
	"fmt"
	"github.com/Centny/Cny4go/routing/httptest"
	"github.com/Centny/igtest"
	"github.com/Centny/igtest_ex/pkga"
	"testing"
)

func TestPkgb(t *testing.T) {
	//mock pkga
	pkga.Vars["abc"] = 111
	//
	ts := httptest.NewMuxServer()
	ts.Mux.HFunc("^/addb", AddB)
	ts.Mux.HFunc("^/getb", GetB)
	jme := igtest.NewJME("pkgb.json", "pkgb.ig")
	jme.ShowLog = true
	jme.SET("PKGA_KEY_V", "abc")
	jme.SET("URL_ADD_PKGB", fmt.Sprintf("%v/addb", ts.URL))
	jme.SET("URL_GET_PKGB", fmt.Sprintf("%v/getb", ts.URL))
	err := jme.Exec()
	if err != nil {
		t.Error(err.Error())
		return
	}
}
