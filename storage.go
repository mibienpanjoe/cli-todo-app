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

func (s *Storage[T]) save(data *T) error {
	fileData , err := json.MarshalIndent(data , "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.Filename , fileData , 0644)
}