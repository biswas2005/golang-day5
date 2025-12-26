package practice

import "fmt"

type Student struct {
	Name  string
	SID   int
	mark1 int
	mark2 int
	mark3 int
	total int
}

func total(s *Student) {
	s.total = s.mark1 + s.mark2 + s.mark3

}

func isPass(s Student) bool {
	if s.mark1 >= 40 && s.mark2 >= 40 && s.mark3 >= 40 {
		return true
	}
	return false
}

func Grade(Total int) string {
	if Total >= 250 {
		return "A"
	} else if Total >= 200 {
		return "B"
	} else if Total >= 120 {
		return "C"
	}
	return "Fail"
}

func Final() {
	var s Student

	fmt.Print("Enter student name: ")
	fmt.Scan(&s.Name)

	fmt.Print("Enter student ID: ")
	fmt.Scan(&s.SID)

	fmt.Print("Enter marks<Subject 1>:")
	fmt.Scan(&s.mark1)
	fmt.Print("Enter marks<Subject 2>:")
	fmt.Scan(&s.mark2)
	fmt.Print("Enter marks<Subject 3>:")
	fmt.Scan(&s.mark3)

	total(&s)

	fmt.Println("<Student Results>")
	fmt.Println("Name:", s.Name)
	fmt.Println("Student ID:", s.SID)
	fmt.Println("Total Marks:", s.total)
	if isPass(s) {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
	fmt.Println("<Subject wise marks>")
	fmt.Println("Subject1:", s.mark1)
	fmt.Println("Subject2:", s.mark2)
	fmt.Println("Subject3:", s.mark3)
}
