package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// N computers in the network
const N = 20

// Message struct constitutes
type Message struct {
	message     string
	recipientID int
}

// Computer struct handles the message processes
type Computer struct {
	ID         int
	prev, next chan Message
	hasMessage bool
	msg        Message
}

func (c *Computer) bootup() {
	go c.generateMessage()
	go c.receiveMessage()
}

func (c *Computer) generateMessage() {
	for {
		num1 := rand.Intn(30)
		num2 := rand.Intn(30)
		if num1 == num2 {
			c.hasMessage = true
			c.msg.recipientID = rand.Intn(N-1) + 1
			c.msg.message = "h" + strconv.Itoa(rand.Int())
			c.sendMessage()
		}
	}
}

func (c *Computer) sendMessage() {
	fmt.Print("Computer ", c.ID, " -> ", c.msg.recipientID, ": ", c.msg.message, "\n")
}

func (c *Computer) receiveMessage() {

}

func (c *Computer) receiveToken(tokenChan chan bool) bool {
	<-tokenChan

	return false
}

func passToken(network [N]Computer) {
	var tokenChan = make(chan bool, 1)
	for {
		for i := 0; i < N; i++ {
			tokenChan <- true
			if !network[i].receiveToken(tokenChan) {

			}
		}
	}
}

func main() {
	var communicationChannels [N]chan Message
	for i := 0; i < N; i++ {
		communicationChannels[i] = make(chan Message, 1)
	}
	var network [N]Computer
	for i := 0; i < N; i++ {
		network[i].ID = i + 1
		if i == 0 {
			network[i].prev = communicationChannels[N-1]
		} else {
			network[i].prev = communicationChannels[i-1]
		}
		if i == N-1 {
			network[i].next = communicationChannels[0]
		} else {
			network[i].next = communicationChannels[i+1]
		}
		go network[i].bootup()
	}
	time.Sleep(1 * time.Second)
}
