
# Arrays 與 Slices

在 golang 中，陣列類型資料類型主要有兩種 array 與 slice 。

## array
array 是一種固定長度的資料陣列。在初始化時，就設定好長度，並且無法後續修改。

```golang=
arr := [4]int{1,1,1,1}
arr := [...]int{1,1,1,1}
```

![image](https://hackmd.io/_uploads/HkY6BHIbex.png)

## slice

slice 是則是一種動態長度的資料陣列。在初始化，可以設定好容量 capacity 。可以透過 append 新增資料到 slice 上。而實際上， slice 內部是指向一個 array 資料型態的參考，遇到超過預設 capacity 的資料量時，就會產生新的二倍於原本 capacity 的陣列實體，然後把 slice 參考更新。

```golang=
slice := []int{1,1,1,1}
// slice 的預設的 capcity 會是當下資料量長度，當不指定時
// slice 基本有3個屬性 ptr 指向目前資料陣列參照， len 代表當下資料長度， cap 代表當下最多容量。
```

![image](https://hackmd.io/_uploads/rJixLSI-xg.png)

# iteration

```golang=
stockAmountByDay := []int{200,210, 300,311}
```
## loop over by index

```golang=
sliceLength := len(stockAmountByDay)
for idx := 0; idx < sliceLength; idx++ {
   stockAmount := stockAmountByDay[idx]
   fmt.Printf("stockAmount = %d\n",stockAmount)
}  
```
[iteration-slice-by-idx](https://go.dev/play/p/w_iYwtB6u6c)

## loop over by range

```golang=
for _, stockAmount := range stockAmountByDay {
   fmt.Printf("stockAmount = %d\n",stockAmount)
}
```

[iteration-slice-by-range](https://go.dev/play/p/CYLQDp2p4mY)

# operation

## append

```golang=
stockAmountByDay = append(stockAmountByDay, 170)
```
![image](https://hackmd.io/_uploads/S1pQsT8-xl.png)
[append-example](https://go.dev/play/p/843TDDp3N6I)

## sub range reference
```golang=
stockAmountByDayRange := stockAmountByDay[1:4]
```

![image](https://hackmd.io/_uploads/rJow3a8Wge.png)
[sub-range-reference](https://go.dev/play/p/K9x8MOBA3gh)


## copy
```golang=
stockAmountByRangeClone := make([]int, len(stockAmountByRange))
copy(stockAmountByRangeClone, stockAmountByRange)
```

![image](https://hackmd.io/_uploads/rkGH8C8Zge.png)

[copy-example](https://go.dev/play/p/4v2qE0jikrb)


# 參考文件

[slice-intro](https://go.dev/blog/slices-intro)

# install debug tool

```shell
go install github.com/go-delve/delve/cmd/dlv@latest
```