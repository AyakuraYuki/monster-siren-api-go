package model

type AlbumsRsp struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []*Album `json:"data"`
}

func (r *AlbumsRsp) GetCode() int   { return r.Code }
func (r *AlbumsRsp) GetMsg() string { return r.Msg }

type AlbumRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *Album `json:"data"`
}

func (r *AlbumRsp) GetCode() int   { return r.Code }
func (r *AlbumRsp) GetMsg() string { return r.Msg }
