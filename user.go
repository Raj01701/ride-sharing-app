package main

import (
	"fmt"
	"strings"
)

type Vehicle struct {
	Name   string
	Number string
}

type User struct {
	ID           int
	Name         string
	Gender       string
	Age          int
	Vehicles     []Vehicle
	Roles        []string
	RidesOffered int
	RidesTaken   int
}

var users []User
var nextUserID int

func addUser(name, gender string, age int, roles []string) {
	nextUserID++
	user := User{
		ID:      nextUserID,
		Name:    name,
		Gender:  gender,
		Age:     age,
		Roles:   roles,
	}
	users = append(users, user)
}

func addVehicle(userID int, vehicleName, vehicleNumber string) {
	for i, user := range users {
		if user.ID == userID {
			users[i].Vehicles = append(users[i].Vehicles, Vehicle{Name: vehicleName, Number: vehicleNumber})
			return
		}
	}
}

func listUsers() {
	fmt.Println("User ID\tName\tGender\tAge\tRoles")
	for _, user := range users {
		fmt.Printf("%d\t%s\t%s\t%d\t%s\n", user.ID, user.Name, user.Gender, user.Age, strings.Join(user.Roles, ","))
	}
}

func incrementRidesTaken(userID int) {
	for i, user := range users {
		if user.ID == userID {
			users[i].RidesTaken++
			return
		}
	}
}

func incrementRidesOffered(userID int) {
	for i, user := range users {
		if user.ID == userID {
			users[i].RidesOffered++
			return
		}
	}
}

func printRideStats() {
	for _, user := range users {
		fmt.Printf("User %s: Rides Offered = %d, Rides Taken = %d\n", user.Name, user.RidesOffered, user.RidesTaken)
	}
}
