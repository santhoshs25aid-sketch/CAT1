package payroll
import (
	"errors"
	"fmt"
	"sync"
)
var (
	ErrEmployeeNotFound  = errors.New("employee not found")
	ErrEmployeeExists    = errors.New("employee with this ID already exists")
	ErrInvalidDepartment = errors.New("department cannot be empty")
	ErrInvalidSalary     = errors.New("base salary must be greater than zero")
	ErrInvalidSalaryBump = errors.New("salary bump must be greater than or equal to zero")
	ErrInvalidRole       = errors.New("new role cannot be empty")
)
type Employee struct {
	EmpID      int
	Name       string
	Department string
	BaseSalary float64
	Role       string
}
type Registry struct {
	mu        sync.RWMutex
	employees map[int]Employee
}

func NewRegistry() *Registry {
	return &Registry{
		employees: make(map[int]Employee),
	}
}
func (r *Registry) AddEmployee(e Employee) error {
	if e.Department == "" {
		return ErrInvalidDepartment
	}
	if e.BaseSalary <= 0 {
		return ErrInvalidSalary
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.employees[e.EmpID]; exists {
		return fmt.Errorf("%w: ID %d", ErrEmployeeExists, e.EmpID)
	}

	r.employees[e.EmpID] = e
	return nil
}
func (r *Registry) FindByID(id int) (*Employee, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	emp, exists := r.employees[id]
	if !exists {
		return nil, fmt.Errorf("%w: ID %d", ErrEmployeeNotFound, id)
	}

	empCopy := emp
	return &empCopy, nil
}

// PromoteEmployee updates an employee's role and adds a salary bump.
func (r *Registry) PromoteEmployee(id int, newRole string, salaryBump float64) error {
	if newRole == "" {
		return ErrInvalidRole
	}
	if salaryBump < 0 {
		return ErrInvalidSalaryBump
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	emp, exists := r.employees[id]
	if !exists {
		return fmt.Errorf("%w: ID %d", ErrEmployeeNotFound, id)
	}

	emp.Role = newRole
	emp.BaseSalary += salaryBump
	r.employees[id] = emp

	return nil
}

func (r *Registry) TerminateEmployee(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.employees[id]; !exists {
		return fmt.Errorf("%w: ID %d", ErrEmployeeNotFound, id)
	}

	delete(r.employees, id)
	return nil
}

func (r *Registry) ListAll() []Employee {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]Employee, 0, len(r.employees))
	for _, e := range r.employees {
		list = append(list, e)
	}
	return list
}