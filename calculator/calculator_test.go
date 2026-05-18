package calculator

import (
	"testing"
)

func TestCalculate_Addition(t *testing.T) {
	svc := NewService()
	resp := svc.Calculate(CalculateRequest{A: 10, B: 2, Operation: "+"})
	if resp.Result != 12 {
		t.Errorf("expected 12, got %v", resp.Result)
	}
}

func TestCalculate_Subtraction(t *testing.T) {
	svc := NewService()
	resp := svc.Calculate(CalculateRequest{A: 10, B: 2, Operation: "-"})
	if resp.Result != 8 {
		t.Errorf("expected 8, got %v", resp.Result)
	}
}

func TestCalculate_Multiplication(t *testing.T) {
	svc := NewService()
	resp := svc.Calculate(CalculateRequest{A: 10, B: 2, Operation: "*"})
	if resp.Result != 20 {
		t.Errorf("expected 20, got %v", resp.Result)
	}
}

func TestCalculate_Division(t *testing.T) {
	svc := NewService()
	resp := svc.Calculate(CalculateRequest{A: 10, B: 2, Operation: "/"})
	if resp.Result != 5 {
		t.Errorf("expected 5, got %v", resp.Result)
	}
}

// BUG: Division durch Null gibt +Inf zurück statt Fehler
func TestCalculate_DivisionByZero(t *testing.T) {
	svc := NewService()
	resp := svc.Calculate(CalculateRequest{A: 10, B: 0, Operation: "/"})
	t.Logf("Division by zero result: %v", resp.Result) // gibt +Inf
}

// BUG: Unbekannter Operator gibt 0 zurück statt Fehler
func TestCalculate_UnknownOperator(t *testing.T) {
	svc := NewService()
	resp := svc.Calculate(CalculateRequest{A: 10, B: 2, Operation: "^"})
	t.Logf("Unknown operator result: %v", resp.Result) // gibt 0
}

func TestGetHistory(t *testing.T) {
	svc := NewService()
	svc.Calculate(CalculateRequest{A: 1, B: 2, Operation: "+"})
	svc.Calculate(CalculateRequest{A: 3, B: 4, Operation: "*"})
	history := svc.GetHistory()
	if len(history) != 2 {
		t.Errorf("expected 2 history entries, got %d", len(history))
	}
}
