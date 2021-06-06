package repository

type Storage interface {
	CreateUserTypeTable() error
	CreateUserTable() error
	CreateTransactionTable() error
	CreateReversalTable() error
	executeInsert(query string) error
	CreateUsersTypes() error
	CreateUsers() error
}
