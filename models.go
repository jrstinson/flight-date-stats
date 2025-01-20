package main

import (
	"github.com/jackc/pgx/v5"
    "context"
	"time"
)

type FlightHour struct {
    ID int `json:"id"`
    Facility string `json:"facility"`
    Date time.Time `json:"date"`
    Hour int `json:"hour"`
    GMTHour int `json:"gmt_hour"`
    ScheduledDepartures int `json:"scheduled_departures"`
    ScheduledArrivals int `json:"scheduled_arrivals"`
    PercentOnTimeGateDepartures float64 `json:"percent_on_time_gate_departures"`
    PercentOnTimeAirportDepartures float64 `json:"percent_on_time_airport_departures"`
    PercentOnTimeGateArrivals float64 `json:"percent_on_time_gate_arrivals"`
    AverageGateDepartureDelay float64 `json:"average_gate_departure_delay"`
    AverageTaxiOutTime float64 `json:"average_taxi_out_time"`
    AverageTaxiOutDelay float64 `json:"average_taxi_out_delay"`
    AverageAirportDepartureDelay float64 `json:"average_airport_departure_delay"`
    AverageAirborneDelay float64 `json:"average_airborne_delay"`
    AverageTaxiInDelay float64 `json:"average_taxi_in_delay"`
    AverageBlockDelay float64 `json:"average_block_delay"`
    AverageGateArrivalDelay float64 `json:"average_gate_arrival_delay"`
}

func initDBTable(db *pgx.Conn) error {
    _, err := db.Exec(context.Background(),("CREATE TABLE IF NOT EXISTS flight_hours " +
        "(ID SERIAL PRIMARY KEY," + 
        "Facility VARCHAR(3)," + 
        "Date DATE," +
        "Hour INT," +
        "GMTHour INT," + 
        "ScheduledDepartures INT," +
        "ScheduledArrivals INT," + 
        "PercentOnTimeGateDepartures FLOAT," + 
        "PercentOnTimeAirportDepartures FLOAT," +
        "PercentOnTimeGateArrivals FLOAT, " +
        "AverageGateDepartureDelay FLOAT," +
        "AverageTaxiOutTime FLOAT," +
        "AverageTaxiOutDelay FLOAT," +
        "AverageAirportDepartureDelay FLOAT," +
        "AverageAirborneDelay FLOAT," +
        "AverageTaxiInDelay FLOAT," +
        "AverageBlockDelay FLOAT," +
        "AverageGateArrivalDelay FLOAT)"))

    return err
}

func GetFlightHours(db *pgx.Conn) ([]FlightHour, error) {
    rows, err := db.Query(context.Background(), "select * from flight_hours")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var flightHours []FlightHour
    for rows.Next() {
        var flightHour FlightHour
        err = rows.Scan(
            &flightHour.ID,
            &flightHour.Facility,
            &flightHour.Date,
            &flightHour.Hour,
            &flightHour.GMTHour,
            &flightHour.ScheduledDepartures,
            &flightHour.ScheduledArrivals,
            &flightHour.PercentOnTimeGateDepartures,
            &flightHour.PercentOnTimeAirportDepartures,
            &flightHour.PercentOnTimeGateArrivals,
            &flightHour.AverageGateDepartureDelay,
            &flightHour.AverageTaxiOutTime,
            &flightHour.AverageTaxiOutDelay,
            &flightHour.AverageAirportDepartureDelay,
            &flightHour.AverageAirborneDelay,
            &flightHour.AverageTaxiInDelay,
            &flightHour.AverageBlockDelay,
            &flightHour.AverageGateArrivalDelay,
        )
        if err != nil {
            return nil, err
        }
        flightHours = append(flightHours, flightHour)
    }

    return flightHours, nil
}
