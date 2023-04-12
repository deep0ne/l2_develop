/*
Паттерн "Посетитель" позволяет добавлять новые операции к объектами, не изменяя эти объекты.
Создаётся класс посетителя, который содержит методы для каждого типа объекта, которые он может посетить.
Он позволяет избежать изменений классов, но усложняет код программы.

Визитер может использоваться для сериализации объектов в различные форматы.
*/

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Visitor interface {
	VisitDocuments(*Document) ([]byte, error)
	VisitInfo(*Info) ([]byte, error)
}

type Object interface {
	Accept(Visitor)
}

type Document struct {
	Number int
	Title  string
}

func (d *Document) Accept(v Visitor) ([]byte, error) {
	return v.VisitDocuments(d)
}

type Info struct {
	Author string
}

func (i *Info) Accept(v Visitor) ([]byte, error) {
	return v.VisitInfo(i)
}

type SerializerVisitor struct {
}

func (sv *SerializerVisitor) VisitDocuments(d *Document) ([]byte, error) {
	bytes, err := json.Marshal(d)
	return bytes, err
}

func (sv *SerializerVisitor) VisitInfo(i *Info) ([]byte, error) {
	bytes, err := xml.Marshal(i)
	return bytes, err
}

func main() {
	document := Document{
		Number: 15,
		Title:  "First Scientific Work",
	}
	info := Info{
		Author: "Kozhukhov",
	}

	serializer := &SerializerVisitor{}
	doc, _ := document.Accept(serializer)
	inf, _ := info.Accept(serializer)
	fmt.Println(string(doc))
	fmt.Println(string(inf))
}
