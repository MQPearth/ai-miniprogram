package vo

type NamingVO struct {
	// 性别
	Gender string `form:"gender" binding:"required"`
	// 姓氏
	Surname string `form:"surname" binding:"required"`
	// 风格
	Style string `form:"style" binding:"required"`
	// 名字长度
	Length uint `form:"length" binding:"required,max=4,min=2"`
	// 结果集数量
	Qty uint `form:"qty" binding:"required,max=10,min=1"`
}
