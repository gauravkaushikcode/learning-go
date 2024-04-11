package main

import "fmt"

type CoffeeMachine struct {
	beanAmount       float32
	grinderLevel     int
	waterTemperature int
	waterAmt         float32
	milkAmount       float32
	addFoam          bool
}

func (c *CoffeeMachine) startCoffeeMachine(beanAmount float32, grinderLevel int) {
	c.beanAmount = beanAmount
	c.grinderLevel = grinderLevel
	fmt.Println("Coffee machine starting with bean amount:", beanAmount, "and grinderLevel:", grinderLevel)
}

func (c *CoffeeMachine) stopCoffeeMachine() {
	fmt.Println("Coffee machine stopping...")
}

func (c *CoffeeMachine) grindBeans() bool {
	fmt.Println("Grinding Beans", c.beanAmount, "beans at", c.grinderLevel, "granularity")
	return true
}

func (c *CoffeeMachine) setMilkAmount(milkAmt float32) float32 {
	c.milkAmount = milkAmt
	fmt.Println("Adding Milk Amount:", c.milkAmount, "oz")
	return milkAmt
}

func (c *CoffeeMachine) setWaterTemperature(waterTemperature int) {
	c.waterTemperature = waterTemperature
	fmt.Println("Adding water (temperature):", c.waterTemperature, "'C")
}

func (c *CoffeeMachine) setWaterAmt(waterAmt float32) float32 {
	c.waterAmt = waterAmt
	fmt.Println("Adding Hot Water :", c.waterAmt, "oz")
	return waterAmt
}

func (c *CoffeeMachine) doFoam(addFoam bool) {
	c.addFoam = addFoam
	fmt.Println("Adding Foam:", c.addFoam)
}
