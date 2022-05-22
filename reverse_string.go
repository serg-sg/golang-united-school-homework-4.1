package reverse_string

func ReverseString(input string) (output string) {

	a := []rune(input) // Сделаем копию input с преобразованием в rune

	b := make([]rune, len(a)) // Создадим пустую rune для output

	x := len(a) - 1 // Инициализируем переменную
	//для обратного отсчёта индексов rune для возврата результата

	// Перевернём rune
	for _, value := range a {

		b[x] = value
		x--
	}

	return string(b) // Преобразуем rune в string
}
