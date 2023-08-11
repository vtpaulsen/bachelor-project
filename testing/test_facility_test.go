package secretSharingSchemes

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	. "gitlab.au.dk/vtpaulsen/bachelor/v2/src/secretSharingSchemes"
)

// Testcase for Shamir Secret Sharing
// This function generate random values for t and n, and then generate a random key.
// The shares used to generate are in the interval [t, n].
func TestShamir(t *testing.T) {
	cases := []struct {
		t              int
		n              int
		input          []byte
		expectedOutput []byte
	}{}

	for i := 0; i < 10; i++ {
		var min = 2
		var max = 1000

		rand.Seed(time.Now().UnixNano())
		n := min + rand.Intn(max-min+1)

		t := min + rand.Intn((n-1)-min+1)

		key := GenerateAESKey()
		cases = append(cases, struct {
			t              int
			n              int
			input          []byte
			expectedOutput []byte
		}{t: t, n: n, input: key, expectedOutput: key})
	}
	for _, c := range cases {
		scheme := &ShamirSecretSharing{}

		shares := scheme.SecretShare(c.input, c.t, c.n)

		if !(len(shares) > 0) {
			t.Errorf("No shares were created")
		}

		reconstructed := scheme.SecretReconstruct(shares[0 : c.t+1])

		for i := 0; i < len(reconstructed); i++ {
			if !bytes.Equal(reconstructed, c.expectedOutput) {
				t.Errorf("SecretReconstruct() == %v, want %v", reconstructed, c.expectedOutput)
			}
		}
	}
}

// Testcase for Shamir Secret Sharing
// This function generate random values for t and n, and then generate a random key.
// The shares used to generate are chosen random from shares.
func TestShamirWithRandomShares(t *testing.T) {
	cases := []struct {
		t              int
		n              int
		input          []byte
		expectedOutput []byte
	}{}

	for i := 0; i < 10; i++ {
		var min = 2
		var max = 1000

		rand.Seed(time.Now().UnixNano())
		n := min + rand.Intn(max-min+1)

		t := min + rand.Intn((n-1)-min+1)

		key := GenerateAESKey()
		cases = append(cases, struct {
			t              int
			n              int
			input          []byte
			expectedOutput []byte
		}{t: t, n: n, input: key, expectedOutput: key})
	}

	for _, c := range cases {
		scheme := &ShamirSecretSharing{}

		shares := scheme.SecretShare(c.input, c.t, c.n)

		if !(len(shares) > 0) {
			t.Errorf("No shares were created")
		}

		randomShares := make([][]Share, c.t+1)

		for i := 0; i < c.t+1; i++ {
			randomShares[i] = shares[rand.Intn(c.n)]
		}

		reconstructed := scheme.SecretReconstruct(shares)

		for i := 0; i < len(reconstructed); i++ {
			if !bytes.Equal(reconstructed, c.expectedOutput) {
				t.Errorf("SecretReconstruct() == %v, want %v", reconstructed, c.expectedOutput)
			}
		}
	}
}

// Testcase for t-out-of-n Secret Sharing
// This function generate random values for t and n, and then generate a random key.
// The shares used to generate are in the interval [t, n].
func TestReplicatedSecretSharing(t *testing.T) {
	cases := []struct {
		t              int
		n              int
		input          []byte
		expectedOutput []byte
	}{}

	for i := 0; i < 10; i++ {
		var min = 2
		var max = 11

		rand.Seed(time.Now().UnixNano())
		n := min + rand.Intn(max-min+1)

		t := min + rand.Intn(n-min+1)

		key := GenerateAESKey()
		cases = append(cases, struct {
			t              int
			n              int
			input          []byte
			expectedOutput []byte
		}{t: t, n: n, input: key, expectedOutput: key})
	}

	for _, c := range cases {
		scheme := &ReplicatedSecretSharing{}
		shares := scheme.SecretShare(c.input, c.t, c.n)

		if !(len(shares) > 0) {
			t.Errorf("No shares were created")
		}

		reconstructed := scheme.SecretReconstruct(shares[0:c.t])

		for i := 0; i < len(reconstructed); i++ {
			if !bytes.Equal(reconstructed, c.expectedOutput) {
				t.Errorf("SecretReconstruct() == %v, want %v", reconstructed, c.expectedOutput)
			}
		}
	}

}

// Testcase for t-out-of-n Secret Sharing
// This function generate random values for t and n, and then generate a random key.
// The shares used to generate are chosen random from shares.
func TestReplicatedWithRandomShares(t *testing.T) {
	cases := []struct {
		t              int
		n              int
		input          []byte
		expectedOutput []byte
	}{}

	for i := 0; i < 10; i++ {
		var min = 2
		var max = 11

		rand.Seed(time.Now().UnixNano())
		n := min + rand.Intn(max-min+1)

		t := min + rand.Intn(n-min+1)

		key := GenerateAESKey()
		cases = append(cases, struct {
			t              int
			n              int
			input          []byte
			expectedOutput []byte
		}{t: t, n: n, input: key, expectedOutput: key})
	}

	for _, c := range cases {
		scheme := &ReplicatedSecretSharing{}
		shares := scheme.SecretShare(c.input, c.t, c.n)

		if !(len(shares) > 0) {
			t.Errorf("No shares were created")
		}

		randomShares := make([][]Server, c.t+1)

		for i := 0; i < c.t+1; i++ {
			randomShares[i] = shares[rand.Intn(c.n)]
		}

		reconstructed := scheme.SecretReconstruct(shares[0:c.t])

		for i := 0; i < len(reconstructed); i++ {
			if !bytes.Equal(reconstructed, c.expectedOutput) {
				t.Errorf("SecretReconstruct() == %v, want %v", reconstructed, c.expectedOutput)
			}
		}
	}
}

// Testcase for n-out-of-n Secret Sharing
// This function generate random values for t and n, and then generate a random key.
// The shares used to generate are in the interval [t, n].
func TestAdditiveSecretSharing(t *testing.T) {
	cases := []struct {
		n              int
		input          []byte
		expectedOutput []byte
	}{}

	for i := 0; i < 10; i++ {
		var min = 2
		var max = 1000

		rand.Seed(time.Now().UnixNano())
		n := min + rand.Intn(max-min+1)

		key := GenerateAESKey()
		cases = append(cases, struct {
			n              int
			input          []byte
			expectedOutput []byte
		}{n: n, input: key, expectedOutput: key})
	}

	for _, c := range cases {
		scheme := &AdditiveSecretSharing{}
		shares := scheme.SecretShare(c.input, c.n, c.n)

		if !(len(shares) > 0) {
			t.Errorf("No shares were created")
		}

		reconstructed := scheme.SecretReconstruct(shares[0:c.n])

		for i := 0; i < len(reconstructed); i++ {
			if !bytes.Equal(reconstructed, c.expectedOutput) {
				t.Errorf("SecretReconstruct() == %v, want %v", reconstructed, c.expectedOutput)
			}
		}
	}
}

func Test2NSecretSharing(t *testing.T) {
	cases := []struct {
		n              int
		input          []byte
		expectedOutput []byte
	}{}

	for i := 0; i < 10; i++ {
		var min = 2
		var max = 1000

		rand.Seed(time.Now().UnixNano())
		n := min + rand.Intn(max-min+1)

		key := GenerateAESKey()
		cases = append(cases, struct {
			n              int
			input          []byte
			expectedOutput []byte
		}{n: n, input: key, expectedOutput: key})
	}

	for _, c := range cases {
		scheme := &TwoNSecretSharing{}
		shares := scheme.SecretShare(string(c.input), c.n)

		if !(len(shares) > 0) {
			t.Errorf("No shares were created")
		}

		reconstructed := scheme.SecretReconstruct(shares[0], shares[1])

		for i := 0; i < len(reconstructed); i++ {
			if !(strings.Compare(string(reconstructed), string(c.expectedOutput)) == 0) {
				t.Errorf("SecretReconstruct() == %v, want %v", string(reconstructed), string(c.expectedOutput))
			}
		}
	}
}

func Benchmark_Replicated(b *testing.B) {
	cases := []struct {
		t int
		n int
	}{{3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}, {10, 10}, {11, 11}}

	scheme := &ReplicatedSecretSharing{}

	for _, c := range cases {
		b.Run(fmt.Sprintf("%d-out-of-%d", c.t, c.n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				key := GenerateAESKey()
				shares := scheme.SecretShare(key, c.t, c.n)
				reconstruction := scheme.SecretReconstruct(shares[0:c.t])

				if !bytes.Equal(key, reconstruction) {
					b.Errorf("SecretReconstruct() == %v, want %v", reconstruction, key)
				}
			}
		})
	}

}

func Benchmark_Additive(b *testing.B) {
	cases := []struct {
		t int
		n int
	}{{3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}, {10, 10}, {11, 11}}

	scheme := &AdditiveSecretSharing{}

	for _, c := range cases {
		b.Run(fmt.Sprintf("%d-out-of-%d", c.t, c.n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				key := GenerateAESKey()
				shares := scheme.SecretShare(key, c.t, c.n)
				reconstruction := scheme.SecretReconstruct(shares[0:c.t])

				if !bytes.Equal(key, reconstruction) {
					b.Errorf("SecretReconstruct() == %v, want %v", reconstruction, key)
				}

			}
		})
	}

}

func Benchmark_Shamir(b *testing.B) {
	cases := []struct {
		t int
		n int
	}{{3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}, {10, 10}, {11, 11}}

	scheme := &ShamirSecretSharing{}

	for _, c := range cases {
		b.Run(fmt.Sprintf("%d-out-of-%d", c.t, c.n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				key := GenerateAESKey()
				shares := scheme.SecretShare(key, c.t-1, c.n)
				reconstruction := scheme.SecretReconstruct(shares[0:c.t])

				if !bytes.Equal(key, reconstruction) {
					b.Errorf("SecretReconstruct() == %v, want %v", reconstruction, key)
				}

			}
		})
	}
}
