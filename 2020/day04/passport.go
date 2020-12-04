package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func (p Passport) IsSet() bool {
	// Note: CountryID is optional
	return p.BirthYear != "" && p.IssueYear != "" && p.ExpirationYear != "" &&
		p.Height != "" && p.HairColor != "" && p.EyeColor != "" && p.PassportID != ""
}

func (p Passport) IsValid() bool {
	return p.ValidateBirthYear() && p.ValidateIssueYear() && p.ValidateExpirationYear() &&
		p.ValidateHeight() && p.ValidateHairColor() && p.ValidateEyeColor() &&
		p.ValidatePassportID()
}

func (p Passport) ValidateBirthYear() bool {
	year, err := strconv.ParseInt(p.BirthYear, 10, 64)
	if err != nil {
		return false
	}

	return year >= 1920 && year <= 2002
}

func (p Passport) ValidateIssueYear() bool {
	year, err := strconv.ParseInt(p.IssueYear, 10, 64)
	if err != nil {
		return false
	}

	return year >= 2010 && year <= 2020
}

func (p Passport) ValidateExpirationYear() bool {
	year, err := strconv.ParseInt(p.ExpirationYear, 10, 64)
	if err != nil {
		return false
	}

	return year >= 2020 && year <= 2030
}

func (p Passport) ValidateHeight() bool {
	if len(p.Height) < 3 {
		return false
	}

	unit := p.Height[len(p.Height)-2:]
	value := p.Height[:len(p.Height)-2]

	switch unit {
	case "cm":
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return false
		}

		return v >= 150 && v <= 193
	case "in":
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return false
		}

		return v >= 59 && v <= 76
	}

	return false
}

var haircolorRegexp = regexp.MustCompile("^#[0-9a-z]{6}$") // nolint:gochecknoglobals

func (p Passport) ValidateHairColor() bool {
	return haircolorRegexp.MatchString(p.HairColor)
}

func (p Passport) ValidateEyeColor() bool {
	switch p.EyeColor {
	case "amb", "blu", "brn", "grn", "gry", "hzl", "oth":
		return true
	}

	return false
}

var passportIDRegexp = regexp.MustCompile(`^\d{9}$`) // nolint:gochecknoglobals

func (p Passport) ValidatePassportID() bool {
	return passportIDRegexp.MatchString(p.PassportID)
}

func parse(data []byte) ([]*Passport, error) {
	passports := []*Passport{}

	x := bytes.Split(data, []byte("\n\n"))
	for i, input := range x {
		line := bytes.ReplaceAll(input, []byte("\n"), []byte{' '})
		fields := bytes.Split(bytes.TrimSpace(line), []byte{' '})

		passport := &Passport{}

		for j, field := range fields {
			parts := strings.Split(string(field), ":")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid input %s at passport %d, field %d", field, i, j)
			}

			switch parts[0] {
			case "byr":
				passport.BirthYear = parts[1]
			case "iyr":
				passport.IssueYear = parts[1]
			case "eyr":
				passport.ExpirationYear = parts[1]
			case "hgt":
				passport.Height = parts[1]
			case "hcl":
				passport.HairColor = parts[1]
			case "ecl":
				passport.EyeColor = parts[1]
			case "pid":
				passport.PassportID = parts[1]
			case "cid":
				passport.CountryID = parts[1]
			default:
				return nil, fmt.Errorf("unknown field type %s at passport %d, field %d", parts[0], i, j)
			}
		}

		passports = append(passports, passport)
	}

	return passports, nil
}
