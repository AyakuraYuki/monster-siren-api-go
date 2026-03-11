package client

import (
	"context"
	"net/url"

	"github.com/AyakuraYuki/monster-siren-api-go/model"
)

// Search 混合搜索专辑和新闻，该接口不支持分页
func (c *Client) Search(ctx context.Context, keyword string) (*model.SearchRspData, error) {
	query := url.Values{}
	query.Set("keyword", keyword)
	rsp, err := doGet[*model.SearchRsp](ctx, c, "/api/search", query)
	if err != nil {
		return nil, err
	}
	return rsp.Data, nil
}

// SearchAlbum 搜索专辑
// 获取第一页时，lastCid 传空字符串；获取后续页时，lastCid 传 [model.Album.Cid]。
func (c *Client) SearchAlbum(ctx context.Context, keyword, lastCid string) (albums []*model.Album, end bool, err error) {
	query := url.Values{}
	query.Set("keyword", keyword)
	if lastCid != "" {
		query.Set("lastCid", lastCid)
	}
	rsp, err := doGet[*model.SearchAlbumRsp](ctx, c, "/api/search/album", query)
	if err != nil {
		return nil, true, err
	}
	return rsp.Data.List, rsp.Data.End, nil
}

// SearchNews 搜索新闻
// 获取第一页时，lastCid 传空字符串；获取后续页时，lastCid 传 [model.News.Cid]。
func (c *Client) SearchNews(ctx context.Context, keyword, lastCid string) (news []*model.News, end bool, err error) {
	query := url.Values{}
	query.Set("keyword", keyword)
	if lastCid != "" {
		query.Set("lastCid", lastCid)
	}
	rsp, err := doGet[*model.SearchNewsRsp](ctx, c, "/api/search/news", query)
	if err != nil {
		return nil, true, err
	}
	return rsp.Data.List, rsp.Data.End, nil
}
