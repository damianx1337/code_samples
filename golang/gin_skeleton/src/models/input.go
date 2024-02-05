package inputs

type TestInput struct {
	Name	string	`form:"name" binding:"required,max=15"`
}
