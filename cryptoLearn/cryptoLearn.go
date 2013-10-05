// cryptoLearn project cryptoLearn.go
package cryptoLearn

import (
	"crypto"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"math/big"
)

func Jiangyou() {
	println("打酱油~")
}

func Md5_crypto() {
	md5 := md5.New()

	md5.Write([]byte("ab"))
	md5.Write([]byte("cd"))

	fmt.Printf("%x\n", md5.Sum(nil))
}

func Sha256_crypto() {
	sha256 := sha256.New()

	sha256.Write([]byte("ab"))
	sha256.Write([]byte("cd"))

	fmt.Printf("%x\n", sha256.Sum(nil))
}

func Sha512_crypto() {
	sha512 := sha512.New()

	sha512.Write([]byte("ab"))
	sha512.Write([]byte("cd"))

	fmt.Printf("%x\n", sha512.Sum(nil))
}

func HMAC_crypto() {
	key := []byte("123456")

	hmac := hmac.New(sha256.New, key)

	hmac.Write([]byte("abcd"))

	fmt.Printf("%x\n", hmac.Sum(nil))

}

func Rand_crypto() {
	x, _ := rand.Int(rand.Reader, big.NewInt(100))
	println(x.Int64())

	y, _ := rand.Prime(rand.Reader, 100)
	fmt.Printf("% x\n", y.Bytes())

	c := make([]byte, 10)
	rand.Read(c)
	fmt.Printf("% x\n", c)
}

func DES_crypto() {
	//注意确保 DES key 长度是 8 字节，TDES key 是 24 字节。原数据长度则必须是 8 的倍数
	key := []byte("12345678")
	s := "1234567887654321"

	block, _ := des.NewCipher(key)
	enc := cipher.NewCBCEncrypter(block, key[:8])
	dec := cipher.NewCBCDecrypter(block, key[:8])

	eb := make([]byte, len(s))
	db := make([]byte, len(eb))
	enc.CryptBlocks(eb, []byte(s))
	fmt.Println(eb)
	dec.CryptBlocks(db, eb)
	fmt.Println(string(db))
}

func RAS_crypto() {
	prv, _ := rsa.GenerateKey(rand.Reader, 256)
	pub := &prv.PublicKey

	fmt.Printf("%#v\n%#v\n", prv, pub)

	out, _ := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte("Hello"))
	fmt.Printf("%#v\n", out)

	out, _ = rsa.DecryptPKCS1v15(rand.Reader, prv, out)
	fmt.Printf("%#v\n", string(out))
}

func FromBase10(base10 string) *big.Int {
	i := new(big.Int)
	i.SetString(base10, 10)
	return i
}
func GetKeys() (prv rsa.PrivateKey, pub rsa.PublicKey) {
	pub = rsa.PublicKey{
		N: FromBase10("77382..."),
		E: 65537,
	}
	prv = rsa.PrivateKey{
		PublicKey: pub,
		D:         FromBase10("11276..."),
		Primes: []*big.Int{
			FromBase10("61637..."),
			FromBase10("12554..."),
		},
	}
	return
}

/**
*FromBase10  GetKeys
 */
func Key_crypto() {
	prv, _ := GetKeys()
	out, _ := rsa.DecryptPKCS1v15(rand.Reader, &prv, []byte{161, 72, 236, 67})
	fmt.Println(string(out))
}

func Check_Crypto() {
	hash := md5.New()
	hash.Write([]byte("hello"))
	data := hash.Sum(nil)
	prv, _ := rsa.GenerateKey(rand.Reader, 1024)
	pub := &prv.PublicKey
	fmt.Printf("key: %#v\n", pub)
	sig, _ := rsa.SignPKCS1v15(rand.Reader, prv, crypto.MD5, data)
	fmt.Printf("sig: %#v\n", sig)
	if rsa.VerifyPKCS1v15(pub, crypto.MD5, data, sig) == nil {
		println("OK")
	} else {
		println("Error")
	}
}
