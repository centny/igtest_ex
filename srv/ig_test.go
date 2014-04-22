package srv

import (
	"fmt"
	"github.com/Centny/Cny4go/routing/httptest"
	"github.com/Centny/igtest"
	"testing"
)

func TestSrv(t *testing.T) {
	//
	ts := httptest.NewMuxServer2(NewSrvMux(""))
	jme := igtest.NewJME("srv.json", "srv.ig")
	jme.ShowLog = true
	jme.SET("URL_ADD_PKGA", fmt.Sprintf("%v/adda", ts.URL))
	jme.SET("URL_DEL_PKGA", fmt.Sprintf("%v/dela", ts.URL))
	jme.SET("URL_ADD_PKGB", fmt.Sprintf("%v/addb", ts.URL))
	jme.SET("URL_GET_PKGB", fmt.Sprintf("%v/getb", ts.URL))
	err := jme.Exec()
	if err != nil {
		t.Error(err.Error())
		return
	}
}
