package helpers
import (
  "encoding/hex"
  "github.com/ethereum/go-ethereum/crypto/sha3"
  "log"
)

//Helper functions
func ToHex(inputBytes []byte) string {
  return hex.EncodeToString(inputBytes)
}
func FromHex(hexString string) (decoded []byte) {
  decoded, err := hex.DecodeString(hexString)
  if err != nil {
		log.Fatal(err)
	}
  return decoded
}

func Utf8ToBytes(inputString string) []byte {
  var length int = len(inputString)
  outputBytes := make([]byte, length)
  asRunes := []rune(inputString)
  for i := 0; i < length; i++ {
    outputBytes[i] = byte(asRunes[i])
  }
  return outputBytes
}

func Sha3FromRawString(inputString string) string {
  asBytes := Utf8ToBytes(inputString)
  buf := toSha3(asBytes)
  return ToHex(buf)
}

func toSha3(inputBytes []byte) []byte {
  hash := sha3.NewKeccak256()
  hash.Write(inputBytes)
  var buf []byte
  buf = hash.Sum(buf)
  return buf
}

func Sha3FromHex(hexStringInput string) string {
  decoded := FromHex(hexStringInput)
  buf := toSha3(decoded)
  return hex.EncodeToString(buf)
}

func HashToInt(hash string) int64 {
  var o int64 = 0
  //hashAsBytes := Utf8ToBytes(hash)
  //for c := range hashAsBytes {
  asRunes := []rune(hash)
  for _, k := range asRunes {
    c := int64(k)
    o = (o << 8) + c
  }
  return o
}
