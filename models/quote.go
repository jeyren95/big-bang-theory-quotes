package models

type Quote struct {
	Quote     string `json:"quote"`
	Character string `json:"character"`
	Season    int    `json:"season"`
	Episode   int    `json:"episode"`
	Title     string `json:"title"`
}
