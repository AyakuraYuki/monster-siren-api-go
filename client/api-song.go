package client

import (
	"context"
	"fmt"

	"github.com/AyakuraYuki/monster-siren-api-go/model"
)

// Songs 获取所有歌曲列表信息，以及官网自动播放的歌曲的ID
func (c *Client) Songs(ctx context.Context) (songs []*model.Song, autoplay string, err error) {
	rsp, err := doGet[*model.SongsRsp](ctx, c, "/api/songs")
	if err != nil {
		return make([]*model.Song, 0), "", err
	}
	return rsp.Data.List, rsp.Data.Autoplay, nil
}

// Song 获取单首歌曲的信息
func (c *Client) Song(ctx context.Context, songID string) (*model.Song, error) {
	api := fmt.Sprintf("/api/song/%s", songID)
	rsp, err := doGet[*model.SongRsp](ctx, c, api)
	if err != nil {
		return nil, err
	}
	return rsp.Data, nil
}
