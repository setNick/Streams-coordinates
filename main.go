package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    // Открыть файл для чтения
    file, err := os.Open("coordinates.txt")
    if err != nil {
        fmt.Println("Ошибка открытия файла:", err)
        return
    }
    defer file.Close()

    // Сканировать файл построчно
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        // Разделить строку на координаты
        parts := strings.Split(line, ",")
        if len(parts) != 2 {
            // Пропустить строку, если она не содержит две координаты
            continue
        }

        // Преобразовать координаты в числа
        latitude, err := strconv.ParseFloat(parts[0], 64)
        if err != nil {
            fmt.Println("Ошибка преобразования широты:", err)
            continue
        }

        longitude, err := strconv.ParseFloat(parts[1], 64)
        if err != nil {
            fmt.Println("Ошибка преобразования долготы:", err)
            continue
        }

        // Проверить, находится ли координата в заданном квадрате
        if isInSquare(latitude, longitude) {
            fmt.Printf("%s\n", line)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Ошибка чтения файла:", err)
        return
    }
}

// Функция проверки координаты в квадрате
func isInSquare(latitude, longitude float64) bool {
    return 50 <= latitude && latitude <= 80 && 20 <= longitude && longitude <= 45
}
