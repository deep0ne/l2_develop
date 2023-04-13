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
}

type Query struct {
	query     string
	queryType string
}

func (q *Query) setQuery(query string) {
	q.query = query
}

func (q *Query) setQueryType(queryType string) {
	q.queryType = queryType
}

type GetSQLQuery struct {
	Query
}

type DeleteSQLQuery struct {
	Query
}

func newGET(id int, db string) SQLQuery {
	return &GetSQLQuery{
		Query: Query{
			query:     fmt.Sprintf("SELECT * FROM %s WHERE order_uid = %d", db, id),
			queryType: "GET",
		},
	}
}

func newDELETE(id int, db string) SQLQuery {
	return &DeleteSQLQuery{
		Query: Query{
			query:     fmt.Sprintf("DELETE FROM %s WHERE order_uid = %d", db, id),
			queryType: "DELETE",
		},
	}
}

func getQuery(queryType, db string, id int) (SQLQuery, error) {
	if queryType == "GET" {
		return newGET(id, db), nil
	}

	if queryType == "DELETE" {
		return newDELETE(id, db), nil
	}

	return nil, fmt.Errorf("Wrong type of query passed")
}

func main() {
	getOrder, _ := getQuery("GET", "orders", 12)
	deleteOrder, _ := getQuery("DELETE", "orders", 12)

	fmt.Println(getOrder)
	fmt.Println(deleteOrder)
}
