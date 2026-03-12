# SDK文档

# client

```go
import "github.com/AyakuraYuki/monster-siren-api-go/client"
```

## Index

- [Variables](<#variables>)
- [type Client](<#Client>)
    - [func NewClient\(\) \*Client](<#NewClient>)
    - [func \(c \*Client\) Album\(ctx context.Context, albumID string\) \(\*model.Album, error\)](<#Client.Album>)
    - [func \(c \*Client\) AlbumDetail\(ctx context.Context, albumID string\) \(\*model.Album, error\)](<#Client.AlbumDetail>)
    - [func \(c \*Client\) Albums\(ctx context.Context\) \(albums \[\]\*model.Album, err error\)](<#Client.Albums>)
    - [func \(c \*Client\) FontSet\(ctx context.Context\) \(\*model.FontSet, error\)](<#Client.FontSet>)
    - [func \(c \*Client\) News\(ctx context.Context, lastCid string\) \(news \[\]\*model.News, end bool, err error\)](<#Client.News>)
    - [func \(c \*Client\) NewsDetail\(ctx context.Context, cid string\) \(\*model.News, error\)](<#Client.NewsDetail>)
    - [func \(c \*Client\) Recommends\(ctx context.Context\) \(\[\]\*model.Recommends, error\)](<#Client.Recommends>)
    - [func \(c \*Client\) Search\(ctx context.Context, keyword string\) \(\*model.SearchRspData, error\)](<#Client.Search>)
    - [func \(c \*Client\) SearchAlbum\(ctx context.Context, keyword, lastCid string\) \(albums \[\]\*model.Album, end bool, err error\)](<#Client.SearchAlbum>)
    - [func \(c \*Client\) SearchNews\(ctx context.Context, keyword, lastCid string\) \(news \[\]\*model.News, end bool, err error\)](<#Client.SearchNews>)
    - [func \(c \*Client\) Song\(ctx context.Context, songID string\) \(\*model.Song, error\)](<#Client.Song>)
    - [func \(c \*Client\) Songs\(ctx context.Context\) \(songs \[\]\*model.Song, autoplay string, err error\)](<#Client.Songs>)

## Variables

<a name="ErrKeywordRequired"></a>

```go
var (
// ErrKeywordRequired 搜索接口需要关键词参数，关键词不能为空
ErrKeywordRequired = errors.New("keyword is required")
)
```

<a name="Client"></a>

## type [Client](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/client/client.go#L24-L27>)

Client 是 塞壬唱片\-MSR 后端接口的客户端封装，用于发起接口请求以及管理HTTP客户端实例

```go
type Client struct {
// contains filtered or unexported fields
}
```

<a name="NewClient"></a>

### func [NewClient](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/client/client.go#L30>)

```go
func NewClient() *Client
```

NewClient 创建一个 塞壬唱片\-MSR 后端接口客户端实例

<a name="Client.Album"></a>

### func \(\*Client\) [Album](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/client/api-album.go#L20>)

```go
func (c *Client) Album(ctx context.Context, albumID string) (*model.Album, error)
```

Album 获取专辑的基本信息，包含专辑封面和一张大图封面，以及创作专辑的艺术家

<a name="Client.AlbumDetail"></a>

### func \(\*Client\) [AlbumDetail](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/client/api-album.go#L30>)

```go
func (c *Client) AlbumDetail(ctx context.Context, albumID string) (*model.Album, error)
```

AlbumDetail 获取专辑的基本信息，包含专辑封面和一张大图封面，以及专辑相关的所有歌曲

<a name="Client.Albums"></a>

### func \(\*Client\) [Albums](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/client/api-album.go#L11>)

```go
func (c *Client) Albums(ctx context.Context) (albums []*model.Album, err error)
```

Albums 获取所有专辑，包含专辑封面以及专辑

<a name="Client.FontSet"></a>

### func \(\*Client\) [FontSet](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/client/api-fontset.go#L10>)

```go
func (c *Client) FontSet(ctx context.Context) (*model.FontSet, error)
```

FontSet 获取网页用字体文件

<a name="Client.News"></a>

### func \(\*Client\) [News](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/client/api-trend.go#L24>)

```go
func (c *Client) News(ctx context.Context, lastCid string) (news []*model.News, end bool, err error)
```

News 获取动向右侧新闻无限滚动列表信息

```
[param] lastCid: 上一次请求中列表内最后一项的cid，请求第一页传空字符串
[return] end: 如果为 true 则代表没有更多新闻了
```

<a name="Client.NewsDetail"></a>

### func \(\*Client\) [NewsDetail](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/client/api-trend.go#L37>)

```go
func (c *Client) NewsDetail(ctx context.Context, cid string) (*model.News, error)
```

NewsDetail 获取一个新闻的详细信息

<a name="Client.Recommends"></a>

### func \(\*Client\) [Recommends](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/client/api-trend.go#L12>)

```go
func (c *Client) Recommends(ctx context.Context) ([]*model.Recommends, error)
```

Recommends 获取动向左侧轮播图的数据

<a name="Client.Search"></a>

### func \(\*Client\) [Search](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/client/api-search.go#L12>)

```go
func (c *Client) Search(ctx context.Context, keyword string) (*model.SearchRspData, error)
```

Search 搜索专辑与动向 该接口不支持分页

<a name="Client.SearchAlbum"></a>

### func \(\*Client\) [SearchAlbum](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/client/api-search.go#L31>)

```go
func (c *Client) SearchAlbum(ctx context.Context, keyword, lastCid string) (albums []*model.Album, end bool, err error)
```

SearchAlbum 通过关键词获取专辑列表

```
[param] keyword: 搜索关键词
[param] lastCid: 上一次请求中列表内最后一项的cid，请求第一页传空字符串
[return] end: 如果为 true 则代表没有更多新闻了
```

<a name="Client.SearchNews"></a>

### func \(\*Client\) [SearchNews](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/client/api-search.go#L53>)

```go
func (c *Client) SearchNews(ctx context.Context, keyword, lastCid string) (news []*model.News, end bool, err error)
```

SearchNews 通过关键词获取动向列表

```
[param] keyword: 搜索关键词
[param] lastCid: 上一次请求中列表内最后一项的cid，请求第一页传空字符串
[return] end: 如果为 true 则代表没有更多新闻了
```

<a name="Client.Song"></a>

### func \(\*Client\) [Song](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/client/api-song.go#L20>)

```go
func (c *Client) Song(ctx context.Context, songID string) (*model.Song, error)
```

Song 获取一首歌的基本信息，包含音频URL和歌词等文件，但不包含专辑图片

<a name="Client.Songs"></a>

### func \(\*Client\) [Songs](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/client/api-song.go#L11>)

```go
func (c *Client) Songs(ctx context.Context) (songs []*model.Song, autoplay string, err error)
```

Songs 获取所有歌曲，和官网自动播放的歌曲的ID

# model

```go
import "github.com/AyakuraYuki/monster-siren-api-go/model"
```

## Index

- [type Album](<#Album>)
    - [func \(a \*Album\) Exists\(\) bool](<#Album.Exists>)
- [type AlbumRsp](<#AlbumRsp>)
    - [func \(r \*AlbumRsp\) GetCode\(\) int](<#AlbumRsp.GetCode>)
    - [func \(r \*AlbumRsp\) GetMsg\(\) string](<#AlbumRsp.GetMsg>)
- [type AlbumsRsp](<#AlbumsRsp>)
    - [func \(r \*AlbumsRsp\) GetCode\(\) int](<#AlbumsRsp.GetCode>)
    - [func \(r \*AlbumsRsp\) GetMsg\(\) string](<#AlbumsRsp.GetMsg>)
- [type BaseRsp](<#BaseRsp>)
- [type FontSansBold](<#FontSansBold>)
- [type FontSansRegular](<#FontSansRegular>)
- [type FontSet](<#FontSet>)
- [type FontSetRsp](<#FontSetRsp>)
    - [func \(r \*FontSetRsp\) GetCode\(\) int](<#FontSetRsp.GetCode>)
    - [func \(r \*FontSetRsp\) GetMsg\(\) string](<#FontSetRsp.GetMsg>)
- [type News](<#News>)
    - [func \(n \*News\) Exists\(\) bool](<#News.Exists>)
- [type NewsDetailRsp](<#NewsDetailRsp>)
    - [func \(r \*NewsDetailRsp\) GetCode\(\) int](<#NewsDetailRsp.GetCode>)
    - [func \(r \*NewsDetailRsp\) GetMsg\(\) string](<#NewsDetailRsp.GetMsg>)
- [type NewsRsp](<#NewsRsp>)
    - [func \(r \*NewsRsp\) GetCode\(\) int](<#NewsRsp.GetCode>)
    - [func \(r \*NewsRsp\) GetMsg\(\) string](<#NewsRsp.GetMsg>)
- [type NewsRspData](<#NewsRspData>)
- [type Recommends](<#Recommends>)
- [type RecommendsCover](<#RecommendsCover>)
- [type RecommendsRsp](<#RecommendsRsp>)
    - [func \(r \*RecommendsRsp\) GetCode\(\) int](<#RecommendsRsp.GetCode>)
    - [func \(r \*RecommendsRsp\) GetMsg\(\) string](<#RecommendsRsp.GetMsg>)
- [type SearchAlbumRsp](<#SearchAlbumRsp>)
    - [func \(r \*SearchAlbumRsp\) GetCode\(\) int](<#SearchAlbumRsp.GetCode>)
    - [func \(r \*SearchAlbumRsp\) GetMsg\(\) string](<#SearchAlbumRsp.GetMsg>)
- [type SearchAlbumRspData](<#SearchAlbumRspData>)
- [type SearchNewsRsp](<#SearchNewsRsp>)
    - [func \(r \*SearchNewsRsp\) GetCode\(\) int](<#SearchNewsRsp.GetCode>)
    - [func \(r \*SearchNewsRsp\) GetMsg\(\) string](<#SearchNewsRsp.GetMsg>)
- [type SearchNewsRspData](<#SearchNewsRspData>)
- [type SearchRsp](<#SearchRsp>)
    - [func \(r \*SearchRsp\) GetCode\(\) int](<#SearchRsp.GetCode>)
    - [func \(r \*SearchRsp\) GetMsg\(\) string](<#SearchRsp.GetMsg>)
- [type SearchRspData](<#SearchRspData>)
- [type Song](<#Song>)
    - [func \(s \*Song\) Exists\(\) bool](<#Song.Exists>)
- [type SongRsp](<#SongRsp>)
    - [func \(r \*SongRsp\) GetCode\(\) int](<#SongRsp.GetCode>)
    - [func \(r \*SongRsp\) GetMsg\(\) string](<#SongRsp.GetMsg>)
- [type SongsRsp](<#SongsRsp>)
    - [func \(r \*SongsRsp\) GetCode\(\) int](<#SongsRsp.GetCode>)
    - [func \(r \*SongsRsp\) GetMsg\(\) string](<#SongsRsp.GetMsg>)
- [type SongsRspData](<#SongsRspData>)

<a name="Album"></a>

## type [Album](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/model.go#L21-L30>)

Album 专辑

```go
type Album struct {
Cid        string   `json:"cid"`
Name       string   `json:"name"`
Intro      string   `json:"intro,omitempty"`
Belong     string   `json:"belong,omitempty"`
CoverUrl   string   `json:"coverUrl,omitempty"`
CoverDeUrl string   `json:"coverDeUrl,omitempty"`
Artistes   []string `json:"artistes,omitempty"`
Songs      []*Song  `json:"songs,omitempty"`
}
```

<a name="Album.Exists"></a>

### func \(\*Album\) [Exists](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/model.go#L32>)

```go
func (a *Album) Exists() bool
```

<a name="AlbumRsp"></a>

## type [AlbumRsp](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-album.go#L12-L16>)

```go
type AlbumRsp struct {
Code int    `json:"code"`
Msg  string `json:"msg"`
Data *Album `json:"data"`
}
```

<a name="AlbumRsp.GetCode"></a>

### func \(\*AlbumRsp\) [GetCode](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-album.go#L18>)

```go
func (r *AlbumRsp) GetCode() int
```

<a name="AlbumRsp.GetMsg"></a>

### func \(\*AlbumRsp\) [GetMsg](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-album.go#L19>)

```go
func (r *AlbumRsp) GetMsg() string
```

<a name="AlbumsRsp"></a>

## type [AlbumsRsp](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-album.go#L3-L7>)

```go
type AlbumsRsp struct {
Code int      `json:"code"`
Msg  string   `json:"msg"`
Data []*Album `json:"data"`
}
```

<a name="AlbumsRsp.GetCode"></a>

### func \(\*AlbumsRsp\) [GetCode](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-album.go#L9>)

```go
func (r *AlbumsRsp) GetCode() int
```

<a name="AlbumsRsp.GetMsg"></a>

### func \(\*AlbumsRsp\) [GetMsg](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-album.go#L10>)

```go
func (r *AlbumsRsp) GetMsg() string
```

<a name="BaseRsp"></a>

## type [BaseRsp](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp.go#L3-L6>)

```go
type BaseRsp interface {
GetCode() int
GetMsg() string
}
```

<a name="FontSansBold"></a>

## type [FontSansBold](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/model.go#L81-L86>)

FontSansBold 字体集粗体字体

```go
type FontSansBold struct {
Tt   string `json:"tt"`
Eot  string `json:"eot"`
Svg  string `json:"svg"`
Woff string `json:"woff"`
}
```

<a name="FontSansRegular"></a>

## type [FontSansRegular](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/model.go#L73-L78>)

FontSansRegular 字体集常规字体

```go
type FontSansRegular struct {
Tt   string `json:"tt"`
Eot  string `json:"eot"`
Svg  string `json:"svg"`
Woff string `json:"woff"`
}
```

<a name="FontSet"></a>

## type [FontSet](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/model.go#L67-L70>)

FontSet 字体集

```go
type FontSet struct {
SansRegular FontSansRegular `json:"Sans-Regular"`
SansBold    FontSansBold    `json:"Sans-Bold"`
}
```

<a name="FontSetRsp"></a>

## type [FontSetRsp](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-fontset.go#L3-L7>)

```go
type FontSetRsp struct {
Code int      `json:"code"`
Msg  string   `json:"msg"`
Data *FontSet `json:"data"`
}
```

<a name="FontSetRsp.GetCode"></a>

### func \(\*FontSetRsp\) [GetCode](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-fontset.go#L9>)

```go
func (r *FontSetRsp) GetCode() int
```

<a name="FontSetRsp.GetMsg"></a>

### func \(\*FontSetRsp\) [GetMsg](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-fontset.go#L10>)

```go
func (r *FontSetRsp) GetMsg() string
```

<a name="News"></a>

## type [News](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/model.go#L53-L60>)

News 动向

```go
type News struct {
Cid     string `json:"cid"`
Title   string `json:"title"`
Cate    int    `json:"cate"`
Author  string `json:"author,omitempty"`
Content string `json:"content,omitempty"`
Date    string `json:"date"`
}
```

<a name="News.Exists"></a>

### func \(\*News\) [Exists](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/model.go#L62>)

```go
func (n *News) Exists() bool
```

<a name="NewsDetailRsp"></a>

## type [NewsDetailRsp](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-trend.go#L26-L30>)

```go
type NewsDetailRsp struct {
Code int    `json:"code"`
Msg  string `json:"msg"`
Data *News  `json:"data"`
}
```

<a name="NewsDetailRsp.GetCode"></a>

### func \(\*NewsDetailRsp\) [GetCode](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-trend.go#L32>)

```go
func (r *NewsDetailRsp) GetCode() int
```

<a name="NewsDetailRsp.GetMsg"></a>

### func \(\*NewsDetailRsp\) [GetMsg](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-trend.go#L33>)

```go
func (r *NewsDetailRsp) GetMsg() string
```

<a name="NewsRsp"></a>

## type [NewsRsp](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-trend.go#L12-L16>)

```go
type NewsRsp struct {
Code int          `json:"code"`
Msg  string       `json:"msg"`
Data *NewsRspData `json:"data"`
}
```

<a name="NewsRsp.GetCode"></a>

### func \(\*NewsRsp\) [GetCode](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-trend.go#L18>)

```go
func (r *NewsRsp) GetCode() int
```

<a name="NewsRsp.GetMsg"></a>

### func \(\*NewsRsp\) [GetMsg](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-trend.go#L19>)

```go
func (r *NewsRsp) GetMsg() string
```

<a name="NewsRspData"></a>

## type [NewsRspData](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-trend.go#L21-L24>)

```go
type NewsRspData struct {
List []*News `json:"list"`
End  bool    `json:"end"`
}
```

<a name="Recommends"></a>

## type [Recommends](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/model.go#L37-L44>)

Recommends 轮播动向

```go
type Recommends struct {
Title       string          `json:"title"`
CoverUrl    string          `json:"coverUrl"`
Cover       RecommendsCover `json:"cover"`
Description string          `json:"description"`
Type        int             `json:"type"`
Data        string          `json:"data"`
}
```

<a name="RecommendsCover"></a>

## type [RecommendsCover](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/model.go#L47-L50>)

RecommendsCover 轮播动向的图片

```go
type RecommendsCover struct {
Private bool   `json:"private"`
Path    string `json:"path"`
}
```

<a name="RecommendsRsp"></a>

## type [RecommendsRsp](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-trend.go#L3-L7>)

```go
type RecommendsRsp struct {
Code int           `json:"code"`
Msg  string        `json:"msg"`
Data []*Recommends `json:"data"`
}
```

<a name="RecommendsRsp.GetCode"></a>

### func \(\*RecommendsRsp\) [GetCode](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-trend.go#L9>)

```go
func (r *RecommendsRsp) GetCode() int
```

<a name="RecommendsRsp.GetMsg"></a>

### func \(\*RecommendsRsp\) [GetMsg](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-trend.go#L10>)

```go
func (r *RecommendsRsp) GetMsg() string
```

<a name="SearchAlbumRsp"></a>

## type [SearchAlbumRsp](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-search.go#L17-L21>)

```go
type SearchAlbumRsp struct {
Code int                 `json:"code"`
Msg  string              `json:"msg"`
Data *SearchAlbumRspData `json:"data"`
}
```

<a name="SearchAlbumRsp.GetCode"></a>

### func \(\*SearchAlbumRsp\) [GetCode](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-search.go#L23>)

```go
func (r *SearchAlbumRsp) GetCode() int
```

<a name="SearchAlbumRsp.GetMsg"></a>

### func \(\*SearchAlbumRsp\) [GetMsg](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-search.go#L24>)

```go
func (r *SearchAlbumRsp) GetMsg() string
```

<a name="SearchAlbumRspData"></a>

## type [SearchAlbumRspData](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-search.go#L26-L29>)

```go
type SearchAlbumRspData struct {
List []*Album `json:"list"`
End  bool     `json:"end"`
}
```

<a name="SearchNewsRsp"></a>

## type [SearchNewsRsp](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-search.go#L30-L34>)

```go
type SearchNewsRsp struct {
Code int                `json:"code"`
Msg  string             `json:"msg"`
Data *SearchNewsRspData `json:"data"`
}
```

<a name="SearchNewsRsp.GetCode"></a>

### func \(\*SearchNewsRsp\) [GetCode](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-search.go#L36>)

```go
func (r *SearchNewsRsp) GetCode() int
```

<a name="SearchNewsRsp.GetMsg"></a>

### func \(\*SearchNewsRsp\) [GetMsg](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-search.go#L37>)

```go
func (r *SearchNewsRsp) GetMsg() string
```

<a name="SearchNewsRspData"></a>

## type [SearchNewsRspData](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-search.go#L39-L42>)

```go
type SearchNewsRspData struct {
List []*News `json:"list"`
End  bool    `json:"end"`
}
```

<a name="SearchRsp"></a>

## type [SearchRsp](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-search.go#L3-L7>)

```go
type SearchRsp struct {
Code int            `json:"code"`
Msg  string         `json:"msg"`
Data *SearchRspData `json:"data"`
}
```

<a name="SearchRsp.GetCode"></a>

### func \(\*SearchRsp\) [GetCode](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-search.go#L9>)

```go
func (r *SearchRsp) GetCode() int
```

<a name="SearchRsp.GetMsg"></a>

### func \(\*SearchRsp\) [GetMsg](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-search.go#L10>)

```go
func (r *SearchRsp) GetMsg() string
```

<a name="SearchRspData"></a>

## type [SearchRspData](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-search.go#L12-L15>)

```go
type SearchRspData struct {
Albums *SearchAlbumRspData `json:"albums"`
News   *SearchNewsRspData  `json:"news"`
}
```

<a name="Song"></a>

## type [Song](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/model.go#L4-L14>)

Song 歌曲

```go
type Song struct {
Cid        string   `json:"cid"`
Name       string   `json:"name"`
AlbumCid   string   `json:"albumCid,omitempty"`
SourceUrl  string   `json:"sourceUrl,omitempty"`
LyricUrl   string   `json:"lyricUrl,omitempty"`
MvUrl      string   `json:"mvUrl,omitempty"`
MvCoverUrl string   `json:"mvCoverUrl,omitempty"`
Artists    []string `json:"artists,omitempty"`
Artistes   []string `json:"artistes,omitempty"`
}
```

<a name="Song.Exists"></a>

### func \(\*Song\) [Exists](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/model.go#L16>)

```go
func (s *Song) Exists() bool
```

<a name="SongRsp"></a>

## type [SongRsp](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-song.go#L17-L21>)

```go
type SongRsp struct {
Code int    `json:"code"`
Msg  string `json:"msg"`
Data *Song  `json:"data"`
}
```

<a name="SongRsp.GetCode"></a>

### func \(\*SongRsp\) [GetCode](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-song.go#L23>)

```go
func (r *SongRsp) GetCode() int
```

<a name="SongRsp.GetMsg"></a>

### func \(\*SongRsp\) [GetMsg](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-song.go#L24>)

```go
func (r *SongRsp) GetMsg() string
```

<a name="SongsRsp"></a>

## type [SongsRsp](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-song.go#L3-L7>)

```go
type SongsRsp struct {
Code int           `json:"code"`
Msg  string        `json:"msg"`
Data *SongsRspData `json:"data"`
}
```

<a name="SongsRsp.GetCode"></a>

### func \(\*SongsRsp\) [GetCode](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-song.go#L9>)

```go
func (r *SongsRsp) GetCode() int
```

<a name="SongsRsp.GetMsg"></a>

### func \(\*SongsRsp\) [GetMsg](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-song.go#L10>)

```go
func (r *SongsRsp) GetMsg() string
```

<a name="SongsRspData"></a>

## type [SongsRspData](<https://github.com/AyakuraYuki/monster-siren-api-go/blob/main/model/rsp-song.go#L12-L15>)

```go
type SongsRspData struct {
List     []*Song `json:"list"`
Autoplay string  `json:"autoplay"`
}
```
