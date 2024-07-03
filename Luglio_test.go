package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var prog = "05099A_CaronePolettini_Luca.go"
var verbose = true

func LanciaGenericaConFileInOutAtteso(t *testing.T, prog string, inFile string, outFile string, verbose bool) {
	// Read input file
	input, err := os.ReadFile(inFile)
	if err != nil {
		t.Fatalf("Error reading input file %s: %v", inFile, err)
	}

	// Read expected output file
	expectedOutput, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("Error reading output file %s: %v", outFile, err)
	}

	// Create a command to run the program
	cmd := exec.Command("go", "run", prog)
	cmd.Stdin = bytes.NewReader(input)

	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Error running program %s: %v\nStderr: %s", prog, err, errBuffer.String())
	}

	// Compare the actual output with the expected output
	actualOutput := outBuffer.Bytes()
	if !bytes.Equal(actualOutput, expectedOutput) {
		t.Errorf("Output mismatch for %s:\nExpected:\n%s\nGot:\n%s", inFile, string(expectedOutput), string(actualOutput))
	}

	// If verbose, print the actual output
	if verbose {
		fmt.Printf("Test %s:\nExpected:\n%s\nGot:\n%s\n", inFile, string(expectedOutput), string(actualOutput))
	}
}

func TestBaseColoraStato(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./inFiles/BaseColoraStato",
		"./outFiles/BaseColoraStato",
		verbose,
	)
}

func TestBaseRegole(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./inFiles/BaseRegole",
		"./outFiles/BaseRegole",
		verbose,
	)
}

func TestBaseBlocco(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./inFiles/BaseBlocco",
		"./outFiles/BaseBlocco",
		verbose,
	)
}

func TestBasePista(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./inFiles/BasePista",
		"./outFiles/BasePista",
		verbose,
	)
}
