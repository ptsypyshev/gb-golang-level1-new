package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/ptsypyshev/gb-golang-level1-new/hw03/urls/internal/models"
)

func main() {
	defer func() {
		// Завершаем работу с клавиатурой при выходе из функции
		_ = keyboard.Close()
	}()

	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для выхода и приложения нажмите Esc")

	storage := make(map[string]models.URL)

OuterLoop:
	for {
		// Подключаем отслеживание нажатия клавиш
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Введите 'a', 'l' или 'r' для выбора команды")

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			// Добавление нового url в список хранения
			fmt.Println("Введите новую запись в формате <url описание теги>")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(strings.TrimSpace(text))
			if len(args) < 3 {
				fmt.Println("Введите правильный аргументы в формате url описание теги")
				continue OuterLoop
			}

			storage[args[0]] = models.URL{
				Date: time.Now(),
				Link: args[0],
				Name: args[1],
				Tags: args[2:],
			}
		case 'l':
			// Вывод списка добавленных url. Выведите количество добавленных url и список с данными url
			// Вывод в формате
			// Имя: <Описание>
			// URL: <url>
			// Теги: <Теги>
			// Дата: <дата>

			// Напишите свой код здесь
			for k := range storage {
				fmt.Printf(
					"Имя: %s\nURL: %s\nТеги: %v\nДата: %s\n",
					storage[k].Name,
					storage[k].Link,
					storage[k].Tags,
					storage[k].Date.Format(time.DateTime),
				)
			}
		case 'r':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Удаление url из списка хранения
			fmt.Println("Введите имя ссылки, которое нужно удалить")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			text = strings.TrimSpace(text)

			// Напишите свой код здесь
			if _, ok := storage[text]; !ok {
				fmt.Printf("URL %s отсутствует в списке, не могу удалить\n", text)
			} else {
				delete(storage, text)
			}
		default:
			// Если нажата Esc выходим из приложения
			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}
