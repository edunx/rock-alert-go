package alert

import (
	"github.com/edunx/lua"
	pub "github.com/edunx/rock-public-go"
)

const (
	MT string = "ROCK_ALERT_GO_MT"
)

func CreateAlertUserdata( L *lua.LState ) int {
	opt := L.CheckTable(1)
	v := &Alert{
		C: Config{
			Url: opt.CheckString("url", "null"),
			Origin: opt.CheckString("origin" , "security-alert"),
			Notifier: opt.CheckString("notifier" , "[]"),
			Resolver: opt.CheckString("resolver"  , "114.114.114.114:53"),
		},
	}

	if e := v.Start(); e != nil {
		L.RaiseError("start alert fail , err: %s" , e)
		return 0
	}

	ud := L.NewUserDataByInterface( v , MT)
	L.Push(ud)

	pub.Out.Debug("create alert successful , info: %s" , v.C )
	return 1

}

func LuaInjectApi(L *lua.LState , parent *lua.LTable) {
	mt := L.NewTypeMetatable( MT )

	L.SetField(mt , "__index" , L.NewFunction(Get))
	L.SetField(mt , "__newindex" , L.NewFunction(Set))

	L.SetField(parent , "alert" , L.NewFunction(CreateAlertUserdata))
}

func Get(L *lua.LState) int {
    self := CheckAlertUserData(L , 1)
    name := L.CheckString(2)
    switch name {
    case "send":
        L.Push(L.NewFunction(func(L *lua.LState) int {
            opt := L.CheckTable(1)
            self.Do( opt.CheckString("severity" , "middle"),
				opt.CheckString("type" , "security"),
				opt.CheckString("object" , "security-honey-pot"),
				opt.CheckString("attribute" , "null"),
				opt.CheckString("subject" , "null"),
				opt.CheckString("body" , "null"),
				opt.CheckString("tags" , "null"),
            )
            return 0
        }))
        return 1
    default:
    }
	return 0
}


func Set(L *lua.LState) int {
	return 0
}

func (self *Alert) ToUserData(L *lua.LState) *lua.LUserData {
	return L.NewUserDataByInterface( self , MT )
}
