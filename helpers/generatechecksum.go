package helpers

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

// GenerateChecksum ...
func GenerateChecksum(base string, weight interface{}) int {
	digits := OnlyNumbers(base)

	value := reflect.ValueOf(weight)
	vtype := value.Type()
	kind := value.Kind()

	if kind != reflect.Int && kind != reflect.Slice {
		log.Fatal("Expected weights to be of type int or []int{}, but got ", vtype.Kind())
	}

	var w []int

	if kind == reflect.Slice {
		el := vtype.Elem()

		if el.Kind() != reflect.Int {
			log.Fatal("Expected weights to be of type int or []int{}, but got slice of ", el.Kind())
		} else {
			for i := 0; i < value.Len(); i++ {
				w = append(w, int(value.Index(i).Int()))
			}
		}
	}

	if kind == reflect.Int {
		for i := 0; i < len(digits); i++ {
			w = append(w, int(value.Int())-i)
		}
	}

	arr := strings.Split(digits, "")
	acc := 0
	for i, d := range arr {
		v, err := strconv.Atoi(d)

		if err != nil {
			log.Fatal(err)
		}

		acc = acc + v*w[i]
	}

	return acc
}
