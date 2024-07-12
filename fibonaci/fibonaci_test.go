package main

import "testing"

func TestFeb1(t *testing.T){
  result := fib1(6)

  if result != 8{
    t.Error("Expected value 8 got", result)
  }
}

func TestFeb2(t *testing.T){
  result := fib2(7, nil)

  if result != 13{
    t.Error("Expected value 13 got", result)
  }
}
