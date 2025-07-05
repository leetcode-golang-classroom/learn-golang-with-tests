# concurrency

所謂的併發是指同時間處理多件事情的能力，但不一定是指同時一起進行。也就是可以透過類似 buffer 的方式來接收起來，然後透過一個 worker 逐步執行

## 說明

golang 採用 user level 的執行緒，並且由 goroutine scheduler 來分配執行任務給 os level 的執行序來執行

