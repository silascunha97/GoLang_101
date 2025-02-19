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

func Preco_Produtos() {
	/*
		Fácil 8: Faça um programa que pergunte o preço de três produtos e informe qual produto você
			deve comprar, sabendo que a decisão é sempre o mais barato.
	*/
	var prod1, prod2, prod3, produto_mais_barato int
	fmt.Println("Digite o valor do primeiro produto")
	fmt.Scanln(&prod1)
	fmt.Println("Digite o valor do segundo produto")
	fmt.Scanln(&prod2)
	fmt.Println("Digite o valor do terceira produto")
	fmt.Scanln(&prod3)
	if prod1 < prod2 && prod2 < prod3 {
		produto_mais_barato = prod1
	} else {
		if prod2 < prod1 && prod2 < prod3 {
			produto_mais_barato = prod2
		} else {
			produto_mais_barato = prod3
		}
	}
	fmt.Println("O pruto mais para é o de valor: ", produto_mais_barato)
}

func saudacao_aluno() {
	/*
		Fácil 10- Faça um programa que pergunte em que turno você estuda.
		 Peça para digitar M-matutino ou V-vespertino ou N-noturno.
		 Imprima a mensagem “Bom dia!” ou  “Boa Noite” ou “Valor inválido”, conforme o caso.
	*/
	var turno string
	fmt.Println("Digite o seu turno M para matutino V para vespentino e N para noturno")
	fmt.Scanln(&turno)
	if turno == "m" || turno == "M" {
		fmt.Println("Seu turno é matutino")
	} else if turno == "v" || turno == "V" {
		fmt.Println("Seu turno é vespentino")
	} else if turno == "n" || turno == "N" {
		fmt.Println("Seu turno é noturno")
	} else {
		fmt.Println("Insira um character valido!")
	}
}

func salario() {
	/*

		Difícil 11- As organizações CSM resolveram dar um aumento de salário aos seus colaboradores e lhe contrataram para desenvolver o programa que calculará os reajustes.

			a. Faça um programa que recebe o salário de um colaborador e o reajuste segundo o seguinte critério, baseado no salário atual;
			b. Salários até R$ 280,00 (incluindo): aumento de 20%;
			c. Salários entre R$ 280,00 e R$700,00: aumento de 15%;
			d. Salários entre R$ 700,00 e R$ 1500,00: aumento de 10%;
			e. Salários de R$ 1500,00 em diante: aumento de 5%
			Após o aumento ser realizado; informe na tela;

			a. O salário antes do reajuste;
			b. O percentual de aumento aplicado;
			c. O valor do aumento;
			d. O novo salário, após o aumento.
	*/
	var salario_atual, novo_salario, valor_aumento float64
	var percentual_aumento int

	fmt.Println("Digite o salario para calcular o reajuste: ")
	fmt.Scanln(&salario_atual)
	if salario_atual <= 280.00 {
		percentual_aumento = 20
	} else if salario_atual <= 700.00 {
		percentual_aumento = 15
	} else {
		if salario_atual <= 1500.00 {
			percentual_aumento = 10
		} else {
			percentual_aumento = 5
		}
	}
	valor_aumento = salario_atual * (float64(percentual_aumento) / 100.0)
	novo_salario = salario_atual + valor_aumento

	// Exibe o salário antes do reajuste
	fmt.Println("Salário antes do reajuste: R$ ", salario_atual)
	// Exibe o percentual de aumento aplicado
	fmt.Println("Percentual de aumento aplicado: ", percentual_aumento)
	// Exibe o valor do aumento
	fmt.Println("Valor do aumento: R$ ", valor_aumento)
	// Exibe o novo salário após o aumento
	fmt.Println("Novo salário após o aumento: R$ ", novo_salario)

}

func main() {
	salario()
}
