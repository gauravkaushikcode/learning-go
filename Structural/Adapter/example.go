package main

import "fmt"

func main() {
	tv1 := &AlphaTV{
		currentChannel: 21,
		currentVolume:  23,
		tvOn:           true,
	}
	tv2 := &BetaTV{
		volume:  20,
		channel: 8,
		isOn:    true,
	}
	tv2.turnOn()
	tv2.volumeUp()
	tv2.volumeDown()
	tv2.channelUp()
	tv2.channelDown()
	tv2.goToChannel(45)
	tv2.turnOff()

	fmt.Println("--------------------")

	alphaAdapt := &AlphatvAdapter{alphatvAdapter: tv1}
	alphaAdapt.turnOn()
	alphaAdapt.volumeUp()
	alphaAdapt.volumeUp()
	alphaAdapt.channelUp()
	alphaAdapt.channelDown()
	alphaAdapt.goToChannel(55)
	alphaAdapt.turnOff()
}
