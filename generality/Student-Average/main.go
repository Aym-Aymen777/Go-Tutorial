package main

import (
	"fmt"
)

//declaring the student struct
type student struct {
	name string
	age int
	grades []module
	average float64
	isPassed bool
}

type module struct {
	name string
	grade float64
}

func FillStudents() []student {
	students := []student{}
	numberOfStudents := 0
	fmt.Println("Enter the number of students: ")
	fmt.Scanf("%d", &numberOfStudents)
	for i:=0;i<numberOfStudents;i++{
		name := ""
		age := 0
		grades:=[]module{}
		fmt.Printf("======= Student %d ======== \n",i+1)
		fmt.Printf("Name : ")
		fmt.Scan(&name)
		fmt.Printf("Age : ")
		fmt.Scan(&age)
		for j:=0;j<3;j++{
			moduleName := ""
			moduleGrade := 0.0
			fmt.Printf("Module %d : ",j+1)
			fmt.Printf("Name : ")
			fmt.Scan(&moduleName)
			fmt.Printf("Grade : ")
			fmt.Scan(&moduleGrade)
			if moduleGrade > 20 || moduleGrade < 0 {
				fmt.Println("Invalid grade, grade is between 0 and 20")
				fmt.Scan(&moduleGrade)
			}
			grades = append(grades,module{moduleName,moduleGrade})
		}
		students = append(students,student{name,age,grades,0.0,false})
	}
	return students	
}

func CalcAverageAndCheckIfPassed(students []student)[]student{
	for i:=0;i<len(students);i++{
      average := 0.0
	  for j:=0;j<len(students[i].grades);j++{
		average += students[i].grades[j].grade
	  }
	  average /= float64(len(students[i].grades))
	  students[i].average = average
	  if average >= 10 {
		students[i].isPassed = true
	  }
	}
	return students
}

func CleanPrint(students []student) {
	for i:=0;i<len(students);i++{
		fmt.Println("Name : ",students[i].name + " Age : " + fmt.Sprintf("%d",students[i].age) + " Average : " + fmt.Sprintf("%.2f",students[i].average) + " Passed : " + fmt.Sprintf("%t",students[i].isPassed))
}
}

//TODO: emplement the logic of while true loop that will keep the program running until the user exits and add the feature of undo to the previous scan

func main()  {
	students := FillStudents()
	students = CalcAverageAndCheckIfPassed(students) 
	CleanPrint(students)
}