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
