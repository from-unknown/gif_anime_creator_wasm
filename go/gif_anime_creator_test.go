package gifanimecreator

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestSimple(t *testing.T) {

	filePath := "picture/Thinking_Face_Emoji.jpg"

	//imageFile, err := os.Open(filePath)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer imageFile.Close()

	//b := make([]byte, imageFile.)
	//len, err := imageFile.Read(b)
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(b))

	buf, err := ConvertImage(b, "jpg", "left", "right", "false")
	if err != nil {
		log.Fatal(err)
	}

	tmpFile, err := os.Create("tmp.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer tmpFile.Close()

	tmpFile.Write(buf)
}
