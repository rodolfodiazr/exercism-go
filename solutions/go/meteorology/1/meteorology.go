package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (t TemperatureUnit) String() string {
    units := []string{"°C", "°F"}
    return units[t]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
    return fmt.Sprintf("%v %v", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (u SpeedUnit) String() string {
    units := []string{"km/h", "mph"}
    return units[u]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
    return fmt.Sprintf("%v %v", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (d MeteorologyData) String() string {
    return fmt.Sprintf(
        "%s: %s, Wind %s at %s, %d%s Humidity", 
        d.location, 
        d.temperature, 
        d.windDirection, 
        d.windSpeed,
        d.humidity,
        "%",
    )
}