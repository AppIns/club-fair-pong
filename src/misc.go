package main

// Misc.go files always contain functions that are universal-- other programs
// could ultilize them if needed

func xZeros(x int) string {
  var str string
  for i:=0; i<x; i++ {
    str += "0"
  }
  return str
}

func xOnes(x int) string {
  var str string
  for i:=0; i<x; i++ {
    str += "1"
  }
  return str
}
