package post

import "math/rand"

var animals [12]string = [12]string{"Duck", "Dog", "Elephant", "Aardvark", "Affenpinscher", "African Bush Elephant", "African Forest Elephant",
	"African Tree Toad", "Airedale Terrier", "Akita", "Alaskan Husky", "Albacore Tuna"}

var adjective [12]string = [12]string{"Elegant", "Exquisite", "Glorious", "Aardvark", "Junoesque", "Magnificent", "Resplendent",
	"Splendid", "Statuesque", "Blue-eyed", "Busy", "Brave"}

func randomNickname() string {
	randomAnimalsNum := rand.Intn(len(animals))
	randomAdjectivesNum := rand.Intn(len(adjective))
	pickAnimal := animals[randomAnimalsNum]
	pickAdjective := adjective[randomAdjectivesNum]
	vName := pickAdjective + " " + pickAnimal
	return vName
}
