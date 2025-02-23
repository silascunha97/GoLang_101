package main

import (
	"fmt"
	//"net/http"
)

//	func helloHandler(w http.ResponseWriter, r *http.Request) {
//	   fmt.Fprintf(w, "Hello, World")
//	}
var (
	Pessoa1 = Pessoa{"Arthur", "Felipe", 20, true}
	Pessoa2 = Pessoa{"Amanda", "Silva", 20, true}
	Pessoa3 = Pessoa{"Felipe", "Silva", 20, true}
)

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
}

type Aluno struct {
	Aluno     *Pessoa
	Matricula string
	Curso     string
}

type Document interface {
	Doc() string
}

func (p Pessoa) GetNome() string {
	return fmt.Sprintf("Olá, meu nome é %s %d", p.Nome, p.Idade)
}

type categoria struct {
	nome string
	pai  *categoria
}

func (c categoria) HasParent() bool {
	return c.pai != nil

}

func main() {
	// http.HandleFunc("/hello", helloHandler)
	// http.ListenAndServe(":8080", nil)

	//idades := make(map[string]uint8)

	//cat := categoria{nome: "categoria 1"}
	//if !cat.HasParent() {
	//	fmt.Println("Não tem pai")
	//} else {
	//	fmt.Println("Tem pai")
	//}

	//  p := Pessoa{
	//	Nome:      "Pericles \n",
	//	Sobrenome: "Silva \n",
	//	Idade:     20,
	//	Status:    true,
	//	}

	//p := Pessoa{"jurandir", "Silva", 20, true}

	//fmt.Println(p, p1)

}
