# Ride-Sharing Application

## Overview

This is a simple command-line ride-sharing application that allows users to offer and take shared rides. The application includes functionalities for user management, vehicle management, ride offering, ride selection, ride management, and statistics. Additionally, it provides an option to find multiple rides if a direct route is not available.

## Features

- **User Management**: Add users dynamically.
- **Vehicle Management**: Add vehicles for users dynamically.
- **Ride Offering**: Users can offer rides with specific details.
- **Ride Selection**: Users can select from multiple offered rides using a selection strategy.
- **Ride Management**: End rides and ensure users can only offer rides with vehicles that are not already in use.
- **Statistics**: Print ride stats showing total rides offered and taken by all users.
- **Multiple Rides Output**: If a direct route is not available, output multiple rides that can facilitate the journey.

## Setting Up the Project

1. **Install Go**: Ensure you have Go installed on your machine. You can download it from [here](https://golang.org/dl/).
2. **Clone the Repository**: Clone this repository to your local machine using:
   ```sh
   git clone https://github.com/lekhrajsaa/ride-sharing-app
   ```
3. **Navigate to the Project Directory**:
   ```sh
   cd ride-sharing-app
   ```
4. **Ensure All Files Are in Place**: Ensure the following files are in the directory:
   - `main.go`
   - `user.go`
   - `ride.go`

## Running the Application

To run the application, use the following command:
```sh
go run *.go
```

### Interaction

Once you run the application, you will be prompted with a menu to select various options. Follow the on-screen prompts to interact with the application.

### Menu Options

1. **Add User**:
   - Prompts for user details: Name, Gender, Age, Roles (comma-separated for multiple roles).
   - Example input: `John, M, 30, Driver, Passenger`

2. **Add Vehicle**:
   - Prompts for user ID and vehicle details: UserID, VehicleName, VehicleNumber.
   - Example input: `1, Swift, KA-01-12345`

3. **Offer Ride**:
   - Prompts for ride details: DriverID, Source, Destination, Seats, VehicleNumber.
   - Example input: `1, Bangalore, Mysore, 3, KA-01-12345`

4. **Select Ride**:
   - Prompts for ride selection details: PassengerID, Source, Destination, Seats, SelectionStrategy.
   - Example input: `2, Bangalore, Mysore, 1, most_vacant`

5. **End Ride**:
   - Prompts for ride ID to end the ride.
   - Example input: `1`

6. **Print Ride Stats**:
   - Prints the total rides offered and taken by all users.

7. **List Users**:
   - Lists all users with their IDs, names, genders, ages, and roles.

8. **Find Multiple Rides**:
   - Prompts for source and destination to find multiple rides.
   - Example input: `Bangalore, Mysore`

9. **Exit**:
   - Exits the application.

### Example Usage

#### Adding a User
```
Select an option:
1. Add User
...
Enter user details (Name, Gender, Age, Roles (comma-separated for multiple roles)): 
John, M, 30, Driver, Passenger
```

#### Adding a Vehicle
```
Select an option:
2. Add Vehicle
...
Enter user ID and vehicle details (UserID, VehicleName, VehicleNumber): 
1, Swift, KA-01-12345
```

#### Offering a Ride
```
Select an option:
3. Offer Ride
...
Enter ride details (DriverID, Source, Destination, Seats, VehicleNumber): 
1, Bangalore, Mysore, 3, KA-01-12345
```

#### Selecting a Ride
```
Select an option:
4. Select Ride
...
Enter ride selection details (PassengerID, Source, Destination, Seats, SelectionStrategy): 
2, Bangalore, Mysore, 1, most_vacant
```

#### Ending a Ride
```
Select an option:
5. End Ride
...
Enter ride ID to end: 
1
```

#### Printing Ride Stats
```
Select an option:
6. Print Ride Stats
...
```

#### Listing Users
```
Select an option:
7. List Users
...
```

#### Finding Multiple Rides
```
Select an option:
8. Find Multiple Rides
...
Enter source and destination to find multiple rides (Source, Destination): 
Bangalore, Mysore
```

## Conclusion

This ride-sharing application provides a simple and dynamic way to manage users, vehicles, and rides. It includes all the required functionalities and allows for easy interaction through a command-line interface.

If you have any questions or encounter any issues, please feel free to reach out.
