package client

import (
	"context"
	"net/url"

	"github.com/AyakuraYuki/monster-siren-api-go/model"
)

// Search 搜索专辑与动向
// 该接口不支持分页
func (c *Client) Search(ctx context.Context, keyword string) (*model.SearchRspData, error) {
	if keyword == "" {
		return nil, ErrKeywordRequired
	}

	query := url.Values{}
	query.Set("keyword", keyword)
	rsp, err := doGet[*model.SearchRsp](ctx, c, "/api/search", query)
	if err != nil {
		return nil, err
	}
	return rsp.Data, nil
}

// SearchAlbum 通过关键词获取专辑列表
//
//	[param] keyword: 搜索关键词
//	[param] lastCid: 上一次请求中列表内最后一项的cid，请求第一页传空字符串
//	[return] end: 如果为 true 则代表没有更多新闻了
func (c *Client) SearchAlbum(ctx context.Context, keyword, lastCid string) (albums []*model.Album, end bool, err error) {
	if keyword == "" {
		return make([]*model.Album, 0), true, ErrKeywordRequired
	}

	query := url.Values{}
	query.Set("keyword", keyword)
	if lastCid != "" {
		query.Set("lastCid", lastCid)
	}
	rsp, err := doGet[*model.SearchAlbumRsp](ctx, c, "/api/search/album", query)
	if err != nil {
		return make([]*model.Album, 0), true, err
	}
	return rsp.Data.List, rsp.Data.End, nil
}

// SearchNews 通过关键词获取动向列表
//
//	[param] keyword: 搜索关键词
//	[param] lastCid: 上一次请求中列表内最后一项的cid，请求第一页传空字符串
//	[return] end: 如果为 true 则代表没有更多新闻了
func (c *Client) SearchNews(ctx context.Context, keyword, lastCid string) (news []*model.News, end bool, err error) {
	if keyword == "" {
		return make([]*model.News, 0), true, ErrKeywordRequired
	}

	query := url.Values{}
	query.Set("keyword", keyword)
	if lastCid != "" {
		query.Set("lastCid", lastCid)
	}
	rsp, err := doGet[*model.SearchNewsRsp](ctx, c, "/api/search/news", query)
	if err != nil {
		return make([]*model.News, 0), true, err
	}
	return rsp.Data.List, rsp.Data.End, nil
}
