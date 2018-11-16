package model

type (
	Search struct {
		Config *SearchConfig `json:"config" yaml:"config"`
		Docs   []*SearchDoc  `json:"docs" yaml:"docs"`
	}

	SearchConfig struct {
		Lang          []string `json:"lang" yaml:"lang"`
		PrebuildIndex bool     `json:"prebuild_index" yaml:"prebuild_index"`
		Separator     string   `json:"separator" yaml:"separator"`
	}

	SearchDoc struct {
		Location string `json:"location" yaml:"location"`
		Text     string `json:"text" yaml:"text"`
		Title    string `json:"title" yaml:"title"`
	}
)

func NewSearch() *Search {
	return &Search{}
}

func (this *Search) AddDoc(doc *SearchDoc) {
	this.Docs = append(this.Docs, doc)
}
