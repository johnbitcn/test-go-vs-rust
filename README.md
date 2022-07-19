# Goland vs Rust 性能测试

这个测试是计算规划求解问题。
一共要执行 $2^{30}=1,073,741,824$个循环，每个循环还要依次判断本次循环数的二进制位中的1。
如果位数为1则执行`sum+=1`的动作。

本测试是使用了多线程cpu密集型计算的设计方法，给每一个物理cpu内核都分配一个线程，并且平均分配所有的循环。

## 测试结果

### 07/19/2022
好吧我错误了，问题找到了。
Rust 是在debug模式下运行的，所以效率降低非常多。
解决办法：把Rust 按照发布编译以后就去掉了Debug模式。
速度快的飞起…… 真的！

```shell
❯ target/release/for_rust
函数运行开始！
Max Loop:1073741824
Have CPUs:8
PerLoop:134217728
AddLoop:0

Cpu:0 Start:0 End:134217728
Cpu:2 Start:268435456 End:402653184
Cpu:3 Start:402653184 End:536870912
Cpu:1 Start:134217728 End:268435456
Cpu:5 Start:671088640 End:805306368
Cpu:4 Start:536870912 End:671088640
Cpu:6 Start:805306368 End:939524096
Cpu:7 Start:939524096 End:1073741824
函数运行时间: 0.000426秒
```

### 07/19/2022

可能是我的程序设计有问题，能力有限。
测试结果让我无法接受，Rust居然比Go慢了近10倍的速度。

**测试机型**

- 机型: MacBook air m1
- cpu: 8个物理核心

**for Golang**

```shell
函数运行开始！
Got CPUs: 8
Max Loops:1073741824
PerLoops:134217728
AddLoops:0
cpu:0 start:0 end:134217728
cpu:1 start:134217728 end:268435456
cpu:2 start:268435456 end:402653184
cpu:3 start:402653184 end:536870912
cpu:4 start:536870912 end:671088640
cpu:5 start:671088640 end:805306368
cpu:6 start:805306368 end:939524096
cpu:7 start:939524096 end:1073741823
函数运行时间：2.094695625s⏎
```

**for Rust**

```shell
函数运行开始！
Max Loop:1073741824
Have CPUs:8
PerLoop:134217728
AddLoop:0
Cpu:0 Start:0 End:134217728
Cpu:1 Start:134217728 End:268435456
Cpu:2 Start:268435456 End:402653184
Cpu:3 Start:402653184 End:536870912
Cpu:4 Start:536870912 End:671088640
Cpu:6 Start:805306368 End:939524096
Cpu:5 Start:671088640 End:805306368
Cpu:7 Start:939524096 End:1073741824
函数运行时间: 21.701秒
```

