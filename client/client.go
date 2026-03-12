package client

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"

	apiError "github.com/AyakuraYuki/monster-siren-api-go/internal/api-error"
	"github.com/AyakuraYuki/monster-siren-api-go/internal/version"
	"github.com/AyakuraYuki/monster-siren-api-go/model"
)

var (
	// ErrKeywordRequired 搜索接口需要关键词参数，关键词不能为空
	ErrKeywordRequired = errors.New("keyword is required")
)

// Client 是 塞壬唱片-MSR 后端接口的客户端封装，用于发起接口请求以及管理HTTP客户端实例
type Client struct {
	baseURL string
	cli     *resty.Client
}

// NewClient 创建一个 塞壬唱片-MSR 后端接口客户端实例
func NewClient() *Client {
	client := &Client{
		baseURL: "https://monster-siren.hypergryph.com",
	}

	cli := resty.New().
		SetBaseURL(client.baseURL).
		SetTimeout(30*time.Second).
		SetRetryCount(3).
		SetHeader("Accept", "*/*").
		SetHeader("Accept-Language", "zh-CN,zh;q=0.9,ja;q=0.8,en;q=0.7,en-GB;q=0.6,en-US;q=0.5").
		SetHeader("Referer", client.baseURL+"/").
		SetHeader("User-Agent", fmt.Sprintf("monster-siren-api-go/%s go/%s", version.Version, runtime.Version()))
	client.cli = cli

	return client
}

func doGet[R model.BaseRsp](ctx context.Context, client *Client, path string, query ...url.Values) (R, error) {
	var rsp R

	request := client.cli.R().SetContext(ctx)
	if len(query) > 0 && len(query[0]) > 0 {
		request = request.SetQueryParamsFromValues(query[0])
	}

	response, err := request.SetResult(&rsp).Get(path)

	if err != nil {
		return rsp, err
	}

	if response.IsError() {
		return rsp, fmt.Errorf("response error: %v", response.Error())
	}

	if rsp.GetCode() != 0 {
		return rsp, apiError.NewError(rsp.GetCode(), rsp.GetMsg())
	}

	return rsp, nil
}

func doPost[R model.BaseRsp](ctx context.Context, client *Client, path string, body any) (R, error) {
	var rsp R

	response, err := client.cli.R().
		SetContext(ctx).
		SetBody(body).
		SetResult(&rsp).
		Post(path)

	if err != nil {
		return rsp, err
	}

	if response.IsError() {
		return rsp, fmt.Errorf("response error: %v", response.Error())
	}

	if rsp.GetCode() != 0 {
		return rsp, apiError.NewError(rsp.GetCode(), rsp.GetMsg())
	}

	return rsp, nil
}
