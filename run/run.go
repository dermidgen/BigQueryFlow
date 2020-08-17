package run

import (
	"flag"
	"fmt"
	"os"
)

var operation string
var verbose string

func init() {
	flag.StringVar(&operation, "o", "", "operation to perform; apply | create | delete | get | status")
	flag.StringVar(&verbose, "v", "", "verbose output")
}

// GetMain returns operation
func GetMain() string {
	flag.Parse()

	var argOp string

	if len(os.Args) >= 2 {
		argOp = os.Args[1]
	} else {
		argOp = ""
	}

	if operation == "" && argOp != "" {
		operation = argOp
	}

	if err := checkFlags(operation); err != nil {
		fmt.Printf(err.Error() + "\n")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(operation)
	return operation
}

func checkFlags(operation string) error {
	ops := map[string]bool{
		"apply":  true,
		"create": true,
		"delete": true,
		"get":    true,
		"status": true,
	}

	if !ops[operation] {
		return fmt.Errorf("invalid operation:" + operation)
	}

	return nil
}
