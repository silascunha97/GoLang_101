package main

import "fmt"

func Compare_Valores() {
	//Fácil 1- Faça um programa que peça dois números e verifique (usando if e else) e imprima o maior deles
	a, b := 10, 10
	//a, b = b, a
	if a > b {
		fmt.Println(" A é maior do que B")
	} else if a < b {
		fmt.Println(" B é maior do que A")
	} else {
		fmt.Println("Os valores são iguais: ")
	}
}

func ValorPositivo() {
	//Fácil 2- Faça um programa que peça um valor e mostre na tela se o valor é positivo ou negativo
	var valor_is_positive int = 10
	if valor_is_positive > 0 {
		fmt.Println("O valor é positivo : ")
	} else if valor_is_positive < 0 {
		fmt.Println("O valor é negativo : ")
	} else {
		fmt.Println("O valor é zero não é nem positivo nem negativo: ")
	}
}

func genderUser() {
	//Fácil 3- Faça um programa que verifique (usando if e else) se uma letra digitada é “F” ou “M”. Conforme a letra escrever: F – Feminino, M- Masculino, Sexo inválido.
	fmt.Println("Digite o seu sexo F para feminino M para masculino: ")
	var genderuser string
	fmt.Scanln(&genderuser)
	if genderuser == "f" || genderuser == "F" {
		fmt.Println("Você é do sexo feminino: ")
	} else if genderuser == "m" || genderuser == "M" {
		fmt.Println("Você é do sexo masculino: ")
	}
}
func is_vogal() {
	//Fácil 4- Faça um programa que verifique (usando if e else) se uma letra digitada é vogal ou consoante.
	var char_is_vogal string
	fmt.Println("Digite um caracterer para comparação: ")
	fmt.Scanln(&char_is_vogal)
	if char_is_vogal == "a" || char_is_vogal == "e" || char_is_vogal == "o" || char_is_vogal == "u" {
		println("É uma vogal")
	} else if char_is_vogal == "A" || char_is_vogal == "E" || char_is_vogal == "O" || char_is_vogal == "U" {
		println("É uma vogal")
	} else {
		println("É uma consoante")
	}
}

func notas_Aluno_intervalo() {
	//Fácil 5- Faça um programa para a leitura de duas notas parciais de um aluno.
	var nota1, nota2 int
	fmt.Println("Digite a sua primeira nota")
	fmt.Scanln(&nota1)
	fmt.Println("Digite a sua segunda nota")
	fmt.Scanln(&nota2)
	if nota1 >= 0 && nota1 <= 10 && nota2 >= 0 && nota2 <= 10 {
		fmt.Println("O aluno tirou as notas", nota1, nota2)
	} else {
		fmt.Println("A nota é invalida!!")
	}
}

func notas_aluno_is_aprovado() {
	//Fácil 5- Faça um programa para a leitura de duas notas parciais de um aluno.
	//A mensagem “Aprovado”, se a média alcançada for maior ou igual a sete;
	//A mensagem “Aprovado com Distinção”, se a média for igual a dez;
	//A mensagem “Reprovado” se a média for menor de do que sete;
	var nota1, nota2, result int
	fmt.Println("Digite a sua primeira nota")
	fmt.Scanln(&nota1)
	fmt.Println("Digite a sua segunda nota")
	fmt.Scanln(&nota2)
	result = (nota1 + nota2) / 2
	if result > 7 && result < 10 {
		fmt.Println()
	} else if result == 10 {
		fmt.Println("Aprovado com honras ao merito")
	} else {
		fmt.Println("Aluno reprovado")
	}
}

func compare_o_maior_valor() {
	//Intermediário 6- Faça um programa que leia três números, verifique (usando if e else), e mostre o maior deles.
	var a, b, c int
	fmt.Println("Digite a sua primeira valor")
	fmt.Scanln(&a)
	fmt.Println("Digite a sua segunda valor")
	fmt.Scanln(&b)
	fmt.Println("Digite a sua terceiro valor")
	fmt.Scanln(&c)
	if a > b && a > c {
		fmt.Println("O primeiro valor é maior")
	} else if b > a && b > c {
		fmt.Println("O segundo valor é maior")
	} else if c > b && c > a {
		fmt.Println("O terceiro valor é maior")
	} else {
		fmt.Println("Os valores são iguais")
	}
}

func compare_o_maior_valor_e_o_menor() {
	//Intermediário 7- Faça um programa que leia três números, verifique (usando if e else) e mostre o maior e o menor deles;
	var a, b, c int
	fmt.Println("Digite a sua primeira valor")
	fmt.Scanln(&a)
	fmt.Println("Digite a sua segunda valor")
	fmt.Scanln(&b)
	fmt.Println("Digite a sua terceiro valor")
	fmt.Scanln(&c)
	if a > b && a > c {
		if b < c {
			fmt.Println("O primeiro valor é o maior e segundo é o menor")
		} else {
			fmt.Println("O primeiro valor é o maior e terceiro é o menor")
		}
	} else if b > a && b > c {
		if a < c {
			fmt.Println("O segundo valor é o maior e primeiro é o menor")
		} else {
			fmt.Println("O segundo valor é o maior e terceiro é o menor")
		}
	} else if c > a && c > b {
		if a < b {
			fmt.Println("O terceiro valor é o maior e primeiro é o menor")
		} else {
			fmt.Println("O terceiro valor é o maior e segundo é o menor")
		}
	} else {
		fmt.Println("Todos os valores são iguais")
	}
}

func main() {

}
