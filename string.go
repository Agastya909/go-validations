package validation

import (
	"errors"
	"reflect"
	"regexp"
)

type ValidationValue struct {
	value any
}

func (s *ValidationValue) IsString() error {
	strType := reflect.TypeOf(s.value).Kind()

	if strType != reflect.String {
		return errors.New("value not a string")
	}

	return nil
}

func (s *ValidationValue) ValidateLength(minMax ...int) error {
	if len(minMax) == 0 {
		return errors.New("no length provided")
	}
	if err := s.IsString(); err != nil {
		return err
	}
	strLen := len(s.value.(string))
	if len(minMax) == 1 {
		if strLen < minMax[0] {
			return errors.New("string too short")
		}
	} else {
		if strLen > minMax[1] {
			return errors.New("string too long")
		}
		if strLen < minMax[0] {
			return errors.New("string too short")
		}
	}
	return nil
}

func (s *ValidationValue) IsEmail() error {
	if err := s.IsString(); err != nil {
		return err
	}
	re, err := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if err != nil {
		return errors.New("invalid email")
	}
	if !re.MatchString(s.value.(string)) {
		return errors.New("invalid email")
	}
	return nil
}

func (s *ValidationValue) HasEmptySpaces() error {
	if err := s.IsString(); err != nil {
		return err
	}
	re, err := regexp.Compile(`\s`)
	if err != nil {
		return errors.New("invalid regex")
	}
	if !re.MatchString(s.value.(string)) {
		return errors.New("no space found")
	}
	return nil
}

func (s *ValidationValue) IsAlphaOnly() error {
	if err := s.IsString(); err != nil {
		return err
	}

	re, err := regexp.Compile(`^[a-zA-Z]+$`)
	if err != nil {
		return errors.New("invalid regex")
	}
	if !re.MatchString(s.value.(string)) {
		return errors.New("string contains non-alphabetic characters")
	}

	return nil
}

func (s *ValidationValue) IsAlphaNumeric() error {
	if err := s.IsString(); err != nil {
		return err
	}

	re, err := regexp.Compile(`^[a-zA-Z0-9]+$`)
	if err != nil {
		return errors.New("invalid regex")
	}
	if !re.MatchString(s.value.(string)) {
		return errors.New("string contains non-alphanumeric characters")
	}

	return nil
}
