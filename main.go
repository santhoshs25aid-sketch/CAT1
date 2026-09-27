package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Employee represents an employee record in the registry
type Employee struct {
	EmpID      int
	Name       string
	Department string
	BaseSalary float64
	Role       string
}

// PayrollRegistry manages the collection of employees
type PayrollRegistry struct {
	employees map[int]Employee
}

// NewPayrollRegistry initializes a new registry
func NewPayrollRegistry() *PayrollRegistry {
	return &PayrollRegistry{
		employees: make(map[int]Employee),
	}
}

// AddEmployee validates and adds an employee
func (p *PayrollRegistry) AddEmployee(e Employee) error {
	if _, exists := p.employees[e.EmpID]; exists {
		return errors.New("employee ID already exists")
	}
	if strings.TrimSpace(e.Department) == "" {
		return errors.New("department cannot be empty")
	}
	if e.BaseSalary <= 0 {
		return errors.New("base salary must be greater than zero")
	}
	if strings.TrimSpace(e.Name) == "" {
		return errors.New("employee name cannot be empty")
	}

	p.employees[e.EmpID] = e
	return nil
}

// FindByID retrieves an employee record
func (p *PayrollRegistry) FindByID(id int) (*Employee, error) {
	emp, exists := p.employees[id]
	if !exists {
		return nil, errors.New("employee not found")
	}
	return &emp, nil
}

// PromoteEmployee updates an employee's role and bumps their salary
func (p *PayrollRegistry) PromoteEmployee(id int, newRole string, salaryBump float64) error {
	emp, exists := p.employees[id]
	if !exists {
		return errors.New("employee not found")
	}
	if strings.TrimSpace(newRole) == "" {
		return errors.New("new role cannot be empty")
	}
	if salaryBump < 0 {
		return errors.New("salary bump cannot be negative")
	}

	emp.Role = newRole
	emp.BaseSalary += salaryBump
	p.employees[id] = emp
	return nil
}

// TerminateEmployee deletes an employee from the directory
func (p *PayrollRegistry) TerminateEmployee(id int) error {
	if _, exists := p.employees[id]; !exists {
		return errors.New("employee not found")
	}
	delete(p.employees, id)
	return nil
}

func main() {
	registry := NewPayrollRegistry()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n==============================")
		fmt.Println("   EMPLOYEE PAYROLL REGISTRY  ")
		fmt.Println("==============================")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Find Employee By ID")
		fmt.Println("3. Promote Employee")
		fmt.Println("4. Terminate Employee")
		fmt.Println("5. List All Employees")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option (1-6): ")

		if !scanner.Scan() {
			break
		}
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Enter Emp ID (int): ")
			scanner.Scan()
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Error: Invalid ID format.")
				continue
			}

			fmt.Print("Enter Name: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Enter Department: ")
			scanner.Scan()
			dept := scanner.Text()

			fmt.Print("Enter Base Salary: ")
			scanner.Scan()
			salary, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
			if err != nil {
				fmt.Println("Error: Invalid salary format.")
				continue
			}

			fmt.Print("Enter Role/Title: ")
			scanner.Scan()
			role := scanner.Text()

			emp := Employee{
				EmpID:      id,
				Name:       name,
				Department: dept,
				BaseSalary: salary,
				Role:       role,
			}

			if err := registry.AddEmployee(emp); err != nil {
				fmt.Printf("Failed to add employee: %v\n", err)
			} else {
				fmt.Println("Success: Employee added successfully!")
			}

		case "2":
			fmt.Print("Enter Emp ID to search: ")
			scanner.Scan()
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Error: Invalid ID format.")
				continue
			}

			emp, err := registry.FindByID(id)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("\n--- Employee Found ---\n")
				fmt.Printf("ID:         %d\n", emp.EmpID)
				fmt.Printf("Name:       %s\n", emp.Name)
				fmt.Printf("Department: %s\n", emp.Department)
				fmt.Printf("Role:       %s\n", emp.Role)
				fmt.Printf("Base Salary: $%.2f\n", emp.BaseSalary)
			}

		case "3":
			fmt.Print("Enter Emp ID to promote: ")
			scanner.Scan()
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Error: Invalid ID format.")
				continue
			}

			fmt.Print("Enter New Role: ")
			scanner.Scan()
			newRole := scanner.Text()

			fmt.Print("Enter Salary Bump Amount: ")
			scanner.Scan()
			bump, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
			if err != nil {
				fmt.Println("Error: Invalid amount format.")
				continue
			}

			if err := registry.PromoteEmployee(id, newRole, bump); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Success: Employee promoted and salary updated!")
			}

		case "4":
			fmt.Print("Enter Emp ID to terminate: ")
			scanner.Scan()
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Error: Invalid ID format.")
				continue
			}

			if err := registry.TerminateEmployee(id); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Success: Employee removed from registry.")
			}

		case "5":
			if len(registry.employees) == 0 {
				fmt.Println("\nRegistry is empty.")
				continue
			}
			fmt.Println("\n--- All Employees ---")
			for _, emp := range registry.employees {
				fmt.Printf("ID: %d | Name: %s | Dept: %s | Role: %s | Salary: $%.2f\n",
					emp.EmpID, emp.Name, emp.Department, emp.Role, emp.BaseSalary)
			}

		case "6":
			fmt.Println("Exiting application. Goodbye!")
			return

		default:
			fmt.Println("Invalid option. Please choose between 1 and 6.")
		}
	}
}
