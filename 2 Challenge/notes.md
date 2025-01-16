	Rolling mean challenge
	

	The mean of a data set is calculated by first adding all the points in the data set and then dividing their sum by the number of points in the data set. 
	How would you do this for a very large stream of data? 
	

    Calculating the mean

    Sum all the elements in a dataset and dividie it by the number of elemenets in the dataset.

    window size = m
    sum(xm) = x0 + x1 + ... + xm
    mean(xm) = sum(xm) / m


    First, the sum of this data set could cause overflow errors due to its size. 
	Secondly, this calculation might take a long time. So how can we get around this problem and get some meaningful averages for our data streams? 

    A common solution to this problem is to instead calculate the rolling mean or average, which makes use of the sliding windows technique that you saw earlier. 

        Windowing this technique divides a continuous stream of data into smaller subsets called windows. 
        These manageable subsets make it possible for us to run calculations and aggregations instead of processing the entire stream at once. 
        These windows can be split in a few different ways. The first way of splitting data is through non-overlapping windows. 
	
	    Two common ways to split windows are count-based or time-based. 
		Count-based windows produce subsets of N sample size but unequal time durations. 
		Time-based windows produce subsets of the same time duration of N seconds, but unequal sample size. 
		
        Sliding windows are a variation of both time and count-based windows with the key difference that they allow overlapping data subsets. 
        Sliding windows change their starting point as new data points arrive, but maintain their length, either time-based or count-based. 


    The calculation of the rolling mean is done by dividing the stream into sliding windows and calculating the mean for each of them. 
	The example demonstrates this technique for a window of size three.

    Full data set -> (5, 7, 9, 6, 8, 10) = 7.5

    Window n = [v1, v2, v3] = sum(v1 + v2 + v3) / len(Window n)

    Window 1 [5,7,9] = 7
    Window 2 [6,7,9] = 7.33
    Window 3 [6,8,9] = 7.67
    Window 4 [6,8,10] = 8


    Create a window by creating a slice of the data set and incrementing the index 

    Window n = Full data set {index : (index + length window)}

    index = 0
    length = 2 

    Window 1 = Full data set(index:(index + length))
    index++
    Window 2 = Full data set(index:(index + length))
    index++
    Window 3 = Full data set(index:(index + length))
    index++
    Window 4 = Full data set(index:(index + length))
    

    Begin with the first three elements and calculate the first mean as seven. Then drop the first element and create a new window by moving one element to the right. 
	Next, calculate the new mean as 7.33. Then repeat these steps, calculating the average for each window until we reach the end of the data set. 
	This technique has the advantage of preferring new data as we slowly phase out older elements in preference of new ones. 