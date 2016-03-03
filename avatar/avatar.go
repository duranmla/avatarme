package avatar

import (
	"fmt"
	"github.com/duranmla/avatarme/user"
	"github.com/lucasb-eyer/go-colorful"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	background color.Color = color.RGBA{240, 240, 240, 255} // make the background a light white color
)

type Avatar struct {
	*user.User
	Ink    string
	Pixels [][]string
}

func New(email string) *Avatar {
	avatar := &Avatar{User: user.New(email)}
	avatar._getAvatarColor()
	avatar._getAvatarPixels()
	fmt.Println(avatar.Pixels)
	return avatar
}

func (avatar *Avatar) _getAvatarColor() {
	avatar.Ink = avatar.Hash[26:] // last 6 characteres (Hexcolor)
}

func (avatar *Avatar) _getAvatarPixels() {
	source := avatar.Hash[:26] // get first 26 characteres

	for i := 0; i < (len(source) - 2); i = i + 2 {
		avatar.Pixels = append(avatar.Pixels, []string{avatar.Hash[i:(i + 1)], avatar.Hash[(i + 1):(i + 2)]})
	}
}

func (avatar *Avatar) GenerateImage() {
	hexColor, _ := colorful.Hex(avatar.Ink)
	r, g, b := hexColor.RGB255()
	pixelColor := color.RGBA{r, g, b, 255}
	canvas := image.NewRGBA(image.Rect(0, 0, 420, 420))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{background}, image.ZP, draw.Src) // fill canvas with background
	position := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			hexValue := avatar.Pixels[position]
			value, _ := strconv.ParseInt(strings.Join(hexValue, ""), 10, 0)

			rect := image.Rect((i*90)+30, (j*90)+75, (i*90)+120, (j*90)+165)

			if value%2 == 0 {
				draw.Draw(canvas, rect, &image.Uniform{pixelColor}, image.ZP, draw.Src)
			}

			position++
		}
	}

	fileImage, _ := os.Create("identicons.png")
	defer fileImage.Close()
	png.Encode(fileImage, canvas)
	_showImage(fileImage.Name())
}

func _showImage(name string) {
	command := "open"
	arg1 := "-a"
	arg2 := "/Applications/Preview.app"
	cmd := exec.Command(command, arg1, arg2, name)
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func (avatar *Avatar) String() string {
	return fmt.Sprintf("Hi!\nyour email is: %s\nwe've generated an image with: %s\n\ncolor assigned was: #%s\n", avatar.Email, avatar.Hash, avatar.Ink)
}
