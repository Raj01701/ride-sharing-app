package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Select an option:")
		fmt.Println("1. Add User")
		fmt.Println("2. Add Vehicle")
		fmt.Println("3. Offer Ride")
		fmt.Println("4. Select Ride")
		fmt.Println("5. End Ride")
		fmt.Println("6. Print Ride Stats")
		fmt.Println("7. List Users")
		fmt.Println("8. Find Multiple Rides")
		fmt.Println("9. Exit")

		scanner.Scan()
		option, _ := strconv.Atoi(scanner.Text())

		switch option {
		case 1:
			fmt.Println("Enter user details (Name, Gender, Age, Roles (comma-separated for multiple roles)): ")
			scanner.Scan()
			userDetails := strings.Split(scanner.Text(), ",")
			name := strings.TrimSpace(userDetails[0])
			gender := strings.TrimSpace(userDetails[1])
			age, _ := strconv.Atoi(strings.TrimSpace(userDetails[2]))
			roles := strings.Split(strings.TrimSpace(userDetails[3]), ",")
			for i := range roles {
				roles[i] = strings.TrimSpace(roles[i])
			}
			addUser(name, gender, age, roles)

		case 2:
			fmt.Println("Enter user ID and vehicle details (UserID, VehicleName, VehicleNumber): ")
			scanner.Scan()
			vehicleDetails := strings.Split(scanner.Text(), ",")
			userID, _ := strconv.Atoi(strings.TrimSpace(vehicleDetails[0]))
			vehicleName := strings.TrimSpace(vehicleDetails[1])
			vehicleNumber := strings.TrimSpace(vehicleDetails[2])
			addVehicle(userID, vehicleName, vehicleNumber)

		case 3:
			fmt.Println("Enter ride details (DriverID, Source, Destination, Seats, VehicleNumber): ")
			scanner.Scan()
			rideDetails := strings.Split(scanner.Text(), ",")
			driverID, _ := strconv.Atoi(strings.TrimSpace(rideDetails[0]))
			source := strings.TrimSpace(rideDetails[1])
			destination := strings.TrimSpace(rideDetails[2])
			seats, _ := strconv.Atoi(strings.TrimSpace(rideDetails[3]))
			vehicleNumber := strings.TrimSpace(rideDetails[4])
			offerRide(driverID, source, destination, seats, vehicleNumber)

		case 4:
			fmt.Println("Enter ride selection details (PassengerID, Source, Destination, Seats, SelectionStrategy): ")
			scanner.Scan()
			selectionDetails := strings.Split(scanner.Text(), ",")
			passengerID, _ := strconv.Atoi(strings.TrimSpace(selectionDetails[0]))
			source := strings.TrimSpace(selectionDetails[1])
			destination := strings.TrimSpace(selectionDetails[2])
			seats, _ := strconv.Atoi(strings.TrimSpace(selectionDetails[3]))
			selectionStrategy := strings.TrimSpace(selectionDetails[4])
			selectRide(passengerID, source, destination, seats, selectionStrategy)

		case 5:
			fmt.Println("Enter ride ID to end: ")
			scanner.Scan()
			rideID, _ := strconv.Atoi(scanner.Text())
			endRide(rideID)

		case 6:
			printRideStats()

		case 7:
			listUsers()

		case 8:
			fmt.Println("Enter source and destination to find multiple rides (Source, Destination): ")
			scanner.Scan()
			routeDetails := strings.Split(scanner.Text(), ",")
			source := strings.TrimSpace(routeDetails[0])
			destination := strings.TrimSpace(routeDetails[1])
			routes := findMultipleRides(source, destination)
			printMultipleRides(routes)

		case 9:
			return

		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
