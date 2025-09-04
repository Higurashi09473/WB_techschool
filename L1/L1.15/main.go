package main

func createHugeString(i int) []byte {
	return make([]byte, i)
}

func someFunc() string{
	// При создании большой строки justString она имеет указатель на тот же массив что и v, это не дает сборщику мусора удалить v (утечка памяти)
	v := createHugeString(1 << 10)

	//Мы возвращаем данные и после GC может удалить v
	return string(v[:100])
}


func main() {
  someFunc()
}