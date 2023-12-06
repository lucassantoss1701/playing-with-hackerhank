package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	positives := 0
	negatives := 0
	zeros := 0
	quantityOfValues := len(arr)

	for _, value := range arr {
		if value > 0 {
			positives++
		}

		if value < 0 {
			negatives++
		}

		if value == 0 {
			zeros++
		}

	}

	fmt.Println(positives)
	fmt.Println(negatives)
	fmt.Println(zeros)
	fmt.Println(quantityOfValues)

	fmt.Printf("%.6f\n", float64(positives)/float64(quantityOfValues))
	fmt.Printf("%.6f\n", float64(negatives)/float64(quantityOfValues))
	fmt.Printf("%.6f\n", float64(zeros)/float64(quantityOfValues))

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	fmt.Println("ue")

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
