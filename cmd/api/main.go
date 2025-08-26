package main

func main() {
	//declare string containing application version number
	const version = "1.0.0"

	//define config struct to hold config settings
	type config struct {
		port int
		env  string
	}

}
