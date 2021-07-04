package model

type (
	CreateProductRequest struct {
		Name     string
		Price    float64
		Quantity int
	}

	UpdateProductRequest struct {
		Name     string
		Price    float64
		Quantity int
	}
)
