package model

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

func (s *Song) Exists() bool {
	return s != nil && s.Cid != ""
}

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

func (a *Album) Exists() bool {
	return a != nil && a.Cid != ""
}

type Recommends struct {
	Title       string          `json:"title"`
	CoverUrl    string          `json:"coverUrl"`
	Cover       RecommendsCover `json:"cover"`
	Description string          `json:"description"`
	Type        int             `json:"type"`
	Data        string          `json:"data"`
}

type RecommendsCover struct {
	Private bool   `json:"private"`
	Path    string `json:"path"`
}

type News struct {
	Cid     string `json:"cid"`
	Title   string `json:"title"`
	Cate    int    `json:"cate"`
	Author  string `json:"author,omitempty"`
	Content string `json:"content,omitempty"`
	Date    string `json:"date"`
}

func (n *News) Exists() bool {
	return n != nil && n.Cid != ""
}

type FontSet struct {
	SansRegular FontSansRegular `json:"Sans-Regular"`
	SansBold    FontSansBold    `json:"Sans-Bold"`
}

type FontSansRegular struct {
	Tt   string `json:"tt"`
	Eot  string `json:"eot"`
	Svg  string `json:"svg"`
	Woff string `json:"woff"`
}

type FontSansBold struct {
	Tt   string `json:"tt"`
	Eot  string `json:"eot"`
	Svg  string `json:"svg"`
	Woff string `json:"woff"`
}
