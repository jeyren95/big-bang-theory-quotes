package models

type Quote struct {
	Id        int    `json:"id"`
	Quote     string `json:"quote"`
	Character string `json:"character"`
	Season    int    `json:"season"`
	Episode   int    `json:"episode"`
	Title     string `json:"title"`
}
