package main

import (
	"fmt"
	"sort"
)

type Ride struct {
	ID           int
	DriverID     int
	Source       string
	Destination  string
	Seats        int
	VacantSeats  int
	Vehicle      string
}

var rides []Ride
var nextRideID int

func offerRide(driverID int, source, destination string, seats int, vehicleNumber string) {
	for _, ride := range rides {
		if ride.DriverID == driverID && ride.Vehicle == vehicleNumber && ride.VacantSeats > 0 {
			fmt.Printf("User %d already has an active ride with vehicle %s.\n", driverID, vehicleNumber)
			return
		}
	}

	nextRideID++
	ride := Ride{
		ID:          nextRideID,
		DriverID:    driverID,
		Source:      source,
		Destination: destination,
		Seats:       seats,
		VacantSeats: seats,
		Vehicle:     vehicleNumber,
	}
	rides = append(rides, ride)
	incrementRidesOffered(driverID)
}

func selectRide(passengerID int, source, destination string, seats int, selectionStrategy string) bool {
	availableRides := []Ride{}
	for _, ride := range rides {
		if ride.Source == source && ride.Destination == destination && ride.VacantSeats >= seats {
			availableRides = append(availableRides, ride)
		}
	}

	if len(availableRides) == 0 {
		fmt.Println("No available rides found.")
		return false
	}

	switch selectionStrategy {
	case "most_vacant":
		sort.Slice(availableRides, func(i, j int) bool {
			return availableRides[i].VacantSeats > availableRides[j].VacantSeats
		})
	case "preferred_vehicle":
		sort.Slice(availableRides, func(i, j int) bool {
			return availableRides[i].Vehicle < availableRides[j].Vehicle
		})
	}

	selectedRide := availableRides[0]
	for i, ride := range rides {
		if ride.ID == selectedRide.ID {
			rides[i].VacantSeats -= seats
			incrementRidesTaken(passengerID)
			fmt.Printf("Ride selected: %v\n", selectedRide)
			return true
		}
	}

	return false
}

func endRide(rideID int) {
	for i, ride := range rides {
		if ride.ID == rideID {
			rides = append(rides[:i], rides[i+1:]...)
			fmt.Printf("Ride ended: ID %d\n", rideID)
			return
		}
	}
	fmt.Printf("Ride ID %d not found.\n", rideID)
}

func findMultipleRides(source, destination string) [][]Ride {
	var results [][]Ride
	visited := make(map[int]bool)
	var dfs func(current string, path []Ride)
	dfs = func(current string, path []Ride) {
		if current == destination {
			temp := make([]Ride, len(path))
			copy(temp, path)
			results = append(results, temp)
			return
		}
		for _, ride := range rides {
			if ride.Source == current && !visited[ride.ID] {
				visited[ride.ID] = true
				dfs(ride.Destination, append(path, ride))
				visited[ride.ID] = false
			}
		}
	}
	dfs(source, nil)
	return results
}

func printMultipleRides(rides [][]Ride) {
	for _, ridePath := range rides {
		for _, ride := range ridePath {
			fmt.Printf("Ride ID: %d, Driver ID: %d, Source: %s, Destination: %s, Seats: %d\n",
				ride.ID, ride.DriverID, ride.Source, ride.Destination, ride.Seats)
		}
		fmt.Println("-----")
	}
}
