package main

import "fmt"

func main() {
	nickNames := make(map[string]string)
	nickNames["Mohamed"] = "mo"
	fmt.Println(nickNames["Mohamed"])

	cityCode := map[string]string {
		"Cairo" : "CA",
		"Alexandria" : "AL", 
		"Suez" : "SU",  // take care of this ending comma
	}
	cityCode["Cairo"] = "CAI"
	delete(cityCode,"Suez")
	for key,value := range cityCode {
		fmt.Println(key,",",value)
	}

	cityC, exists := cityCode["XYZ"]
	fmt.Println(cityC,exists)
}
