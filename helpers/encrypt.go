package helpers

import (
	"api/models"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func GetTradeInfo(params map[string]string) string {
	key := os.Getenv("PAY_HASH_KEY")
	iv := os.Getenv("PAY_HASH_IV")
	param_str := HttpBuildQuery(params)
	return Aes256(param_str, key, iv, aes.BlockSize)
}

func GetTradeSha(info string) string {
	key := os.Getenv("PAY_HASH_KEY")
	iv := os.Getenv("PAY_HASH_IV")
	hash := "HashKey=" + key + "&" + info + "&HashIV=" + iv
	h := sha256.New()
	h.Write([]byte(hash))
	sha := hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(sha)
}

func DecodeUnicode(input string) string {
	var jsonRawEscaped json.RawMessage
	var jsonRawUnescaped json.RawMessage
	sentence := fmt.Sprintf(`{"code": "%s"}`, input)
	jsonRawEscaped = []byte(sentence)
	jsonRawUnescaped, _ = _UnescapeUnicodeCharactersInJSON(jsonRawEscaped) // "☺"
	dest := models.Code{}
	json.Unmarshal(jsonRawUnescaped, &dest)
	return dest.Code
}

func _UnescapeUnicodeCharactersInJSON(_jsonRaw json.RawMessage) (json.RawMessage, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(_jsonRaw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func DecodeTradeInfo(cipherText string) string {
	key := os.Getenv("PAY_HASH_KEY")
	iv := os.Getenv("PAY_HASH_IV")
	return Aes256Decode(cipherText, key, iv)
}

func Aes256(plaintext string, key string, iv string, blockSize int) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}

func Aes256Decode(cipherText string, encKey string, iv string) (decryptedString string) {
	bKey := []byte(encKey)
	bIV := []byte(iv)
	cipherTextDecoded, err := hex.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, bIV)
	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	return string(cipherTextDecoded)
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func HttpBuildQuery(params map[string]string) (param_str string) {
	params_arr := make([]string, 0, len(params))
	for k, v := range params {
		params_arr = append(params_arr, fmt.Sprintf("%s=%s", k, v))
	}
	param_str = strings.Join(params_arr, "&")
	return param_str
}
