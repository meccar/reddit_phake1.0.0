package util

// import (
// 	"crypto/rand"
// 	"crypto/ed25519"
// 	"encoding/base64"
// 	"fmt"
// 	"os"
// 	"time"
// )

// func main() {
// 	// Start a goroutine to generate keys at 00:00 every day
// 	go func() {
// 		for {
// 			// Calculate the duration until 00:00 of the next day
// 			now := time.Now()
// 			next := now.Add(24 * time.Hour)
// 			nextMidnight := time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
// 			duration := nextMidnight.Sub(now)

// 			// Sleep until 00:00 of the next day
// 			time.Sleep(duration)

// 			// Generate and save keys
// 			err := generateAndSaveKeys()
// 			if err != nil {
// 				fmt.Println("Error generating and saving keys:", err)
// 			}
// 		}
// 	}()

// 	// Block the main goroutine
// 	select {}
// }

// func generateAndSaveKeys() error {
// 	// Generate a new private key
// 	privateKey, publicKey, err := ed25519.GenerateKey(rand.Reader)
// 	if err != nil {
// 		return fmt.Errorf("error generating key pair: %w", err)
// 	}

// 	// Encode keys to base64 strings
// 	privateKeyBase64 := base64.StdEncoding.EncodeToString(privateKey)
// 	publicKeyBase64 := base64.StdEncoding.EncodeToString(publicKey)

// 	// Write keys to environment file
// 	err = writeKeysToEnvFile(privateKeyBase64, publicKeyBase64)
// 	if err != nil {
// 		return fmt.Errorf("error writing keys to environment file: %w", err)
// 	}

// 	fmt.Println("Keys generated and saved successfully")
// 	return nil
// }

// func writeKeysToEnvFile(privateKeyBase64, publicKeyBase64 string) error {
// 	// Open or create the environment file
// 	file, err := os.OpenFile("app.env", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Write the keys to the file
// 	_, err = file.WriteString("PRIVATE_KEY=" + privateKeyBase64 + "\n")
// 	if err != nil {
// 		return err
// 	}
// 	_, err = file.WriteString("PUBLIC_KEY=" + publicKeyBase64 + "\n")
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }