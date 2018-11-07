package search

type Doc struct {
	From float64 `json:"from"`
	To float64 `json:"to"`
	At float64 `json:"at"`
	// Number, "OVA/OAD", "Special", ""
	Episode string `json:"episode"`
	Similarity float64 `json:"similarity"`
	TitleChinese string `json:"title_chinese"`
	TitleEnglish string `json:"title_english"`
	
	MalID int64 `json:"mal_id"`
}

type SearchResult struct {
	Docs []Doc `json:"docs,omitempty"`
	Trial int64 `json:"trial"`
}