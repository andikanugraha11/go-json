package main

import (
	"encoding/json"
	"fmt"
)

type Biodata struct {
	FirstName 	string `json:"first_name"`
	LastName	string
	Age 		int
	SocialMedia SocialMedia
	Hobbies		[]string
}

type SocialMedia struct {
	Linkedin string `json:"linked_id"`
	Instagram string
}

func main()  {

	// Struct Encode
	//data := Biodata{
	//	FirstName: "Bagus",
	//	LastName:  "Aji",
	//	Age:       17,
	//	SocialMedia: SocialMedia{
	//		Linkedin:  "www.linkedin.com",
	//		Instagram: "www.instagram.com",
	//	},
	//	Hobbies: []string{
	//		"Menggambar",
	//		"Membaca",
	//	},
	//}
	//PrintJSON(data)

	// Struct Decode
	//dataString := `{"first_name":"Bagus","LastName":"Aji","Age":17,
	//				"SocialMedia":{"linked_id":"www.linkedin.com","Instagram":"www.instagram.com"},
	//				"Hobbies":["Menggambar","Membaca"]}`
	//
	//dataBytes := []byte(dataString)
	//
	//var biodata Biodata
	//
	//json.Unmarshal(dataBytes, &biodata)
	//
	//fmt.Println(biodata)
	//fmt.Println(biodata.Age)
	//fmt.Println(biodata.SocialMedia.Linkedin)

	// Interface Encode
	//data := map[string]interface{}{
	//	"FirstName": "Bagus",
	//	"LastName":  "Aji",
	//	"Email": "bagus.aji@mail.com",
	//	"Age":       17,
	//	"SocialMedia": SocialMedia{
	//		Linkedin:  "www.linkedin.com",
	//		Instagram: "www.instagram.com",
	//	},
	//	"Hobbies": []string{
	//		"Menggambar",
	//		"Membaca",
	//	},
	//}
	//
	//PrintJSON(data)


	//Interface Decode
	dataString := `{"first_name":"Bagus","LastName":"Aji","Age":17,
					"SocialMedia":{"linked_id":"www.linkedin.com","Instagram":"www.instagram.com"},
					"Hobbies":["Menggambar","Membaca"]}`

	dataBytes := []byte(dataString)

	var biodata map[string]interface{}
	json.Unmarshal(dataBytes, &biodata)

	fmt.Println(biodata)
	fmt.Println(biodata["Age"])
	fmt.Println(biodata["SocialMedia"].(map[string]interface{})["Instagram"])
	//fmt.Println(biodata.SocialMedia.Linkedin)

}

func PrintJSON(data interface{})  {
	// Encode: Go -> JSON
	bytesData, _ := json.Marshal(data)
	fmt.Println(string(bytesData))
}
