package request

func Request(id <-chan int, body chan<- string) {
	for i := range id {
		url := generatorUrl(i)
		body <- get(url)
	}
}
