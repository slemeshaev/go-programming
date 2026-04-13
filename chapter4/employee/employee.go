package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

func main() {
	var dilbert Employee
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code

	position := &dilbert.Position
	*position = "Senior" + *position // promoted, for outsourcing to Elbonia

	fmt.Println(position)

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += "(proactive team player)"
}
