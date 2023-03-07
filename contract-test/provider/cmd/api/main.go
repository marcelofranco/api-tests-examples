package main

type Client struct{}

const port = ":80"

func main() {
	DBConnect()
	HandleRequests(port)
}
