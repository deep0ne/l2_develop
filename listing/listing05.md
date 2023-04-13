Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```
Функция `test()` возвращает указатель на `customError`, которая в свою очередь реализует интерфейс ошибки. Выведем следующее:
```go
fmt.Printf("Underlying Type: %T\n", err)
fmt.Printf("Underlying Value: %v\n", err)
```
И увидим следующее:
```go
Underlying Type: *main.customError
Underlying Value: <nil>
```
Вместо интерфейса ошибки из функции возвращается указатель на customError. Поэтому мы и заходим в блок с условием `if err != nil`, потому что для того, чтобы `err` являлась `nil`, у неё должен быть `nil` underlaying type. 