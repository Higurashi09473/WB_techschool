package main

// Определяем структуру для кастомной ошибки
type customError struct {
	msg string
}

// Реализуем метод Error() — это делает *customError совместимым с интерфейсом error
func (e *customError) Error() string {
	return e.msg
}

// Функция, которая может возвращать ошибку (или nil)
func test() *customError {
	// ... здесь могла бы быть логика, которая генерирует ошибку
	return nil  // Возвращаем nil-указатель на customError
}

func main() {
	var err error  // Объявляем переменную типа интерфейс error

	// Присваиваем результат test() в интерфейс error
	err = test()

	// Что происходит под капотом:
	// 1. test() возвращает (*customError)(nil) — nil-указатель конкретного типа
	// 2. При присваивании в интерфейс error:
	//    • itab (тип + методы) = *customError (не nil!)
	//    • data (значение)     = nil
	//
	// Итог: интерфейс err содержит nil внутри, но сам интерфейс != nil,
	//       потому что имеет информацию о типе (*customError).

	// Проверка: err != nil → true! (хотя внутри nil)
	if err != nil {
		println("error")  // Это выполнится
		return
	}

	println("ok")  // Это НЕ выполнится
}