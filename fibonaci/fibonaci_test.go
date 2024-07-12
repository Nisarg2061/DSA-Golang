package main

import "testing"

func TestFeb1(t *testing.T){
  result := fib1(6)

  if result != 8{
    t.Error("Expected value 8 got", result)
  }
}
