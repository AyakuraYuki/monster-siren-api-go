package model

type AlbumsRsp struct {
	Msg  string   `json:"msg"`
	Data []*Album `json:"data"`
	Code int      `json:"code"`
}

func (r *AlbumsRsp) GetCode() int   { return r.Code }
func (r *AlbumsRsp) GetMsg() string { return r.Msg }

type AlbumRsp struct {
	Data *Album `json:"data"`
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func (r *AlbumRsp) GetCode() int   { return r.Code }
func (r *AlbumRsp) GetMsg() string { return r.Msg }
