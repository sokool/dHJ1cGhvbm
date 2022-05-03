package main

func main() {
	s, err := NewService(Config{})
	if err != nil {
		panic(err)
	}

	if err = s.Run(); err != nil {
		panic(err)
	}
}
