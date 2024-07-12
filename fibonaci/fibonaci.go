package main

func fib1(p int) int{
  seq := []int {0,1}
  for len(seq) <= p {
    seq = append(seq, seq[len(seq)-1] + seq[len(seq)-2])
  }
  return seq[len(seq)-1]
}
