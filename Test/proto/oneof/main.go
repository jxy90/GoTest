package main

import (
	"fmt"
	"github.com/jxy90/GoTest/Test/proto/oneof/account"
)

func main() {
	p1 := &account.Profile{
		Avatar: &account.Profile_ImageUrl{ImageUrl: "https://example.com/image.png"},
	}
	check(p1)
	// imageData is []byte
	imageData := getImageData()
	p2 := &account.Profile{
		Avatar: &account.Profile_ImageData{ImageData: imageData},
	}
	check(p2)
}

func getImageData() []byte {
	return []byte{}
}

func check(m *account.Profile) error {
	switch x := m.Avatar.(type) {
	case *account.Profile_ImageUrl:
		// Load profile image based on URL
		// using x.ImageUrl
		fmt.Println("is Profile_ImageUrl")
	case *account.Profile_ImageData:
		// Load profile image based on bytes
		// using x.ImageData
		fmt.Println("is Profile_ImageData")
	case nil:
		// The field is not set.
	default:
		return fmt.Errorf("Profile.Avatar has unexpected type %T", x)
	}
	return nil
}
