package main

import "fmt"
import "math/rand"


func main() {
	StudentMax := 37
	Student := make([]int,StudentMax)
	for index := range Student {
		Student[index] = index + 1
	}

	var seed int64
	fmt.Print("Input Seed:")
	fmt.Scan(&seed)
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(StudentMax, func(i int, j int){
		Student[i], Student[j] = Student[j], Student[i]
	})

	fmt.Println("Student[] = ", Student)

	Committe := map[string]int{
		"会計": 1,
		"体育": 2,
		"環境衛生": 2,
		"交通安全": 2,
		"監査": 1,
	}

	for Key, Role := range Committe {
		fmt.Println(Key,": ", Student[0:Role])
		Student = Student[Role:]
	}
}
