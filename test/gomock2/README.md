
**Reference**
- https://juejin.im/post/6844903853532381198

```
mockgen -source user.go -destination user_mock.go -package user
```

---

gomock 使用流程
1. 用 mockgen 为要模拟的接口生成模拟；
2. 在测试中通过 gomock.NewController() 闯将一个 ctrl 实例，并将其传递给模拟对象构造函数以获取模拟对象(mockObject)；
3. 调用 mockObject.EXPECT() 为你的模拟设置他们的期望值和返回值；
4. 调用 ctrl.Finish() 模拟控制器来断言模拟的期望；

关注点
- 期望输入值
- 期望输出值
- 期望调用次数
- 期望调用顺序
