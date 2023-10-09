package main 

import (
  "fmt"
  "time"
  "net"
)

func Check(destination string, port string) string {
  address := destination + ":" + port
  timeout := time.Duration(5 + time.Second)
  conn, err := 
}
