package main

import (
	"context"
	"fmt"
	"go-micro.dev/v4/metadata"
)

type Cmk struct {
}

//type CommonCtx struct {
//	comctx context.Context
//	mdctx  context.Context
//}

func main() {
	tmap := map[string]string{
		"locate":   "md", //md,sys
		"playerID": "21334455",
		"areaID":   "123",
	}
	ctx := context.TODO()
	//ctxCommon := NewContext(ctx, tmap)
	//ctxMetaData := NewMetaDataCtx(ctx, tmap)

	cCtx := NewCommonCtx(ctx, tmap)

	areaID := CommonGet(cCtx, "areaID")

	//resp, _ := Get(ctxMetaData, `areaID`)
	//mdresp, _ := Mget(ctxCommon, "areaID")

	fmt.Println("======")
	fmt.Println(areaID)
	fmt.Println("======")
	fmt.Println("======")

}

func NewContext(ctx context.Context, v map[string]string) context.Context {
	return context.WithValue(ctx, Cmk{}, v)
}

func NewMetaDataCtx(ctx context.Context, v map[string]string) context.Context {
	return metadata.NewContext(ctx, v)
}

func NewCommonCtx(ctx context.Context, v map[string]string) context.Context {
	var commonCtx = context.TODO()
	locate := v["locate"]
	if locate == "md" {
		md := metadata.NewContext(commonCtx, v)
		commonCtx = context.WithValue(ctx, Cmk{}, md)
	} else {
		commonCtx = context.WithValue(ctx, Cmk{}, v)
		metadata.NewContext(commonCtx, nil)
	}
	return commonCtx
}

func CommonGet(ctx context.Context, k string) string {
	amap, ok := ctx.Value(Cmk{}).(map[string]string)
	var resp string
	//无metadata数据
	if ok {
		for k1, v := range amap {
			if k == k1 {
				resp = v
			}
		}
	}
	metaData, ok := ctx.Value(Cmk{}).(context.Context)
	if ok {
		mdresp, ok := metadata.Get(metaData, k)
		if !ok {
			return ""
		}
		resp = mdresp
	}
	return resp
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
