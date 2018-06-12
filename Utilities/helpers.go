package helpers

import ("os")

//error helpers
func check(e error) {
    if e != nil {
        panic(e)
    }
}
//file helpers
func fileExists(path string) bool {
    if _, err := os.Stat(path); err == nil {
        return true
    }
    return false
}
