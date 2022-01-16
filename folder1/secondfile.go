package folder1

import (
	"fmt"
	"math"
)

func calculateWater(waterMachine int) (bool, int) {
	const waterPerCup int = 200
	return calculateIfPossible(waterMachine, waterPerCup)
}

func calculateMilk(milkMachine int) (bool, int) {
	const milkPerCup int = 50
	return calculateIfPossible(milkMachine, milkPerCup)
}

func calculateBeans(beansMachine int) (bool, int) {
	const beansPerCup int = 15
	return calculateIfPossible(beansMachine, beansPerCup)
}

func calculateIfPossible(inMachine int, perCup int) (bool, int) {
	if inMachine < perCup {
		return false, 0
	} else if inMachine == perCup {
		return true, 1
	} else {
		return true, inMachine / perCup
	}
}

func howManyCupsCanWeMake(cupsWater int, cupsMilk int, cupsBeans int) int {
	if cupsWater == cupsMilk && cupsMilk == cupsBeans {
		return cupsWater
	} else if float64(cupsWater) > math.Min(float64(cupsMilk), float64(cupsBeans)) {
		return int(math.Min(float64(cupsMilk), float64(cupsBeans)))
	} else {
		return cupsWater
	}
}

func calculateAnswerMessage(waterMachine int, milkMachine int, beansMachine int, cups int) {
	var cupsCanMake int
	resultWater, cupsWater := calculateWater(waterMachine)
	fmt.Println("resultWater", resultWater, "cupsWater", cupsWater)
	resultMilk, cupsMilk := calculateMilk(milkMachine)
	fmt.Println("resultMilk", resultMilk, "cupsMilk", cupsMilk)
	resultBeans, cupsBeans := calculateBeans(beansMachine)
	fmt.Println("resultBeans", resultBeans, "cupsBeans", cupsBeans)

	if resultWater && resultMilk && resultBeans {
		if cups == 1 && (cupsWater == cups || cupsMilk == cups || cupsBeans == cups) {
			fmt.Println("Yes, I can make that amount of coffee")
		} else {
			cupsCanMake = howManyCupsCanWeMake(cupsWater, cupsMilk, cupsBeans)
			fmt.Println("cupsCanMake", cupsCanMake, "cups", cups)
			switch {
			case cups > cupsCanMake:
				fmt.Println("No, I can make only", cupsCanMake, "cups of coffee")
			case cups <= cupsCanMake:
				fmt.Println("Yes, I can make that amount of coffee (and even", cupsCanMake-cups, "more than that)")
			}
		}

	} else {
		if cups == 0 {
			fmt.Println("Yes, I can make that amount of coffee")
		} else {
			fmt.Println("No, I can make only 0 cups of coffee")
		}
	}
}

func Calculate() {
	var cups int
	var waterMachine int
	var milkMachine int
	var beansMachine int

	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scan(&waterMachine)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scan(&milkMachine)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scan(&beansMachine)
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scan(&cups)
	calculateAnswerMessage(waterMachine, milkMachine, beansMachine, cups)
}
