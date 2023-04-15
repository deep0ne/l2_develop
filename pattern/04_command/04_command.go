/*
Паттерн команда относится к поведенческим паттернам
Он позволяет представить запрос в виде объекта. Из этого следует, что команда - это объект.
Такие запросы можно ставить в очередь, отменять или возобновлять.

Реализуем этот паттерн на примере взаимодействия с бд с операциями из L0
*/

package main

import (
	"database/sql"
	"log"
)

type Command interface {
	Execute() error
}

type AddCommand struct {
	db          *sql.DB
	order_uid   string
	order_price float64
}

func (c *AddCommand) Execute() error {
	_, err := c.db.Exec("INSERT INTO orders (order_uid, order_price) VALUES (?, ?)", c.order_uid, c.order_price)
	if err != nil {
		return err
	}
	return nil
}

type DeleteCommand struct {
	db        *sql.DB
	order_uid string
}

func (c *DeleteCommand) Execute() error {
	_, err := c.db.Exec("DELETE FROM orders WHERE order_uid = ?", c.order_uid)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := sql.Open("postgres", "user:password@tcp(1.1.1.1)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	addCommand := AddCommand{
		db:          db,
		order_uid:   "1234",
		order_price: 5000.12,
	}
	deleteCommand := DeleteCommand{
		db:        db,
		order_uid: "1234",
	}

	err = addCommand.Execute()
	if err != nil {
		log.Fatal(err)
	}

	err = deleteCommand.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
