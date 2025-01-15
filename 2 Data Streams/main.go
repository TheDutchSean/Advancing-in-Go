/*  

	2.1 Data Steams Basic

	Data streams have two main characteristics. First, they are continuous, potentially infinite flows of data.
  	The values are discrete, ordered, and sequential. Therefore, the order in which they are processed must be respected. 
	Often, values will contain a key, like an ID, which will indicate order. 
	Secondly, due to their size, data streams require specialized technologies and programming techniques. 
	As you can imagine, we can't save infinite data sets into memory or aggregate them in any meaningful way.

	Specific Steaming Technologies

	Apache Kafka
	Google Cloud Pub/Sub
	RabbitMQ
	Amazon Kinesis


	2.2 Data streams processing techniques

	Here are two common stream processing techniques, creating append-only logs and windowing. 
	
	Append-only logs are data structures which are designed to record a series of sequential entries. 
	They're also known as transaction logs or write-ahead logs. Records are added at the end when they are processed and are considered immutable once they are added. 
	They also preserve the order of events, ensuring that the traversal of the list remains consistent.

	Append-only logs are optimized for quick writes, which makes them ideal for processing streams with quickly occurring values. 
	If a new value is received, then another record is written instead of the old record being updated. 
	The new version of the record can be found and used during the log reconciliation process. 
	Append-only logs are often used by database systems for storing their data. They're also used by event sourcing architectures for capturing and replaying events. 
	
	
	Windowing this technique divides a continuous stream of data into smaller subsets called windows. 
	These manageable subsets make it possible for us to run calculations and aggregations instead of processing the entire stream at once. 
	These windows can be split in a few different ways. The first way of splitting data is through non-overlapping windows. 
	
	Two common ways to split windows are count-based or time-based. 
		Count-based windows produce subsets of N sample size but unequal time durations. 
		Time-based windows produce subsets of the same time duration of N seconds, but unequal sample size. 
		
	Sliding windows are a variation of both time and count-based windows with the key difference that they allow overlapping data subsets. 
	Sliding windows change their starting point as new data points arrive, but maintain their length, either time-based or count-based. 
	This technique allows us to aggregate historical data together with incoming data, slowly moving preference towards new data as historical data is phased out. 
	
	As is common in most algorithm problems, windowing is all about breaking the problem into manageable sub-problems. 
	In the case of streams, the amount of data is the biggest challenge. So the purpose of windowing is to break down the data sets into manageable subsets. 


	3.1  Concurrent processing

	Concurrency is a way for us to make the most of our computing resources and aggregate multiple data subsets or windows as we called them in the previous video. 
	You can easily leverage concurrency in our data processing by using goroutines and channels. 
	
	As we slice and subset our data, we might want to speed up processing by taking advantage of concurrency. 
	This is one of Go's great strengths. In this example, we create a number slice, loop over each number and start a new goroutine to process it. 
	Once processing is completed, the goroutines send their computed value to the results channel. 
	After the for loop, once all the goroutines are started, the main function receives all the computed values from the results channel and prints them. 
	Due to their nature, channels will maintain the order of the results. 
	
	(see line 95)

	While the concurrent solution we've just seen is correct and will make use of concurrency to calculate each element of the number slice, let's consider its use of resources. 
	We'll start as many goroutines as we have elements in the number slice. 
	These goroutines will then only be able to complete once the main goroutine has been able to receive their values. 
	This will make the process slow and will not be suitable for the large data streams that we have been discussing so far. 
	
	An alternative solution which addresses these shortcomings. Begin by creating a work function which takes in three channels. 
	One for input, one for the results, and a done channel. Use unidirectional channel types to ensure type safety inside the work function. 
	This function makes use of a for loop, together with the select keyword, which makes it possible for one single goroutine to process multiple inputs. 
	As soon as a message is received on the done channel, shut down the goroutine using the return statement. 
	Otherwise, the function will receive values from the in channel and process them, sending the value to the out channel continuously. 
	The only return statement is on received values from the done channel. 
	
	Let's have a look at the correct invocation of the workers from the main function. 
	Initialize the in and results channels as buffered channels of the same size as the worker count. 
	This will allow values to be immediately sent through these channels if they have the capacity, reducing the blocking of the workers. 
	We start the correct number of goroutines by using a for loop and passing them the channels as parameters. 
	Next, loop through the numbers and send each of them to the in channel. 
	We do this in another goroutine. Once all the values are sent, close the done channel, signaling to our workers that they should shut down. 
	Finally, once everything is completed, the main goroutine is ready to receive all the computed values and print them to the terminal as we have done before. 
	The pattern demonstrated by this simple example is fairly common. 
	Reusing goroutines makes it easier for us to control the scale of our processing logic which is essential for large data streams. 
	It's also common to use another signal channel which signals to the workers when it is time to shut down. 

*/

package main

import (
	"log"
)


// Correct invocation of the workers

const workers int = 2

func main(){

	/*  Correct invocation of the workers
	
		Both inputs and results channels are bufferd to unblock workers immediatly.

		We write all the numbers to the in channel, clossing the done channel once they have all been written.

	*/ 

	numbers := []int{0, 1, 2, 3}
	in := make(chan int, workers)
	results := make(chan int, workers)
	done := make(chan struct{})

	for i := 0; i < workers; i++ {
		go work(in, results, done)
	}

	go func(){
		for _, n := range numbers {
			in <- n
		}
		close(done)
	}()

	for i := 0; i < len(numbers); i++{
		log.Printf("Result %d = %d \n", i , <-results)
	}


}


func Concurrent_solution(){

	/*  Concurrent solution -> example
	
		Goroutines are used for the concurrent processing of data entries.

		Channels are used to consolidate the results together

		Due to the properties of channels results will remain orderd

	*/ 

	numbers := []int{0, 1, 2, 3}
	results := make(chan int)
	for _, n := range numbers {
		go func(n int){
			results <- i * n
		}(n)
	}

	for i := 0; i < len(numbers); i++{
		log.Printf("Result %d = %d \n", i , <-results)
	}

}

// alternative solution

/* 

	Channels are used for inpute, output and completion communication.

	The select statement together with a for loop allows us to process mutiple input with the same goroutine

*/

func work(in <- chan int, out chan<-int, done <-chan struct{}){
	for {
		select {
			case <- done:
				return
			case n := <-in:
				out <- 2 * n	
		}
	}
}

