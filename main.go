package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// data := "cat,bird,dog"

	// // Split on comma.
	// result := strings.Split(data, ",")

	T := []struct {
		Version string
		Type    string
		Result  string
	}{
		{"1.0.1", "MAJOR", ""},
		{"1.2.3", "PATCH", ""},
		{"2.9.4", "MINOR", ""},
	}

	for i, v := range T {

		//var result_version string

		tpe := ""

		fmt.Println("i = ", i, "  =   ", T[i].Type)

		tpe = T[i].Type

		TT := strings.Split(v.Version, ".")

		fmt.Println(TT)
		v0 := TT[0]
		v1 := TT[1]
		v2 := TT[2]

		//i1, err := strconv.Atoi(str1)
		vv0, _ := strconv.Atoi(TT[0])
		vv1, _ := strconv.Atoi(TT[1])
		vv2, _ := strconv.Atoi(TT[2])

		fmt.Println(v0)
		fmt.Println(v1)
		fmt.Println(v2)

		fmt.Println(vv0)
		fmt.Println(vv1)
		fmt.Println(vv2)

		vv0 += 1
		vv1 += 1
		vv2 += 1

		fmt.Println(vv0)
		fmt.Println(vv1)
		fmt.Println(vv2)

		vvv0 := strconv.Itoa(vv0)
		vvv1 := strconv.Itoa(vv1)
		vvv2 := strconv.Itoa(vv2)

		fmt.Println(vvv0)
		fmt.Println(vvv1)
		fmt.Println(vvv2)

		if tpe == "MAJOR" {
			fmt.Println("===  Major")

			result_version := vvv0 + "." + v1 + ".0" //+ v2
			fmt.Println("-----   v.version[0] = ", result_version)
			T[i].Result = result_version
		}

		if tpe == "PATCH" {

			result_version := v0 + "." + vvv1 + "." + v2
			fmt.Println("-----   v.version[1] = ", result_version)
			T[i].Result = result_version
		}

		if tpe == "MINOR" {
			result_version := v0 + "." + v1 + "." + vvv2
			fmt.Println("-----   v.version[2] = ", result_version)

			T[i].Result = result_version

		}
	}

	fmt.Println(T)

}
