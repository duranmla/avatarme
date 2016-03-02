package user

import (
  "crypto/md5"
  "io"
  "fmt"
)

// the idea is that the same identicon is going to be retrieved using email, ip or token
type User struct {
  Email         string `json:"Email"`
  Hash          string
}

func New(email string) *User {
	user := &User{Email: email}
  user.Hash = user._getStringMD5(email)
  return user
}

func (user *User) _getStringMD5(str string) string {
	Hash := md5.New()
	io.WriteString(Hash, str)

	return fmt.Sprintf("%x", Hash.Sum(nil))
}

func (user *User) String() string  {
  return fmt.Sprintf("Hi!\nyour email is: %s\nwe've generated an image with: %s\n", user.Email, user.Hash)
}
