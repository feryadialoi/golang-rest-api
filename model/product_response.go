package model

type (
	GetProductResponse struct {
		Name     string  `json:"name,omitempty"`
		Price    float64 `json:"price,omitempty"`
		Quantity int     `json:"quantity,omitempty"`
	}

	CreateProductResponse struct {
		Name     string  `json:"name,omitempty"`
		Price    float64 `json:"price,omitempty"`
		Quantity int     `json:"quantity,omitempty"`
	}

	UpdateProductResponse struct {
		Name     string  `json:"name,omitempty"`
		Price    float64 `json:"price,omitempty"`
		Quantity int     `json:"quantity,omitempty"`
	}
)
