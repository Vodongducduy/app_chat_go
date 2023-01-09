package router

type (
	IUserRouter interface {
		SetUserRouter()
	}
)

func (s router) SetUserRouter() {
	s.routerSetting.GET("/user", s.controller.GetUserProfile)
}
