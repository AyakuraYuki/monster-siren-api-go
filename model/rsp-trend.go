package model

type RecommendsRsp struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data []*Recommends `json:"data"`
}

func (r *RecommendsRsp) GetCode() int   { return r.Code }
func (r *RecommendsRsp) GetMsg() string { return r.Msg }

type NewsRsp struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data *NewsRspData `json:"data"`
}

func (r *NewsRsp) GetCode() int   { return r.Code }
func (r *NewsRsp) GetMsg() string { return r.Msg }

type NewsRspData struct {
	List []*News `json:"list"`
	End  bool    `json:"end"`
}

type NewsDetailRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *News  `json:"data"`
}

func (r *NewsDetailRsp) GetCode() int   { return r.Code }
func (r *NewsDetailRsp) GetMsg() string { return r.Msg }
