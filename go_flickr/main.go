package main

import (
	"encoding/xml"
	"fmt"
	"github.com/mncaudill/go-flickr"
	"os/exec"
	"runtime"
)

type Auth_Frob struct {
	Frob string `xml:"frob"`
}

func FindSpecifcPhoto(api_key string, api_secret string, photo_id string) {
	r := &flickr.Request{
		ApiKey: api_key,
		Method: "flickr.photos.getInfo",
		Args: map[string]string{
			"photo_id": photo_id,
		},
	}

	// Don't need to sign but might as well since we're testing
	r.Sign(api_secret)

	fmt.Println(r.URL())

	resp, err := r.Execute()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func main() {
	api_key := ""
	api_secret := ""

	FindSpecifcPhoto(api_key, api_secret, "5356343650")
	var respone_frob Auth_Frob

	{
		fmt.Println("Get frob")
		r := &flickr.Request{
			ApiKey: api_key,
			Method: "flickr.auth.getFrob",
			Args: map[string]string{
				"api_key": api_key,
			},
		}

		// Don't need to sign but might as well since we're testing
		r.Sign(api_secret)

		fmt.Println(r.URL())

		resp, err := r.Execute()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)

		err = xml.Unmarshal([]byte(resp), &respone_frob)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("frob=%s\n", respone_frob.Frob)
		}
	}

	//Open login windows
	{
		fmt.Println("Starting to login page.......")
		request_url := fmt.Sprintf("http://flickr.com/services/auth/?api_key=%s&perms=%s&frob=%s&api_sig=%s", api_key, "write", respone_frob.Frob, api_secret)
		var err error
		switch runtime.GOOS {
		case "linux":
			fmt.Println("Mac/Linux")
			err = exec.Command("xdg-open", request_url).Start()
		case "windows", "darwin":
			fmt.Println("Windows")
			fmt.Println(request_url)
			err = exec.Command("cmd", "/c", "start", request_url).Start()
		default:
			err = fmt.Errorf("unsupported platform")
		}

		if err != nil {
			fmt.Println(err)
		}
	}

	{
		fmt.Println("Get token")
		r := &flickr.Request{
			ApiKey: api_key,
			Method: "flickr.auth.getToken",
			Args: map[string]string{
				"api_key": api_key,
				"frob":    respone_frob.Frob,
			},
		}

		// Don't need to sign but might as well since we're testing
		r.Sign(api_secret)

		fmt.Println(r.URL())

		resp, err := r.Execute()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)

		err = xml.Unmarshal([]byte(resp), &respone_frob)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("frob=%s\n", respone_frob.Frob)
		}
	}
	// fmt.Println("Starting to upload file.")
	// resp2, err2 := r.Upload("sample1.jpg", "image/jpeg")
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	// fmt.Println(resp2)
	fmt.Println("Done....")
}
