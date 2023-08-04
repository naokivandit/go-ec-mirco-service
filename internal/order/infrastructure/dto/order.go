package dto

type Order struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Orders []Order

type OrderItem struct {
	ID       int    `json:"id"`
	OrderID  int    `json:"order_id"`
	ImageURL string `json:"image_url"`
}
