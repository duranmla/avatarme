package avatar

import (
  "fmt"
  "github.com/duranmla/avatarme/user"
)

type Avatar struct {
  *user.User
  Ink     string
  Pixels  [12][2]int // [[0,1], [4,5]...]
}

func New(email string) *Avatar {
  avatar := &Avatar{User: user.New(email)}
  avatar._getAvatarColor()
  avatar._getAvatarPixels()

  fmt.Println(avatar)
  return avatar
}

func (avatar *Avatar) _getAvatarColor() {
  avatar.Ink = avatar.Hash[26:] // last 6 characteres (Hexcolor)
}

func (avatar *Avatar) _getAvatarPixels() {
  source := avatar.Hash[:26] // get first 26 characteres

  for i:=0; i<=(len(source)-2); i+2 {
    if i%2 == 0 {
      avatar.Pixels[i/2] = [2]int{avatar.Hash[i:(i+1)], avatar.Hash[(i+1):(i+2)]}
    }
  }
}

func (avatar *Avatar) String() string  {
  return fmt.Sprintf("Hi!\nyour email is: %s\nwe've generated an image with: %s\n\n color assigned was: %s\n", avatar.Email, avatar.Hash, avatar.Ink)
}
