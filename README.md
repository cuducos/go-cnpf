# Go CNPF [![Tests](https://github.com/cuducos/go-cnpf/workflows/Tests/badge.svg)]() [![GoDoc](https://godoc.org/github.com/cuducos/go-cnpf?status.svg)](https://godoc.org/github.com/cuducos/go-cnpf) ![Go version](https://img.shields.io/github/go-mod/go-version/cuducos/go-cnpj)

A Go module to validate CPF and CNPJ numbers (Brazilian people and companies unique identifier for the Federal Revenue).

> The pseudo-acronym _CNPF_ is a sort of tong-twister and a common typo when developers discuss the implementation of objects that could hold either a CPF or a CNPJ numbers.

```go
package main

import "github.com/cuducos/go-cnpf"


func main() {
	// these return true
	cnpf.IsValid("23858488135")
	cnpf.IsValid("238.584.881-35")
	cnpf.IsValid("11222333000181")
	cnpf.IsValid("11.222.333/0001-81")
	cnpj.IsValid("12.ABC.345/01DE-35")
	cnpj.IsValid("12ABC34501DE35")

	// these return false
	cnpf.IsValid("111.111.111-11")
	cnpf.IsValid("11.111.111/1111-11")
	cnpf.IsValid("123.456.769/01")
	cnpf.IsValid("12.345.678 9012-34")
	cnpf.IsValid("ABC.DEF.GHI-JK")
	cnpf.IsValid("AB.CDE.FGH/IJKL-MN")
	cnpf.IsValid("123")

	// these return 11111111111, 11111111111111 and 12ABC34501DE35
	cnpf.Unmask("111.111.111-11")
	cnpf.Unmask("11.111.111/1111-11")
	cnpj.Unmask("12.ABC.345/01DE-35")

	// this return 111.111.111-11, 11.111.111/1111-11 and 12.ABC.345/01DE-35
	cnpf.Mask("11111111111")
	cnpf.Mask("11111111111111")	
	cnpj.Mask("12ABC34501DE35")
}
```
> [!IMPORTANT]
> Starting in July 2026 [the CNPJ number will be alphanumeric](https://www.gov.br/receitafederal/pt-br/acesso-a-informacao/acoes-e-programas/programas-e-atividades/cnpj-alfanumerico). This package **already supports the new format**. If you **do not** want to support the new format, tag this package to [`v0.0.1`](https://github.com/cuducos/go-cnpf/releases/tag/v0.0.1).

Based on [Go CPF](https://github.com/cuducos/go-cpf) and [Go CNPJ](https://github.com/cuducos/go-cnpj) ❤️
