package weather

import (
	"encoding/json"
	"math"
	"time"

	"github.com/gsainz/airborne/utils/random"
)

var summaries = []string{"Freezing", "Bracing", "Chilly", "Cool", "Mild", "Warm", "Balmy", "Hot", "Sweltering", "Scorching"}

type forecast struct {
	Date       time.Time `json:"date"`
	Celsius    float64   `json:"temperatureC"`
	Fahrenheit float64   `json:"temperatureF"`
	Summary    string    `json:"summary"`
}

func (w *forecast) ToFahrenheit() {
	w.Fahrenheit = math.Round(32 + (w.Celsius / 0.5556))
}

type Predictions []forecast

func Predict(days int) []byte {
	temps := make(Predictions, 0, days-1)
	for i := 1; i <= days; i++ {
		temp := forecast{}
		temp.Date = time.Now().Add(time.Duration(i) * 24 * time.Hour)
		temp.Celsius = math.Round(float64(random.Int(-50, 120)))
		temp.Summary = summaries[random.Int(0, int64(len(summaries)-1))]
		temp.ToFahrenheit()
		temps = append(temps, temp)
	}
	bytes, err := json.Marshal(temps)
	if err != nil {
		return []byte("")
	}
	return bytes
}
