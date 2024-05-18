package weather

import "errors"

type Weather struct {
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

func NewWeather(celsius float64) (*Weather, error) {
	weather := &Weather{
		Celsius: celsius,
	}
	err := weather.isValid()
	if err != nil {
		return nil, err
	}

	weather.ConvertFahrenheit()
	weather.ConvertKelvin()

	return weather, nil
}

func (t *Weather) isValid() error {
	if t.Celsius <= -273.15 {
		return errors.New("invalid celsius")
	}
	return nil
}

func (t *Weather) ConvertFahrenheit() {
	t.Fahrenheit = t.Celsius*1.8 + 32
}

func (t *Weather) ConvertKelvin() {
	t.Kelvin = t.Celsius + 273.15
}
