package users

import "math/rand"

func getRandomNickname() string {
	return getRandomAdjective() + " " + getRandomAnimal()
}

func getRandomElementFromArray(elements []string) string {
	index := rand.Intn(len(elements))
	return elements[index]
}

func getRandomAdjective() string {
	adjective := []string{
		"Busy",
		"Lazy",
		"Careless",
		"Clumsy",
		"Nimble",
		"Brave",
		"Mighty",
		"Clever",
		"Afraid",
		"Bashful",
		"Proud",
		"Fair",
		"Greedy",
		"Wise",
		"Foolish",
		"Tricky",
		"Truthful",
		"Loyal",
		"Happy",
		"Cheerful",
		"Joyful",
		"Carefree",
		"Friendly",
		"Moody",
		"Cranky",
		"Gloomy",
		"Worried",
		"Excited",
		"Calm",
		"Silly",
		"Wild",
		"Crazy",
		"Odd",
		"Alert",
		"Sleepy",
		"Surprised",
		"Tense",
		"Rude",
		"Selfish",
		"Strict",
		"Tough",
		"Polite",
		"Amusing",
		"Kind",
		"Gentle",
		"Quiet",
		"Caring",
		"Hopeful",
		"Rich",
		"Thrifty",
		"Stingy",
		"Generous",
		"Quick",
		"Speedy",
		"Swift",
		"Fantastic",
		"Splendid",
		"Wonderful",
		"Freezing",
		"Icy",
		"Steaming",
		"Sizzling",
		"Huge",
		"Sturdy",
		"Grand",
		"Heavy",
		"Plump",
		"Deep",
		"Small",
		"Tiny",
		"Beautiful",
		"Adorable",
		"Shining",
		"Sparkling",
		"Glowing",
		"Sloppy",
		"Messy",
		"Spiky",
		"Rusty",
		"Fuzzy",
		"Plush",
		"Smooth",
		"Glassy",
		"Stiff",
		"Loud",
	}
	return getRandomElementFromArray(adjective)
}

func getRandomAnimal() string {
	animals := []string{
		"Aardvark",
		"Alligator",
		"Alpaca",
		"Anaconda",
		"Ant",
		"Antelope",
		"Armadillo",
		"Baboon",
		"Badger",
		"Barracuda",
		"Bat",
		"Beaver",
		"Bee",
		"Bird",
		"Bison",
		"Bobcat",
		"Buffalo",
		"Butterfly",
		"Camel",
		"Cat",
		"Catfish",
		"Cheetah",
		"Chicken",
		"Chimpanzee",
		"Chipmunk",
		"Cobra",
		"Cougar",
		"Coyote",
		"Crab",
		"Crocodile",
		"Crow",
		"Deer",
		"Dinosaur",
		"Dolphin",
		"Dove",
		"Dragonfly",
		"Duck",
		"Eagle",
		"Eel",
		"Elephant",
		"Emu",
		"Falcon",
		"Ferret",
		"Fish",
		"Flamingo",
		"Fox",
		"Frog",
		"Goat",
		"Goose",
		"Gopher",
		"Gorilla",
		"Hamster",
		"Hawk",
		"Horse",
		"Hummingbird",
		"Husky",
		"Iguana",
		"Impala",
		"Kangaroo",
		"Lemur",
		"Leopard",
		"Lion",
		"Lizard",
		"Llama",
		"Lobster",
		"Lynx",
		"Monkey",
		"Moose",
		"Mouse",
		"Octopus",
		"Ostrich",
		"Otter",
		"Owl",
		"Ox",
		"Oyster",
		"Panda",
		"Panther",
		"Parrot",
		"Peacock",
		"Pelican",
		"Penguin",
		"Perch",
		"Pigeon",
		"Rabbit",
		"Raccoon",
		"Seal",
		"Sheep",
		"Sloth",
		"Tiger",
		"Whale",
		"Wolf",
		"Zebra",
	}
	return getRandomElementFromArray(animals)
}
