package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	firstName, lastName, department1, department2, department3       string
	physicsScore, chemistryScore, mathScore, compScore, specialScore float32
}

func main() {
	var M int

	fmt.Scan(&M)

	file, err := os.Open("applicants.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var allStudents []Student

	var biotechStudents []Student
	var chemistryStudents []Student
	var engineeringStudents []Student
	var mathStudents []Student
	var physicsStudents []Student

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		studentData := strings.Split(scanner.Text(), " ")

		firstName := studentData[0]
		lastName := studentData[1]
		physicsScore, err := strconv.ParseFloat(studentData[2], 32)
		chemistryScore, err := strconv.ParseFloat(studentData[3], 32)
		mathScore, err := strconv.ParseFloat(studentData[4], 32)
		compScore, err := strconv.ParseFloat(studentData[5], 32)
		specialScore, err := strconv.ParseFloat(studentData[6], 32)

		if err != nil {
			log.Fatal(err)
		}
		department1 := studentData[7]
		department2 := studentData[8]
		department3 := studentData[9]

		student := Student{firstName, lastName, department1, department2, department3, float32(physicsScore), float32(chemistryScore), float32(mathScore), float32(compScore), float32(specialScore)}

		switch student.department1 {
		case "Biotech":
			biotechStudents = append(biotechStudents, student)
		case "Chemistry":
			chemistryStudents = append(chemistryStudents, student)
		case "Engineering":
			engineeringStudents = append(engineeringStudents, student)
		case "Mathematics":
			mathStudents = append(mathStudents, student)
		case "Physics":
			physicsStudents = append(physicsStudents, student)
		}

		allStudents = append(allStudents, student)
	}

	allStudents = sortList(allStudents, "Biotech")
	biotechStudents = selectStudents1(allStudents, M, "Biotech")

	allStudents = sortList(allStudents, "Chemistry")
	chemistryStudents = selectStudents1(allStudents, M, "Chemistry")

	allStudents = sortList(allStudents, "Engineering")
	engineeringStudents = selectStudents1(allStudents, M, "Engineering")

	allStudents = sortList(allStudents, "Mathematics")
	mathStudents = selectStudents1(allStudents, M, "Mathematics")

	allStudents = sortList(allStudents, "Physics")
	physicsStudents = selectStudents1(allStudents, M, "Physics")

	allStudents = sortList(allStudents, "Biotech")
	biotechStudents = selectStudents2(allStudents, biotechStudents, M, "Biotech")

	allStudents = sortList(allStudents, "Chemistry")
	chemistryStudents = selectStudents2(allStudents, chemistryStudents, M, "Chemistry")

	allStudents = sortList(allStudents, "Engineering")
	engineeringStudents = selectStudents2(allStudents, engineeringStudents, M, "Engineering")

	allStudents = sortList(allStudents, "Mathematics")
	mathStudents = selectStudents2(allStudents, mathStudents, M, "Mathematics")

	allStudents = sortList(allStudents, "Physics")
	physicsStudents = selectStudents2(allStudents, physicsStudents, M, "Physics")

	allStudents = sortList(allStudents, "Biotech")
	biotechStudents = selectStudents3(allStudents, biotechStudents, M, "Biotech")

	allStudents = sortList(allStudents, "Chemistry")
	chemistryStudents = selectStudents3(allStudents, chemistryStudents, M, "Chemistry")

	allStudents = sortList(allStudents, "Engineering")
	engineeringStudents = selectStudents3(allStudents, engineeringStudents, M, "Engineering")

	allStudents = sortList(allStudents, "Mathematics")
	mathStudents = selectStudents3(allStudents, mathStudents, M, "Mathematics")

	allStudents = sortList(allStudents, "Physics")
	physicsStudents = selectStudents3(allStudents, physicsStudents, M, "Physics")

	biotechStudents = sortList(biotechStudents, "Biotech")
	chemistryStudents = sortList(chemistryStudents, "Chemistry")
	engineeringStudents = sortList(engineeringStudents, "Engineering")
	mathStudents = sortList(mathStudents, "Mathematics")
	physicsStudents = sortList(physicsStudents, "Physics")

	printStudents(biotechStudents, "Biotech")
	printStudents(chemistryStudents, "Chemistry")
	printStudents(engineeringStudents, "Engineering")
	printStudents(mathStudents, "Mathematics")
	printStudents(physicsStudents, "Physics")

	writeToFile(biotechStudents, "Biotech")
	writeToFile(chemistryStudents, "Chemistry")
	writeToFile(engineeringStudents, "Engineering")
	writeToFile(mathStudents, "Mathematics")
	writeToFile(physicsStudents, "Physics")
}

func sortList(students []Student, priority string) []Student {
	switch priority {
	case "Biotech":
		sort.Slice(students, func(i, j int) bool {

			if getStudentScore(students[i], "Biotech") != getStudentScore(students[j], "Biotech") {
				return getStudentScore(students[i], "Biotech") > getStudentScore(students[j], "Biotech")
			}
			return students[i].firstName+" "+students[i].lastName < students[j].firstName+" "+students[j].lastName
		})
	case "Chemistry":
		sort.Slice(students, func(i, j int) bool {

			if getStudentScore(students[i], "Chemistry") != getStudentScore(students[j], "Chemistry") {
				return getStudentScore(students[i], "Chemistry") > getStudentScore(students[j], "Chemistry")
			}
			return students[i].firstName+" "+students[i].lastName < students[j].firstName+" "+students[j].lastName
		})
	case "Engineering":
		sort.Slice(students, func(i, j int) bool {

			if getStudentScore(students[i], "Engineering") != getStudentScore(students[j], "Engineering") {
				return getStudentScore(students[i], "Engineering") > getStudentScore(students[j], "Engineering")
			}
			return students[i].firstName+" "+students[i].lastName < students[j].firstName+" "+students[j].lastName
		})
	case "Mathematics":
		sort.Slice(students, func(i, j int) bool {

			if getStudentScore(students[i], "Mathematics") != getStudentScore(students[j], "Mathematics") {
				return getStudentScore(students[i], "Mathematics") > getStudentScore(students[j], "Mathematics")
			}
			return students[i].firstName+" "+students[i].lastName < students[j].firstName+" "+students[j].lastName
		})
	case "Physics":
		sort.Slice(students, func(i, j int) bool {

			if getStudentScore(students[i], "Physics") != getStudentScore(students[j], "Physics") {
				return getStudentScore(students[i], "Physics") > getStudentScore(students[j], "Physics")
			}
			return students[i].firstName+" "+students[i].lastName < students[j].firstName+" "+students[j].lastName
		})

	}
	return students
}

func getStudentScore(student Student, priority string) float32 {
	switch priority {
	case "Biotech":
		if (student.chemistryScore+student.physicsScore)/2 > student.specialScore {
			return (student.chemistryScore + student.physicsScore) / 2
		}
		return student.specialScore
	case "Chemistry":
		if student.chemistryScore > student.specialScore {
			return student.chemistryScore
		}
		return student.specialScore
	case "Engineering":
		if (student.compScore+student.mathScore)/2 > student.specialScore {
			return (student.compScore + student.mathScore) / 2
		}
		return student.specialScore
	case "Mathematics":
		if student.mathScore > student.specialScore {
			return student.mathScore
		}
		return student.specialScore
	case "Physics":
		if (student.mathScore+student.physicsScore)/2 > student.specialScore {
			return (student.mathScore + student.physicsScore) / 2
		}
		return student.specialScore
	}
	return 0
}

func selectStudents1(studentList []Student, studentNumber int, priority string) []Student {
	var selectedStudents []Student

	for i := 0; len(selectedStudents) < studentNumber; i++ {
		if i > len(studentList)-1 {
			break
		}

		student := studentList[i]

		if student.department1 == priority {
			selectedStudents = append(selectedStudents, student)
			studentList[i] = Student{"", "", "", "", "", 0, 0, 0, 0, 0}
		}
	}

	return selectedStudents
}

func selectStudents2(studentList []Student, selectedStudents []Student, studentNumber int, priority string) []Student {
	if len(selectedStudents) < studentNumber {
		for i := 0; len(selectedStudents) < studentNumber; i++ {
			if i > len(studentList)-1 {
				break
			}

			student := studentList[i]

			if student.department2 == priority {
				selectedStudents = append(selectedStudents, student)

				studentList[i] = Student{"", "", "", "", "", 0, 0, 0, 0, 0}
			}
		}
	}

	return selectedStudents
}

func selectStudents3(studentList []Student, selectedStudents []Student, studentNumber int, priority string) []Student {
	if len(selectedStudents) < studentNumber {
		for i := 0; len(selectedStudents) < studentNumber; i++ {
			if i > len(studentList)-1 {
				break
			}

			student := studentList[i]

			if student.department3 == priority {
				selectedStudents = append(selectedStudents, student)

				studentList[i] = Student{"", "", "", "", "", 0, 0, 0, 0, 0}
			}
		}
	}

	return selectedStudents
}

func printStudents(studentList []Student, priority string) {
	fmt.Println(priority)

	for _, student := range studentList {
		var score float32
		switch priority {
		case "Biotech":
			score = getStudentScore(student, "Biotech")
		case "Chemistry":
			score = getStudentScore(student, "Chemistry")
		case "Engineering":
			score = getStudentScore(student, "Engineering")
		case "Mathematics":
			score = getStudentScore(student, "Mathematics")
		case "Physics":
			score = getStudentScore(student, "Physics")
		}
		GPA := fmt.Sprintf("%.2f", score)
		fmt.Println(student.firstName, student.lastName, GPA)
	}

	fmt.Println()
}

func writeToFile(studentList []Student, priority string) {
	switch priority {
	case "Biotech":
		file, err := os.Create("biotech.txt")
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		for _, student := range studentList {
			score := getStudentScore(student, "Biotech")
			lineToWrite := student.firstName + " " + student.lastName + " " + fmt.Sprintf("%.2f", score)
			fmt.Fprintln(file, lineToWrite)
		}
	case "Chemistry":
		file, err := os.Create("chemistry.txt")
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		for _, student := range studentList {
			score := getStudentScore(student, "Chemistry")
			lineToWrite := student.firstName + " " + student.lastName + " " + fmt.Sprintf("%.2f", score)
			fmt.Fprintln(file, lineToWrite)
		}
	case "Engineering":
		file, err := os.Create("engineering.txt")
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		for _, student := range studentList {
			score := getStudentScore(student, "Engineering")
			lineToWrite := student.firstName + " " + student.lastName + " " + fmt.Sprintf("%.2f", score)
			fmt.Fprintln(file, lineToWrite)
		}
	case "Mathematics":
		file, err := os.Create("mathematics.txt")
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		for _, student := range studentList {
			score := getStudentScore(student, "Mathematics")
			lineToWrite := student.firstName + " " + student.lastName + " " + fmt.Sprintf("%.2f", score)
			fmt.Fprintln(file, lineToWrite)
		}
	case "Physics":
		file, err := os.Create("physics.txt")
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		for _, student := range studentList {
			score := getStudentScore(student, "Physics")
			lineToWrite := student.firstName + " " + student.lastName + " " + fmt.Sprintf("%.2f", score)
			fmt.Fprintln(file, lineToWrite)
		}
	}
}
