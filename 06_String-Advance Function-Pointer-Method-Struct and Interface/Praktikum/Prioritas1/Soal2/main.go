package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (e Student) Average() float64 {
	sum := 0.0
    for _, num := range e.score {
        sum += float64(num)
    }
    average := sum / float64(len(e.score)) 

	return average
}

func (e Student) Min() int {
	min := e.score[0] 
    for _, num := range e.score {
        if num < min {
            min = num 
        }
	}
	return min
}

func (e Student) Max() int {
	max := e.score[0] 
    for _, num := range e.score {
        if num > max {
            max = num 
        }
	}
	return max
}

func main() {
	students := []Student {
		{name: []string{"Rizky"}, score: []int{80}},
		{name: []string{"Kobar"}, score: []int{75}},
		{name: []string{"Ismail"}, score: []int{70}},
		{name: []string{"Umam"}, score: []int{75}},
		{name: []string{"Kobar"}, score: []int{60}},
	}

	minStudent := students[0]
	maxStudent := students[0]
	totalScore := 0

	for _, student := range students {
		totalScore += student.Min()
		
		if student.Max() > maxStudent.Max() {
			maxStudent = student
		}
		if student.Min() < minStudent.Min() {
			minStudent = student
		}
	}

	averageScore := float64(totalScore) / float64(len(students))

	fmt.Printf("Average score: %.f\n", averageScore)
	fmt.Printf("Max score student: %s (%d)\n", maxStudent.name[0], maxStudent.Max())
	fmt.Printf("Min score student: %s (%d)\n", minStudent.name[0], minStudent.Min())
}