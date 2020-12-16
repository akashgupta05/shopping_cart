package items

type Item struct {
	Category  string `json:"category"`
	Inventory string `json:"inventory"`
	Rating    string `json:"rating"`
	Price     string `json:"price"`
	Origin    string `json:"origin"`
	Product   string `json:"product"`
}

type Items struct {
	Items []Item
}

func NewItems() *Items {
	return &Items{}
}
