package entity

type Video struct {
	Title       string `json:"title" binding:"min=2,max=32,required" validate:"StartsWithCapital"` // custom validator
	Description string `json:"description" binding:"max=128"`
	URL         string `json:"url" binding:"url,required"`
	Author      Person `json:"author" binding:"required"`
}

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"email,required"`
}
