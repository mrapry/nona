package usecase

//UsecaseUser interface for user users
type UsecaseUser interface {
	FindAll()
	Find()
	Insert()
	Update()
	Delete()
}
