package model

type SongsRsp struct {
	Data *SongsRspData `json:"data"`
	Msg  string        `json:"msg"`
	Code int           `json:"code"`
}

func (r *SongsRsp) GetCode() int   { return r.Code }
func (r *SongsRsp) GetMsg() string { return r.Msg }

type SongsRspData struct {
	Autoplay string  `json:"autoplay"`
	List     []*Song `json:"list"`
}

type SongRsp struct {
	Data *Song  `json:"data"`
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func (r *SongRsp) GetCode() int   { return r.Code }
func (r *SongRsp) GetMsg() string { return r.Msg }
