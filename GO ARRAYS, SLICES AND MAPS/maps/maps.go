package maps

import "fmt"

func Maps() {
	websites := map[string]string{"google": "google.com", "amazon": "aws.com"}

	//Creating a new value
	websites["telegram"] = "telegram.com"

	//Deleta a key in a map
	delete(websites, "google")
	fmt.Println(websites)
}
