/*
Паттерн "Фабричный метод" относится к классу порождающих паттернов проектирования и представляет собой
Способ создания объектов без явного указания их класса. Вместо этого используется метод-фабрика, который
Создаёт объекты определенного типа в зависимости от переданных параметров. Он упрощает создание объектов и
Уменьшает связность классов, но усложняет код за счёт добавления дополнительных классов и методов.

Хотелось попробовать реализовать не "игрушечный" пример для создания различных запросов к БД.
*/

package main

import "fmt"

type SQLQuery interface {
	setQuery(query string)
	setQueryType(queryType string)
	setDB(db string)
}

type Query struct {
	query     string
	queryType string
	db        string
}

func (q *Query) setQuery(query string) {
	q.query = query
}

func (q *Query) setQueryType(queryType string) {
	q.queryType = queryType
}

func (q *Query) setDB(db string) {
	q.db = db
}

type GetSQLQuery struct {
	Query
}

type DeleteSQLQuery struct {
	Query
}

func newGET(id int) SQLQuery {
	return &GetSQLQuery{
		Query: Query{
			query:     fmt.Sprintf("SELECT * FROM orders WHERE order_uid = %d", id),
			queryType: "GET",
			db:        "orders",
		},
	}
}

func newDELETE(id int) SQLQuery {
	return &DeleteSQLQuery{
		Query: Query{
			query:     fmt.Sprintf("DELETE FROM orders WHERE order_uid = %d", id),
			queryType: "DELETE",
			db:        "orders",
		},
	}
}

func getQuery(queryType string, id int) (SQLQuery, error) {
	if queryType == "GET" {
		return newGET(id), nil
	}

	if queryType == "DELETE" {
		return newDELETE(id), nil
	}

	return nil, fmt.Errorf("Wrong type of query passed")
}

func main() {
	getOrder, _ := getQuery("GET", 12)
	deleteOrder, _ := getQuery("DELETE", 12)

	fmt.Println(getOrder)
	fmt.Println(deleteOrder)
}
