package traintickets

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type problem struct {
	fields   []field
	myTicket ticket

	// nearbyTickets is a map, which makes removing entries easier
	nearbyTickets map[int]ticket

	// parsingTickets is used for parsing purposes: false = parsing fields; true = parsing tickets (starting with mine)
	parsingTickets bool
}

func newProblem() *problem {
	p := problem{}
	p.fields = make([]field, 0)
	p.nearbyTickets = make(map[int]ticket)
	return &p
}

// parseLine implements the LineParser interface
func (p *problem) ParseLine(line string) error {
	if len(line) == 0 {
		// nothing to be done on empty lines
		return nil
	}

	if strings.Contains(line, "ticket") {
		p.parsingTickets = true
		return nil
	}

	// at this point, we're dealing with data that's rawer than Elisabeta Lipa
	if p.parsingTickets {
		return p.parseTicket(line)
	} else {
		return p.parseField(line)
	}
}

func (p *problem) parseTicket(line string) error {
	// a tickets consists in a series of numbers.
	values := strings.Split(line, ",")
	t := ticket{make([]int, len(values))}
	for i, v := range values {
		value, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("unable to parse value '%s' as a number: %s", v, err.Error())
		}
		t.values[i] = value
	}

	// if this is the first ticket we see, it's mine.
	if len(p.myTicket.values) == 0 {
		p.myTicket = t
	} else {
		p.nearbyTickets[len(p.nearbyTickets)] = t
	}

	return nil
}

func (p *problem) parseField(line string) error {
	const fieldRegularExpression = `([a-z ]*): ([0-9]*)-([0-9]*) or ([0-9]*)-([0-9]*)$`
	fieldRE := regexp.MustCompile(fieldRegularExpression)
	matches := fieldRE.FindStringSubmatch(line)

	const numFields = 6
	if len(matches) != numFields {
		// need to take into account that FindStringSubmatch returns the whole string as [0].
		return fmt.Errorf("expected %d fields, got %d instead", numFields-1, len(matches)-1)
	}

	// this is by far not the nicest code ever.
	f := field{name: matches[1], ranges: make([]minMax, 2)}
	min1, err := strconv.Atoi(matches[2])
	if err != nil {
		return fmt.Errorf("unable to convert '%s' to an integer: %s", matches[2], err.Error())
	}
	max1, err := strconv.Atoi(matches[3])
	if err != nil {
		return fmt.Errorf("unable to convert '%s' to an integer: %s", matches[3], err.Error())
	}
	min2, err := strconv.Atoi(matches[4])
	if err != nil {
		return fmt.Errorf("unable to convert '%s' to an integer: %s", matches[4], err.Error())
	}
	max2, err := strconv.Atoi(matches[5])
	if err != nil {
		return fmt.Errorf("unable to convert '%s' to an integer: %s", matches[5], err.Error())
	}

	// finally save these
	f.ranges[0] = minMax{min1, max1}
	f.ranges[1] = minMax{min2, max2}
	p.fields = append(p.fields, f)
	return nil
}

// sumInvalidValues sums the unexpected values of the tickets, with regards to the allowed values in the fields.
//  my ticket is not to be taken into consideration here.
func (p *problem) sumInvalidValues() uint64 {
	sum := 0
	for _, t := range p.nearbyTickets {
		s, _ := p.checkValues(t)
		sum += s
	}
	return uint64(sum)
}

func (p *problem) multiplyDepartures() uint64 {
	// get rid of tickets that don't match rules
	p.removeInvalidTickets()

	// at this point, we assume both the existence and the unicity of a solution
	p.assignColumnsToFields()

	const departureKey = "departure"
	return p.multiplyFieldsOfMyTicket(departureKey)
}

func (p *problem) assignColumnsToFields() {
	// start by creating a structure for the results
	p.buildSieve()

	// use tickets to restrict possible matches
	p.sieveWithTickets()

	// finish solving the problem
	p.removeMatchedFields()
}

// buildSieve prepares the structure for a solving
func (p *problem) buildSieve() {
	// let's count entries
	nbEntries := len(p.myTicket.values)
	// by default, all fields are eligible
	for index := range p.fields {
		p.fields[index].possibleEntries = make(map[int]struct{})
		for i := 0; i < nbEntries; i++ {
			p.fields[index].possibleEntries[i] = struct{}{}
		}
	}
}

// sieveWithTickets checks the tickets against the fields: if a ticket doesn't match a field's definition, that rules it out
func (p *problem) sieveWithTickets() {
	for _, f := range p.fields {
		for _, t := range p.nearbyTickets {
			f.sieve(t)
		}
	}
}

// removeMatchedFields cleans the eligible columns for fields that are only mapped to one column
func (p *problem) removeMatchedFields() {
	// sort by increasing number of potential entries
	sort.Slice(p.fields, func(i, j int) bool { return len(p.fields[i].possibleEntries) < len(p.fields[j].possibleEntries) })

	for i, f := range p.fields[:len(p.fields)-1] {
		if len(f.possibleEntries) != 1 {
			panic(fmt.Sprintf("unsolvable problem, can't choose column for field %s: %v", f.name, f.possibleEntries))
		}
		// this possibleEntries loop is the only way to get the key. we don't actually loop (cf. the panic above).
		for k := range f.possibleEntries {
			for _, f2 := range p.fields[i+1:] {
				delete(f2.possibleEntries, k)
			}
		}
	}
}

// multiplyFieldsOfMyTicket multiplies the fields with name containing the key
func (p *problem) multiplyFieldsOfMyTicket(key string) uint64 {
	product := uint64(1)
	for _, f := range p.fields {
		if strings.Contains(f.name, key) {
			// this for loop only iterates over one element, at this point (see removeMatchedFields)
			for k := range f.possibleEntries {
				product *= uint64(p.myTicket.values[k])
			}
		}
	}
	return product
}

// checkValues returns the sum of invalid values of a ticket, and a boolean indicating whether a rule was found for a field of the ticket
func (p *problem) checkValues(t ticket) (int, bool) {
	sum, found := 0, false
	for _, v := range t.values {
		if !p.checkValue(v) {
			sum += v
			found = true
		}
	}
	return sum, found
}

// checkValue returns whether a value is valid against all fields.
func (p *problem) checkValue(value int) bool {
	for _, fs := range p.fields {
		for _, f := range fs.ranges {
			if f.min <= value && value <= f.max {
				// we're in the range
				return true
			}
		}
	}
	return false
}

// removeInvalidTickets checks the tickets against all the rules, and removes those that can't match a field.
func (p *problem) removeInvalidTickets() {
	invalidTickets := make([]int, 0)
	for i, t := range p.nearbyTickets {
		if _, found := p.checkValues(t); found {
			invalidTickets = append(invalidTickets, i)
		}
	}

	for _, i := range invalidTickets {
		delete(p.nearbyTickets, i)
	}
}
