package main

/*
struct A {
	int i;
	float f;
	int type;
	float _type; // 将屏蔽CGO对 type 成员的访问
};
*/
import "C"
import "fmt"

func main() {
	var a C.struct_A
	a.i = 10
	a.f = 10.2
	a._type = 10
	fmt.Println(a._type)
}
