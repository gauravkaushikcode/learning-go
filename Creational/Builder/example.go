package main

import "fmt"

func main() {
	var builder = newNotificationBuilder()

	builder.SetTitle("New Notification")
	builder.SetSubtitle("Notification Subtitle")
	builder.SetIcon("icon.png")
	builder.SetImage("image.jpg")
	builder.SetPriority(5)
	builder.SetMessage("This is first notification")
	builder.SetType("alert")

	notification, err := builder.Build()
	if err != nil {
		fmt.Println("Error creating notification:", err)
	} else {
		fmt.Printf("Notification: %+v", notification)
	}
}
