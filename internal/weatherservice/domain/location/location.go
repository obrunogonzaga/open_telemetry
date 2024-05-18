package location

import "errors"

type Location struct {
	CEP  string
	City string
}

func NewLocation(city string) (*Location, error) {
	location := &Location{
		City: city,
	}
	err := location.IsValid()
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (l *Location) IsValid() error {

	// TODO: Implement the validation logic for CEP
	if l.City == "" {
		return errors.New("invalid city")
	}
	return nil
}
