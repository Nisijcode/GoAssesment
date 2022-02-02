package models

type Pet struct {
	ID    uint    `json:"id"`
	Type  string  `json:"type"`
	Price float32 `json:"price"`
}
type SavePet struct {
	Type  string  `json:"type"`
	Price float32 `json:"price"`
}
