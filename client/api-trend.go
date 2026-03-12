package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/AyakuraYuki/monster-siren-api-go/model"
)

// Recommends 获取动向左侧轮播图的数据
func (c *Client) Recommends(ctx context.Context) ([]*model.Recommends, error) {
	rsp, err := doGet[*model.RecommendsRsp](ctx, c, "/api/recommends")
	if err != nil {
		return nil, err
	}
	return rsp.Data, nil
}

// News 获取动向右侧新闻无限滚动列表信息
//
//	[param] lastCid: 上一次请求中列表内最后一项的cid，请求第一页传空字符串
//	[return] end: 如果为 true 则代表没有更多新闻了
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

// NewsDetail 获取一个新闻的详细信息
func (c *Client) NewsDetail(ctx context.Context, cid string) (*model.News, error) {
	api := fmt.Sprintf("/api/news/%s", cid)
	rsp, err := doGet[*model.NewsDetailRsp](ctx, c, api)
	if err != nil {
		return nil, err
	}
	return rsp.Data, nil
}
