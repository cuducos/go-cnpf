package cnpf

import "fmt"

func ExampleIsValid_unmaskedValidCpf() {
	fmt.Println(IsValid("23858488135"))
	// Output: true
}

func ExampleIsValid_maskedValidCpf() {
	fmt.Println(IsValid("238.584.881-35"))
	// Output: true
}

func ExampleIsValid_unmaskedValidCnpjWithNewFormat() {
	fmt.Println(IsValid("12ABC34501DE35"))
	// Output: true
}

func ExampleIsValid_maskedValidCnpj() {
	fmt.Println(IsValid("11.222.333/0001-81"))
	// Output: true
}

func ExampleIsValid_maskedValidCnpjWithNewFormat() {
	fmt.Println(IsValid("12.ABC.345/01DE-35"))
	// Output: true
}

func ExampleIsValid_maskedInvalidCpf() {
	fmt.Println(IsValid("111.111.111-11"))
	// Output: false
}

func ExampleIsValid_maskedInvalidCnpj() {
	fmt.Println(IsValid("11.111.111/1111-11"))
	// Output: false
}

func ExampleIsValid_maskedUnknownStringForCpf() {
	fmt.Println(IsValid("123.456.769/01"))
	// Output: false
}

func ExampleIsValid_maskedUnknownStringCnpj() {
	fmt.Println(IsValid("12.345.678 9012-34"))
	// Output: false
}

func ExampleIsValid_maskedLettersCpf() {
	fmt.Println(IsValid("ABC.DEF.GHI-JK"))
	// Output: false
}

func ExampleIsValid_maskedLettersCnpj() {
	fmt.Println(IsValid("AB.CDE.FGH/IJKL-MN"))
	// Output: false
}

func ExampleIsValid_IncompleteNumber() {
	fmt.Println(IsValid("123"))
	// Output: false
}

func ExampleUnmask_maskedCpf() {
	fmt.Println(Unmask("111.111.111-11"))
	// Output: 11111111111
}

func ExampleUnmask_maskedCnpj() {
	fmt.Println(Unmask("11.111.111/1111-11"))
	// Output: 11111111111111
}

func ExampleUnmask_maskedCnpjWithNewFormat() {
	fmt.Println(Unmask("12.ABC.345/01DE-35"))
	// Output: 12ABC34501DE35
}

func ExampleMask_unmaskedCpf() {
	fmt.Println(Mask("11111111111"))
	// Output: 111.111.111-11
}

func ExampleMask_unmaskedCnpj() {
	fmt.Println(Mask("11111111111111"))
	// Output: 11.111.111/1111-11
}

func ExampleMask_unmaskedCnpjWithNewFormat() {
	fmt.Println(Mask("12ABC34501DE35"))
	// Output: 12.ABC.345/01DE-35
}
