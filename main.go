package main

import (
	"fmt"
	"log"
	"CAT_1/payroll"
)
func main() {
	reg := payroll.NewRegistry()
	fmt.Println("=== Employee Payroll Registry Demo ===")
	emp1 := payroll.Employee{
		EmpID:      101,
		Name:       "Alice Smith",
		Department: "Engineering",
		BaseSalary: 85000.00,
		Role:       "Software Engineer",
	}
	if err := reg.AddEmployee(emp1); err != nil {
		log.Fatalf("Failed to add employee: %v", err)
	}
	fmt.Printf("[ADDED] ID: %d | Name: %s | Dept: %s | Salary: $%.2f | Role: %s\n",
		emp1.EmpID, emp1.Name, emp1.Department, emp1.BaseSalary, emp1.Role)
	invalidDeptEmp := payroll.Employee{
		EmpID:      102,
		Name:       "Bob Jones",
		Department: "",
		BaseSalary: 60000.00,
		Role:       "Analyst",
	}
	if err := reg.AddEmployee(invalidDeptEmp); err != nil {
		fmt.Printf("[VALIDATION ERROR EXPECTED] Empty Dept: %v\n", err)
	}
	invalidSalaryEmp := payroll.Employee{
		EmpID:      103,
		Name:       "Charlie Brown",
		Department: "Marketing",
		BaseSalary: -500.00,
		Role:       "Coordinator",
	}
	if err := reg.AddEmployee(invalidSalaryEmp); err != nil {
		fmt.Printf("[VALIDATION ERROR EXPECTED] Salary <= 0: %v\n", err)
	}
	found, err := reg.FindByID(101)
	if err != nil {
		log.Fatalf("FindByID failed: %v", err)
	}
	fmt.Printf("[FOUND] %s working as %s, earning $%.2f\n", found.Name, found.Role, found.BaseSalary)
	fmt.Println("\nPromoting employee 101...")
	if err := reg.PromoteEmployee(101, "Lead Engineer", 20000.00); err != nil {
		log.Fatalf("Promotion failed: %v", err)
	}
	promoted, _ := reg.FindByID(101)
	fmt.Printf("[PROMOTED] %s is now %s with base salary $%.2f\n",
		promoted.Name, promoted.Role, promoted.BaseSalary)
	fmt.Println("\nTerminating employee 101...")
	if err := reg.TerminateEmployee(101); err != nil {
		log.Fatalf("Termination failed: %v", err)
	}
	fmt.Println("[TERMINATED] Employee 101 removed successfully.")
	_, err = reg.FindByID(101)
	if err != nil {
		fmt.Printf("[VERIFIED GONE] Lookup ID 101 returned: %v\n", err)
	}
}