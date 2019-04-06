package main

import (
	"os"
)

const (
	MODE os.FileMode = os.FileMode(0700)
)

type Entity string

const (
	Client  Entity = ".client"
	Project Entity = ".project"
	Task    Entity = ".task"
)
