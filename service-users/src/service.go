package src

type ServiceList interface {
	InsertNewUser(data []byte) (int, string, error)
	InsertUser(name, cpf_cnpj, email, password string, userType int) (int, error)
}
