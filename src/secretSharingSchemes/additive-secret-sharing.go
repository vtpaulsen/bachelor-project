package secretSharingSchemes

import (
	"bytes"
	"encoding/binary"
	"math/bits"
)

type AdditiveSecretSharing struct {
	// Create a t-out-of-n secret sharing scheme object
}

func (l *AdditiveSecretSharing) SecretShare(secret []byte, n int, t int) [][]Share {
	var values = make([][]int, 8)
	for i := 0; i < 8; i++ {
		tmp := secret[i*2 : (i+1)*2]
		secret := int(binary.BigEndian.Uint16(tmp))
		values[i] = createAdditiveShares(secret, n, t)

	}
	values = Transpose(values)
	return ToShares(values)
}

func (l *AdditiveSecretSharing) SecretReconstruct(sharesArray [][]Share) []byte {
	CheckArgument(len(sharesArray) > 0, "Invalid argument: No shares provided")

	reconstruction := new(bytes.Buffer)
	for i := 0; i < 8; i++ {
		shares := make([]Share, len(sharesArray))
		for j := 0; j < len(sharesArray); j++ {
			shares[j] = sharesArray[j][i]
		}
		secret := shares[0].Share
		for j := 1; j < len(sharesArray); j++ {
			secret = secret ^ shares[j].Share
		}

		err := (binary.Write(reconstruction,
			binary.BigEndian, uint16(secret)))
		CheckError(err, "The secret was not in the range of 0 through 65535")
	}
	return reconstruction.Bytes()
}

func createAdditiveShares(secret int, n int, t int) []int {
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
