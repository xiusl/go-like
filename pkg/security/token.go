package security

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	AuthorizationHeaderKey = "Authorization"
	ContextUserKey         = "authorization_user"
)

var duration = time.Hour * 24 * 7
var tokenErr = fmt.Errorf("令牌无效")
var defaultPwd = "klsajdionaksndsmkdl"

func hmacSha256(key, message string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(message))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

func GenerateToken(uid int64, password string) string {
	if password == "" {
		password = defaultPwd
	}
	exp := time.Now().Add(duration).Unix()
	msg := fmt.Sprintf("%d:%d", uid, exp)
	sign := hmacSha256(password, msg)
	signMsg := fmt.Sprintf("%s$%s", msg, sign)
	token := base64.StdEncoding.EncodeToString([]byte(signMsg))
	return token
}
func VerifyToken(token string, password string) (uid int64, err error) {
	uid = -1
	err = tokenErr
	if len(token) <= 0 {
		return
	}
	tokenStr, _ := base64.StdEncoding.DecodeString(token)

	msgD := strings.Split(string(tokenStr), "$")
	if len(msgD) < 2 {
		return
	}

	msg, sign := msgD[0], msgD[1]
	msgArr := strings.Split(msg, ":")
	if len(msgArr) < 2 {
		return
	}

	userId, exp := msgArr[0], msgArr[1]
	expInt, err := strconv.ParseInt(exp, 10, 64)
	if err != nil {
		return
	}
	expTime := time.Unix(expInt, 0)
	if expTime.Before(time.Now()) {
		return
	}

	if password == "" {
		password = defaultPwd
	}
	if sign != hmacSha256(password, msg) {
		return
	}
	uid, err = strconv.ParseInt(userId, 10, 64)
	return
}
func VerifyTokenSign(sign, msg, password string) error {
	if password == "" {
		password = defaultPwd
	}
	if sign != hmacSha256(password, msg) {
		return fmt.Errorf("token 校验错误")
	}
	return nil
}
func ParseToken(token string) (int64, string, string, error) {
	if len(token) <= 0 {
		return -1, "", "", fmt.Errorf("token len is zero")
	}
	tokenStr, _ := base64.StdEncoding.DecodeString(token)

	msgD := strings.Split(string(tokenStr), "$")
	if len(msgD) < 2 {
		return -1, "", "", fmt.Errorf("token 格式错误，$")
	}

	msg, sign := msgD[0], msgD[1]
	msgArr := strings.Split(msg, ":")
	if len(msgArr) < 2 {
		return -1, "", "", fmt.Errorf("token 格式错误")
	}

	userId, exp := msgArr[0], msgArr[1]
	expInt, err := strconv.ParseInt(exp, 10, 64)
	if err != nil {
		return -1, "", "", fmt.Errorf("token 格式错误, expInt")
	}
	expTime := time.Unix(expInt, 0)
	if expTime.Before(time.Now()) {
		return -1, "", "", fmt.Errorf("token 过期, expInt")
	}

	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return -1, "", "", fmt.Errorf("token 解释错误, uid is not int64")
	}
	return uid, sign, msg, nil
}

var encryptKey = []byte("ssssssss9876")

func EncryptToken(token string) (string, error) {
	block, err := des.NewCipher(encryptKey)
	if err != nil {
		return "", err
	}
	originalData := []byte(token)
	originalData = PKCS5Padding(originalData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, encryptKey)
	cryptData := make([]byte, len(originalData))
	blockMode.CryptBlocks(cryptData, originalData)
	return string(cryptData), nil
}

func DecryptToken(token string) (string, error) {
	if token == "" {
		return "", errors.New("invalid token")
	}
	block, err := des.NewCipher(encryptKey)
	if err != nil {
		return "", err
	}
	cryptData := []byte(token)
	blockMode := cipher.NewCBCDecrypter(block, encryptKey)
	originalData := make([]byte, len(cryptData))
	blockMode.CryptBlocks(originalData, cryptData)
	data := PKCS5UnPadding(originalData)
	return string(data), nil
}

func PKCS5Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, paddingText...)
}
func PKCS5UnPadding(text []byte) []byte {
	length := len(text)
	unPadding := int(text[length-1])
	return text[:(length - unPadding)]
}
