package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Функция для переворота слов в строке
func reverseWords(input string) string {
	// Разделяем строку на слова
	words := strings.Fields(input)
	// Переворачиваем слова в массиве
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	// Соединяем слова обратно в строку и возвращаем
	return strings.Join(words, " ")
}

func main() {
	fmt.Print("Введите строку: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	result := reverseWords(strings.TrimSpace(input))

	fmt.Println("Перевернутая строка:", result)
}
