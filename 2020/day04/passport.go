package main

import (
	"bytes"
	"fmt"
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
