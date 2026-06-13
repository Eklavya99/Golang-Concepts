package golangconcepts

import "fmt"

type BaseStruct struct {
	Name string
}

func (b *BaseStruct) Log(msg string) {
	fmt.Printf("[%s] Log: %s\n", b.Name, msg)
}

type OrderStruct struct{
	BaseStruct // composition via embedding, Now OrderStruct has access to property Name and method Log
	DbConnStr string
}
func GoStructs() {
	svc := OrderStruct{
		BaseStruct: BaseStruct{Name: "OrderProcessor"},
		DbConnStr: "postgres://.../.../",
	}
	fmt.Println(svc.Name)
	svc.Log("Processing Order#15633")
}