package alert

import (
	"github.com/edunx/lua"
)

func CheckAlertUserDataByTable(L *lua.LState , opt *lua.LTable , key string ) *Alert {
	var obj *Alert
	var ud  *lua.LUserData
	var ok bool

	ud , ok = opt.RawGetString(key).(*lua.LUserData)
	if !ok { goto ERR }

	obj , ok = ud.Value.(*Alert)
	if !ok { goto ERR } else { return obj }

ERR:
	L.RaiseError("expect invalid type , must be *Alert")
	return nil
}

func CheckAlertUserData(L *lua.LState , idx int) *Alert {
	ud := L.CheckUserData( idx )

	switch v := ud.Value.(type) {
	case *Alert:
		return ud.Value.(*Alert)
	default:
		L.RaiseError("expect invalid type , must be *Alert, got %T" , v )
		return nil
	}

}
