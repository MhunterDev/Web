package main

import (
	api "github.com/MhunterDev/Web/api"
	easy "github.com/MhunterDev/Web/encryption"
)

func main() {

	easy.BuildFS()
	api.Router()
}
