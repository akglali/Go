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

var colors = [16]string{
	"bg-blue-100", "bg-red-100", "bg-purple-100", "bg-gray-100", "bg-green-100", "bg-yellow-100", "bg-indigo-100", "bg-pink-100", "bg-blue-300", "bg-red-300", "bg-purple-300", "bg-gray-300", "bg-green-300", "bg-yellow-300", "bg-indigo-300", "bg-pink-300"}

func randomColor() string {
	randomColorsNum := rand.Intn(len(colors))
	pickColor := colors[randomColorsNum]
	return pickColor
}
