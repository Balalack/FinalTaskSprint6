package service

import (
	"fmt"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Conv(text string) (string, error) {

	// Переменная для подсчёта вхождений символов
	// азбуки Морзе в строку
	var count int
	// количество рун в строке
	var size int

	for i, simv := range text {
		if simv == '-' || simv == '.' || simv == ' ' {
			count += 1
		}
		size = i
	}

	if size == 0 {
		return "", fmt.Errorf("Строка для конвертации не должна быть пустой!")
	}
	// return morse.ToMorse(text), nil
	// Если количество рун в строке совпало с количеством символов
	// азбуки Морзе в строке, значит эта строка написана в Морзе
	if size+1 == count {
		return morse.ToText(text), nil
	} else {
		return morse.ToMorse(text), nil
	}

}
