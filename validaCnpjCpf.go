package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	printValidDocuments   = false
	printInvalidDocuments = true
	filePathName          = "documents.csv"
)

func main() {
	fmt.Println("STARTING PROCESS TO VALIDATE DOCUMENTS (CNPJs and CPFs)...\n")
	// fmt.Println()
	readFile()
}

func readFile() {
	file, error := os.Open(filePathName)

	if error != nil {
		panic(error)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	if printValidDocuments {
		fmt.Println("VALID DOCUMENTS")
	}
	if printInvalidDocuments {
		fmt.Println("INVALID DOCUMENTS")
	}

	var qtdDocuments, qtdValid, qtdInvalid (int)

	for scanner.Scan() {
		document := scanner.Text()

		qtdDocuments++

		var validDoc bool

		if len(document) == 11 {
			validDoc = validCpf(document)
		} else {
			validDoc = validCnpj(document)
		}

		if validDoc == false {
			qtdInvalid++
			if printInvalidDocuments {
				fmt.Println(document)
			}
		} else {
			qtdValid++
			if printValidDocuments {
				fmt.Println(document)
			}
		}
	}

	fmt.Println()
	fmt.Println("FINAL REPORT")
	fmt.Println("Number of documents processed: ", qtdDocuments)
	fmt.Println("Number of valid documents: ", qtdValid)
	fmt.Println("Number of invalid documents: ", qtdInvalid)
	fmt.Println()
	fmt.Println("GOD BLESS YOU")
}

func validCnpj(cnpj string) bool {
	cnpj = strings.Replace(cnpj, ".", "", -1)
	cnpj = strings.Replace(cnpj, "-", "", -1)
	cnpj = strings.Replace(cnpj, "/", "", -1)
	if len(cnpj) != 14 {
		return false
	}

	algs := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var algProdCpfDig1 = make([]int, 12, 12)
	for key, val := range algs {
		intParsed, _ := strconv.Atoi(string(cnpj[key]))
		sumTmp := val * intParsed
		algProdCpfDig1[key] = sumTmp
	}
	sum := 0
	for _, val := range algProdCpfDig1 {
		sum += val
	}
	digit1 := sum % 11
	if digit1 < 2 {
		digit1 = 0
	} else {
		digit1 = 11 - digit1
	}
	char12, _ := strconv.Atoi(string(cnpj[12]))
	if char12 != digit1 {
		return false
	}
	algs = append([]int{6}, algs...)

	var algProdCpfDig2 = make([]int, 13, 13)
	for key, val := range algs {
		intParsed, _ := strconv.Atoi(string(cnpj[key]))

		sumTmp := val * intParsed
		algProdCpfDig2[key] = sumTmp
	}
	sum = 0
	for _, val := range algProdCpfDig2 {
		sum += val
	}

	digit2 := sum % 11
	if digit2 < 2 {
		digit2 = 0
	} else {
		digit2 = 11 - digit2
	}
	char13, _ := strconv.Atoi(string(cnpj[13]))
	if char13 != digit2 {
		return false
	}
	return true
}

func validCpf(cpf string) bool {
	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)
	if len(cpf) != 11 {
		return false
	}
	var eq bool
	var dig string
	for _, val := range cpf {
		if len(dig) == 0 {
			dig = string(val)
		}
		if string(val) == dig {
			eq = true
			continue
		}
		eq = false
		break
	}
	if eq {
		return false
	}

	i := 10
	sum := 0
	for index := 0; index < len(cpf)-2; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}

	prod := sum * 10
	mod := prod % 11
	if mod == 10 {
		mod = 0
	}
	digit1, _ := strconv.Atoi(string(cpf[9]))
	if mod != digit1 {
		return false
	}
	i = 11
	sum = 0
	for index := 0; index < len(cpf)-1; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}
	prod = sum * 10
	mod = prod % 11
	if mod == 10 {
		mod = 0
	}
	digit2, _ := strconv.Atoi(string(cpf[10]))
	if mod != digit2 {
		return false
	}

	return true
}
