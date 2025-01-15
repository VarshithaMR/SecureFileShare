package controllers

// holds all the controllers
type Controllers struct {
	UserController *UserController
	FileController *FileController
}

// NewControllers is a constructor function to create new instances of controllers
func NewControllers() *Controllers {
	return &Controllers{
		UserController: &UserController{},
		FileController: &FileController{},
	}
}
