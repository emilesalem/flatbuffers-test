package main

import (
	"fmt"
	"test/flatbuffers/internal/fb/a"
	"test/flatbuffers/internal/fb/b"

	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
	builder := flatbuffers.NewBuilder(256)

	field1 := builder.CreateString(`some string`)

	a.AStart(builder)

	a.AAddFieldOne(builder, field1)

	builder.Finish(a.AEnd(builder))

	msgA := builder.FinishedBytes()

	msgB := b.GetRootAsB(msgA, 0)

	fmt.Println(string(msgB.Field1()))
}
