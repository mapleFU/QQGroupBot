package search

import "fmt"

type Doc struct {
	//From float64 `json:"from"`
	//To float64 `json:"to"`
	//At float64 `json:"at"`
	// Number, "OVA/OAD", "Special", ""
	Episode    interface{} `json:"episode"`
	Season     string      `json:"season"`
	Similarity float64     `json:"similarity"`
	//TitleChinese string `json:"title_chinese"`
	TitleEnglish string `json:"title_english"`
	FileName     string `json:"file_name"`
	MalID        int64  `json:"mal_id"`
}

func (doc *Doc) String() string {
	return fmt.Sprintf("来自动画%s(%s), 出处是%v. MAL的动画 ID 是 %d, 相似度 %v",
		doc.TitleEnglish, doc.FileName, doc.Episode, doc.MalID, doc.Similarity)
}

type SearchResult struct {
	Docs  []Doc `json:"docs"`
	Trial int64 `json:"trial"`
}

func (result *SearchResult) String() string {
	retString := fmt.Sprintf("搜索尝试%v次\n", result.Trial)
	for i, v := range result.Docs {
		retString += fmt.Sprintf("(%d): %s\n", i, v.String())
	}
	return retString
}
