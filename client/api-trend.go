package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/AyakuraYuki/monster-siren-api-go/model"
)

// Recommends 获取官网「动向」页面的轮播信息
func (c *Client) Recommends(ctx context.Context) ([]*model.Recommends, error) {
	rsp, err := doGet[*model.RecommendsRsp](ctx, c, "/api/recommends")
	if err != nil {
		return nil, err
	}
	return rsp.Data, nil
}

// News 获取官网「动向」页面的新闻列表
// 获取第一页时，lastCid 传空字符串；获取后续页时，lastCid 传 [model.News.Cid]。
func (c *Client) News(ctx context.Context, lastCid string) (news []*model.News, end bool, err error) {
	query := url.Values{}
	if lastCid != "" {
		query.Set("lastCid", lastCid)
	}
	rsp, err := doGet[*model.NewsRsp](ctx, c, "/api/news", query)
	if err != nil {
		return make([]*model.News, 0), true, err
	}
	return rsp.Data.List, rsp.Data.End, nil
}

// NewsDetail 获取新闻详情
// cid 是 [model.News.Cid]，来自 News 接口返回的新闻列表。
func (c *Client) NewsDetail(ctx context.Context, cid string) (*model.News, error) {
	api := fmt.Sprintf("/api/news/%s", cid)
	rsp, err := doGet[*model.NewsDetailRsp](ctx, c, api)
	if err != nil {
		return nil, err
	}
	return rsp.Data, nil
}
