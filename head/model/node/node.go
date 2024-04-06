package node

type GetAllNodeRequest struct {
	Token string `form:"token" json:"token" xml:"token"  binding:"required"`
}

type GetAllNodeResponse struct {
	Code  int      `form:"code" json:"code" xml:"code"`
	Msg   string   `form:"msg" json:"msg" xml:"msg"`
	Nodes []string `form:"nodes" json:"nodes" xml:"nodes"`
}
