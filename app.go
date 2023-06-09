package main

import (
	"changeme/internal/logic"
	"changeme/internal/model/tables"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
)

// App struct
type App struct {
	Ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.Ctx = ctx
}

func (a *App) domReady(ctx context.Context) {

}

func (a *App) shutdown(ctx context.Context) {

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type Form struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Enabled bool   `json:"enabled"`
}

func (a *App) GetFiles() []tables.Qc {
	qcLogic := logic.NewQcLogic()
	qcs, err := qcLogic.GetQCList()
	if err != nil {
		log.Error(err)
		return nil
	}
	return qcs
}

func (a *App) AddFile(form Form) (error string) {
	if form.Name == "" {
		return "The file name cannot be empty!"
	}
	qcLogic := logic.NewQcLogic()
	err := qcLogic.CreateQC(tables.Qc{
		Name:     form.Name,
		Filepath: QcPath,
		Content:  form.Content,
		Enabled:  form.Enabled,
	})
	if err != nil {
		return err.Error()
	}
	return ""
}

func (a *App) EditFile(data tables.Qc) (error string) {
	qcLogic := logic.NewQcLogic()
	err := qcLogic.UpdateQC(data)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (a *App) RemoveFile(data tables.Qc) (error string) {
	qcLogic := logic.NewQcLogic()
	err := qcLogic.DeleteQC(data)
	if err != nil {
		return err.Error()
	}
	return ""
}
