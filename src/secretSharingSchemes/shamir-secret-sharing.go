package secretSharingSchemes

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"math/big"
)

type ShamirSecretSharing struct {
	// A Shamir Secret Sharing object
}

// An important note about this implementation of Shamir's secret-sharing is that
// the value of t is 1 higher than in the thesis. This implementation work with
// a polynomial of degree t, while the thesis works with a polynomial of degree t-1.
func (s *ShamirSecretSharing) SecretShare(secret []byte, t int, n int) [][]Share {
	CheckArgument(len(secret) == 16, "Invalid argument: secret must be 16 bytes")
	CheckArgument(t > 0, "Invalid argument: t must be greater than 0")
	CheckArgument(n >= t, "Invalid argument: n must be greater than or equal to t")

	var values = make([][]int, 8)

	// The 128-bit AES-key is split into 8 pieces of 16 bytes
	for i := 0; i < 8; i++ {
		prime := createPrime()
		slice := ToInt(secret[i*2 : (i+1)*2])
		coefficients := createCoefficients(t)
		values[i] = createShamirShares(slice, n, t, prime, coefficients)
	}
	values = Transpose(values)
	return ToShares(values)
}

func (s *ShamirSecretSharing) SecretReconstruct(shares [][]Share) []byte {
	CheckArgument(len(shares) > 0, "Invalid argument: No shares provided")

	reconstruction := new(bytes.Buffer)
	for i := 0; i < 8; i++ {
		sharesArray := make([]Share, len(shares))
		for j := 0; j < len(shares); j++ {
			sharesArray[j] = shares[j][i]
		}
		err := (binary.Write(reconstruction,
			binary.BigEndian, uint16(interpolate(sharesArray))))
		CheckError(err, "The secret was not in the range of 0 through 65535")
	}
	return reconstruction.Bytes()
}

func createPrime() int {
	return 15307073
}

func createCoefficients(t int) []int {
	coefficients := make([]int, t)
	// for each secret slice, generate t random coefficients
	for i := 0; i < t; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(10000))
		coefficients[i] = int(num.Int64())
	}
	return coefficients
}

func createShamirShares(secret int, n int, t int, prime int, coefficients []int) []int {
	shares := make([]int, n)
	for i := 0; i < n; i++ {
		shares[i] = secret
		for j := 0; j < t; j++ {
			var x, exp = big.NewInt(int64(i + 1)), big.NewInt(int64(j + 1))
			value := x.Exp(x, exp, nil)
			mod := value.Mod(value, big.NewInt(int64(prime)))
			shares[i] = shares[i] + coefficients[j]*int(mod.Int64())
		}
		shares[i] = Mod(shares[i], prime)
	}
	return shares
}

func interpolate(shares []Share) int {
	reconstruction := 0
	prime := createPrime()

	// Initialize the deltas with value 1
	deltas := make([]int, len(shares))
	for i := range deltas {
		deltas[i] = 1
	}

	// Calulate f(0) of the gives shares using Langrange interpolation
	for i := 0; i < len(shares); i++ {
		denominator := 1
		for j := 0; j < len(shares); j++ {
			if i != j {
				deltas[i] = Mod(-shares[j].Index*deltas[i], prime)
				denominator = denominator * (shares[i].Index - shares[j].Index)
				denominator = Mod(denominator, prime)
			}
		}

		// Calculate the modulo multiplicative inverse of the denominator with math/big library
		denominatorBig := big.NewInt(int64(denominator))
		primeBig := big.NewInt(int64(prime))
		modInvBig := denominatorBig.ModInverse(denominatorBig, primeBig)
		deltas[i] = Mod(deltas[i]*int(modInvBig.Int64()), prime)

		// Reconstruct the secret
		reconstruction += deltas[i] * shares[i].Share
	}
	return Mod(reconstruction, prime)
}
