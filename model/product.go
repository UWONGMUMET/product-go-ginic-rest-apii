package model

type Product struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Price    uint32 `json:"price"`
	Quantity uint   `json:"quantity"`
}

type PostProduct struct {
	Name     string `json:"name"`
	Price    uint32 `json:"price"`
	Quantity uint   `json:"quantity"`
}

type ProductUri struct {
	ID uint `uri:"id" binding:"required,number"`
}
