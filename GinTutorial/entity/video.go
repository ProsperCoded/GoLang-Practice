package entity

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       int    `json:"age" binding:"gte=0,lte=130"`
}
type Video struct {
	Title       string `json:"title" xml:"title" validate:"required"`
	URL         string `json:"url" binding:"required"`
	Description string `json:"description"`
	Author      Person `json:"author" binding:"required"`
}