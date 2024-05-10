package main

import (
	"fmt"
	"listener/event"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	//try to connect  to rabbit mq
	rabbitConn, err := connect()

	if err!= nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	//start listening for messages
	log.Println("Listening for and consuming rabbitmq messages")

	//create consumer 
	consumer,err := event.NewConsumer(rabbitConn)
	if err!=nil{
		panic(err)
	}
	
	//watch the queue and consume events
	err = consumer.Listen([]string{"log.INFO","log.WARNING","log.ERROR"})
}

func connect() (*amqp.Connection,error){
	var counts int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	//dont continue until rabit is ready
	for{
		c,err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err!=nil{
			fmt.Println("rabbitmq not ready")
			counts++
		}else{
			connection = c
			log.Println("connected to RabbitMQ!!!!")
			break
		}

		if counts>5{
			fmt.Println(err)
			return nil,err
		}

		backoff = time.Duration(math.Pow(float64(counts),2))* time.Second
		log.Println("backing off...")
		time.Sleep(backoff)
		continue
	}

	return connection,nil
}