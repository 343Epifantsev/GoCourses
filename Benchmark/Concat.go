// Concat
package main

import (
	"bytes"
	"fmt"
	"strings"
)

type Record struct {
	A string
	B string
	C string
	D string
	E string
	F string
}

func ConcatSprintf(obj Record) (string, error) {
	return fmt.Sprintf("%s_%s_%s_%s_%s_%s", obj.A, obj.B, obj.C, obj.D, obj.E, obj.F), nil
}
func ConcatJoin(obj Record) (string, error) {
	return strings.Join([]string{obj.A, obj.B, obj.C, obj.D, obj.E, obj.F}, "_"), nil
}
func ConcatBuffer(obj Record) (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString(obj.A)
	buffer.WriteString(obj.B)
	buffer.WriteString(obj.C)
	buffer.WriteString(obj.D)
	buffer.WriteString(obj.E)
	buffer.WriteString(obj.F)
	return buffer.String(), nil
}

func main() {
	fmt.Println("Hello World!")
}
