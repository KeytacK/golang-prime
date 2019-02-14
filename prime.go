/*

1-100までの素数を表示する

参考：
[お気楽 Go 言語プログラミング入門](http://www.geocities.jp/m_hiroi/golang/abcgo02.html)

*/
package main

import (
	"fmt"
)

var (
	// 見つけた素数を保存する(最大100個)
	primeTable [100]int

	// 配列サイズ(素数配列に値が入っている個数)
	primeSize int
)

// 素数のチェックを行う
func checkPrime(n int) (isPrime bool) {
	isPrime = true

	// 既に素数と確定している配列を(配列サイズ)-1回走査する
	for i := 1; i < primeSize; i++ {
		p := primeTable[i]

		// ループ回数を減らすため、√nより小さい数のみ参照する
		if p*p > n {
			break
		}

		// 素数判定 調査対象の数が以前の素数で割り切れる場合false
		if n%p == 0 {
			isPrime = false
			break
		}

	}

	return
}

func main() {
	primeTable[0] = 2
	primeSize = 1

	// 3から始めて奇数番目の数を調査する(偶数は2で割り切れる事が自明のため)
	for n := 3; n <= len(primeTable); n += 2 {
		// 調査対象の数が素数の場合、素数配列に追加し配列サイズを増やす
		if checkPrime(n) {
			primeTable[primeSize] = n
			primeSize++
		}
	}

	// 素数配列を出力する
	for i := 0; i < primeSize; i++ {
		fmt.Print(primeTable[i], " ")
	}
	fmt.Print("\n")
}
