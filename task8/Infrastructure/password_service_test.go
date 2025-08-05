package infrastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/abrishk26/a2sv-project-track/task8/Domain"
)

func TestPasswordService_HashAndVerify(t *testing.T) {
	ps := NewPasswordService()

	password := "mySecret123!"

	// Test hashing
	hash, err := ps.Hash(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, password, hash) // hash should not be the same as password

	// Test successful verification
	err = ps.Verify(password, hash)
	assert.NoError(t, err)

	// Test verification failure on wrong password
	err = ps.Verify("wrongPassword", hash)
	assert.ErrorIs(t, err, domain.ErrPasswordVerificationFailed)

	// Test hashing empty password (bcrypt allows it)
	emptyHash, err := ps.Hash("")
	assert.NoError(t, err)
	assert.NotEmpty(t, emptyHash)

	// Verify empty password against its hash
	err = ps.Verify("", emptyHash)
	assert.NoError(t, err)
}
