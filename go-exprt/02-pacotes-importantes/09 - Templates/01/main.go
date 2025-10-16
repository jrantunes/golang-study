package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 60}
	tmp := template.New("CursoTemplate")
	tmp, err := tmp.Parse("Curso: {{.Nome}}, Carga horária: {{.CargaHoraria}}")
	if err != nil {
		panic(err)
	}
	err = tmp.Execute(os.Stdout, curso) // manda o resultado para Stdout
	if err != nil {
		panic(err)
	}
}
