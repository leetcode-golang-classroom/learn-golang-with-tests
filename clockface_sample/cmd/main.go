package main

import (
	"os"
	"time"

	"github.com/leetcode-golang-classroom/learn-golang-with-tests/clockface_sample/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
