package main 


import ( "fmt" 
		amqp "github.com/rabbitmq/amqp091-go"
		//"math/rand"
		//"time"
		//"strconv"
		//"sync"
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

	PS: The imports you may need are already imported for you. Run go mod tidy to download them.
	

*/


func main(){

	channel_name := "sensor_data" // name of the queue we will be sending messages to

	/*

	Start by analysing the function below.
	It creates a connection to the rabbitmq server and creates a channel to send messages to.
	It then returns this channel so that we can use it to send messages to the queue and receive messages .

	*/

	channel := createRabbitMQQueue(channel_name) // create a rabbitmq queue to send messages

	/* 
	
	We now have a rabbitmq queue to where we can send messages.
	Start off by completing the function publishMessage(string message, *amqp.Channel channel, channel_name ) to send a message to the queue.
	This may be useful -> https://www.rabbitmq.com/tutorials/tutorial-one-go.html

	PS : You dont need to follow the instructions completely, its just a guide to help you. If you find a better way to do it, go for it !

	*/

	//publishMessage("Hello World", channel, channel_name) // send a message to the queue


	/*

	After testing your function, lets now create a consumer that will read from the queue.
	Complete the function consumeMessage( *amqp.Channel channel, channel_name string ) to read from the queue.
	This operation is blocking so you will only be able to send messages before you read them ( for now ).

	*/


	//consumeMessage(channel, channel_name) // read messages in the queue



	/*

	Now lets create a function that will generate random sensor data. 
	Complete the function randomSensorData() to return a random string that will represent sensor data.
	You can basically do anything here =) 


	*/


	/*

	To end this exercise, lets enable consumer and producer to run at the same time using Go Routines.
	To do this, we will use a WaitGroup from our sync import to run these functions asynchronously and tell main to wait for them to finish.

	You can do this by declaring a WaitGroup and adding 2 go routines to it.

	var wg sync.WaitGroup

	wg.Add(2)

	Now we will tell main to wait for these 2 go routines to finish:

	wg.Wait()

	( 
	  In this case we are telling main to wait for routines that never finish. 
	  We could program the routines to finish, but we dont want to in this case, they should finish when we CTRL-C 
	)

	*/


	/*

	Now we can use our randomSensorData() function in our publishMessage() function to send random data to the queue.

	We can also make the publishMessage() run on a timer.
	Since GO doesn't have a while Loop, we can use just for { expressions } to make it run forever.
	Dont forget to give it a sleep timer so that it doesnt send messages too fast.
	time.Sleep( 1 * time.Second ) // sleep for 1 second for example


	You should now have a producer and a consumer running at the same time sending data to eachother.
	Congratulations !

	*/








}


func randomSensorData() string {

	// Your code here

	// This function will return a random string that will represent sensor data
	
	return "something"

}




func consumeMessage( channel *amqp.Channel, channel_name string ){


}



func publishMessage( message string , channel *amqp.Channel, channel_name string ){



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


