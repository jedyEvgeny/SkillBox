package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("*****Задание 2. Поиск символов в нескольких строках*****")
	newSliceString, newSliceRune := createAndFillSlices()
	matrix(newSliceString, newSliceRune)
}

func createAndFillSlices() (newSliceString []string, newSliceRune []rune) {
	sizeSlice := getSizeSlice()
	newSliceString = createSliceString(sizeSlice)
	fillSliceString(newSliceString)
	printSliceString(newSliceString)
	sizeSlice = getSizeSlice()
	newSliceRune = createSliceRune(sizeSlice)
	fillSliceRune(newSliceRune)
	printSliceRune(newSliceRune)
	cutSliceString(newSliceString)
	printSliceString(newSliceString)
	return
}

func matrix(sliceString []string, sliceRune []rune) {
	colsMatrix, rowsMatrix := defineSizeMatrix(sliceString, sliceRune)
	newMatrix := createMatrix(colsMatrix, rowsMatrix)
	fillMatrix(newMatrix, sliceString, sliceRune, colsMatrix, rowsMatrix)
	printMatrix(newMatrix, colsMatrix, rowsMatrix)
}

func getSizeSlice() int {
	fmt.Println("\n---Задаём размер нового среза---")
	var size int
	fmt.Print("Введите размер среза: ")
	fmt.Scan(&size)
	return size
}

func createSliceString(size int) []string {
	fmt.Println("\n---Создаём пустой срез для строк---")
	return make([]string, size)
}

func fillSliceString(slice []string) {
	fmt.Println("\n---Заполняем срез предложениями---")
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Введите фразу %d: ", i+1)
		scanner.Scan()
		slice[i] = scanner.Text()
	}
}

func printSliceString(slice []string) {
	fmt.Println("\n---Печатаем срез---")
	fmt.Println("Элементы среза строк:")
	for i, phrase := range slice {
		fmt.Printf("%d: %s\n", i+1, phrase)
	}
}

func createSliceRune(size int) []rune {
	fmt.Println("\n---Создаём срез из рун---")
	return make([]rune, size)
}

func fillSliceRune(slice []rune) {
	fmt.Println("\n---Заполняем срез рун символами---")
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Введите %d-й символ: ", i+1)
		scanner.Scan()
		slice[i] = []rune(scanner.Text())[0]
	}
}

func printSliceRune(slice []rune) {
	fmt.Println("\n---Печатаем срез---")
	fmt.Println("Элементы среза рун:")
	for _, phrase := range slice {
		fmt.Printf("%c\t", phrase)
		//спецификаток %c выводит символы, представленные числовым кодом
	}
	fmt.Println()
}

func cutSliceString(slice []string) []string {
	fmt.Println("\n---Обрабатываем срез строк---")
	for i := 0; i < len(slice); i++ {
		//Создаём подсрез из фрагментов фразы. Разделитель - пробел
		sliceWords := strings.Split(slice[i], " ")
		//Добавляем в исходный срез последнее слово фразы
		slice[i] = sliceWords[len(sliceWords)-1]
	}
	return slice
}

func defineSizeMatrix(sliceString []string, sliceRune []rune) (int, int) {
	fmt.Println("\n---Определяем размеры матрицы---")
	cols := len(sliceRune)
	rows := len(sliceString)
	fmt.Printf("У матрицы %d рядов и %d столбцов.", rows, cols)
	return cols, rows
}

func createMatrix(cols, rows int) [][]string {
	fmt.Println("\n---Создаём матрицу--")
	outMatrix := make([][]string, rows)
	for i := range outMatrix {
		outMatrix[i] = make([]string, cols)
	}
	return outMatrix
}

func fillMatrix(matrix [][]string, sliceString []string, sliceRune []rune, cols, rows int) [][]string {
	fmt.Println("\n---Заполняем матрицу---")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			idx := strings.Index(sliceString[i], string(sliceRune[j]))
			str := fmt.Sprintf("'%c' в поз. %d", sliceRune[j], idx)
			matrix[i][j] = str
		}
	}
	return matrix
}

func printMatrix(matrix [][]string, cols, rows int) {
	fmt.Println("\n---Матрица---")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%s\t", matrix[i][j])
		}
		fmt.Println()
	}
}
