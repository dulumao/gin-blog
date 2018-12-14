package forms

type TagAddForm struct {
	Name      string `form:"name" json:"name" binding:"required"`
	CreatedBy string `form:"created_by" json:"created_by" binding:"required"`
	Status    byte   `form:"status" json:"status" binding:"required"`
}
