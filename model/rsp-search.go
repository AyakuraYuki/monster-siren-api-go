package model

type SearchRsp struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data *SearchRspData `json:"data"`
}

func (r *SearchRsp) GetCode() int   { return r.Code }
func (r *SearchRsp) GetMsg() string { return r.Msg }

type SearchRspData struct {
	Albums *SearchAlbumRspData `json:"albums"`
	News   *SearchNewsRspData  `json:"news"`
}

type SearchAlbumRsp struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data *SearchAlbumRspData `json:"data"`
}

func (r *SearchAlbumRsp) GetCode() int   { return r.Code }
func (r *SearchAlbumRsp) GetMsg() string { return r.Msg }

type SearchAlbumRspData struct {
	List []*Album `json:"list"`
	End  bool     `json:"end"`
}
type SearchNewsRsp struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data *SearchNewsRspData `json:"data"`
}

func (r *SearchNewsRsp) GetCode() int   { return r.Code }
func (r *SearchNewsRsp) GetMsg() string { return r.Msg }

type SearchNewsRspData struct {
	List []*News `json:"list"`
	End  bool    `json:"end"`
}
