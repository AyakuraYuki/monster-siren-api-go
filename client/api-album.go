package client

import (
	"context"
	"fmt"

	"github.com/AyakuraYuki/monster-siren-api-go/model"
)

// Albums 获取所有专辑，包含专辑封面以及专辑
func (c *Client) Albums(ctx context.Context) (albums []*model.Album, err error) {
	rsp, err := doGet[*model.AlbumsRsp](ctx, c, "/api/albums")
	if err != nil {
		return make([]*model.Album, 0), err
	}
	return rsp.Data, nil
}

// Album 获取专辑的基本信息，包含专辑封面和一张大图封面，以及创作专辑的艺术家
func (c *Client) Album(ctx context.Context, albumID string) (*model.Album, error) {
	api := fmt.Sprintf("/api/album/%s/data", albumID)
	rsp, err := doGet[*model.AlbumRsp](ctx, c, api)
	if err != nil {
		return nil, err
	}
	return rsp.Data, nil
}

// AlbumDetail 获取专辑的基本信息，包含专辑封面和一张大图封面，以及专辑相关的所有歌曲
func (c *Client) AlbumDetail(ctx context.Context, albumID string) (*model.Album, error) {
	api := fmt.Sprintf("/api/album/%s/detail", albumID)
	rsp, err := doGet[*model.AlbumRsp](ctx, c, api)
	if err != nil {
		return nil, err
	}
	return rsp.Data, nil
}
