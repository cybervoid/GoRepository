package beacon

import (
  "../helpers"
  "encoding/binary"
  "bytes"
)

const NOTARIES int = 40
const BASE_TS_DIFF int = 1
const SKIP_TS_DIFF int = 6
const SAMPLE int = 9
const MIN_SAMPLE int = 5
const POWDIFF int = 30 * NOTARIES
const SHARDS int = 4


func checkPow(work int, nonce int) {
  // discrete log PoW
  //quadratic nonresidues only
  return pow(work, nonce, 65537) * POWDIFF < 65537 * 2
}

func pow(work string, nonce int, input int) int {
  return 65536;//placeholder
}

// type Block struct {
//   Contents []string
//   Parent_Hash, Hash string
//   Ts, PowNonce, Number int
// }

type Block struct {
  Contents []string
  Parent_Hash, Hash string
  Ts, PowNonce, Number int
}

type BeaconBlock struct {
  MainBlockRef Block
  BeaconBlock Block
}

type Blockchain interface {
  Init() Block
  CheckPOW() Block
}

//Main chain Blockchain interface implementation
func (ts int, pownonce int, previousBlock Block) Init() Block {
  var b = new(Block)
  b.Contents = helpers.RandomStringArray(5)
  if b.Parent_Hash != nil {
    b.Parent_Hash = previousBlock.Hash
    previousBlock.CheckPOW()
  } else {
    b.Parent_Hash = byte('\x00') * 32
  }

  //b.Hash =
  b.Ts = ts
  b.PowNonce = pownonce
  return b
}

func (block Block) CheckPOW() Block {

}

//type ChainBlock interface {
//  Init() *Block
//  CheckPOW() *Block
//}



// type MainChainBlock struct {
//   Parent_hash string
//   Hash string
//   ts int
//
// }
// type BeaconBlock struct {
//
// }
