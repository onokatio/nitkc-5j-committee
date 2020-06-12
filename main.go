package main

import "fmt"
import "math/rand"

func main() {
	StudentMax := 37
	Student := []int{}
	Roled := []int{27, 4, 7, 26, 8, 5, 28, 13, 30, 23}

	for index := 1; index <= StudentMax; index++ {
		skip := 0
		for _, roled := range Roled {
			if(index == roled){
				skip = 1
			}
		}
		if(skip == 0){
			Student = append(Student, index)
		}
	}

	var seed int64
	fmt.Print("Input Seed:")
	fmt.Scan(&seed)
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(Student), func(i int, j int) {
		Student[i], Student[j] = Student[j], Student[i]
	})

	fmt.Println("Student[] = ", Student)

	Committe := map[string]int{
		"会計":   1,
		"体育":   2,
		"環境衛生": 2,
		"交通安全": 2,
		"監査":   1,
	}

	for Key, Role := range Committe {
		fmt.Println(Key, ": ", Student[0:Role])
		Student = Student[Role:]
	}
}
