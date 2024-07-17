package crypter

import (
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

func New() *Service {
	return &Service{}
}

type Service struct{}

func (*Service) HashPassword(password string) string {
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPW)
}

func (*Service) CompareHashAndPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (*Service) UID() string {
	return ksuid.New().String()
}
