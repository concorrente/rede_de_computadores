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
	senderID    int
}

// Computer struct handles the message processes
type Computer struct {
	ID         int
	prev, next chan Message
	hasMessage bool
	hasToken   bool
	msg        Message
}

func (c *Computer) bootup() {
	go c.generateMessage()
	go c.receiveMessage()
}

func (c *Computer) generateMessage() {
	for {
		c.hasMessage = false
		num1 := rand.Intn(5)
		num2 := 3
		if num1 == num2 {
			c.hasMessage = true
			c.msg.message = "a" + strconv.Itoa(rand.Intn(300))
			c.msg.recipientID = rand.Intn(N-1) + 1
			c.msg.senderID = c.ID
		}
	}
}

func (c *Computer) sendMessage() {
	fmt.Print("Computer ", c.ID, " -> ", c.msg.recipientID, ": ", c.msg.message, "\n")
	c.hasMessage = false
	c.next <- c.msg
}

func (c *Computer) receiveMessage() {
	for {
		packet := <-c.prev
		if c.ID == packet.recipientID {
			fmt.Print("[", c.ID, "] Message received from ", packet.senderID, "\n")
		} else {
			c.next <- packet
		}
	}
}

func (c *Computer) receiveToken(tokenChan chan bool) {
	c.hasToken = <-tokenChan
	if c.hasMessage {
		// Keep the token and send the message
		c.sendMessage()
	}
	c.hasToken = false
	tokenChan <- true
}

func passToken(network [N]Computer) {
	var tokenChan = make(chan bool, 1)
	tokenChan <- true
	for {
		for i := 0; i < N; i++ {
			network[i].receiveToken(tokenChan)
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
	go passToken(network)
	time.Sleep(10 * time.Second)
}
