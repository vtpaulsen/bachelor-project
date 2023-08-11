package secretSharingSchemes

import (
	"bytes"
	"encoding/binary"
	"math/bits"
)

type ReplicatedSecretSharing struct {
	// A t-out-of-n Secret Sharing object
}

func (l *ReplicatedSecretSharing) SecretShare(secret []byte, t int, n int) [][]Server {
	var shares = make([][]Server, 8)
	for i := 0; i < 8; i++ {
		slice := secret[i*2 : (i+1)*2]
		secret := int(binary.BigEndian.Uint16(slice))
		shares[i] = secretShare(secret, t, n)
	}
	shares = TransposePlayers(shares)
	return shares
}

func (l *ReplicatedSecretSharing) SecretReconstruct(servers [][]Server) []byte {
	CheckArgument(len(servers) > 0, "Invalid argument: No shares provided")

	reconstruction := new(bytes.Buffer)
	for i := 0; i < 8; i++ {
		shares := make([]int, len(servers)*len(servers[0][i].Shares))
		numberOfShares := len(servers[0][i].Shares)
		for j := 0; j < len(servers); j++ {
			copy(shares[j*numberOfShares:], servers[j][i].Shares)
		}
		shares = RemoveDuplicateIntegers(shares)
		err := (binary.Write(reconstruction,
			binary.BigEndian, uint16(XorShares(shares))))
		CheckError(err, "The share was not in the range of 0 through 65535")
	}
	return reconstruction.Bytes()
}

func secretShare(secret int, t int, n int) []Server {
	servers := []Server{}
	k := FindKPermutations(t, n)

	// Create n servers
	for i := 1; i <= n; i++ {
		servers = append(servers, Server{Index: i})
	}
	shares := createTNShares(secret, t, k)

	// Creat a string of length n with t-1 0's, where 0 indicates parties that do not get the share
	delegate := zeroOneString(n, t-1)

	// Find all permutations of the string
	temp := []string{}
	Perm([]rune(delegate), func(a []rune) {
		temp = append(temp, string(a))
	})
	permutations := RemoveDuplicateStrings(temp)

	// Assigning shares to servers
	for i := 0; i < len(shares); i++ {
		for j := 0; j < len(permutations[i]); j++ {
			if permutations[i][j] == '1' {
				servers[j].Shares = append(servers[j].Shares, shares[i])
			}
		}
	}
	return servers
}

func createTNShares(secret int, t int, n int) []int {
	shares := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		shares[i] = Generate32BitInteger()
	}
	last := secret
	for i := range shares {
		last = last ^ shares[i]
	}
	shares = append(shares, last)
	max := 0
	for i := 0; i < len(shares); i++ {
		if bits.Len(uint(shares[i])) > max {
			max = bits.Len(uint(shares[i]))
		}
	}
	return shares
}

func zeroOneString(length int, numberOfZeroes int) string {
	result := ""
	counter := numberOfZeroes
	for i := 0; i < length; i++ {
		if counter > 0 {
			result += "0"
			counter--
		} else {
			result += "1"
		}
	}
	return result
}
