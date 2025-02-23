package main

import "fmt"

type Document interface {
	Doc() string
}
type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
}

type Aluno struct {
	Aluno     *Pessoa
	Matricula int64
	Curso     string
}

func (p Pessoa) Doc() string {
	return fmt.Sprintf("Nome: %s %s\nIdade: %d\nStatus: %t\n", p.Nome, p.Sobrenome, p.Idade, p.Status)
}

func Show(d Document) any {
	switch d.(type) {
	case Pessoa:
		fmt.Println(d.(Pessoa).Doc())
	case Aluno:
		fmt.Println(d.(Aluno).Doc())
	}
	fmt.Println(d)
	fmt.Println(d.Doc())
	return nil
}

func (a Aluno) Doc() string {
	Nome := a.Aluno.Nome
	Sobrenome := a.Aluno.Sobrenome
	Idade := a.Aluno.Idade
	Status := a.Aluno.Status
	Matricula := a.Matricula
	Curso := a.Curso
	return fmt.Sprintf("Nome: %s %s\nIdade: %d\nStatus: %t\nMatricula: %s\nCurso: %s\n", Nome, Sobrenome, Idade, Status, Matricula, Curso)
}

func main() {
	a := Aluno{
		&Pessoa{Nome: "Silas", Sobrenome: "Silva", Idade: 20, Status: true},
		123456789,
		"Engenharia de Software",
	}
	fmt.Println(Show(a))
	Show(a)
}
