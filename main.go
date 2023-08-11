package main

import (
	"fmt"

	. "gitlab.au.dk/vtpaulsen/bachelor/v2/src/encryption"
	. "gitlab.au.dk/vtpaulsen/bachelor/v2/src/secretSharingSchemes"
)

func main() {
	(fmt.Println("The following examples will show how the implemented secret-sharing schemes works with a threshold\n of 4 out of 9 for both Shamir's secret-sharing and Replicated secret-sharing.\nFor additive secret-sharing the threshold is 9 out of 9."))
	/**
	 *	This is the setup for the encryption and secret sharing
	 */
	fmt.Println("\nThis is an example of how the encryption and secret sharing works for Shamir's secret-sharing")
	var key = GenerateAESKey()
	var message = "*Secret Message*"
	fmt.Println("\nThe key is:", key)
	fmt.Println("The message is:", message)

	fmt.Println("\nEncryption-phase begins..")
	aes := &AES{}
	var mac, cipher = aes.Encrypt(key, message)
	fmt.Println("\nThe mac is:", mac)
	fmt.Println("The cipher is:", cipher)

	/**
	 *	This is an example with Shamir's secret-sharing
	 */
	fmt.Println("\nSharing-phase begins..")
	shamir := &ShamirSecretSharing{}
	shamir_shares := shamir.SecretShare(key, 4, 9)
	fmt.Println("\nSome of the shares are:", shamir_shares[0][:])

	fmt.Println("\nReconstruction-phase for Shamir's secret-sharing begins with 4 parties..")
	shamir_reconstructed := shamir.SecretReconstruct(shamir_shares[0:5])
	fmt.Println("\nThe reconstructed key is:", shamir_reconstructed)

	fmt.Println("\nDecryption-phase begins with with 4 parties..")
	shamir_decrypted := aes.Decrypt(shamir_reconstructed, mac, cipher)
	fmt.Println("\nThe decrypted message is:", shamir_decrypted)

	fmt.Println("\nReconstruction-phase for Shamir's secret-sharing begins with 3 parties..")
	shamir_wrong_reconstruction := shamir.SecretReconstruct(shamir_shares[0:4])
	fmt.Println("\nThe reconstructed key is:", shamir_wrong_reconstruction)

	fmt.Println("\nDecryption-phase begins with 3 parties..")
	shamir_wrong_decrypted := aes.Decrypt(shamir_wrong_reconstruction, mac, cipher)
	fmt.Println("\nThe decrypted message is:", shamir_wrong_decrypted)

	/**
	 *	This is the setup for the encryption and secret sharing
	 */
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("\nThis is an example of how the encryption and secret sharing works for replicated secret-sharing")
	key = GenerateAESKey()
	message = "*Secret Message*"
	fmt.Println("\nThe key is:", key)
	fmt.Println("The message is:", message)

	fmt.Println("\nEncryption-phase begins..")
	mac, cipher = aes.Encrypt(key, message)
	fmt.Println("\nThe mac is:", mac)
	fmt.Println("The cipher is:", cipher)

	/**
	 *	This is an example with replicated secret-sharing
	 */
	fmt.Println("\nSharing-phase begins..")
	replicated := &ReplicatedSecretSharing{}
	replicated_shares := replicated.SecretShare(key, 4, 9)
	fmt.Println("\nSome of the shares are:", replicated_shares[0][1])

	fmt.Println("\nReconstruction-phase for replicated secret-sharing begins with 4 parties..")
	replicated_reconstructed := replicated.SecretReconstruct(replicated_shares[0:4])
	fmt.Println("\nThe reconstructed key is:", replicated_reconstructed)

	fmt.Println("\nDecryption-phase begins with 4 parties..")
	replicated_decrypted := aes.Decrypt(replicated_reconstructed, mac, cipher)
	fmt.Println("\nThe decrypted message is:", replicated_decrypted)

	fmt.Println("\nReconstruction-phase for replicated secret-sharing begins with 3 parties..")
	replicated_wrong_reconstruction := replicated.SecretReconstruct(replicated_shares[0:3])
	fmt.Println("\nThe reconstructed key is:", replicated_wrong_reconstruction)

	fmt.Println("\nDecryption-phase begins with 3 parties..")
	replicated_wrong_decrypted := aes.Decrypt(replicated_wrong_reconstruction, mac, cipher)
	fmt.Println("\nThe decrypted message is:", replicated_wrong_decrypted)

	/**
	 *	This is the setup for the encryption and secret sharing
	 */
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("\nThis is an example of how the encryption and secret sharing works for additive secret-sharing")
	key = GenerateAESKey()
	message = "*Secret Message*"
	fmt.Println("\nThe key is:", key)
	fmt.Println("The message is:", message)

	fmt.Println("\nEncryption-phase begins..")
	mac, cipher = aes.Encrypt(key, message)
	fmt.Println("\nThe mac is:", mac)
	fmt.Println("The cipher is:", cipher)

	/**
	 *	This is an example with additive secret-sharing
	 */
	fmt.Println("\nSharing-phase begins..")
	additive := &AdditiveSecretSharing{}
	additive_shares := additive.SecretShare(key, 4, 9)
	fmt.Println("\nSome of the shares are:", additive_shares[0][:])

	fmt.Println("\nReconstruction-phase for additive secret-sharing begins with 9 parties..")
	additive_reconstructed := additive.SecretReconstruct(additive_shares[0:4])
	fmt.Println("\nThe reconstructed key is:", additive_reconstructed)

	fmt.Println("\nDecryption-phase begins with 9 parties..")
	additive_decrypted := aes.Decrypt(additive_reconstructed, mac, cipher)
	fmt.Println("\nThe decrypted message is:", additive_decrypted)

	fmt.Println("\nReconstruction-phase for additive secret-sharing begins with 3 parties..")
	additive_wrong_reconstruction := additive.SecretReconstruct(additive_shares[0:3])
	fmt.Println("\nThe reconstructed key is:", additive_wrong_reconstruction)

	fmt.Println("\nDecryption-phase begins with 3 parties..")
	additive_wrong_decrypted := aes.Decrypt(additive_wrong_reconstruction, mac, cipher)
	fmt.Println("\nThe decrypted message is:", additive_wrong_decrypted)
}
