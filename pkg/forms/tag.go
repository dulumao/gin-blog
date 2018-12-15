package forms

// 添加tag表单验证
type AddTagForm struct {
	Name      string `form:"name" json:"name" validate:"required"`
	CreatedBy string `form:"created_by" json:"created_by" validate:"required"`
	Status    int    `form:"status" json:"status" validate:"number"`
}

// 修改tag表单验证
type EditTagForm struct {
	Name       string `json:"name" form:"name" validate:"max=10"`
	ModifiedBy string `json:"modified_by" form:"modified_by" validate:""`
	Status     uint   `json:"status" form:"status" validate:"gte=0,lte=1,required"`
}
