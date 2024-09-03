package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

var ErrCannotUnmarshal = "json: cannot unmarshal array into Go value of type main.NumbersRequest"

type NumbersRequest struct {
	Numbers []int64 `json:"numbers"`
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}

// Define a handler function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var buf bytes.Buffer
	_, err := io.Copy(&buf, r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Ensure the request body is closed
	defer r.Body.Close()

	// Convert buffer content to a slice of int32
	intSlice, err := bufferToInt32Slice(&buf)
	if err != nil {
		http.Error(w, "Invalid integer values in request body", http.StatusBadRequest)
		return
	}

	result := sumNumbers(intSlice)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	slog.Info("msg", result)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}

func sumNumbers(numbers []int32) int32 {
	var sum int32 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Function to convert buffer content to a slice of int32
func bufferToInt32Slice(buf *bytes.Buffer) ([]int32, error) {
	var intSlice []int32

	// Decode JSON array from buffer content
	decoder := json.NewDecoder(buf)
	if err := decoder.Decode(&intSlice); err != nil {
		return nil, err
	}

	return intSlice, nil
}
