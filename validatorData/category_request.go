package validatorData

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
