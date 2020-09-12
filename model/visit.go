package model

type Visit struct {
	ID         uint `json:"id" gorm:"priamry_key"`
	Visitcount int  `json:"visitcount" gorm:"not null;unique"`
	CreatedAt  Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt  Time `json:"updated_at" gorm:"type:timestamp"`
}
