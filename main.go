package main

import (
	"fmt"
	"log"
)

type struct Menu {
	ID int `jsong:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Available   bool    `json:"available"`
}


type Order struct {
	ID         string      `json:"id"`
	CustomerID string      `json:"customer_id"`
	Items      []MenuItem  `json:"items"`
	Total      float64     `json:"total"`
	Status     string      `json:"status"` // e.g., pending, completed
	Timestamp  string      `json:"timestamp"`
}

type Customer struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Contact    string `json:"contact"`
}

type Reservation struct {
	ID           string `json:"id"`
	CustomerID   string `json:"customer_id"`
	Date         string `json:"date"`
	Time         string `json:"time"`
	NumPeople    int    `json:"num_people"`
}

func main() {
	fmt.Println("Building of Restful Api for Cafe Manager Application")
}