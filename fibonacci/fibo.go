package main

//import(
// "fmt"
// "math"
//)

func fibo(n int) int {
  if n < 2 {
   return 1
  }
  return fibo(n-2) + fibo(n-1)
}

func main() {
  // rang = 5
  println(fibo(5))
}
