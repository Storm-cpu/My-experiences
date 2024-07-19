package crypter

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"social_blog/internal/model"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/uniplaces/carbon"
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

// GenRefreshToken generate refresh token
func (s *Service) GenRefreshToken(secret string) (string, error) {
	now, err := carbon.NowInLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		return "", err
	}

	endOfDay := now.AddDays(1)
	endOfDay.SetHour(3)
	endOfDay.SetMinute(0)
	endOfDay.SetSecond(0)

	data := model.RefreshToken{
		ExpiredAt: endOfDay.Time.Unix(),
	}
	dataStr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	encryptedData, err := s.EncryptTripleDES([]byte(secret), string(dataStr))
	if err != nil {
		return "", err
	}

	if encryptedData == nil {
		return "", errors.New("encrypted data is nil")
	}

	return *encryptedData, nil
}

// ValidateRefreshToken validate refresh token
func (s *Service) ValidateRefreshToken(token, secret string) bool {
	d, err := s.DecryptTripleDES([]byte(secret), token)
	if err != nil {
		fmt.Println("===== err Decrypt", err.Error())
	}

	fmt.Println("==== d", *d)

	result := new(model.RefreshToken)
	if err := json.Unmarshal([]byte(*d), result); err != nil {
		fmt.Println("===== err Unmarshal", err.Error())
		return false
	}

	t := time.Unix(result.ExpiredAt, 0)

	return t.After(time.Now())
}

// EncryptTripleDES encrypt a string
func (*Service) EncryptTripleDES(key []byte, text string) (*string, error) {
	plaintext := []byte(text)

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, des.BlockSize+len(plaintext))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[des.BlockSize:], plaintext)

	result := base64.URLEncoding.EncodeToString(ciphertext)
	return &result, nil
}

// DecryptTripleDES decrypt a string
func (*Service) DecryptTripleDES(key []byte, cryptoText string) (*string, error) {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:des.BlockSize]
	ciphertext = ciphertext[des.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	result := string(ciphertext)
	return &result, nil
}
