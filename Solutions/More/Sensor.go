package main 


import ( "fmt" 
		amqp "github.com/rabbitmq/amqp091-go"
		"math/rand"
		"time"
		"strconv"
		"sync"
	)

/*

	In this exercise, we will be creating a simple sensor that will send data to a RabbitMQ queue.
	If you dont know what RabbitMQ is, you can read about it here -> https://www.rabbitmq.com/
	We will also create a consumer that will read the data from the queue and print it to the terminal.
	To learn a bit about concurrency, we will also be using Go Routines to run a consumer and a producer at the same time.

	To start off, you need to have rabbitmq installed on your machine.
	If you dont, you can run this command to install it ( on Ubuntu ) -> sudo apt install rabbitmq-server
	If it doesnt work, try checking your distro's documentation on how to install rabbitmq
	After that boot it up, usually with this command -> sudo service rabbitmq-server start ( or ) sudo rabbitmq-server

*/


func main(){

	channel_name := "sensor_data" // name of the queue we will be sending messages to

	channel := createRabbitMQQueue(channel_name) // create a rabbitmq queue to send messages

	var wg sync.WaitGroup // create a wait group to handle go routines

	wg.Add(2) // add 2 go routines to the wait group







	/* 
	
	We now have a rabbitmq queue to where we can send messages.
	Start off by completing the function publishMessage(string message, *amqp.Channel channel, channel_name ) to send a message to the queue.
	This may be useful -> https://www.rabbitmq.com/tutorials/tutorial-one-go.html

	*/

	//publishMessage(randomSensorData(), channel, channel_name) // send a message to the queue


	/*

	After testing your function, lets now create a consumer that will read from the queue.
	Complete the function consumeMessage( *amqp.Channel channel, channel_name string ) to read from the queue.
	This operation is blocking ence why we will use GO routines to handle multiple producers and consumers at the same time.

	*/

	//consumeMessage(channel, channel_name) // read the message from the queue



	/*

	Now lets create a function that will generate random sensor data. 
	You can basically do anything here =) 


	*/


	/*

	To end this exercise, lets create the go routines that will run the producer and the consumer at the same time.

	*/


	go publishMessage(randomSensorData(), channel, channel_name) // send a message to the queue
	go consumeMessage(channel, channel_name) // read the message from the queue

	wg.Wait() // wait for the go routines to finish





}


func randomSensorData() string {

	// Your code here

	// This function will return a random string that will represent sensor data
	// You can use the function below to generate a random string

	sensor_type := "temperature"

	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between 20 and 25
	randomInt := rand.Intn(6) + 20

	return sensor_type + " " + strconv.Itoa(randomInt)



	


}

func consumeMessage( channel *amqp.Channel, channel_name string ){

	// Your code here

	messages, err := channel.Consume(
		channel_name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {

		panic(err)

	}

	for message := range messages {
		
		fmt.Println("Received message -> " , string(message.Body) , " from queue -> " , channel_name)

	}



}



func publishMessage( message string , channel *amqp.Channel, channel_name string ){

	// Your code here

	for  {

	err := channel.Publish( "" , channel_name , false , false , amqp.Publishing{ContentType: "text/plain", Body: []byte(message)})

	if err != nil {
		panic(err)
	}

	fmt.Println("Published message -> " , message , " to queue -> " , channel_name)
	time.Sleep(3 * time.Second)
	
}


}

func createRabbitMQQueue(channel_name string) *amqp.Channel {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // should be the default connection string for RabbitMQ
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to RabbitMQ instance")

	channel, err := connection.Channel() // create a RabbitMQ channel
	if err != nil {
		panic(err)
	}

	err = channel.ExchangeDeclare(
		channel_name, // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		panic(err)
	}

	// Declare a queue
	_, err = channel.QueueDeclare(
		channel_name, // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully declared an exchange and a queue")

	return channel
}