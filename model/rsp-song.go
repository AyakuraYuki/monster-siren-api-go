package model

type SongsRsp struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Data *SongsRspData `json:"data"`
}

func (r *SongsRsp) GetCode() int   { return r.Code }
func (r *SongsRsp) GetMsg() string { return r.Msg }

type SongsRspData struct {
	List     []*Song `json:"list"`
	Autoplay string  `json:"autoplay"`
}

type SongRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *Song  `json:"data"`
}

func (r *SongRsp) GetCode() int   { return r.Code }
func (r *SongRsp) GetMsg() string { return r.Msg }
