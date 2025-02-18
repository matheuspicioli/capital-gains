package main

import (
	"bufio"
	"capital-gains/operation"
	"capital-gains/types"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	stat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if stat.Size() <= 0 {
		fmt.Println("Stdin not provided")
		return
	}

	reader := bufio.NewReaderSize(os.Stdin, 32)
	var cliOperations []types.OperationInputCLI

	for {
		operationsFromFile, err := reader.ReadString(']')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		unmarshalErr := json.Unmarshal([]byte(operationsFromFile), &cliOperations)
		if unmarshalErr != nil {
			panic(unmarshalErr)
		}

		operations := operation.FromOperationsInputCLI(cliOperations)
		_taxes := operation.Process(operations)
		ProcessTaxesToStdout(_taxes)
	}
}

func ProcessTaxesToStdout(_taxes []float64) {
	taxes := "[{"
	for i, tax := range _taxes {
		taxes += fmt.Sprintf("%.2f", tax)
		if i != len(_taxes)-1 {
			taxes += ","
		}
	}
	taxes += "}]\n"

	os.Stdout.WriteString(taxes)
}
