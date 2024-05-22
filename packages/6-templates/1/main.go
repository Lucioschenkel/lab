package main

import (
	"os"
	"text/template"
)

type Curso struct {
  Nome string
  CargaHoraria int
}

func main()  {
  curso := Curso{"Go", 40}  
  tmpl := template.New("CursoTemplate")
  tmpl, _ = tmpl.Parse("Curso: {{.Nome}} - Carga horaria: {{.CargaHoraria}}")
  err := tmpl.Execute(os.Stdout, curso)
  if err != nil {
    panic(err)
  }
}
