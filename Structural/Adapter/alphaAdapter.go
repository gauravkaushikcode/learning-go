package main

import "C"

type AlphatvAdapter struct {
	alphatvAdapter *AlphaTV
}

func (alphaTV *AlphatvAdapter) turnOn() {
	alphaTV.alphatvAdapter.setOnState(true)
}

func (alphaTV *AlphatvAdapter) turnOff() {
	alphaTV.alphatvAdapter.setOnState(false)
}

func (alphaTV *AlphatvAdapter) volumeUp() int {
	volume := alphaTV.alphatvAdapter.getVolume() + 1
	alphaTV.alphatvAdapter.setVolume(volume)
	return volume
}
func (alphaTV *AlphatvAdapter) volumeDown() int {
	volume := alphaTV.alphatvAdapter.getVolume() - 1
	alphaTV.alphatvAdapter.setVolume(volume)
	return volume
}

func (alphaTV *AlphatvAdapter) channelUp() int {
	channel := alphaTV.alphatvAdapter.getChannel() + 1
	alphaTV.alphatvAdapter.setChannel(channel)
	return channel
}

func (alphaTV *AlphatvAdapter) channelDown() int {
	channel := alphaTV.alphatvAdapter.getChannel() - 1
	alphaTV.alphatvAdapter.setChannel(channel)
	return channel
}

func (alphaTV *AlphatvAdapter) goToChannel(channel int) {
	alphaTV.alphatvAdapter.setChannel(channel)
}
