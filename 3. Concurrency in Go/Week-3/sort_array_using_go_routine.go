package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func ReadStringLine() (string, error) {
	fmt.Println("Enter numbers: ")
	stdin := bufio.NewReader(os.Stdin)
	if inputString, err := stdin.ReadString('\n'); err != nil {
		fmt.Println("Invalid Input")
		return "", err
	} else {
		if strings.HasSuffix(inputString, "\n") {
			inputString = strings.TrimSuffix(inputString, "\n")
		}
		inputString = strings.TrimSpace(inputString)

		return inputString, err

	}

}

func ToInteger(inputstring string) ([]int, error) {
	parts := strings.Split(inputstring, " ")
	arr := make([]int, 0, len(parts))
	for _, value := range parts {
		if len(value) == 0 {
			continue
		}
		if num, err := strconv.Atoi(value); err != nil {
			return nil, err
		} else {
			arr = append(arr, num)
		}

	}
	return arr, nil
}

func ReadIntegers() ([]int, error) {
	if inputstring, err := ReadStringLine(); err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	} else {
		return ToInteger(inputstring)
	}

}

func BubbleSort(sli []int, wg *sync.WaitGroup) {
	fmt.Println("Sequence to be sorted: ", sli)
	for i := 0; i < len(sli); i++ {
		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j+1] < sli[j] {
				temp := sli[j]
				sli[j] = sli[j+1]
				sli[j+1] = temp
			}
		}
	}
	if wg != nil {
		wg.Done()
	}
}
func MergeTwoSorted(slice1 []int, slice2 []int, slice12 []int) {
	fmt.Println()
	i, j, k := 0, 0, 0
	for {

		if slice1[i] < slice2[j] {
			slice12[k] = slice1[i]
			k++
			i++
		} else if slice1[i] == slice2[j] {
			slice12[k] = slice1[i]
			k++
			i++

			slice12[k] = slice2[j]
			k++
			j++
		} else if slice1[i] > slice2[j] {
			slice12[k] = slice2[j]
			k++
			j++
		}

		if i == len(slice1) {
			for j < len(slice2) {
				slice12[k] = slice2[j]
				k++
				j++
			}
		}

		if j == len(slice2) {
			for i < len(slice1) {
				slice12[k] = slice1[i]
				k++
				i++
			}
		}

		if k == len(slice12) {
			break
		}
	}

}

func main() {
	//fmt.Println(ReadIntegers())
	if intSlice, err := ReadIntegers(); err != nil {
		fmt.Println("Error: ", err)
	} else {
		count := len(intSlice)
		if count < 4 {
			BubbleSort(intSlice, nil)
			fmt.Println("Slice sorted...")

		} else {
			divCount := len(intSlice) / 4
			firstSlice := intSlice[0:divCount]
			secondSlice := intSlice[divCount : divCount*2]
			thirdSlice := intSlice[divCount*2 : divCount*3]
			fourthSlice := intSlice[divCount*3 : divCount*4]
			var wg sync.WaitGroup
			wg.Add(4)
			go BubbleSort(firstSlice, &wg)
			go BubbleSort(secondSlice, &wg)
			go BubbleSort(thirdSlice, &wg)
			go BubbleSort(fourthSlice, &wg)
			wg.Wait()

			slice12 := make([]int, len(firstSlice)+len(secondSlice))
			MergeTwoSorted(firstSlice, secondSlice, slice12)

			slice34 := make([]int, len(thirdSlice)+len(fourthSlice))
			MergeTwoSorted(thirdSlice, fourthSlice, slice34)

			slice1234 := make([]int, len(slice12)+len(slice34))
			MergeTwoSorted(slice12, slice34, slice1234)

			fmt.Println("All numbers sorted:", slice1234)
		}
	}

}
