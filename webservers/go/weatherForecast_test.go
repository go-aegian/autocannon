package weather_test

import (
	"testing"

	"github.com/sainzg/autocannon/webservers/go"
)

func TestForecast(t *testing.T) {
	bytes := weather.Predict(5)
	t.Log(string(bytes))
	if bytes == nil {
		t.FailNow()
	}
}
