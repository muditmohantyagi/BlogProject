package main

import router "blog.com/route"

func main() {
	r := router.SetupRouter()
	r.Run()
}
