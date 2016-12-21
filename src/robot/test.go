package main

// import (
// 	"fmt"
// 	"net/http"
// 	"reflect"
// )

// func main() {
// 	c := &http.Client{}
// 	fooType := reflect.TypeOf(c)
// 	for i := 0; i < fooType.NumMethod(); i++ {
// 		method := fooType.Method(i)
// 		fmt.Println(method.Name)
// 	}
// }

// import "github.com/markcheno/go-talib"

// type talib struct{
//     . talib
// }

// func main() {
// 	vm := otto.New()

// 	vm.Set("talib.test", test)
// 	vm.Run(`talib.test();`)
// }

// func test() {
// 	log.Print("this is from go")
// }

// func main() {
// 	var hello = func(name string) (string, error) {
// 		return fmt.Sprintf("Hello World, %s", name), errors.New("this is my error")
// 	}

// 	vm := otto.New()

// 	vm.Set("hello", hello)

// 	_, err := vm.Run(`
//     name = hello("otto");
//     console.log("name[0]:",name[0]);
//     if(typeof(name[1]) !="undefined"){
//         console.log("name[1].message:",name[1].Error());
//     }
//     `)

// 	if err != nil {
// 		panic(err)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math"

// 	"github.com/d4l3k/talib"
// )

// func main() {
// 	fmt.Println(talib.Sin([]float64{0, math.Pi / 2}))
// 	// => [0 1]
// }
