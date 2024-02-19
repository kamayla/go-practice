package main

import "fmt"

type Occupation int

const (
	DENTIST Occupation = (iota + 1) * 10
	DENTAL_HYGIENIST
)

func (o Occupation) String() string {
	switch o {
	case DENTIST:
		return "歯科医師"
	case DENTAL_HYGIENIST:
		return "歯科衛生士"
	default:
		return "不明"
	}

}

func main() {
	o := Occupation(10)
	if o == DENTIST {
		fmt.Println(o)
		a := o.String()
		fmt.Println(a)
	}
}
