package client_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"

	"github.com/AyakuraYuki/monster-siren-api-go/client"
)

const (
	testAlbumCID = "6661"             // 好得不能再好了！泰拉投资大师课
	testSongCID  = "306807"           // 好得不能再好了！泰拉投资大师课
	testKeyword  = "operation pyrite" // BURN ME TO THE GROUND!!!
)

var (
	cli *client.Client
	ctx context.Context
)

func TestMain(m *testing.M) {
	cli = client.NewClient()
	ctx = context.Background()
	m.Run()
}

func TestClient_Albums(t *testing.T) {
	albums, err := cli.Albums(ctx)
	assert.NoError(t, err)
	assert.Greater(t, len(albums), 0)
	litter.Dump(albums[0])
}

func TestClient_Album(t *testing.T) {
	album, err := cli.Album(ctx, testAlbumCID)
	assert.NoError(t, err)
	assert.NotNil(t, album)
	litter.Dump(album)
}

func TestClient_AlbumDetail(t *testing.T) {
	album, err := cli.AlbumDetail(ctx, testAlbumCID)
	assert.NoError(t, err)
	assert.NotNil(t, album)
	litter.Dump(album)
}

func TestClient_FontSet(t *testing.T) {
	fontset, err := cli.FontSet(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, fontset)
	litter.Dump(fontset)
}

func TestClient_Search(t *testing.T) {
	search, err := cli.Search(ctx, testKeyword)
	assert.NoError(t, err)
	assert.NotNil(t, search)
	litter.Dump(search)
}

func TestClient_SearchAlbum(t *testing.T) {
	search, end, err := cli.SearchAlbum(ctx, testKeyword, "")
	assert.NoError(t, err)
	assert.Greater(t, len(search), 0)
	assert.True(t, end) // 搜索 operation pyrite 仅有一个结果
	litter.Dump(search)
}

func TestClient_SearchAlbum2(t *testing.T) {
	keyword := "危机合约"

	search, end, err := cli.SearchAlbum(ctx, keyword, "")
	assert.NoError(t, err)
	assert.Greater(t, len(search), 0)

	maxLoops := 0
	for !end {
		search, end, err = cli.SearchAlbum(ctx, keyword, search[len(search)-1].Cid)
		assert.NoError(t, err)
		assert.Greater(t, len(search), 0)

		maxLoops++
		if maxLoops > 100 {
			t.Fatal("too many loops")
		}
		time.Sleep(200*time.Millisecond + time.Duration(rand.Int63n(200))*time.Millisecond)
	}
}

func TestClient_SearchNews(t *testing.T) {
	keyword := "正式上架"

	search, end, err := cli.SearchNews(ctx, keyword, "")
	assert.NoError(t, err)
	assert.Greater(t, len(search), 0)

	maxLoops := 0
	for !end {
		search, end, err = cli.SearchNews(ctx, keyword, search[len(search)-1].Cid)
		assert.NoError(t, err)
		assert.Greater(t, len(search), 0)

		maxLoops++
		if maxLoops > 100 {
			t.Fatal("too many loops")
		}
		time.Sleep(200*time.Millisecond + time.Duration(rand.Int63n(200))*time.Millisecond)
	}
}

func TestClient_Songs(t *testing.T) {
	songs, autoplay, err := cli.Songs(ctx)
	assert.NoError(t, err)
	assert.Greater(t, len(songs), 0)
	t.Log("自动播放:", autoplay)
	litter.Dump(songs[0])
}

func TestClient_Song(t *testing.T) {
	song, err := cli.Song(ctx, testSongCID)
	assert.NoError(t, err)
	assert.NotNil(t, song)
	assert.NotEmpty(t, song.SourceUrl)
	litter.Dump(song)
}

func TestClient_Recommends(t *testing.T) {
	recommends, err := cli.Recommends(ctx)
	assert.NoError(t, err)
	assert.Greater(t, len(recommends), 0)
	litter.Dump(recommends[0])
}

func TestClient_News(t *testing.T) {
	news, end, err := cli.News(ctx, "")
	assert.NoError(t, err)
	assert.Greater(t, len(news), 0)

	maxLoops := 0
	for !end {
		news, end, err = cli.News(ctx, news[len(news)-1].Cid)
		assert.NoError(t, err)
		assert.Greater(t, len(news), 0)

		maxLoops++
		if maxLoops > 100 {
			t.Fatal("too many loops")
		}
		time.Sleep(200*time.Millisecond + time.Duration(rand.Int63n(200))*time.Millisecond)
	}
}

func TestClient_NewsDetail(t *testing.T) {
	news, err := cli.NewsDetail(ctx, "578838")
	assert.NoError(t, err)
	assert.NotNil(t, news)
	litter.Dump(news)
}
