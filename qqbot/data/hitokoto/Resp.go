package hitokoto

import "fmt"

type HitoResp struct {
	From     string `json:"from"`
	Hitokoto string `json:"hitokoto"`
	Type     string `json:"type"`
}

func (resp *HitoResp) String() string {
	return fmt.Sprintf("[%s](%s):%s", resp.From, resp.Type, resp.Hitokoto)
}
