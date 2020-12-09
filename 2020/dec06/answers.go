package dec06


// group holds the answers of a group of people
type group struct {
	people []person
}

// person holds the answers of a person
type person struct {
	response answers
}

type answers map[byte]struct{}

// countPart1 returns the number of questions whereto anyone answered 'yes'.
func countPart1(gs []group) int {
	yesCount := 0
	for _, g := range gs {
		yesCount += len(g.differentYeses())
	}
	return yesCount
}

// differentYeses returns the different questions people of a group answered with Yes at least once.
func (g group) differentYeses() answers {
	yeses := make(answers)
	for _, p := range g.people {
		for a := range p.response {
			yeses[a] = struct{}{}
		}
	}
	return yeses
}

// countPart2 returns the number of questions whereto group had all people answer yes.
func countPart2(gs []group) int {
	yesCount := 0
	for _, g := range gs {
		yesCount += len(g.commonYeses())

	}
	return yesCount
}

// commonYeses returns the list of responses whereto every person of the group answered Yes
func (g group) commonYeses() answers {
	if len(g.people) == 0 {
		// this group is empty
		return nil
	}
	// use the response of the first person as a sieve
	res := g.people[0].answers()
	for _, p := range g.people[1:] {
		res.sieveAnswers(p.response)
	}
	return res
}

// sieveAnswers only keeps in a what is also in others. Notice we don't need to iterate over other, as we don't care.
func (a *answers) sieveAnswers(other answers) {
	for k := range *a {
		if _, ok := other[k]; !ok {
			delete(*a, k)
		}
	}
}

func (p person) answers() answers {
	m := make(answers)
	for k, v := range p.response {
		m[k] = v
	}
	return m
}