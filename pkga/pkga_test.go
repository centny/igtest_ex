package pkga

import (
	"fmt"
	"github.com/Centny/Cny4go/routing/httptest"
	"github.com/Centny/igtest"
	"testing"
)

func TestPkga(t *testing.T) {
	ts := httptest.NewMuxServer()
	ts.Mux.HFunc("^/adda", AddA)
	ts.Mux.HFunc("^/dela", DelA)
	jme := igtest.NewJME("pkga.json", "pkga.ig")
	jme.ShowLog = true
	jme.SET("URL_ADD_PKGA", fmt.Sprintf("%v/adda", ts.URL))
	jme.SET("URL_DEL_PKGA", fmt.Sprintf("%v/dela", ts.URL))
	err := jme.Exec()
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestPkgaErr(t *testing.T) {
	fmt.Println(httptest.Tf2(AddA, "?key="))
}
