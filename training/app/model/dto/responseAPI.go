package dto

type ResponseAPI struct {
	Name  string `json:"name,omitempty"`
	Bio   string `json:"bio,omitempty"`
	Title string `json:"title,omitempty"`
}
