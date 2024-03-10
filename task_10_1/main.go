package main

import "fmt"

func main() {
	test := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10
	fmt.Println(groupByStep(test, step))
}

func groupByStep(temps []float32, step int) map[int][]float32 {
	group := make(map[int][]float32)
	for _, temp := range temps {
		// Получаем ключ группы для данной температуры, если такой еще нет, то создаем
		key := step * int(temp/float32(step))
		if _, isExist := group[key]; !isExist {
			group[key] = make([]float32, 0)
		}
		group[key] = append(group[key], temp)
	}
	return group
}
