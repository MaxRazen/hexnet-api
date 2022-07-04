package hasher

import (
	"testing"
)

var password = "pas@swoRd0* $"

func TestHash(t *testing.T) {
	salt, _ := GenerateSaltBytes(32)
	hash, err := Hash(password, salt)

	if err != nil {
		t.Error("Hashing password failed")
	}

	if len(hash) == 0 {
		t.Error("len(encoded) must be > 0")
	}

	for _, b := range hash {
		if b == 0 {
			t.Error("encoded must not contain 0x00")
		}
	}
}

func TestGenerateSaltBytes(t *testing.T) {
	salt, err := GenerateSaltBytes(32)

	if err != nil {
		t.Error("Salt could not be generated: " + err.Error())
	}

	if len(salt) != 32 {
		t.Error("Salt length is wrong")
	}
}

func TestVerifyHash(t *testing.T) {
	salt, _ := GenerateSaltBytes(32)
	hash, _ := Hash(password, salt)

	ok, err := VerifyHash(password, string(hash))

	if !ok || err != nil {
		t.Error("Truth case failed")
	}

	ok, err = VerifyHash("wrong-password", string(hash))

	if ok || err != nil {
		t.Error("False case failed")
	}
}
