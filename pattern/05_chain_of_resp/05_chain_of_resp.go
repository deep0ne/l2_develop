/*
Паттерн Chain Of Responsibility относится к поведенческим паттернам
Паттерн позволяет избежать привязки объекта-отправителя запроса к объекту-получателю запроса, при этом
Давай шанс обработать этот запрос нескольким объектам. Получатели связаваются в цепочку, и запрос передаётся по цепочке
Я решил реализовать паттер цепочки обязанностей на примере последовательной обработки данных для машинного обучения
*/

package main

import "fmt"

type Data struct {
	dataset                string
	convertingDone         bool
	anomaliesDetectionDone bool
	oneHotEncodingDone     bool
	statAnalysisDone       bool
}

type Handler interface {
	DataProcessing(*Data)
	setNext(Handler)
}

type TypeConverter struct {
	next Handler
}

func (tc *TypeConverter) DataProcessing(d *Data) {
	if d.convertingDone {
		fmt.Println("Convertation of types is already done")
		tc.next.DataProcessing(d)
		return
	}
	fmt.Println("Converting wrong types of data...")
	d.convertingDone = true
	tc.next.DataProcessing(d)
}

func (tc *TypeConverter) setNext(next Handler) {
	tc.next = next
}

type AnomaliesDetector struct {
	next Handler
}

func (ad *AnomaliesDetector) DataProcessing(d *Data) {
	if d.anomaliesDetectionDone {
		fmt.Println("Anomalies are already detected and removed")
		ad.next.DataProcessing(d)
		return
	}
	fmt.Println("Detecting and removing anomalies...")
	d.anomaliesDetectionDone = true
	ad.next.DataProcessing(d)
}

func (ad *AnomaliesDetector) setNext(next Handler) {
	ad.next = next
}

type OneHotEncoder struct {
	next Handler
}

func (ohe *OneHotEncoder) DataProcessing(d *Data) {
	if d.oneHotEncodingDone {
		fmt.Println("OHE is already done")
		ohe.next.DataProcessing(d)
		return
	}
	fmt.Println("Applying OHE...")
	d.oneHotEncodingDone = true
	ohe.next.DataProcessing(d)
}

func (ohe *OneHotEncoder) setNext(next Handler) {
	ohe.next = next
}

type StatAnalyzer struct {
	next Handler
}

func (stats *StatAnalyzer) DataProcessing(d *Data) {
	if d.statAnalysisDone {
		fmt.Println("Stat analysis is already done")
		stats.next.DataProcessing(d)
		return
	}
	fmt.Println("Doing stat analysis...")
	d.statAnalysisDone = true
	// stats.next.DataProcessing(d)
}

func (stats *StatAnalyzer) setNext(next Handler) {
	stats.next = next
}

func main() {
	data := &Data{
		dataset: "titanic",
	}

	stats := &StatAnalyzer{}

	oneHotEncoder := &OneHotEncoder{}
	oneHotEncoder.setNext(stats)

	anomalyDetector := &AnomaliesDetector{}
	anomalyDetector.setNext(oneHotEncoder)

	converter := &TypeConverter{}
	converter.setNext(anomalyDetector)

	converter.DataProcessing(data)

}
