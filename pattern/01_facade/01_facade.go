package main

import "fmt"

/*
при проектировании сложных систем часто применяется принцип декомпозиции, при котором
сложная система разбивается на более мелкие и простые подсистемы. Однако после этого такие подсистемы
довольно сложно использовать. Эту проблему решает паттерн фасад, которые позволяет сделать простой, единый интерфейс
для взаимодействия с подсистемами.
в качестве реального примера использования можно привести упрощенные интерфейсы для работы с БД
*/

type Docker struct{}

func (d Docker) CreateContainer(user, password string) {
	fmt.Printf("Creating postgres container with user '%s' and password '%s'\n", user, password)
}

func (d Docker) CreateDB() {
	fmt.Println("Creating DB...")
}

type GoMigrate struct{}

func (m GoMigrate) MigrateUp() {
	fmt.Println("Migration of data...")
}

type SQLC struct{}

func (s SQLC) CRUDGenerate() {
	fmt.Println("Generating CRUDS on Golang...")
}

type DBFacade struct {
	docker    *Docker
	migration *GoMigrate
	plugin    *SQLC
}

func NewDBFacade() *DBFacade {
	return &DBFacade{
		docker:    &Docker{},
		migration: &GoMigrate{},
		plugin:    &SQLC{},
	}
}

func (facade *DBFacade) DBSetup() {
	facade.docker.CreateContainer("root", "secret")
	facade.docker.CreateDB()
	facade.migration.MigrateUp()
	facade.plugin.CRUDGenerate()
	fmt.Println("Database and CRUDS are set and the system is ready to use!")
}

func main() {
	DB := NewDBFacade()
	DB.DBSetup()
}
