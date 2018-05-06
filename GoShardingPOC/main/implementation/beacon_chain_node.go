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
  b.Parent_Hash = previousBlock.Hash
  b.Ts = ts
  b.PowNonce = pownonce
}

//type ChainBlock interface {
//  Init() *Block
//  CheckPOW() *Block
//}


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
