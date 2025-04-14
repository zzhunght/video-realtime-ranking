package dto

type CreateVideo struct {
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	CategoryId int    `json:"category_id"`
}
