package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"textForge/helpers"
)

func main() {
	if len(os.Args) != 3{
		fmt.Println("✗ USAGE: go run . input.txt output.txt")
		os.Exit(1)
	}

	inFile := os.Args[1]
	outFile := os.Args[2]

	if (!(strings.HasSuffix(inFile, ".txt"))) || (!(strings.HasSuffix(outFile, ".txt"))) {
		fmt.Println("✗ File must end with .txt")
		os.Exit(1)
	}

	if inFile == outFile {
		fmt.Println("✗ Both file name cannot be the same")
		os.Exit(1)
	}

	file, err := os.OpenFile(inFile, os.O_RDONLY, 0664)
	if err != nil {
		log.Fatal("✗ Error opening file: ", err)
	}
	defer file.Close()

	var b strings.Builder

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text() // To maintain structure


		//====== Helper Func ======
		tokens := helpers.TokenizeElements(line)
		tokens = helpers.ApplyHex(tokens)
		tokens = helpers.ApplyBin(tokens)
		tokens = helpers.ApplyCase(tokens)
		tokens = helpers.ApplyQuote(tokens)
		tokens = helpers.ApplyPunct(tokens)
		tokens = helpers.ApplyAn(tokens)

		processed := strings.Join(tokens, " ")
		
		b.WriteString(processed)
		b.WriteString("\n")
	}	
	
	//=========================//
	ProcessedText := b.String()
	//=========================//

	if err := scanner.Err(); err != nil{
		log.Fatal("✗ Error reading file: ", err)
	}	


	file2, err := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		log.Fatal("✗ Error opening file: ", err)
	}
	defer file2.Close()

	_, err = file2.WriteString(ProcessedText)
	if err != nil {
		log.Fatal("✗ Error writing file: ", err)
	}

	fmt.Println("Text Forge Processed.")
}