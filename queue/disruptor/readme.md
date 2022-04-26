并发安全的队列，简单实现，加上锁即可；但是锁粒度大并发度会比较低，同一时刻仅允许一个存或者取操作。
优化手段：
- 可以使用分片锁，减少锁粒度。
- 基于 cas
这里参考 `java Disruptor` 和 `go chan` 实现一个基于 cas 的高性能并发安全队列