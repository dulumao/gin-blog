package forms

// 添加tag表单验证
type AddTagForm struct {
	Name      string `form:"name" json:"name" validate:"required"`
	CreatedBy string `form:"created_by" json:"created_by" validate:"required"`
	Status    int    `form:"status" json:"status" validate:"number"`
}

// 修改tag表单验证
type EditTagForm struct {
	Name       string `json:"name" form:"name" binding:"max=10"`
	ModifiedBy string `json:"modified_by" form:"modified_by"`
	Status     uint   `json:"status" form:"status" binding:"gte=0,lte=1,required"`
}

// 用户登录表单
type AuthForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterUserForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required,gte=4"`
	Name     string `form:"name" binding:"required"`
	Phone    string `form:"phone"`
}
