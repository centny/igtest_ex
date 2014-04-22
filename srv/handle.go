package srv

import (
	"github.com/Centny/Cny4go/routing"
	"github.com/Centny/igtest_ex/pkga"
	"github.com/Centny/igtest_ex/pkgb"
)

func NewSrvMux(pre string) *routing.SessionMux {
	sb := routing.NewSrvSessionBuilder("", "/", "ckey", 600000, 1000)
	mux := routing.NewSessionMux(pre, sb)
	mux.HFunc("^/adda", pkga.AddA)
	mux.HFunc("^/dela", pkga.DelA)
	mux.HFunc("^/addb", pkgb.AddB)
	mux.HFunc("^/getb", pkgb.GetB)
	return mux
}
