package main

import (
	"os"
	"time"

	"github.com/leetcode-golang-classroom/learn-golang-with-tests/clockface_sample/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
