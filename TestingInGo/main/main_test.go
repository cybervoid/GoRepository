package main

import "testing"

func TestCalulcate(t *testing.T) {
    if Calculate(2) != 4 {
        t.Error("Expected 2 + 2 = 4")
    }
}
