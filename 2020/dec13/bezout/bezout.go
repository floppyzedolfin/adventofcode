package bezout

import (
	"math/big"
)

// BBPair is a Bachet-Bézout pair. The number X we are looking for
//  verifies X == Remainder % Divisor
type BBPair struct {
	Divisor   big.Int
	Remainder big.Int
}

// Solve a Chinese Remainder Theorem and returns the smallest positive value that matches
func Solve(c []BBPair) big.Int {
	switch len(c) {
	case 0:
		// unlikely, but you never know...
		return *big.NewInt(0)
	case 1:
		// shouldn't happen, but this is a valid answer
		return c[0].Remainder
	case 2:
		// this is the end of the task, usually
		// get the Bézout coefficients
		x, y := extendedEuclideanAlgorithm(c[0].Divisor, c[1].Divisor)
		res := add(mul(mul(y, c[0].Remainder), c[1].Divisor), mul(mul(x, c[1].Remainder), c[0].Divisor))
		// let's try and make it the smallest positive result - we can add or subtract c[0].Divisor * c[1].Divisor at will
		c0c1 := mul(c[0].Divisor, c[1].Divisor)

		// at this point, res could be negative, so let's take the smallest positive answer
		if res.Sign() == -1 {
			q := quo(sub(*big.NewInt(0), res), c0c1)
			res = add(res, mul(add(q, *big.NewInt(1)), c0c1))
		}
		return rem(res, c0c1)
	default:
		// we'll solve this complex problem recursively: solve the first two equations,
		// and then use that solution with the remaining equations.

		// merge the first two equations
		mergedRemainder := Solve(c[:2])
		mergedPair := BBPair{Divisor: mul(c[0].Divisor, c[1].Divisor), Remainder: mergedRemainder}

		// this new problem has a complexity (length of input) reduced by 1 when compared to the current problem.
		return Solve(append(c[2:], mergedPair))
	}
}

// extendedEuclideanAlgorithm returns the (x,y) Bézout coefficients that satisfy a*x + b*y = 1, given a and b are coprime
func extendedEuclideanAlgorithm(a, b big.Int) (x, y big.Int) {
	// graciously taken from Wikipedia
	oldr, r := a, b
	olds, s := *big.NewInt(1), *big.NewInt(0)
	oldt, t := *big.NewInt(0), *big.NewInt(1)

	// check if r is 0
	for r.Sign() != 0 {
		quotient := quo(oldr, r)                  // q = oldr / r
		oldr, r = r, sub(oldr, mul(quotient, r))  // oldr, r = r, oldr - q*r
		olds, s = s, sub(olds, mul(quotient, s))  // olds, s = s, olds - q*s
		oldt, t = t, sub(oldt, mul(quotient, t))  // oldt, t = t, oldt - q*t
	}
	return olds, oldt
}

// I didn't like big.Int's operations.
//  It won't let you big.Mul(a, b big.Int) big.Int, for instance. For some reason, you already need a big.Int
//  Hopefully, generics will solve this...

// mul returns the product a*b
func mul(a, b big.Int) big.Int {
	return *big.NewInt(1).Mul(&a, &b)
}

// add returns the sum a+b
func add(a, b big.Int) big.Int {
	return *big.NewInt(1).Add(&a, &b)
}

// sub returns the difference a-b
func sub(a, b big.Int) big.Int {
	return *big.NewInt(1).Sub(&a, &b)
}

// quo returns the integer part of the quotient a/b
func quo(a, b big.Int) big.Int {
	return *big.NewInt(1).Quo(&a, &b)
}

// rem is the remainder of a/b, it's equal to a%b
func rem(a, b big.Int) big.Int {
	// it's me in the corner
	return *big.NewInt(1).Rem(&a, &b)
}
