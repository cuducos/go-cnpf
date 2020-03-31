package cnpf

import (
    cpf "github.com/cuducos/go-cpf"
    cnpj "github.com/cuducos/go-cnpj"
)

//Unmask removes any non-digit (numeric) from a CPF or CNPJ number
func Unmask(n string) string {
	return cpf.Unmask(n)
}

//Mask returns the CPF or CNPJ number formatted
func Mask(n string) string {
	u := Unmask(n)
	if len(u) == 11 {
		return cpf.Mask(u)
	}
	if len(u) == 14 {
		return cnpj.Mask(u)
	}
	return n
}

//IsValid checks whether a number is a valid CPF or CNPJ number
func IsValid(n string) bool {
	return cpf.IsValid(n) || cnpj.IsValid(n)
}
