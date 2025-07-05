# select

當遇到多個 goroutine 需要溝通時，可以透過 channel 來傳遞資料
而當遇到多個資料需要傳遞給一個 goroutine 時

這時可以透過 select 語法來最有效的等候多個 channel 的傳遞資訊

