package model

type SearchRsp struct {
	Data *SearchRspData `json:"data"`
	Msg  string         `json:"msg"`
	Code int            `json:"code"`
}

func (r *SearchRsp) GetCode() int   { return r.Code }
func (r *SearchRsp) GetMsg() string { return r.Msg }

type SearchRspData struct {
	Albums *SearchAlbumRspData `json:"albums"`
	News   *SearchNewsRspData  `json:"news"`
}

type SearchAlbumRsp struct {
	Data *SearchAlbumRspData `json:"data"`
	Msg  string              `json:"msg"`
	Code int                 `json:"code"`
}

func (r *SearchAlbumRsp) GetCode() int   { return r.Code }
func (r *SearchAlbumRsp) GetMsg() string { return r.Msg }

type SearchAlbumRspData struct {
	List []*Album `json:"list"`
	End  bool     `json:"end"`
}
type SearchNewsRsp struct {
	Data *SearchNewsRspData `json:"data"`
	Msg  string             `json:"msg"`
	Code int                `json:"code"`
}

func (r *SearchNewsRsp) GetCode() int   { return r.Code }
func (r *SearchNewsRsp) GetMsg() string { return r.Msg }

type SearchNewsRspData struct {
	List []*News `json:"list"`
	End  bool    `json:"end"`
}
