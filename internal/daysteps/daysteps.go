package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	splitData := strings.Split(data, ",")

	if len(splitData) != 2 {
		return 0, 0, fmt.Errorf("неверное количество аргументов")
	}

	steps, err := strconv.Atoi(splitData[0])

	if err != nil {
		return 0, 0, err
	}

	duration, err := time.ParseDuration(splitData[1])

	if err != nil {
		return 0, 0, err
	}

	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distance := float64(steps) * StepLength
	distance /= 1000

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км. \nВы сожгли %.2f ккал. \n", steps, distance, spentcalories.WalkingSpentCalories(steps, weight, height, duration))
}
