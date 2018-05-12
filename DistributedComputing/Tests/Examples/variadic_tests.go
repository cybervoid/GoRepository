package main

import "testing"

func TestSimpleVariadicToSlice(t *testing.T) {
    //test for no arguments
    if val := simpleVariadicToSlice(); val != nil {
        t.Error("value should be nil")
    }
}
