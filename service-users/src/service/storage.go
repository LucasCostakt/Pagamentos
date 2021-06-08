package service

import "service-users/src"

type Service struct {
	Repository src.RepositoryList
}

type User struct {
	Name      string `json:"name"`
	Cpf_cnpj  string `json:"cpf_cnpj"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	User_type int    `json:"user_type"`
}
