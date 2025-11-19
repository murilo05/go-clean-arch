package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go-clean-arch/internal/core/domain"
)

func BuildIdempotencyKey(user *domain.User) {
	base := fmt.Sprintf("%s:%s:%s", user.Document, user.Name, user.Email)

	sum := sha256.Sum256([]byte(base))

	user.ID = hex.EncodeToString(sum[:])
}
