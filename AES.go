package method

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"github.com/15226124477/define"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func KeepLogin(strTime string) bool {
	stamp, err := strconv.Atoi(strTime)
	if err != nil {
		log.Error(err)
	}
	// 将时间戳转换为time.Time类型的值
	t := time.Unix(int64(stamp), 0)
	// 格式化时间
	if time.Now().Sub(t).Seconds() > define.WebKeepLoginSecond {
		return false
	} else {
		log.Info("Offline Time:", define.WebKeepLoginSecond-time.Now().Sub(t).Seconds(), "Last Time:", t.Format("2006-01-02 15:04:05"))
		return true
	}

}

// AESEncrypt AES加密函数
func AESEncrypt(plaintext []byte) (string, error) {
	// 补充明文，使其长度为块大小的倍数
	if len(plaintext)%aes.BlockSize != 0 {
		padding := aes.BlockSize - len(plaintext)%aes.BlockSize
		plaintext = append(plaintext, bytes.Repeat([]byte{byte(padding)}, padding)...)
	}

	key := []byte(define.WebAES16ByteKey)   // 16字节的密钥
	iv := []byte(define.WebAES16ByteVector) // 16字节的初始向量
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(plaintext)%aes.BlockSize != 0 {
		return "", fmt.Errorf("plaintext is not a multiple of the block size")
	}

	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	return hex.EncodeToString(ciphertext), nil
}

// AESDecrypt AES解密函数
func AESDecrypt(text string) ([]string, int) {
	ciphertext, _ := hex.DecodeString(text)
	key := []byte(define.WebAES16ByteKey)   // 16字节的密钥
	iv := []byte(define.WebAES16ByteVector) // 16字节的初始向量
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, 0
	}
	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, 0
	}
	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)
	result := strings.Split(string(plaintext), "@")
	if len(result) == 3 {
		rs := make([]string, 0)
		for i := 0; i < 3; i++ {
			re := regexp.MustCompile(`\d+`)
			// 使用FindAllString方法来查找所有匹配的数字
			matches := re.FindAllString(result[i], -1)
			rs = append(rs, matches[0])
		}

		return rs, len(rs)
	} else {
		return result, len(result)
	}

}
