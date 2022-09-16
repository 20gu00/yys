package main

import (
	"crypto/sha256"
	"fmt"
)
func main(){
	str:=Saltpassword("admin","123456")
	fmt.Println(str)
	return
}

func Saltpassword(salt, password string) string {
	s1 := sha256.New() //hash.Hash
	s1.Write([]byte(password))
	str1 := fmt.Sprintf("%x", s1.Sum(nil)) //[]byte
	s2 := sha256.New()
	s2.Write([]byte(str1 + salt))
	return fmt.Sprintf("%x", s2.Sum(nil))
}
