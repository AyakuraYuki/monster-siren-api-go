package model

type RecommendsRsp struct {
	Msg  string        `json:"msg"`
	Data []*Recommends `json:"data"`
	Code int           `json:"code"`
}

func (r *RecommendsRsp) GetCode() int   { return r.Code }
func (r *RecommendsRsp) GetMsg() string { return r.Msg }

type NewsRsp struct {
	Data *NewsRspData `json:"data"`
	Msg  string       `json:"msg"`
	Code int          `json:"code"`
}

func (r *NewsRsp) GetCode() int   { return r.Code }
func (r *NewsRsp) GetMsg() string { return r.Msg }

type NewsRspData struct {
	List []*News `json:"list"`
	End  bool    `json:"end"`
}

type NewsDetailRsp struct {
	Data *News  `json:"data"`
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func (r *NewsDetailRsp) GetCode() int   { return r.Code }
func (r *NewsDetailRsp) GetMsg() string { return r.Msg }
