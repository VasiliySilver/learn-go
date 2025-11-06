package main

type Message struct {
	Author string
	Text   string
}

func main() {
	messageChan1 := make(chan Message)
	messageChan2 := make(chan Message)

	go func() {
		for {
			messageChan1 <- Message{
				Author: "Friend 1",
				Text:   "Hello!",
			}
		}
	}()

	go func() {
		for {
			messageChan2 <- Message{
				Author: "Friend 2",
				Text:   "How are you?",
			}
		}
	}()

}
