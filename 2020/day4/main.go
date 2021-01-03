package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := strings.Split(input, "\n\n")
	var passports []Passport
	for _, d := range data {
		passports = append(passports, generatePassport(d))
	}
	validPassports := 0
	for _, passport := range passports {
		fmt.Printf("%+v \n", passport)
		if passport.validate() {
			fmt.Println("Valid passport")
			validPassports++
		}
	}
	fmt.Println("Number of valid passports", validPassports)

}

/*
	ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929
*/
type Passport struct {
	fields map[string]string
}

func generatePassport(data string) Passport {
	pass := Passport{fields: make(map[string]string)}
	d := strings.ReplaceAll(data, "\n", " ")
	pairs := strings.Split(d, " ")
	for _, pair := range pairs {
		kv := strings.Split(pair, ":")
		if len(kv) != 2 {
			continue
		}
		pass.fields[kv[0]] = kv[1]
	}
	return pass
}
func (p Passport) validate() bool {
	val, ok := p.fields["byr"]
	if !ok {
		return false
	}
	v, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	if v < 1920 || v > 2002 {
		return false
	}

	val, ok = p.fields["iyr"]
	if !ok {
		return false
	}
	v, err = strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	if v < 2010 || v > 2020 {
		return false
	}

	val, ok = p.fields["eyr"]
	if !ok {
		return false
	}
	v, err = strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	if v < 2020 || v > 2030 {
		return false
	}

	val, ok = p.fields["hgt"]
	if !ok {
		return false
	}
	// val 152cm
	if strings.HasSuffix(val, "cm") {
		heightCm := strings.TrimSuffix(val, "cm")
		v, err = strconv.Atoi(heightCm)
		if err != nil {
			panic(err)
		}
		if v < 150 || v > 193 {
			return false
		}
	} else if strings.HasSuffix(val, "in") {
		heightIn := strings.TrimSuffix(val, "in")
		v, err := strconv.Atoi(heightIn)
		if err != nil {
			panic(err)
		}
		if v < 59 || v > 76 {
			return false
		}

	} else {
		fmt.Println("hgt doesn't have a suffix..")
		return false
	}

	val, ok = p.fields["hcl"]
	if !ok {
		return false
	}
	hairRegex := regexp.MustCompile("#([0-9a-f]{6})")
	result := hairRegex.FindAll([]byte(val), -1)
	if result == nil {
		return false
	}

	val, ok = p.fields["ecl"]
	if !ok {
		return false
	}
	switch val {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
	default:
		return false //not a valid type
	}

	val, ok = p.fields["pid"]
	if !ok {
		return false
	}
	// 9 digit number with leading zeros
	val = strings.TrimSpace(val)
	if len(val) != 9 {
		return false
	}

	if _, ok := p.fields["cid"]; !ok {
		fmt.Println("missing country identifier...")
	}
	return true
}
