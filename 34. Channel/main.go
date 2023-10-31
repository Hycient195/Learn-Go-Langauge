package main;

import (
	"time"
	"fmt"
	"sync"
)

func main()  {
	BasicChannel();
	ReadAndWriteOnlyChannels();
	BufferedChannels();
}

func BasicChannel() {
	fmt.Println("Working with Channels in Go");

	wg := &sync.WaitGroup{};

	ch := make(chan int);

	wg.Add(2);

	go func(cha chan int, wg *sync.WaitGroup) {
		val := <-cha;
		fmt.Println(val)
		wg.Done();
	}(ch, wg);

	go func (cha chan int, wg *sync.WaitGroup) {
		cha <- 34;
		wg.Done();
	}(ch, wg);

	wg.Wait();
}


func ReadAndWriteOnlyChannels() {
	wg := &sync.WaitGroup{}
	ch := make(chan int);

	wg.Add(2);

	// Read-Only Channel (This goroutine can only read from a channel)
	go func(cha <-chan int, wg *sync.WaitGroup) {
		val := <-cha;
		fmt.Println(val);
		/* cha <- 78; */  // This operation is prohibited since the goroutine receives a read-only channel, and cannot perform write operations.
		wg.Done();
	}(ch, wg)

	// Write Only Channel (This goroutine can only write to a channel)
	go func(cha chan<- int, wg *sync.WaitGroup) {
		cha <- 56;
		/* fmt.Println(<-chan); */ // This operation is prohibited since the goroutine receives only a write only channel, and cannot perform a read operation. 
		wg.Done();
	}(ch, wg)
	wg.Wait();
}

func BufferedChannels() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 2);
	/* The basic difference in this example is the addition of a buffer to the chanel, (the digit above, as the second argument go "make")
		prescribes how many written values that are not being the channel is able to hold. If the values written to the channel that are not
		read instantly by another goroutine exceeds the number specified, and error would be thrown.
	*/

	wg.Add(2);

	// Read-Only Channel (This goroutine can only read from a channel)
	go func(cha <-chan int, wg *sync.WaitGroup) {
		val, status := <-cha;
		fmt.Println(val);
		fmt.Println(status)
		time.Sleep(2 * time.Second);
		val, status = <- cha;
		fmt.Println(val)
		fmt.Println(status);
		/* cha <- 78; */  // This operation is prohibited since the goroutine receives a read-only channel, and cannot perform write operations.
		wg.Done();
	}(ch, wg)

	// Write Only Channel (This goroutine can only write to a channel)
	go func(cha chan int, wg *sync.WaitGroup) {
		cha <- 56;
		cha <- 45
		close(cha);
		fmt.Println("All data sent")
		/* fmt.Println(<-chan); */ // This operation is prohibited since the goroutine receives only a write only channel, and cannot perform a read operation. 
		// close(cha); // closing the channel;
		wg.Done();
	}(ch, wg)


	wg.Wait();
}




/* It is good to note that else a channel is buffered, every data sent by a goroutine over a chennel must be received by another
	goroutine, else an error would be thrown. It is also good to note that a goroutine that sends data across a channel cannot
	read that data, until the channel is closed.
*/