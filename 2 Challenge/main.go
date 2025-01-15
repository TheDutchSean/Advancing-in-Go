/* 

	Rolling mean challenge
	
	All of our exploration with streams brings us to our second challenge. You'll use the techniques we've learned so far to calculate the rolling mean of a data set. 
	The mean of a data set is calculated by first adding all the points in the data set and then dividing their sum by the number of points in the data set. 
	How would you do this for a very large stream of data? 
	
	First, the sum of this data set could cause overflow errors due to its size. 
	Secondly, this calculation might take a long time. So how can we get around this problem and get some meaningful averages for our data streams? 
	
	A common solution to this problem is to instead calculate the rolling mean or average, which makes use of the sliding windows technique that you saw earlier. 
	The calculation of the rolling mean is done by dividing the stream into sliding windows and calculating the mean for each of them. 
	The example demonstrates this technique for a window of size three. 
	
	Begin with the first three elements and calculate the first mean as seven. Then drop the first element and create a new window by moving one element to the right. 
	Next, calculate the new mean as 7.33. Then repeat these steps, calculating the average for each window until we reach the end of the data set. 
	This technique has the advantage of preferring new data as we slowly phase out older elements in preference of new ones. 
	
	Your challenge will be to implement the rolling mean of a stream of random numbers using the process presented in this video as well as all the knowledge you've gained in this section. 
	Write your solution considering how it will scale for large data sets, including where it will make sense to use concurrency. 
	
	As always, I have written some code to get you started. We'll be processing our rolling averages with a window size of three as defined by the constant on line 58. 
	On lines 60 to 70, define a numeric type interface, which defines the acceptable types of the moving average struct. 
	This struct contains fields which you may find useful for window management, the position of the element, the values of the window, and whether the window is filled. 
	It also contains two channels, one for receiving the stream input values and one for sending the calculated rolling means. 
	Channels are a great way to represent streams as they maintain ordering. 
	
	On lines 92 to 98, we have defined the new moving average function which takes in the two channels required by the moving average struct as parameters. 
	This function initializes the moving average struct and returns it ready for use. 
	
	Further down on lines 100 to 110, 
	I have implemented the calculate mean method for you. It calculates the mean and prints the window together with its mean. 
	
	Your challenge is to implement the rolling mean method defined on lines 112 to 120, which currently prints a message to the terminal and reads values. 
	This method should manage the window values received from the input channel and invoke the calculate mean method accordingly. 
	
	One thing to note is that I have also defined a read results function on line 123 to 129, which reads all the values of the output channel and writes them to a slice. 
	This function will make it easier to verify the behavior of our implementation in this limited test environment. 
	This wouldn't be something you would or should do for a large data sets. 
	
	Take your time to solve this challenge and join me in the next video to see my solution. Good luck. 

*/


// Write your answer here, and then test your code.
// Your job is to implement the RollingMean() method. ( inputs 5, 7, 9, 6, 8, 10, returns [5,7,9] = 7, [6,7,9] = 7.33, [6,8,9] = 7.67, [6,8,10] = 8)

package main

import (
	"fmt"
	"strings"
)


const WindowSize = 3

type NumericType interface {
	int64 | float64
}

type MovingAverage[T NumericType] struct {
	Position     int
	WindowValues [WindowSize]T
	WindowFilled bool
	Input        <-chan T
	Output       chan<- string
}


func main(){

	numbers := []int64{5, 7, 9, 6, 8, 10}
	input := make(chan int64)
	output := make(chan string)
	movingAvg := NewMovingAverage(input, output)
	go movingAvg.RollingMean()
	go func() {
		for _, n := range numbers {
			input <- n
		}
		close(input)
	}()
	
	learnerResult := ReadResults(output)
	fmt.Println(learnerResult)
}


func NewMovingAverage[T NumericType](in <-chan T, 
    out chan<- string) *MovingAverage[T] {
	return &MovingAverage[T]{
		Input:  in,
		Output: out,
	}
}

// CalculateMean prints the window and its calculated mean.
func (ma *MovingAverage[T]) CalculateMean() string {
	values := make([]string, WindowSize)
	var sum float64
	for i, v := range ma.WindowValues {
		values[i] = fmt.Sprint(v)
		sum += float64(v)
	}
	mean := sum / WindowSize
	return fmt.Sprintf("[%s] = %.2f", strings.Join(values, ","), mean)
}

func (ma *MovingAverage[T]) RollingMean() {
	fmt.Println("RollingMean not implemented")
	for {
		if _, ok := <-ma.Input; !ok {
			close(ma.Output)
			return
		}
	}
}


func ReadResults(output <-chan string) []string {
	var results []string
	for r := range output {
		results = append(results, r)
	}
	return results
}

