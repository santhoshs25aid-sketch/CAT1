package payroll

import (
	"errors"
	"testing"
)

func TestAddEmployee(t *testing.T) {
	reg := NewRegistry()

	// Valid add
	err := reg.AddEmployee(Employee{
		EmpID:      1,
		Name:       "Test User",
		Department: "IT",
		BaseSalary: 50000,
		Role:       "Dev",
	})
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	// Duplicate ID
	err = reg.AddEmployee(Employee{
		EmpID:      1,
		Name:       "Duplicate",
		Department: "HR",
		BaseSalary: 60000,
		Role:       "Dev",
	})
	if !errors.Is(err, ErrEmployeeExists) {
		t.Errorf("expected ErrEmployeeExists, got %v", err)
	}

	// Empty department
	err = reg.AddEmployee(Employee{
		EmpID:      2,
		Name:       "No Dept",
		Department: "",
		BaseSalary: 50000,
		Role:       "Dev",
	})
	if !errors.Is(err, ErrInvalidDepartment) {
		t.Errorf("expected ErrInvalidDepartment, got %v", err)
	}

	// Non-positive salary
	err = reg.AddEmployee(Employee{
		EmpID:      3,
		Name:       "Zero Salary",
		Department: "Sales",
		BaseSalary: 0,
		Role:       "Rep",
	})
	if !errors.Is(err, ErrInvalidSalary) {
		t.Errorf("expected ErrInvalidSalary, got %v", err)
	}
}

func TestPromoteAndTerminate(t *testing.T) {
	reg := NewRegistry()
	_ = reg.AddEmployee(Employee{
		EmpID:      10,
		Name:       "Jane",
		Department: "Finance",
		BaseSalary: 70000,
		Role:       "Accountant",
	})

	// Promote
	err := reg.PromoteEmployee(10, "Senior Accountant", 10000)
	if err != nil {
		t.Fatalf("expected promote success, got %v", err)
	}

	emp, _ := reg.FindByID(10)
	if emp.Role != "Senior Accountant" || emp.BaseSalary != 80000 {
		t.Errorf("promotion values mismatch: role=%s, salary=%.2f", emp.Role, emp.BaseSalary)
	}

	// Terminate
	err = reg.TerminateEmployee(10)
	if err != nil {
		t.Fatalf("expected terminate success, got %v", err)
	}

	// Ensure deleted
	_, err = reg.FindByID(10)
	if !errors.Is(err, ErrEmployeeNotFound) {
		t.Errorf("expected ErrEmployeeNotFound, got %v", err)
	}
}