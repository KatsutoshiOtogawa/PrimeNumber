package main

import (
	"errors"
	"fmt"
	"math"
)

// 指定した自然数以下の素数を配列で返す。
func createPrimeNumberArray(naturalNumber uint64) ([]uint, error) {
	if naturalNumber == 0 {
		err := errors.New("this number is not natural Number")
		return nil, err
	}
	// indexは0から始まるので1ずれる。
	checkArray := make([]bool, naturalNumber+1)
	countArray := make([]uint, naturalNumber+1)
	// 素数はtrue, 合成数はfalseとして、
	// 合成数か素数かわからない値はとりあえずfalseとしておく。
	for i := range checkArray {
		checkArray[i] = true
	}
	checkArray[0] = false
	checkArray[1] = false
	// インデックスの値で初期化。
	for i := range countArray {
		countArray[i] = uint(i)
	}
	// 素数の数をカウントする整数
	countPrimeNum := uint(0)
	for i := range checkArray {
		if i == 0 || i == 1 {
			continue
		}
		// 素数かもしれない数字を順番に割っていくことで判定
		for j := range countArray {
			// 素数かもしれない数字以上に割る数が大きくなったのでループを抜ける。
			// その場合は素数となる。
			if i <= j {
				countPrimeNum++
				break
			}
			// 0, 1は素数判定に使えないので飛ばす。
			if j == 0 || j == 1 {
				continue
			}
			// 割り切れたので合成数となる。
			if i%j == 0 {
				checkArray[i] = false
				break
			}
		}
	}
	// 素数の数だけメモリを確保。
	resultArray := make([]uint, countPrimeNum)
	// 素数だった要素をまとめる。
	for i, j := uint(0), 0; i < uint(len(checkArray)); i++ {
		if checkArray[i] {
			resultArray[j] = i
			j++
		}
	}
	return resultArray, nil
}

// 指定した自然数以下の素数を配列で返す。
func createPrimeNumberArray2(naturalNumber uint64) ([]uint, error) {
	// 効率的な方法の場合。
	// 0は自然数でないので飛ばす。
	if naturalNumber == 0 {
		err := errors.New("this number is not natural Number")
		return nil, err
	}
	// indexは0から始まるので1ずれる。
	checkArray := make([]bool, naturalNumber+1)
	countArray := make([]uint, naturalNumber+1)
	// 素数はtrue, 合成数はfalseとして、
	// 合成数か素数かわからない値はとりあえずfalseとしておく。
	for i := range checkArray {
		checkArray[i] = true
	}
	checkArray[0] = false
	checkArray[1] = false
	// インデックスの値で初期化。
	for i := range countArray {
		countArray[i] = uint(i)
	}
	// 素数の数をカウントする整数
	countPrimeNum := uint(0)
	for i := range checkArray {
		if i == 0 || i == 1 {
			continue
		}
		root := math.Sqrt(float64(i))
		// 素数かもしれない数字を順番に割っていくことで判定
		for j := range countArray {
			// 素数かもしない数字nの平方根までの数字で合成数、素数の判定ができるので
			// その数字よりも大きくなったらループを抜ける。
			// その場合は素数となる。
			if int(root) < j {
				countPrimeNum++
				break
			}
			// 0, 1は素数判定に使えないので飛ばす。
			if j == 0 || j == 1 {
				continue
			}
			// 割り切れたので合成数となる。
			if i%j == 0 {
				checkArray[i] = false
				break
			}
		}
	}
	// 素数の数だけメモリを確保。
	resultArray := make([]uint, countPrimeNum)
	// 素数だった要素をまとめる。
	for i, j := uint(0), 0; i < uint(len(checkArray)); i++ {
		if checkArray[i] {
			resultArray[j] = i
			j++
		}
	}
	return resultArray, nil
}

// 指定した自然数以下の素数を配列で返す。
func eratosthenePrimeNumberArray(naturalNumber uint64) ([]uint, error) {
	if naturalNumber == 0 {
		err := errors.New("this number is not natural Number")
		return nil, err
	}
	// indexは0から始まるので1ずれる。
	checkArray := make([]bool, naturalNumber+1)
	countArray := make([]uint, naturalNumber+1)
	// 素数はtrue, 合成数はfalseとして、
	// 合成数か素数かわからない値はとりあえずfalseとしておく。
	for i := range checkArray {
		checkArray[i] = true
	}
	checkArray[0] = false
	checkArray[1] = false
	// インデックスの値で初期化。
	for i := range countArray {
		countArray[i] = uint(i)
	}
	// 素数の数をカウントする整数
	countPrimeNum := uint(0)
	for p := uint(2); math.Pow(float64(p), 2) < float64(len(checkArray)); {
		for q := math.Pow(float64(p), 2); q < float64(len(checkArray)); q += float64(p) {
			// pの２乗+pは合成数
			checkArray[uint(q)] = false
		}
		// forループの最後に
		// マークが終わったら、リストにある次の素数かもしれない数字に移る。
		for i, v := range checkArray {
			if v && p < uint(i) {
				p = uint(i)
				break
			}
		}
	}
	// 素数だった個数カウント
	for _, v := range checkArray {
		if v {
			countPrimeNum++
		}
	}
	// 素数の数だけメモリを確保。
	resultArray := make([]uint, countPrimeNum)
	// 素数だった要素をまとめる。
	for i, j := uint(0), 0; i < uint(len(checkArray)); i++ {
		if checkArray[i] {
			resultArray[j] = i
			j++
		}
	}
	return resultArray, nil
}

func main() {

	num := uint64(100)
	array1, err := createPrimeNumberArray(num)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Printf(fmt.Sprintf("array1: %v\n", array1))

	array2, err := createPrimeNumberArray2(num)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Printf(fmt.Sprintf("array2: %v\n", array2))
	array3, err := eratosthenePrimeNumberArray(num)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Printf(fmt.Sprintf("array3: %v\n", array3))
}
