package crypto

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"path/filepath"
	"time"
)

func GenerateUniqueFileName(originalName string) string {
	timestamp := time.Now().UnixNano()
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	randomPart := hex.EncodeToString(randomBytes)
	ext := filepath.Ext(originalName)
	return fmt.Sprintf("%d-%s%s", timestamp, randomPart, ext)
}

func GenerateJobID(sessionId uint) string {
	return fmt.Sprintf("%d-%d-%d", sessionId, time.Now().UnixNano(), rand.Intn(1000))
}
