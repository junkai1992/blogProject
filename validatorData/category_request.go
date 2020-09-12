package validatorData

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateVisitRequest struct {
	Visitcount int `json:"visitcount" binding:"required"`
}
