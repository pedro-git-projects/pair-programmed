package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type App struct {
	Alunos []Aluno
	rw     *bufio.ReadWriter
}

func NewApp() *App {
	return &App{
		rw: bufio.NewReadWriter(bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)),
	}
}

func (app *App) readName() (string, error) {
	_, err := fmt.Fprintf(app.rw, "Entre o nome do aluno (ou -1 para sair):\n-> ")
	if err != nil {
		log.Fatal(WRITE_ERR)
	}
	app.rw.Writer.Flush()
	nameLine, err := app.rw.Reader.ReadString('\n')
	if err != nil {
		log.Fatal(READ_ERR)
	}
	return strings.TrimSpace(nameLine), nil
}

func (app *App) readGrades() ([]float64, error) {
	var grade float64
	grades := []float64{}
	for {

		_, err := fmt.Fprintf(app.rw, "Insira uma nota ou -1 para sair:\n-> ")
		if err != nil {
			log.Fatal(WRITE_ERR)
		}
		app.rw.Writer.Flush()

		_, err = fmt.Fscan(app.rw.Reader, &grade)
		if err != nil {
			return nil, errors.New(SCAN_ERR)
		}
		_, _ = app.rw.Reader.ReadString('\n')
		if grade == -1 {
			break
		}
		grades = append(grades, grade)
	}
	return grades, nil
}

func (app *App) appendStudent() (bool, error) {
	n, err := app.readName()
	if err != nil {
		return false, errors.New(READ_ERR)
	}
	if n == "-1" {
		return false, nil
	}
	g, err := app.readGrades()
	if err != nil {
		return false, err
	}
	a := NewAluno(n, g...)
	app.Alunos = append(app.Alunos, *a)
	return true, nil
}

func (app App) printResult() {
	for _, v := range app.Alunos {
		fmt.Printf("--- Resultados ---\n")
		fmt.Println(v)
	}
}

func (app *App) MainLoop() {
	for {
		cont, err := app.appendStudent()
		if err != nil {
			log.Fatal(err)
		}
		if !cont {
			break
		}
	}
	app.printResult()
}
