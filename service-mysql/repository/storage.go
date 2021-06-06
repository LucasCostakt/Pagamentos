package repository

type Storage interface {
	CreateUserTypeTable() error
	CreateUserTable() error
	CreateTransationTable() error
	CreateReversalTable() error
	executeInsert(query string) error
	CreateUsersTypes() error
	CreateUsers() error
}
