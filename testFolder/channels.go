//Understanding Channels
/////////////////////////////////////////////

c := make(chan int)
d := make(chan int, 1) //This is a buffered channel, it allows 1 value to sit in it.
//Generally, you want to try to stay away from BufferedChannels

//c <- 42 //Channel blocks the whole program here
/*
	BUT, the below code works, because this channel is now going off on it's own...
	the Main func can run WHILE this func is running by itself.
*/
go func() {
	c <- 42
}()

d <- 45 //This dosen't block because the buffered channel gives 1 value to sit in it.
//d <- 96 //CAN'T DO THIS...ONLY ROOM FOR 1 ON BUFFERED CHANNEL! NEEDS TO BE 'PULLED OFF'
//OR YOU NEED TO HAVE THE VALUE TAKE 2 BUFFERS! d := make(chan int, 2)
fmt.Println(<-c)
fmt.Println(<-d)
d <- 96
fmt.Println(<-d)

//Directional Channels
/////////////////////////////////////////////
/*
This is the regular, bi-directional channel assignment below.
*/
e := make(chan int, 2)
/*
With the below declaration,(<-) you are only allowed to SEND values to this channel.
You cannot take values from this channel.
*/
c := make(chan<- int, 2)
/*
With the below declaration(<-chan) you are only allowed to RECIEVE values to this channel.
You cannot put values into this channel.
*/
d := make(<-chan int, 2)

c <- 42
c <- 43

//d <- 69
//d <- 70 //These would not work...you can't send values to this channel!

//fmt.Println(<-c)
//fmt.Println(<-c) These wouldn't work...they're listed as 'invalid channel'.
fmt.Println("------")
fmt.Printf("Type of channel e: %T\n", e)
fmt.Printf("Type of channel c: %T\n", c)
fmt.Printf("Type of channel d: %T\n", d)

//How/When to use Channels
/////////////////////////////////////////////
testChannel := make(chan int)
//send
go channelSender(testChannel)
//recieve
//go channelReciever(testChannel) This won't work, because they'd
//be running their functions at the same time.
channelReciever(testChannel)

fmt.Println("About to exit...")

//send function
func channelSender(testChan chan<- int) {
testChan <- 69
}

//recieve function
func channelReciever(testChan <-chan int) {
fmt.Println(<-testChan)
}


//Range
//////////////////////////////////////////////////////
testChannel := make(chan int)
testChannel2 := make(chan int)

//send
go sendFunc(testChannel)

/*
	Here's the function written in main as a self running function below
*/

go func() {
	for i := 0; i < 100; i++ {
		testChannel2 <- i
	}
	close(testChannel2) //This closes the channel when finished. Needed!
}()

for v := range testChannel {
	fmt.Println(v)
}

for v := range testChannel2 {
	fmt.Println("The second testChannel has a number of ", v)
}

fmt.Println("About to exit.")

//send
func sendFunc(testChan chan<- int) {
for i := 0; i < 100; i++ {
	testChan <- i
}
close(testChan) //This closes the channel when finished. Needed!
}

//Select Statements
/////////////////////////////////////////////
evenChannel := make(chan int)
oddChannel := make(chan int)
quitChannel := make(chan int)

//send
go send(evenChannel, oddChannel, quitChannel)

//recive
recieve(evenChannel, oddChannel, quitChannel)

fmt.Println("About to exit...")

func send(ec, oc, qc chan<- int) {
for i := 0; i < 100; i++ {
	if i%2 == 0 {
		ec <- i
	} else {
		oc <- i
	}
}
close(ec)
close(oc)
qc <- 0
}

func recieve(ec, oc, qc <-chan int) {
for {
	select {
	case v := <-ec:
		fmt.Println("From the even channel: ", v)
	case v := <-oc:
		fmt.Println("From the odd channel: ", v)
	case v := <-qc:
		fmt.Println("From the quit channel: ", v)
		return
	}
}
}