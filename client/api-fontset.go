package client

import (
	"context"

	"github.com/AyakuraYuki/monster-siren-api-go/model"
)

// FontSet 获取官网使用的字体集文件的下载链接
func (c *Client) FontSet(ctx context.Context) (*model.FontSet, error) {
	rsp, err := doGet[*model.FontSetRsp](ctx, c, "/api/fontset")
	if err != nil {
		return nil, err
	}
	return rsp.Data, nil
}
