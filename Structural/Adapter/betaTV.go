package main

import "fmt"

type BetaTV struct {
	volume  int
	channel int
	isOn    bool
}

func (tv *BetaTV) turnOn() {
	fmt.Println("Beta TV turning on")
	tv.isOn = true
}

func (tv *BetaTV) turnOff() {
	fmt.Println("Beta TV turning off")
	tv.isOn = false
}

func (tv *BetaTV) volumeUp() int {
	tv.volume++
	fmt.Println("Inc. Beta TV volume to", tv.volume)
	return tv.volume
}

func (tv *BetaTV) volumeDown() int {
	tv.volume--
	fmt.Println("Dec. Beta TV volume to", tv.volume)
	return tv.volume
}

func (tv *BetaTV) channelUp() int {
	tv.channel++
	fmt.Println("Inc. Beta TV volume to", tv.channel)
	return tv.channel
}

func (tv *BetaTV) channelDown() int {
	tv.channel--
	fmt.Println("Dec. Beta TV volume to", tv.channel)
	return tv.channel
}

func (tv *BetaTV) goToChannel(i int) {
	tv.channel = i
	fmt.Println("Set Beta TV channel to", tv.channel)
}
