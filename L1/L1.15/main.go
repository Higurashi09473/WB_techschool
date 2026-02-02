package main


var justString string

func createHugeString(i int) string {
	return string(make([]byte, i))
}

func someFunc(){
	// При создании большой строки justString она имеет указатель на тот же массив что и v, это не дает сборщику мусора удалить v (утечка памяти)
	v := createHugeString(1 << 10)

	//Мы возвращаем данные и после GC может удалить v
	justString = string([]byte(v[:100]))
}


func main() {
  someFunc()
}