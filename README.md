# monster-siren-api-go

塞壬唱片-MSR 服务端 API Go SDK

## 版本

Go 1.26 及以上

## 使用方法

```shell
go get -u "github.com/AyakuraYuki/monster-siren-api-go/client"
```

### 示例

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AyakuraYuki/monster-siren-api-go/client"
)

func main() {
	cli := client.NewClient()

	// 用例 1：获取所有歌曲
	songs, autoplay, err := cli.Songs(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("官网自动播放: %s\n", autoplay)
	fmt.Printf("目前已发布歌曲数: %d\n", len(songs))

	// 用例 2：搜索专辑
	albums, end, err := cli.SearchAlbum(context.Background(), "危机合约", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("已找到 %d 张符合关键词的专辑\n", len(albums))
	for !end {
		albums, end, err = cli.SearchAlbum(context.Background(), "危机合约", albums[len(albums)-1].Cid)
		fmt.Printf("已找到 %d 张符合关键词的专辑\n", len(albums))
	}
}

```
