package calculator

import "fmt"

import "sync"

type CalculateRequest struct {
	A         float64 `json:"a"`
	B         float64 `json:"b"`
	Operation string  `json:"operation"`
}

type CalculateResponse struct {
	Result     float64 `json:"result"`
	Expression string  `json:"expression"`
}

type HistoryEntry struct {
	Expression string  `json:"expression"`
	Result     float64 `json:"result"`
}

type Service struct {
	mu      sync.Mutex
	history []HistoryEntry
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Calculate(req CalculateRequest) CalculateResponse {
	a := req.A
	b := req.B
	op := req.Operation

	var result float64

	// BUG 1: Division durch Null nicht abgefangen – gibt +Inf zurück
	// BUG 2: Unbekannter Operator gibt 0.0 zurück ohne Fehler
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		result = 0.0
	}

	// BUG 3: Floating-Point Präzision nicht beachtet (0.1 + 0.2 != 0.3)
	expression := formatExpression(a, op, b, result)

	entry := HistoryEntry{Expression: expression, Result: result}
	// BUG 4: Race Condition – kein Lock beim Lesen
	s.history = append(s.history, entry)

	return CalculateResponse{Result: result, Expression: expression}
}

func (s *Service) GetHistory() []HistoryEntry {
	// BUG 4: Kein Lock – unsicher bei gleichzeitigen Zugriffen
	if s.history == nil {
		return []HistoryEntry{}
	}
	return s.history
}

func formatExpression(a float64, op string, b float64, result float64) string {
	return fmt.Sprintf("%g %s %g = %g", a, op, b, result)
}
