package main

import (
	"GO_code/statusCode"
	"fmt"
	"sync"
)

type Stock struct {
	count int32
	Name  string
	mu    sync.Mutex
}

func (st *Stock) Add(inc int32) string {
	//st.mu.Lock()
	//defer st.mu.Unlock()
	st.count += inc
	return "增加成功"
}

func (st *Stock) Dec(inc int32) (string, int8) {
	//st.mu.Lock()
	//defer st.mu.Unlock()
	if st.count < 0 {
		return "扣减失败，没有库存了", statusCode.STOCK_DEC_FAIL
	} else {
		st.count -= inc
		return "扣减成功", statusCode.STOKC_DEC_SUCCESS
	}
}

func (st *Stock) q(inc int32, dec int32, wg *sync.WaitGroup) {
	defer wg.Done()
	st.Add(inc)
	s, ok := st.Dec(dec)
	if ok == 0 {
		fmt.Println(s)
	}
}

func main() {
	var stock Stock
	stock.count = 5
	var wg sync.WaitGroup
	wg.Add(3)
	go stock.q(4, 1, &wg)
	go stock.q(1, 2, &wg)
	go stock.q(2, 7, &wg)

	wg.Wait()

	fmt.Println(stock.count)
}
