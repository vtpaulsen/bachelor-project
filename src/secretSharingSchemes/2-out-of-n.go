package secretSharingSchemes

import (
	"math"
	"strconv"
)

type TwoNSecretSharing struct {
	// Create a 2-out-of-n secret sharing scheme object
}

func (s *TwoNSecretSharing) SecretShare(message string, n int) [][]string {
	CheckArgument(len(message) == 16, "Invalid argument: secret must be 16 bytes")
	CheckArgument(n >= 2, "Invalid argument: n must be greater than or equal to 2")

	// Iterations is the number of times shares are created
	iterations := int(math.Log2(float64(n)))
	shares := make([][]string, n)
	for i := range shares {
		shares[i] = make([]string, iterations+1)
	}
	values := make([]TwoShares, iterations)
	for k := 0; k < iterations; k++ {
		first, second := create2NShares(message)
		values[k] = TwoShares{first, second}
	}

	// Convert the index to binary
	for i := 0; i <= n-1; i++ {
		binary := strconv.FormatInt(int64(i), 2)
		if len(binary) < iterations {
			for len(binary) < iterations {
				binary = "0" + binary
			}
		}

		// The first column is the index of the share
		shares[i][0] = strconv.Itoa(i)

		// The string 'binary' is used to select the correct shares
		for j := 1; j < iterations; j++ { // We have to avoid overwriting the index in S
			if binary[j] == '0' {
				shares[i][j] = values[j-1].first // We have to start from zero in shares
			} else {
				shares[i][j] = values[j-1].second
			}
		}
	}
	return shares
}

func (s *TwoNSecretSharing) SecretReconstruct(first []string, second []string) string {
	reconstruction := ""
	for i := range first[1:] {
		if first[i] != second[i] {
			reconstruction = xor(first[i], second[i])
		}
	}
	return reconstruction
}

func create2NShares(message string) (string, string) {
	CheckArgument(len(message) > 0, "The length of the message must be 128 bits")

	first := GenerateAESKey()
	second := make([]byte, len(message))
	for i := range first {
		second[i] = message[i] ^ first[i]
	}
	return string(first), string(second)
}

func xor(first string, second string) string {
	result := make([]byte, len(first))
	for i := range result {
		result[i] = first[i] ^ second[i]
	}
	return string(result)
}
