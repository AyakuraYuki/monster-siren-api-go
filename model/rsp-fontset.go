package model

type FontSetRsp struct {
	Data *FontSet `json:"data"`
	Msg  string   `json:"msg"`
	Code int      `json:"code"`
}

func (r *FontSetRsp) GetCode() int   { return r.Code }
func (r *FontSetRsp) GetMsg() string { return r.Msg }
