package data

import (
	"encoding/json"
	"os"
)

type Response struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	HasBeach    bool      `json:"hasBeach"`
	HasMountain bool      `json:"hasMountain"`
	TempC       []float64 `json:"tempC"`
}

//go:generate mockgen -destination=../mocks/mock_reader.go -package=mocks linkedInLearning/tempService/data DataReader
type DataReader interface {
	ReadData() ([]Response, error)
}

type reader struct {
	path string
}

// NewReader initialises a DataReader
func NewReader() DataReader {
	return &reader{
		path: "./data/cities.json",
	}
}

// ReadData is a helper method to read the file at
// the given path and return a response array.
func (r *reader) ReadData() ([]Response, error) {
	file, err := os.ReadFile(r.path)
	if err != nil {
		return nil, err
	}

	var data []Response
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
