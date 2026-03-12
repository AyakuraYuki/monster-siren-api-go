# 接口文档

## 专辑

### GET /api/albums

获取所有专辑，包含专辑封面以及专辑

### GET /api/album/{cid}/data

获取专辑的基本信息，包含专辑封面和一张大图封面，以及创作专辑的艺术家

> 提醒：这个接口拿不到歌曲信息

#### 参数

- `{cid}`: 专辑ID

### GET /api/album/{cid}/detail

获取专辑的基本信息，包含专辑封面和一张大图封面，以及专辑相关的所有歌曲

#### 参数

- `{cid}`: 专辑ID

## 歌曲

### GET /api/songs

获取所有歌曲，和官网自动播放的歌曲的ID

### GET /api/song/{cid}

获取一首歌的基本信息，包含音频URL和歌词等文件，但不包含专辑图片

#### 参数

- `{cid}`: 歌曲ID

## 动向

### GET /api/recommends

获取动向左侧轮播图的数据

### GET /api/news?lastCid={cid}

获取动向右侧新闻无限滚动列表信息

#### 参数

- `lastCid={cid}`: 上一次请求中列表内最后一项的cid，请求第一页传空字符串

### GET /api/news/{cid}

获取一个动向的详细信息

#### 参数

- `{cid}`: 动向ID

## 搜索

### GET /api/search?keyword={keyword}

搜索专辑与动向

> 提醒：这个接口不支持分页

#### 参数

- `keyword={keyword}`: 搜索关键词

### GET /api/search/album?keyword={keyword}&lastCid={cid}

通过关键词获取专辑列表

#### 参数

- `keyword={keyword}`: 搜索关键词
- `lastCid={cid}`: 上一次请求中列表内最后一项的cid，请求第一页传空字符串

### GET /api/search/news?keyword={keyword}&lastCid={cid}

通过关键词获取动向列表

#### 参数

- `keyword={keyword}`: 搜索关键词
- `lastCid={cid}`: 上一次请求中列表内最后一项的cid，请求第一页传空字符串

## 其他

### GET /api/fontset

获取网页用字体文件
