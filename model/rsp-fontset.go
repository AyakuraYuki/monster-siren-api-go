package model

type FontSetRsp struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data *FontSet `json:"data"`
}

func (r *FontSetRsp) GetCode() int   { return r.Code }
func (r *FontSetRsp) GetMsg() string { return r.Msg }
