package secretSharingSchemes

import (
	"crypto/rand"
	"encoding/binary"
)

type Share struct {
	Index int
	Share int
}

type Server struct {
	Index  int
	Shares []int
}

type TwoShares struct {
	first  string
	second string
}

// Generate a random 128-bit AES key
func GenerateAESKey() []byte {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		panic("Error in GenerateAESKey(): " + err.Error())
	}
	return key
}

// Generate a random 32-bit share
func Generate32BitInteger() int {
	key_as_prime, err := rand.Prime(rand.Reader, 32)
	if err != nil {
		panic("The length of the key must be greater than 0: " + err.Error())
	}
	return int(key_as_prime.Int64())
}

// Modulo operation
func Mod(secret, prime int) int {
	return (secret%prime + prime) % prime
}

// Converts a byte array to an array of integers
func ToInt(bytes []byte) int {
	return int(binary.BigEndian.Uint16(bytes))
}

// Transposes an array of integers
func Transpose(intArray [][]int) [][]int {
	result := make([][]int, len(intArray[0]))
	for i := range result {
		result[i] = make([]int, len(intArray))
	}
	for i := range intArray {
		for j := range intArray[i] {
			result[j][i] = intArray[i][j]
		}
	}
	return result
}

// Transpose an array of servers
func TransposePlayers(servers [][]Server) [][]Server {
	// create a struct of sharesArray to put each share into
	result := make([][]Server, len(servers[0]))

	for i := range result { // TODO: this might be redundant
		result[i] = make([]Server, len(servers))
	}

	for i := 0; i < len(servers); i++ {
		for j := 0; j < len(servers[i]); j++ {
			result[j][i] = servers[i][j]
		}
	}
	return result
}

func CheckError(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

// This is used to check arguments given to SecretShare and SecretReconstruct
func CheckArgument(condition bool, msg string) error {
	if !condition {
		panic(msg)
	}
	return nil
}

// Converts a 2D array of integers to a 2D array of shares
func ToShares(array [][]int) [][]Share {
	result := make([][]Share, len(array))
	for i := range result {
		result[i] = make([]Share, len(array[0]))
	}
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[0]); j++ {
			result[i][j] = Share{Index: i + 1, Share: array[i][j]}
		}
	}
	return result
}

// Finds k-permutations of t, n.
// This is primarily used in the t-out-of-n secret sharing scheme
func FindKPermutations(t, n int) int {
	return (factorial(n) / (factorial(t-1) * factorial(n-(t-1))))
}

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func RemoveDuplicateIntegers(array []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range array {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func RemoveDuplicateStrings(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func XorShares(shares []int) int {
	secret := shares[0]
	for i := 1; i < len(shares); i++ {
		secret = secret ^ shares[i]
	}
	return secret
}

func factorial(num int) int {
	if num == 1 || num == 0 {
		return num
	}
	return num * factorial(num-1)
}
