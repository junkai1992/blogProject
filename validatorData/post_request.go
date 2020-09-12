package validatorData

type CreatePostRequetst struct {
	CategoryId uint   `json:"category_id" binding:"required"`
	Title      string `json:"title" binding:"required,max=40"`
	Preface    string `json:"preface" bind:"required"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" binding:"required"`
}
