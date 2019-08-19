package group

type Sender struct {
	UserID int64 `json:"user_id"`

	Nickname string `json:"nickname"`
	Card     string `json:"card"`
	// 性别，male 或 female 或 unknown
	Sex string `json:"sex"`

	Age int32 `json:"age"`
}
