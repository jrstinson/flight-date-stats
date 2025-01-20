package main

import (
    "fmt"
    "github.com/jackc/pgx/v5"
    "os"
    "context"
    "time"
)

func main() {

    db, err := pgx.Connect(context.Background(), "postgresql://localhost:5432/postgres?user=postgres&password=password")

    if err != nil {
        fmt.Println("Unable to connect to the database")
        fmt.Println(err)    
        os.Exit(1)
    }

    defer db.Close(context.Background())

    fmt.Println("Options:")
    fmt.Println("1. Import flight hours from an Excel file")
    fmt.Println("2. Get average flight hours for a specific date")

    var option int

    fmt.Scanln(&option)

    if option == 1 {
        initDBTable(db)
    } else if option == 2 {
        fmt.Println("Enter desired flight date in the format YYYY-MM-DD")

        var date time.Time
        var dateinput string

        fmt.Scanln(&dateinput)

        date, err = time.Parse("2006-01-02", dateinput)

        if err != nil {
            fmt.Println("Invalid date format")
            os.Exit(1)
        }

        all_flight_hours, err := GetFlightHours(db)

        if err != nil {
            fmt.Println("Error getting flight hours")
            os.Exit(1)
        }

        var matching_flight_hours []FlightHour

        for _, flight_hour := range all_flight_hours {
            if flight_hour.Date.Month() == date.Month() && flight_hour.Date.Day() == date.Day() {
                matching_flight_hours = append(matching_flight_hours, flight_hour)
            }
        }

        if len(matching_flight_hours) == 0 {
            fmt.Println("No flight hours found for that day")
        } else {
            sum_percent_on_time_gate_departures := 0.0
            sum_percent_on_time_airport_departures := 0.0
            sum_percent_on_time_gate_arrivals := 0.0
            sum_average_gate_departure_delay := 0.0

            for _, flight_hour := range matching_flight_hours {
                sum_percent_on_time_gate_departures += flight_hour.PercentOnTimeGateDepartures
                sum_percent_on_time_airport_departures += flight_hour.PercentOnTimeAirportDepartures
                sum_percent_on_time_gate_arrivals += flight_hour.PercentOnTimeGateArrivals
                sum_average_gate_departure_delay += flight_hour.AverageGateDepartureDelay
            }

            average_percent_on_time_gate_departures := sum_percent_on_time_gate_departures / float64(len(matching_flight_hours))
            average_percent_on_time_airport_departures := sum_percent_on_time_airport_departures / float64(len(matching_flight_hours))
            average_percent_on_time_gate_arrivals := sum_percent_on_time_gate_arrivals / float64(len(matching_flight_hours))
            average_average_gate_departure_delay := sum_average_gate_departure_delay / float64(len(matching_flight_hours))

            fmt.Printf("Average percent on time gate departures: %.2f\n", average_percent_on_time_gate_departures)
            fmt.Printf("Average percent on time airport departures: %.2f\n", average_percent_on_time_airport_departures)
            fmt.Printf("Average percent on time gate arrivals: %.2f\n", average_percent_on_time_gate_arrivals)
            fmt.Printf("Average average gate departure delay (minutes): %.2f\n", average_average_gate_departure_delay)
        }
        os.Exit(0)
    } else {
        fmt.Println("Invalid option")
        os.Exit(1)
    }
}
