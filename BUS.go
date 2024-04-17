package main

import "fmt"

func main() {
	station1 := "ул. Будапештская"
	station2 := "ул. Славы 4"
	station3 := "Вечный огонь"
	station4 := "Аэропорт"

	tickets := 20
	passengers := 0
	totalSum := 0
	gotOn := 0
	gotOff := 0

	fmt.Println("Остановка", station1)
	fmt.Println("В салоне пассажиров:", passengers)
	fmt.Println("Сколько зашло пассажиров?", gotOn)
	fmt.Scan(&gotOn)
	totalSum += passengers * tickets
	passengers += gotOn
	fmt.Println("Отправляемся с остановки", station1)
	fmt.Println()

	fmt.Println("Остановка", station2)
	fmt.Println("В салоне пассажиров:", passengers)
	fmt.Println("Сколько зашло пассажиров?", gotOff)
	fmt.Scan(&gotOff)
	passengers -= gotOff
	fmt.Println("Сколько зашло пассажиров?", gotOn)
	fmt.Scan(&gotOn)
	totalSum += passengers * tickets
	passengers += gotOn
	fmt.Println("Отправляемся с остановки", station2)
	fmt.Println()

	fmt.Println("Остановка", station3)
	fmt.Println("В салоне пассажиров:", passengers)
	fmt.Println("Сколько зашло пассажиров?", gotOff)
	fmt.Scan(&gotOff)
	passengers -= gotOff
	fmt.Println("Сколько зашло пассажиров?", gotOn)
	fmt.Scan(&gotOn)
	totalSum += passengers * tickets
	passengers += gotOn
	fmt.Println("Отправляемся с остановки", station3)
	fmt.Println()

	fmt.Println("Остановка", station4)
	fmt.Println("В салоне пассажиров:", passengers)
	fmt.Println("Конечная, все пассажиры вышли")
	fmt.Scan(&gotOff)
	passengers = 0
	fmt.Println()

	driverSalary := totalSum / 4
	fuel := totalSum / 5
	tax := totalSum / 5
	repair := totalSum / 5
	profit := totalSum - (driverSalary + fuel + tax + repair)

	fmt.Println("Всего заработали:", totalSum, "руб")
	fmt.Println("Зарплата водителя:", driverSalary, "руб")
	fmt.Println("Расходы на топливо:", fuel, "руб")
	fmt.Println("Налоги:", tax, "руб")
	fmt.Println("Расходы на ремонт машины:", repair, "руб")
	fmt.Println("Итого доход:", profit, "руб")

}
