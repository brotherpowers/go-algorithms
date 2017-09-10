# Linear Search

Goal: Sort an array of numbers from low to high (or high to low).

Gnome sort (or Stupid sort) is a sorting algorithm originally proposed by an Iranian computer engineer [Dr. Hamid Sarbazi-Azad](https://en.wikipedia.org/wiki/Hamid_Sarbazi-Azad) in 2000 and called "stupid sort", and then later on described by [Dick Grune](https://en.wikipedia.org/wiki/Dick_Grune) and named "gnome sort".



## An example

Let's say we have an unsorted array of numbers `[5, 3, 2, 4]`.

We can decompose the problem of sorting n numbers in ascending order into

1. If the current array element is larger or equal to the previous array element then go one step right.

```
	[5, 3, 2, 4]
	    *
```

2. If the current array element is larger or equal to the previous array element then go one step right.

```

	[3, 5, 2, 4]
	    *
```

3. If the current array element is smaller than the previous array element then swap these two elements and go one step backwards.

```

	[3, 5, 2, 4]
	       *
		   
	[3, 2, 5, 4]
	    *	
```

4. Repeat steps 2 and 3 till we reach the end of the array.

```
	[2, 3, 4, 5]
```


## The code

Here is a simple implementation of jump search:

```go
func Sort(arr []int) {
	length := len(arr)
	if length < 1 {
		return
	}

	i := 0
	for i < length {

		if i == 0 || arr[i] >= arr[i-1] {
			i++
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i--
		}
	}
}
```

## Performance

It is conceptually simple, requiring no nested loops. The average, or expected, running time is **O(n2)**, but tends towards **O(n)** if the list is initially almost sorted.


## See also

[Gnome sort on Wikipedia](https://en.wikipedia.org/wiki/Gnome_sort)

*Written by [brotherpowers](https://www.brotherpowers.com/)*
