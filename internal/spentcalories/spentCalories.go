package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	splitData := strings.Split(data, ",")

	if len(splitData) != 3 {
		return 0, "", 0, fmt.Errorf("неверное количество аргументов \n")
	}

	steps, err := strconv.Atoi(splitData[0])
	if err != nil {
		return 0, "", 0, err
	}

	walkingDuration, err := time.ParseDuration(splitData[2])

	if err != nil {
		return 0, "", 0, err
	}

	return steps, splitData[1], walkingDuration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	return (float64(steps) * lenStep) / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	return distance(steps) / duration.Hours()
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	staps, trainingType, duration, err := parseTraining(data)

	if err != nil {
		return ""
	}

	switch trainingType {
	case "Ходьба":
		return fmt.Sprintf("Тип тренировки: %s \nДлительность: %.2f ч. \nДистанция: %.2f км. \nСкорость: %.2f км/ч \nСожгли калорий: %.2f \n", trainingType, duration.Hours(), distance(staps), meanSpeed(staps, duration), WalkingSpentCalories(staps, weight, height, duration))
	case "Бег":
		return fmt.Sprintf("Тип тренировки: %s \nДлительность: %.2f ч. \nДистанция: %.2f км. \nСкорость: %.2f км/ч \nСожгли калорий: %.2f \n", trainingType, duration.Hours(), distance(staps), meanSpeed(staps, duration), RunningSpentCalories(staps, weight, duration))
	default:
		return fmt.Sprintf("неизвестный тип тренировки \n")
	}

}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	return ((runningCaloriesMeanSpeedMultiplier * meanSpeed(steps, duration)) - runningCaloriesMeanSpeedShift) * weight
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	return ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed(steps, duration)*meanSpeed(steps, duration)/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH

}
