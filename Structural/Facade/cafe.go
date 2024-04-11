package main

import "fmt"

func makeAmericano(size float32) {
	fmt.Println("\nMaking Americano\n ..............")
	americano := &CoffeeMachine{}
	beansAmount := (size / 8.0) * 5.0
	americano.startCoffeeMachine(beansAmount, 2)
	americano.grindBeans()
	americano.setWaterAmt(size)
	americano.stopCoffeeMachine()
	fmt.Println("Americano Ready!!")
}

func makeLatte(size float32, foam bool) {
	fmt.Println("\nMaking Latte\n ..............")
	latte := &CoffeeMachine{}
	beansAmount := (size / 8.0) * 5.0
	latte.startCoffeeMachine(beansAmount, 4)
	latte.grindBeans()
	latte.setWaterAmt(size)
	latte.setMilkAmount((size / 8.0) * 2)
	latte.doFoam(foam)
	latte.stopCoffeeMachine()
	fmt.Println("Latte Ready!!")
}
