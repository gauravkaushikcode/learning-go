package main

import "fmt"

type AlphaTV struct {
	currentChannel int
	currentVolume  int
	tvOn           bool
}

func (tv *AlphaTV) getVolume() int {
	fmt.Println("AlphaTV volume is :", tv.currentVolume)
	return tv.currentVolume
}

func (tv *AlphaTV) setVolume(volume int) {
	fmt.Println("setting AlphaTV volume to :", volume)
	tv.currentVolume = volume
}

func (tv *AlphaTV) getChannel() int {
	fmt.Println("AlphaTV channel is :", tv.currentChannel)
	return tv.currentChannel
}

func (tv *AlphaTV) setChannel(channel int) {
	fmt.Println("setting AlphaTV channel to :", channel)
	tv.currentChannel = channel
}

func (tv *AlphaTV) setOnState(tvOn bool) {
	if tvOn == true {
		fmt.Println("AlphaTV is on")
	} else {
		fmt.Println("AlphaTV is off")
	}
	tv.tvOn = tvOn
}
