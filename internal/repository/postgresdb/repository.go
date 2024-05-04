package postgresdb

type Authorization interface {
}

type Account interface {
}

type Product interface {
}

type Reservation interface {
}

type Operation interface {
}

type Repository struct {
	Authorization
	Account
	Product
	Reservation
	Operation
}

func NewRepository() *Repository {
	return &Repository{}
}
