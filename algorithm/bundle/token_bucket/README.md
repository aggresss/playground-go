
- https://zhuanlan.zhihu.com/p/64124008

LeakyBucket 的核心思想就是按固定的速率处理请求，所以不支持突增的流量。因为即使有再多的流量，也是按固定的速率被处理。他与 TokenBucket 的区别是 TokenBucket 是按固定速率产生 Token，请求进来的时候只要有 Token 就能立即被处理，不用等待。只有在无 Token 时才会等待。

