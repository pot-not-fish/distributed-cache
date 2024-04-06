package auth

type LoginRequest struct {
	Username string `form:"username" json:"username" xml:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type LoginResponse struct {
	Code int    `form:"code" json:"code" xml:"code"`
	Msg  string `form:"msg" json:"msg" xml:"msg"`

	Token string `form:"token" json:"token" xml:"token"`
}
