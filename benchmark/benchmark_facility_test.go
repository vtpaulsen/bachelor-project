package secretSharingSchemes

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	. "gitlab.au.dk/vtpaulsen/bachelor/v2/src/secretSharingSchemes"
)

type Runtime struct {
	SecretSharingScheme string
	Algorithm           string
	T, N, Time          int
}

/**
 * This is the test-cases for timing all sharing schemes' sharing phase
 */
func Test_Runtime_for_shamir_sharing_phase(t *testing.T) {
	cases := []struct {
		t int
		n int
	}{{1, 3}, {1, 4}, {2, 5}, {2, 6}, {3, 7}, {3, 8}, {4, 9}, {4, 10}, {5, 11}}

	scheme := &ShamirSecretSharing{}

	result := make([]Runtime, 0)

	for _, c := range cases {
		start := time.Now()
		for i := 0; i < 10; i++ {
			key := GenerateAESKey()
			_ = scheme.SecretShare(key, c.t, c.n)
		}
		elapsed := time.Since(start)
		time := elapsed.Nanoseconds()

		data := Runtime{
			SecretSharingScheme: "Shamir",
			Algorithm:           "SecretShare",
			T:                   c.t + 1,
			N:                   c.n,
			Time:                int(time / 10),
		}

		result = append(result, data)
	}

	// file, _ := json.MarshalIndent(result, "", " ")

	// _ = ioutil.WriteFile("json/all/shamir_sharing.json", file, 0644)
}

func Test_Runtime_for_replicated_sharing_phase(t *testing.T) {
	cases := []struct {
		t int
		n int
	}{{2, 3}, {2, 4}, {3, 5}, {3, 6}, {4, 7}, {4, 8}, {5, 9}, {5, 10}, {6, 11}}

	scheme := &ReplicatedSecretSharing{}

	result := make([]Runtime, 0)

	for _, c := range cases {
		start := time.Now()
		for i := 0; i < 10; i++ {
			key := GenerateAESKey()
			_ = scheme.SecretShare(key, c.t, c.n)
		}
		elapsed := time.Since(start)
		time := elapsed.Nanoseconds()

		data := Runtime{
			SecretSharingScheme: "Replicated",
			Algorithm:           "SecretShare",
			T:                   c.t,
			N:                   c.n,
			Time:                int(time / 10),
		}

		result = append(result, data)
	}

	// file, _ := json.MarshalIndent(result, "", " ")

	// _ = ioutil.WriteFile("json/all/replicated_sharing.json", file, 0644)
}

func Test_Runtime_for_additive_sharing_phase(t *testing.T) {
	cases := []struct {
		t int
		n int
	}{{3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}, {10, 10}, {11, 11}}

	scheme := &AdditiveSecretSharing{}

	result := make([]Runtime, 0)

	for _, c := range cases {
		start := time.Now()
		for i := 0; i < 10; i++ {
			key := GenerateAESKey()
			_ = scheme.SecretShare(key, c.t, c.n)
		}
		elapsed := time.Since(start)
		time := elapsed.Nanoseconds()

		data := Runtime{
			SecretSharingScheme: "Additive",
			Algorithm:           "SecretShare",
			T:                   c.t,
			N:                   c.n,
			Time:                int(time / 10),
		}

		result = append(result, data)
	}

	// file, _ := json.MarshalIndent(result, "", " ")

	// _ = ioutil.WriteFile("json/all/additive_share.json", file, 0644)
}

/**
 * This is the test-cases for timing all sharing schemes' reconstion phase
 */
func Test_Runtime_for_shamir_reconstruction_phase(t *testing.T) {
	cases := []struct {
		t int
		n int
	}{{1, 3}, {1, 4}, {2, 5}, {2, 6}, {3, 7}, {3, 8}, {4, 9}, {4, 10}, {5, 11}}

	scheme := &ShamirSecretSharing{}

	result := make([]Runtime, 0)

	for _, c := range cases {
		key := GenerateAESKey()
		shares := scheme.SecretShare(key, c.t, c.n)
		start := time.Now()
		for i := 0; i < 10; i++ {
			_ = scheme.SecretReconstruct(shares[0 : c.t+1])
		}
		elapsed := time.Since(start)
		time := elapsed.Nanoseconds()

		data := Runtime{
			SecretSharingScheme: "Shamir",
			Algorithm:           "SecretReconstruction",
			T:                   c.t + 1,
			N:                   c.n,
			Time:                int(time / 10),
		}

		result = append(result, data)
	}

	// file, _ := json.MarshalIndent(result, "", " ")

	// _ = ioutil.WriteFile("json/all/shamir_reconstruction.json", file, 0644)
}

func Test_Runtime_for_replicated_reconstruction_phase(t *testing.T) {
	cases := []struct {
		t int
		n int
	}{{2, 3}, {2, 4}, {3, 5}, {3, 6}, {4, 7}, {4, 8}, {5, 9}, {5, 10}, {6, 11}}

	scheme := &ReplicatedSecretSharing{}

	result := make([]Runtime, 0)

	for _, c := range cases {
		key := GenerateAESKey()
		shares := scheme.SecretShare(key, c.t, c.n)
		start := time.Now()
		for i := 0; i < 10; i++ {
			_ = scheme.SecretReconstruct(shares[0:c.t])
		}
		elapsed := time.Since(start)
		time := elapsed.Nanoseconds()

		data := Runtime{
			SecretSharingScheme: "Replicated",
			Algorithm:           "SecretReconstruction",
			T:                   c.t,
			N:                   c.n,
			Time:                int(time / 10),
		}

		result = append(result, data)
	}

	// file, _ := json.MarshalIndent(result, "", " ")

	// _ = ioutil.WriteFile("json/all/replicated_reconstruction.json", file, 0644)
}

func Test_Runtime_for_additive_reconstruction_phase(t *testing.T) {
	cases := []struct {
		t int
		n int
	}{{3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9}, {10, 10}, {11, 11}}

	scheme := &AdditiveSecretSharing{}

	result := make([]Runtime, 0)

	for _, c := range cases {
		key := GenerateAESKey()
		shares := scheme.SecretShare(key, c.t, c.n)
		start := time.Now()
		for i := 0; i < 10; i++ {
			_ = scheme.SecretReconstruct(shares[0:c.t])
		}
		elapsed := time.Since(start)
		time := elapsed.Nanoseconds()

		data := Runtime{
			SecretSharingScheme: "Additive",
			Algorithm:           "SecretReconstruction",
			T:                   c.t,
			N:                   c.n,
			Time:                int(time / 10),
		}

		result = append(result, data)
	}

	// file, _ := json.MarshalIndent(result, "", " ")

	// _ = ioutil.WriteFile("json/all/additive_reconstruction.json", file, 0644)
}

/**
 * This is the test-cases for benchmarking all sharing schemes
 */
func Benchmark_shamir(b *testing.B) {
	cases := []struct {
		t int
		n int
	}{{1, 3}, {1, 4}, {2, 5}, {2, 6}, {3, 7}, {3, 8}, {4, 9}, {4, 10}, {5, 11}}

	scheme := &ShamirSecretSharing{}

	for _, c := range cases {
		b.Run(fmt.Sprintf("%d-out-of-%d", c.t+1, c.n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				key := GenerateAESKey()
				shares := scheme.SecretShare(key, c.t, c.n)
				reconstruction := scheme.SecretReconstruct(shares[0 : c.t+1])

				b.StopTimer()
				if !bytes.Equal(key, reconstruction) {
					b.Errorf("SecretReconstruct() == %v, want %v", reconstruction, key)
				}
				b.StartTimer()
			}
		})
	}
}

func Benchmark_replicated(b *testing.B) {
	cases := []struct {
		t int
		n int
	}{{2, 3}, {2, 4}, {3, 5}, {3, 6}, {4, 7}, {4, 8}, {5, 9}, {5, 10}, {6, 11}}

	scheme := &ReplicatedSecretSharing{}

	for _, c := range cases {
		b.Run(fmt.Sprintf("%d-out-of-%d", c.t, c.n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				key := GenerateAESKey()
				shares := scheme.SecretShare(key, c.t, c.n)
				reconstruction := scheme.SecretReconstruct(shares[0:c.t])

				b.StopTimer()
				if !bytes.Equal(key, reconstruction) {
					b.Errorf("SecretReconstruct() == %v, want %v", reconstruction, key)
				}
				b.StartTimer()
			}
		})
	}
}

func Benchmark_additive(b *testing.B) {
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

				b.StopTimer()
				if !bytes.Equal(key, reconstruction) {
					b.Errorf("SecretReconstruct() == %v, want %v", reconstruction, key)
				}
				b.StartTimer()
			}
		})
	}
}

/**
 * This is the test-cases for benchmarking threshold sharing schemes
 */
func Test_timedShamirShare(t *testing.T) {
	cases := []struct {
		t int
		n int
	}{{1, 3}, {2, 3}, {2, 4}, {3, 4}, {1, 5}, {2, 5}, {3, 5}, {4, 5}, {1, 6}, {2, 6}, {3, 6}, {4, 6},
		{5, 6}, {1, 7}, {2, 7}, {3, 7}, {4, 7}, {5, 7}, {6, 7}, {1, 8}, {2, 8}, {3, 8}, {4, 8},
		{5, 8}, {6, 8}, {7, 8}, {1, 9}, {2, 9}, {3, 9}, {4, 9}, {5, 9}, {6, 9}, {7, 9},
		{8, 9}, {1, 10}, {2, 10}, {3, 10}, {4, 10}, {5, 10}, {6, 10}, {7, 10}, {8, 10},
		{9, 10}, {1, 11}, {2, 11}, {3, 11}, {4, 11}, {5, 11}, {6, 11}, {7, 11},
		{8, 11}, {9, 11}, {10, 11}}

	scheme := &ShamirSecretSharing{}

	results := make([]Runtime, 0)

	for _, c := range cases {
		start := time.Now()
		for i := 0; i < 10; i++ {
			key := GenerateAESKey()
			_ = scheme.SecretShare(key, c.t, c.n)
		}
		elapsed := time.Since(start)
		time := elapsed.Nanoseconds()

		data := Runtime{
			SecretSharingScheme: "Shamir",
			Algorithm:           "SecretShare",
			T:                   c.t + 1,
			N:                   c.n,
			Time:                int(time / 10),
		}

		results = append(results, data)
	}

	// file, _ := json.MarshalIndent(results, "", " ")

	// _ = ioutil.WriteFile("json/threshold/shamir_sharing.json", file, 0644)
}

func Test_timedReplicatedShare(t *testing.T) {
	cases := []struct {
		t int
		n int
	}{{2, 3}, {3, 3}, {3, 4}, {4, 4}, {2, 5}, {3, 5}, {4, 5}, {5, 5}, {2, 6}, {3, 6}, {4, 6}, {5, 6},
		{6, 6}, {2, 7}, {3, 7}, {4, 7}, {5, 7}, {6, 7}, {7, 7}, {2, 8}, {3, 8}, {4, 8}, {5, 8},
		{6, 8}, {7, 8}, {8, 8}, {2, 9}, {3, 9}, {4, 9}, {5, 9}, {6, 9}, {7, 9}, {8, 9},
		{9, 9}, {2, 10}, {3, 10}, {4, 10}, {5, 10}, {6, 10}, {7, 10}, {8, 10}, {9, 10},
		{10, 10}, {2, 11}, {3, 11}, {4, 11}, {5, 11}, {6, 11}, {7, 11}, {8, 11},
		{9, 11}, {10, 11}, {11, 11}}

	scheme := &ReplicatedSecretSharing{}

	arr := make([]Runtime, 0)

	for _, c := range cases {
		start := time.Now()
		for i := 0; i < 10; i++ {
			key := GenerateAESKey()
			_ = scheme.SecretShare(key, c.t, c.n)
		}
		elapsed := time.Since(start)
		nano := elapsed.Nanoseconds()

		data := Runtime{
			SecretSharingScheme: "Replicated",
			Algorithm:           "SecretShare",
			T:                   c.t,
			N:                   c.n,
			Time:                int(nano / 10),
		}

		arr = append(arr, data)
	}

	// file, _ := json.MarshalIndent(arr, "", " ")

	// _ = ioutil.WriteFile("json/threshold/replicated_sharing.json", file, 0644)
}

func Test_timedShamirReconstruction(t *testing.T) {
	cases := []struct {
		t int
		n int
	}{{1, 3}, {2, 3}, {2, 4}, {3, 4}, {1, 5}, {2, 5}, {3, 5}, {4, 5}, {1, 6}, {2, 6}, {3, 6}, {4, 6},
		{5, 6}, {1, 7}, {2, 7}, {3, 7}, {4, 7}, {5, 7}, {6, 7}, {1, 8}, {2, 8}, {3, 8}, {4, 8},
		{5, 8}, {6, 8}, {7, 8}, {1, 9}, {2, 9}, {3, 9}, {4, 9}, {5, 9}, {6, 9}, {7, 9},
		{8, 9}, {1, 10}, {2, 10}, {3, 10}, {4, 10}, {5, 10}, {6, 10}, {7, 10}, {8, 10},
		{9, 10}, {1, 11}, {2, 11}, {3, 11}, {4, 11}, {5, 11}, {6, 11}, {7, 11},
		{8, 11}, {9, 11}, {10, 11}}

	scheme := &ShamirSecretSharing{}

	result := make([]Runtime, 0)

	for _, c := range cases {
		key := GenerateAESKey()
		shares := scheme.SecretShare(key, c.t, c.n)
		start := time.Now()
		for i := 0; i < 10; i++ {
			_ = scheme.SecretReconstruct(shares[0 : c.t+1])
		}
		elapsed := time.Since(start)
		time := elapsed.Nanoseconds()

		data := Runtime{
			SecretSharingScheme: "Shamir",
			Algorithm:           "SecretReconstruct",
			T:                   c.t,
			N:                   c.n,
			Time:                int(time / 10),
		}

		result = append(result, data)
	}

	// file, _ := json.MarshalIndent(result, "", " ")

	// _ = ioutil.WriteFile("json/threshold/shamir_reconstruction.json", file, 0644)
}

func Test_timedReplicatedReconstruction(t *testing.T) {
	cases := []struct {
		t int
		n int
	}{{2, 3}, {3, 3}, {3, 4}, {4, 4}, {2, 5}, {3, 5}, {4, 5}, {5, 5}, {2, 6}, {3, 6}, {4, 6}, {5, 6},
		{6, 6}, {2, 7}, {3, 7}, {4, 7}, {5, 7}, {6, 7}, {7, 7}, {2, 8}, {3, 8}, {4, 8}, {5, 8},
		{6, 8}, {7, 8}, {8, 8}, {2, 9}, {3, 9}, {4, 9}, {5, 9}, {6, 9}, {7, 9}, {8, 9},
		{9, 9}, {2, 10}, {3, 10}, {4, 10}, {5, 10}, {6, 10}, {7, 10}, {8, 10}, {9, 10},
		{10, 10}, {2, 11}, {3, 11}, {4, 11}, {5, 11}, {6, 11}, {7, 11}, {8, 11},
		{9, 11}, {10, 11}, {11, 11}}

	scheme := &ReplicatedSecretSharing{}

	result := make([]Runtime, 0)

	for _, c := range cases {
		key := GenerateAESKey()
		shares := scheme.SecretShare(key, c.t, c.n)
		start := time.Now()
		for i := 0; i < 10; i++ {
			_ = scheme.SecretReconstruct(shares[0:c.t])
		}
		elapsed := time.Since(start)
		time := elapsed.Nanoseconds()

		data := Runtime{
			SecretSharingScheme: "Additive",
			Algorithm:           "SecretReconstruct",
			T:                   c.t,
			N:                   c.n,
			Time:                int(time / 10),
		}

		result = append(result, data)
	}

	// file, _ := json.MarshalIndent(result, "", " ")

	// _ = ioutil.WriteFile("json/threshold/replicated_reconstruction.json", file, 0644)
}

/**
 * This is the test-cases for benchmarking all sharing schemes
 */
func Benchmark_shamir_threshold(b *testing.B) {
	cases := []struct {
		t int
		n int
	}{{1, 3}, {2, 3}, {2, 4}, {3, 4}, {1, 5}, {2, 5}, {3, 5}, {4, 5}, {1, 6}, {2, 6}, {3, 6}, {4, 6},
		{5, 6}, {1, 7}, {2, 7}, {3, 7}, {4, 7}, {5, 7}, {6, 7}, {1, 8}, {2, 8}, {3, 8}, {4, 8},
		{5, 8}, {6, 8}, {7, 8}, {1, 9}, {2, 9}, {3, 9}, {4, 9}, {5, 9}, {6, 9}, {7, 9},
		{8, 9}, {1, 10}, {2, 10}, {3, 10}, {4, 10}, {5, 10}, {6, 10}, {7, 10}, {8, 10},
		{9, 10}, {1, 11}, {2, 11}, {3, 11}, {4, 11}, {5, 11}, {6, 11}, {7, 11},
		{8, 11}, {9, 11}, {10, 11}}

	scheme := &ShamirSecretSharing{}

	for _, c := range cases {
		b.Run(fmt.Sprintf("%d-out-of-%d", c.t+2, c.n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				key := GenerateAESKey()
				shares := scheme.SecretShare(key, c.t, c.n)
				reconstruction := scheme.SecretReconstruct(shares[0 : c.t+1])

				b.StopTimer()
				if !bytes.Equal(key, reconstruction) {
					b.Errorf("SecretReconstruct() == %v, want %v", reconstruction, key)
				}
				b.StartTimer()
			}
		})
	}
}

func Benchmark_replicated_threshold(b *testing.B) {
	cases := []struct {
		t int
		n int
	}{{2, 3}, {3, 3}, {3, 4}, {4, 4}, {2, 5}, {3, 5}, {4, 5}, {5, 5}, {2, 6}, {3, 6}, {4, 6}, {5, 6},
		{6, 6}, {2, 7}, {3, 7}, {4, 7}, {5, 7}, {6, 7}, {7, 7}, {2, 8}, {3, 8}, {4, 8}, {5, 8},
		{6, 8}, {7, 8}, {8, 8}, {2, 9}, {3, 9}, {4, 9}, {5, 9}, {6, 9}, {7, 9}, {8, 9},
		{9, 9}, {2, 10}, {3, 10}, {4, 10}, {5, 10}, {6, 10}, {7, 10}, {8, 10}, {9, 10},
		{10, 10}, {2, 11}, {3, 11}, {4, 11}, {5, 11}, {6, 11}, {7, 11}, {8, 11},
		{9, 11}, {10, 11}, {11, 11}}

	scheme := &ReplicatedSecretSharing{}

	for _, c := range cases {
		b.Run(fmt.Sprintf("%d-out-of-%d", c.t+2, c.n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				key := GenerateAESKey()
				shares := scheme.SecretShare(key, c.t, c.n)
				reconstruction := scheme.SecretReconstruct(shares[0:c.t])

				b.StopTimer()
				if !bytes.Equal(key, reconstruction) {
					b.Errorf("SecretReconstruct() == %v, want %v", reconstruction, key)
				}
				b.StartTimer()
			}
		})
	}
}
