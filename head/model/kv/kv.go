package kv

type Cache struct {
	Key   string `form:"key" json:"key" xml:"key"`
	Value string `form:"value" json:"value" xml:"value"`
}

type GetValueRequest struct {
	Token string `form:"token" json:"token" xml:"token"  binding:"required"`
	Group string `form:"group" json:"group" xml:"group" binding:"required"`
	Key   string `form:"key" json:"key" xml:"key" binding:"required"`
}

type GetValueResponse struct {
	Code  int    `form:"code" json:"code" xml:"code"`
	Msg   string `form:"msg" json:"msg" xml:"msg"`
	Value string `form:"value" json:"value" xml:"value"`
}

type GetAllValueRequest struct {
	Token string `form:"token" json:"token" xml:"token"  binding:"required"`
	Group string `form:"group" json:"group" xml:"group" binding:"required"`
}

type GetAllValueResponse struct {
	Code int    `form:"code" json:"code" xml:"code"`
	Msg  string `form:"msg" json:"msg" xml:"msg"`

	Cache []Cache `form:"cache" json:"cache" xml:"cache"`
}

type DelValueRequest struct {
	Token string `form:"token" json:"token" xml:"token"  binding:"required"`
	Key   string `form:"key" json:"key" xml:"key" binding:"required"`
	Group string `form:"group" json:"group" xml:"group" binding:"required"`
}

type DelValueResponse struct {
	Code int    `form:"code" json:"code" xml:"code"`
	Msg  string `form:"msg" json:"msg" xml:"msg"`
}

type SetValueRequest struct {
	Token string `form:"token" json:"token" xml:"token"  binding:"required"`
	Key   string `form:"key" json:"key" xml:"key" binding:"required"`
	Value string `form:"value" json:"value" xml:"value" binding:"required"`
	Group string `form:"group" json:"group" xml:"group" binding:"required"`
}

type SetValueResponse struct {
	Code int    `form:"code" json:"code" xml:"code"`
	Msg  string `form:"msg" json:"msg" xml:"msg"`
}
