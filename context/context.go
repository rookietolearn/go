package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4/metadata"
)

type Cmk struct {
}

func main() {
	tmap := map[string]string{
		"playerID": "21334455",
		"areaID":   "123",
	}
	ctx := context.TODO()
	ctxCommon := NewContext(ctx, tmap)
	ctxMetaData := NewMetaDataCtx(ctx, tmap)

	resp, _ := Get(ctxMetaData, `areaID`)
	mdresp, _ := Mget(ctxCommon, "areaID")
	fmt.Println("======")
	fmt.Println(resp)
	fmt.Println("======")
	fmt.Println(mdresp)
	fmt.Println("======")

}

func NewContext(ctx context.Context, v map[string]string) context.Context {
	return context.WithValue(ctx, Cmk{}, v)
}

func NewMetaDataCtx(ctx context.Context, v map[string]string) context.Context {
	return metadata.NewContext(ctx, v)
}

func Mget(ctx context.Context, k string) (string, bool) {
	return metadata.Get(ctx, k)
}

func Get(ctx context.Context, k string) (string, bool) {
	var resp string
	amap, ok := ctx.Value(Cmk{}).(map[string]string)
	if !ok {
		return "", false
	}
	for k1, v1 := range amap {
		if k1 == k {
			resp = v1
		}
	}
	return resp, true
}
