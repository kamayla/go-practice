package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}

	// indexが1から(2-1)を取得
	// すなわち[2]が取得される
	fmt.Println(src[1:2])

	// indexが1から最後まで取得
	fmt.Println(src[1:])

	// 先頭から(3-1)まで取得
	fmt.Println(src[:3])

	// 先頭から最後まで取得
	// 固定長の配列をスライスするときによく使う
	fmt.Println(src[:])

	// スライスの長さを取得
	fmt.Println(len(src))
	// スライスの容量を取得
	fmt.Println(cap(src))

	// indexが1から(4-1)まで,capを4で取得
	// capで指定できるのは元のスライスの長さまで
	result := src[1:3:4]
	fmt.Println(result)
	fmt.Println(len(result))
	fmt.Println(cap(result))

	data := []int{1, 2, 3, 4, 5}

	// 容量は５
	fmt.Println(cap(data))

	data = append(data, 10)

	// 容量maxを超えて一つでも追加されると容量は倍になる
	fmt.Println(cap(data))

}
