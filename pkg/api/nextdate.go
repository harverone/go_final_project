package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/harverone/go_final_project/pkg/db"
)

const dateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	startDate, err := time.Parse("20060102", dstart)
	if err != nil {
		return "", errors.New("неверный формат даты")
	}

	if repeat == "" {
		return "", errors.New("нет правила")
	}

	repeat = strings.TrimSpace(repeat)
	parts := strings.Fields(repeat)
	if len(parts) == 0 {
		return "", errors.New("нет правила")
	}

	// Обработка правил
	switch parts[0] {
	case "d":
		return dayRule(now, startDate, parts)
	case "y":
		return yearRule(now, startDate)
	default:
		return "", errors.New("ошибка формата правила")
	}
}

// ежедневные правила
func dayRule(now, startDate time.Time, parts []string) (string, error) {
	if len(parts) != 2 {
		return "", errors.New("ошибка формата правила 'd'")
	}

	days, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", errors.New("неверное количество дней")
	}

	if days < 1 || days > 400 {
		return "", errors.New("интервал дней должен быть от 1 до 400")
	}

	current := startDate
	for {
		current = current.AddDate(0, 0, days)
		if current.After(now) {
			break
		}
	}
	return current.Format("20060102"), nil
}

// ежегодные правила
func yearRule(now, startDate time.Time) (string, error) {

	current := startDate
	for {

		nextYear := current.Year() + 1
		nextDate := time.Date(nextYear, current.Month(), current.Day(), 0, 0, 0, 0, time.UTC)

		// проверка 29 февраля
		if current.Month() == time.February && current.Day() == 29 {

			feb29 := time.Date(nextYear, time.February, 29, 0, 0, 0, 0, time.UTC)
			if feb29.Month() != time.February || feb29.Day() != 29 {

				nextDate = time.Date(nextYear, time.March, 1, 0, 0, 0, 0, time.UTC)
			}
		}

		current = nextDate

		if current.After(now) {
			break
		}
	}
	return current.Format("20060102"), nil
}

func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметры из запроса
	nowParam := r.URL.Query().Get("now")
	dateParam := r.URL.Query().Get("date")
	repeatRule := r.URL.Query().Get("repeat")

	// Обработка now
	nowTime := time.Now().UTC()
	if nowParam != "" {
		parsedNow, err := time.Parse(dateFormat, nowParam)
		if err != nil {
			http.Error(w, "Неверный формат параметра now", http.StatusBadRequest)
			return
		}
		nowTime = parsedNow
	}

	// Проверка параметров
	if dateParam == "" {
		http.Error(w, "Параметр date обязателен", http.StatusBadRequest)
		return
	}
	if repeatRule == "" {
		http.Error(w, "Параметр repeat обязателен", http.StatusBadRequest)
		return
	}

	nextDate, err := NextDate(nowTime, dateParam, repeatRule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, nextDate)
}

func checkDate(task *db.Task) error {
	var next string
	now := time.Now()
	if task.Date == "" {
		task.Date = now.Format(dateFormat)
	}
	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		return fmt.Errorf("ошибка преобразования даты, task.Date: %s, err: %v", task.Date, err)
	}
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return fmt.Errorf("ошибка преобразования даты, task.Date: %s, err: %v", task.Date, err)
		}
	}
	if afterDate(now, t) {
		if len(task.Repeat) == 0 || task.Repeat == "d 1" || t.Day() == now.Day() {
			task.Date = now.Format(dateFormat)
		} else {
			task.Date = next
		}
	}
	return nil
}

func afterDate(dateAfter, dateBefore time.Time) bool {
	return dateAfter.After(dateBefore)
}
