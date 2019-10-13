package main

// func hello(user string) string {
// 	message := "Welcome " + user
// 	return message
// }

func hello(user string) string {
	var message string
	if len(user) == 0 {
		message = "Welcome Boss"
	} else {
		message = "Welcome " + user
	}
	return message
}
