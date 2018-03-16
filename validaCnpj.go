package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	printCnpjsValidos   = false
	printCnpjsInvalidos = true
	filePathName        = "cnpjs.csv"
)

func main() {
	fmt.Printf("PROCESSO VALIDAÇÃO CNPJS INICIADO...")
	fmt.Println()
	readFile()
}

func readFile() {
	file, error := os.Open(filePathName)

	if error != nil {
		panic(error)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	if printCnpjsValidos {
		fmt.Println("CNPJ VÁLIDOS")
	}
	if printCnpjsInvalidos {
		fmt.Println("CNPJ INVÁLIDOS")
	}

	var qtdCnpjs, qtdValidos, qtdInvalidos (int)

	for scanner.Scan() {
		cnpj := scanner.Text()

		qtdCnpjs++

		if cnpjValido(cnpj) == false {
			qtdInvalidos++
			if printCnpjsInvalidos {
				fmt.Println(cnpj)
			}
		} else {
			qtdValidos++
			if printCnpjsValidos {
				fmt.Println(cnpj)
			}
		}
	}

	fmt.Println()
	fmt.Println("REPORT FINAL")
	fmt.Println("Quantidade CNPJs analisados: ", qtdCnpjs)
	fmt.Println("Quantidade CNPJs válidos: ", qtdValidos)
	fmt.Println("Quantidade CNPJs inválidos: ", qtdInvalidos)
	fmt.Println()
	fmt.Println("GOD BLESS YOU")
}

func cnpjValido(cnpj string) bool {
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
