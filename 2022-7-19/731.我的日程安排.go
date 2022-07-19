package main

import "fmt"

type pair struct {
	start, end int
}

type MyCalendarTwo struct {
	booked, overlap []pair
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, i := range this.overlap {
		if start < i.end && end > i.start {
			return false
		}
	}

	for _, i := range this.booked {
		if start < i.end && end > i.start {
			this.overlap = append(this.overlap, pair{max(i.start, start), min(i.end, end)})
		}
	}

	this.booked = append(this.booked, pair{start, end})
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func main() {
	MyCalendar := Constructor()
	fmt.Println(MyCalendar.Book(10, 20))
	fmt.Println(MyCalendar.Book(50, 60))
	fmt.Println(MyCalendar.Book(10, 40))
	fmt.Println(MyCalendar.Book(5, 15))
	fmt.Println(MyCalendar.Book(5, 10))
	fmt.Println(MyCalendar.Book(25, 55))	
}