package dto

type BookResponse struct {
	Id       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Genre    string `json:"genre,omitempty"`
	Price    int    `json:"price,omitempty"`
	SubTitle string `json:"sub_title,omitempty"`
}

type BookResponse2 struct {
}
