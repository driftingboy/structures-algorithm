package disruptor

var a = make(chan int)

// 实现一个无锁队列，性能优于chan
//https://github.com/smarty-prototypes/go-disruptor
//https://github.com/LMAX-Exchange/disruptor/tree/master/src
