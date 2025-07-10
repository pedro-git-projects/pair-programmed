package main

import "fmt"

type Aluno struct {
	Nome  string
	Notas []float64
	Media Optional[float64]
}

func NewAluno(nome string, notas ...float64) *Aluno {
	nSlice := []float64{}
	nSlice = append(nSlice, notas...)

	a := &Aluno{
		Nome:  nome,
		Notas: nSlice,
	}

	if len(notas) > 0 {
		sum := 0.0
		for _, v := range notas {
			sum += v
		}
		mean := sum / float64(len(notas))
		a.Media = Optional[float64]{Value: mean, Set: true}
	}

	return a
}

func (a *Aluno) CalculateMean() {
	sum := 0.0
	l := len(a.Notas)
	for _, v := range a.Notas {
		sum += v
	}
	mean := sum / float64(l)
	a.Media = Optional[float64]{mean, true}
}

func (a Aluno) String() string {
	if a.Media.Set {
		return fmt.Sprintf("Nome: %s, Média: %.2f", a.Nome, a.Media.Value)
	}
	return fmt.Sprintf("Nome: %s, Média (não calculada)", a.Nome)
}
