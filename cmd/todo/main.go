package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)

	//hypot := func(x, y float64) float64 {
	//	return math.Sqrt(x * x) + (y * y)
	//}
	//fmt.Println(hypot(5, 12))
	//
	//fmt.Println(compute(hypot))
	//fmt.Println(compute(math.Pow))
}

const dbPath = "mydb.pb"

func add(text string) error {
	task := &todo.Task
	return nil
}

func list() error {
	return nil
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}
	var err error
	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list()
	case "add":
		err = add(strings.Join(flag.Args()[1:], ""))
	default:
		err = fmt.Errorf("unknown subcommand %s", cmd)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("todo - running")
}
