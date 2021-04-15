package dec04

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/floppyzedolfin/adventofcode/pkg/ptr"
)

func validatePassports(ps []passport) int {
	validPassports := 0
	for _, p := range ps {
		if p.hasMandatoryFields() {
			validPassports++
		}
	}
	return validPassports
}

func (p passport) hasMandatoryFields() bool {
	// missing a CID is fine.
	return p.birthYear != nil &&
		p.issueYear != nil &&
		p.expirationYear != nil &&
		p.height != nil &&
		p.hairColour != nil &&
		p.eyeColour != nil &&
		p.passportID != nil
}

// validatePassportsThoroughly checks that passports have all fields and that each fields is properly filled.
func validatePassportsThoroughly(ps []passport) int {
	validPassports := 0
	for _, p := range ps {
		if p.isThoroughlyValid() {
			validPassports++
		}
	}
	return validPassports
}

func (p passport) isThoroughlyValid() bool {
	return p.birthYearIsValid() &&
		p.issueYearIsValid() &&
		p.expirationYearIsValid() &&
		p.heightIsValid() &&
		p.hairColourIsValid() &&
		p.eyeColourIsValid() &&
		p.passportIDIsValid()
}

// String is a tool we might want to use when debugging
func (p passport) String() string {
	o := strings.Builder{}
	if p.birthYear == nil {
		o.WriteString("byr: nil - ")
	} else {
		o.WriteString(fmt.Sprintf("byr: %d - ", *p.birthYear))
	}
	if p.issueYear == nil {
		o.WriteString("iyr: nil - ")
	} else {
		o.WriteString(fmt.Sprintf("iyr: %d - ", *p.issueYear))
	}
	if p.expirationYear == nil {
		o.WriteString("eyr: nil - ")
	} else {
		o.WriteString(fmt.Sprintf("eyr: %d - ", *p.expirationYear))
	}
	if p.height == nil {
		o.WriteString("hgt: nil - ")
	} else {
		o.WriteString(fmt.Sprintf("hgt: %s - ", *p.height))
	}
	if p.hairColour == nil {
		o.WriteString("hcl: nil - ")
	} else {
		o.WriteString(fmt.Sprintf("hcl: %s - ", *p.hairColour))
	}
	if p.eyeColour == nil {
		o.WriteString("ecl: nil - ")
	} else {
		o.WriteString(fmt.Sprintf("ecl: %s - ", *p.eyeColour))
	}
	if p.passportID == nil {
		o.WriteString("pid: nil - ")
	} else {
		o.WriteString(fmt.Sprintf("pid: %s - ", *p.passportID))
	}
	return o.String()
}

func (p *passport) setField(field string) error {
	const (
		kvSeparator       = ":"
		birthYearKey      = "byr"
		issueYearKey      = "iyr"
		expirationYearKey = "eyr"
		heightKey         = "hgt"
		hairColourKey     = "hcl" // not recommended, unless you're very basic
		eyeColourKey      = "ecl"
		passportIDKey     = "pid"
		countryIDKey      = "cid"
	)
	inputs := strings.Split(field, kvSeparator)
	if len(inputs) != 2 {
		return fmt.Errorf("invalid field '%s', expected 'key:value'", field)
	}

	switch inputs[0] {
	case birthYearKey:
		return p.byr(inputs[1])
	case issueYearKey:
		return p.iyr(inputs[1])
	case expirationYearKey:
		return p.eyr(inputs[1])
	case heightKey:
		return p.hgt(inputs[1])
	case hairColourKey:
		return p.hcl(inputs[1])
	case eyeColourKey:
		return p.ecl(inputs[1])
	case passportIDKey:
		return p.pid(inputs[1])
	case countryIDKey:
		// since neither Part1 nor Part2 performs checks on this, I decided to discard it
		return nil
	default:
		return fmt.Errorf("unknown key '%s'", inputs[0])
	}
	return nil
}

type passport struct {
	birthYear      *int
	issueYear      *int
	expirationYear *int
	height         *string // nope, I'm not converting inches from somewhere to metres.
	hairColour     *string
	eyeColour      *string
	passportID     *string
	countryID      *int
}

func (p *passport) byr(year string) error {
	byr, err := strconv.Atoi(year)
	if err != nil {
		return fmt.Errorf("unable to convert '%s' to birth year", year)
	}
	p.birthYear = ptr.Int(byr)
	return nil
}

func (p *passport) iyr(year string) error {
	iyr, err := strconv.Atoi(year)
	if err != nil {
		return fmt.Errorf("unable to convert '%s' to issue year", year)
	}
	p.issueYear = ptr.Int(iyr)
	return nil
}

func (p *passport) eyr(year string) error {
	eyr, err := strconv.Atoi(year)
	if err != nil {
		return fmt.Errorf("unable to convert '%s' to expiration year", year)
	}
	p.expirationYear = ptr.Int(eyr)
	return nil
}

func (p *passport) hgt(height string) error {
	p.height = &height
	return nil
}

func (p *passport) hcl(hairColour string) error {
	p.hairColour = &hairColour
	return nil
}

func (p *passport) ecl(eyeColour string) error {
	p.eyeColour = &eyeColour
	return nil
}

func (p *passport) pid(passportID string) error {
	p.passportID = &passportID
	return nil
}

func (p passport) birthYearIsValid() bool {
	if p.birthYear == nil {
		return false
	}
	const (
		minBirthYear = 1920
		maxBirthYear = 2002
	)
	return minBirthYear <= *p.birthYear && *p.birthYear <= maxBirthYear
}

func (p passport) issueYearIsValid() bool {
	if p.issueYear == nil {
		return false
	}
	const (
		minIssueYear = 2010
		maxIssueYear = 2020
	)
	return minIssueYear <= *p.issueYear && *p.issueYear <= maxIssueYear
}

func (p passport) expirationYearIsValid() bool {
	if p.expirationYear == nil {
		return false
	}
	const (
		minExpirationYear = 2020
		maxExpirationYear = 2030
	)
	return minExpirationYear <= *p.expirationYear && *p.expirationYear <= maxExpirationYear
}

func (p passport) heightIsValid() bool {
	if p.height == nil {
		return false
	}
	const (
		reHeight    = `^(\d+)([[:alpha:]]+)$`
		inches      = "in"
		centimetres = "cm"
		minHeightCm = 150
		maxHeightCm = 193
		minHeightIn = 59
		maxHeightIn = 76
	)
	re := regexp.MustCompile(reHeight)
	res := re.FindStringSubmatch(*p.height)
	if len(res) != 3 {
		// we need 2 items here - a value and a unit. 3 is because FindStringSubmatch returns the input string as the first item.
		return false
	}
	height, err := strconv.Atoi(res[1])
	if err != nil {
		// shouldn't happen, as we've looked for numbers, but, hey... why wouldn't someone measure 18945684357351964852168 cm ?
		return false
	}

	switch res[2] {
	case centimetres:
		return minHeightCm <= height && height <= maxHeightCm
	case inches:
		// probably "legacy" code we must maintain for reasons
		return minHeightIn <= height && height <= maxHeightIn
	default:
		// unsupported unit
		return false
	}
}

func (p passport) hairColourIsValid() bool {
	if p.hairColour == nil {
		return false
	}
	const reHairColour = `^#[0-9a-f]{6}$`
	re := regexp.MustCompile(reHairColour)
	return re.MatchString(*p.hairColour)
}

func (p passport) eyeColourIsValid() bool {
	validEyeColours := map[string]struct{}{
		"amb": {},
		"blu": {},
		"brn": {},
		"gry": {},
		"grn": {},
		"hzl": {},
		"oth": {},
	}
	if p.eyeColour == nil {
		return false
	}
	_, ok := validEyeColours[*p.eyeColour]
	return ok
}

func (p passport) passportIDIsValid() bool {
	if p.passportID == nil {
		return false
	}
	const rePassportID = `^\d{9}$`
	re := regexp.MustCompile(rePassportID)
	return re.MatchString(*p.passportID)
}
