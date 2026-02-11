package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	Filename string
}

func NewStorage [T any] (filename string) *Storage[T]{
	return &Storage[T]{
		Filename: filename,
	}
}

// converts data to JSON and write it to a file
func (s *Storage[T]) save(data *T) error {
	fileData , err := json.MarshalIndent(data , "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.Filename , fileData , 0644)
}

//reads file converts JSON back and puts it into data
func (s *Storage[T]) load(data *T) error {
	fileData , err := os.ReadFile(s.Filename)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData ,data)
}