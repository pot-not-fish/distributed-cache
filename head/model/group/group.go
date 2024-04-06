package group

type SetGroupRequest struct {
	Token string `form:"token" json:"token" xml:"token"`
	Group string `form:"group" json:"group" xml:"group"`
}

type SetGroupResponse struct {
	Code int    `form:"code" json:"code" xml:"code"`
	Msg  string `form:"msg" json:"msg" xml:"msg"`
}

type GetAllGroupRequest struct {
	Token string `form:"token" json:"token" xml:"token"`
}

type GetAllGroupResponse struct {
	Code int    `form:"code" json:"code" xml:"code"`
	Msg  string `form:"msg" json:"msg" xml:"msg"`

	Group []string `form:"group" json:"group" xml:"group"`
}
