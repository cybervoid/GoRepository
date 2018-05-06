package beacon

//import (
//  "../helpers"
//)

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

type Block struct {
  Contents []string
  Parent_Hash, Hash string
  Ts, PowNonce, Number int
}

type ChainBlock interface {
  Init() *Block
  CheckPOW() *Block
}

func (previousBlock *Block, contents []string) Init() *Block {
  block  := Block { Contents: contents, Parent_Hash: previousBlock.Hash, Number: previousBlock.Number + 1 }
  return block
}



// type MainChainBlock struct {
//   Parent_hash string
//   Hash string
//   ts int
//
// }
// type BeaconBlock struct {
//
// }
