package beacon

import ("go-ethereum/common")

func ToHex(s string) {
  return hexutil.Encode(s)
}
