package filejson

import (
	"testing"
)

func TestReadFromJsonContent_Success(t *testing.T) {
	//PLACEHOLDER: Replace this with real tests
	input := `[{
    "id": "a8cfcb76-7f24-4420-a5ba-d46dd77bdffd",
    "name": "Entry 1",
    "quantity": 2
  	}]
	`
	expectedEntry := Entry{
		ID:       "a8cfcb76-7f24-4420-a5ba-d46dd77bdffd",
		Name:     "Entry 1",
		Quantity: 2,
	}
	bytes := []byte(input)
	entries, err := ReadFromJsonContent(bytes)
	if err != nil {
		t.Fatalf("error should be nil but got %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry but got %d", len(entries))
	}

	if expectedEntry != entries[0] {
		t.Fatalf("expected entry is %v but got %v", expectedEntry, entries[0])
	}
}

func TestReadFromJsonContent_Failure(t *testing.T) {
	input := `[{
    "dish": "Meatball",
  	}]
	`
	bytes := []byte(input)
	_, err := ReadFromJsonContent(bytes)
	if err == nil {
		t.Fatalf("error should not be nil but got nil")
	}
}
