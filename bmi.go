package main

import (
	"fmt"
	"math"
)

var (
	rule   = [5]float64{18.5, 25.0, 28.0, 32.0, 0}
	accept = [5]string{"过轻", "正常", "过重", "肥胖", "严重肥胖"}
	ch     = make(chan int)
)

func main() {

	var weight, height float64

	for {

		fmt.Printf("\n身高(m)=")
		fmt.Scanln(&height)

		fmt.Print("体重(kg)=")
		fmt.Scanln(&weight)

		bmi := weight / math.Pow(height, 2)

		go Bmi(bmi, 0)
		ch <- 1
	}
}

func Bmi(bmi float64, x int) bool {
	if x > 3 {
		return false
	}
	if Bmi(bmi, x+1) {
		return true
	}
	if bmi > rule[x] {
		fmt.Printf("\nbmi=%f,%s\n", bmi, accept[x+1])
		<-ch
		return true

	} else if x == 0 && bmi < rule[x] {
		fmt.Printf("\nbmi=%f,%s\n", bmi, accept[x])
		<-ch

	}
	return false
}
