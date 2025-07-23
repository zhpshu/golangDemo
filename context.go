context 是在 Go 语言 1.7 版本才正式被纳入官方标准库，为什么今天要介绍 context 的使用方式呢？原因很简单：在初学 Go 时，写 API 的过程中，经常会看到 HTTP Handler 的第一个参数就是 ctx context.Context。那么，这个 context 在这里使用的目的是什么？它的含义又是什么？

本篇就是要带大家了解什么是 context，以及它的使用场景和方法。内容不会涉及 context 的源码实现，而是通过几个实际例子来帮助理解。

1. 使用 WaitGroup
学习 Go 时，肯定要掌握如何使用并发（goroutine）。而开发者该如何控制并发呢？其实有两种方式：一种是 WaitGroup，另一种就是 context。

那么什么时候需要用到 WaitGroup 呢？
很简单：当你需要把同一件事拆分成多个任务（Job）并行执行，并且必须等待所有任务都完成后，主程序才能继续执行时，就需要使用 WaitGroup。来看一个实际例子：

go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup

    wg.Add(2)
    go func() {
        time.Sleep(2 * time.Second)
        fmt.Println("job 1 done.")
        wg.Done()
    }()
    go func() {
        time.Sleep(1 * time.Second)
        fmt.Println("job 2 done.")
        wg.Done()
    }()
    wg.Wait()
    fmt.Println("All Done.")
}
从上面的例子可以看到，主程序通过 wg.Wait() 等待所有任务执行完毕后，才打印最后的消息。

但这里会遇到一个问题：虽然我们将任务拆分成多个并在后台运行，但用户如何通过其他方式终止这些正在运行的 goroutine 呢？（比如开发者常写的后台监控程序，需要长时间运行）

例如，UI 上有一个“停止”按钮，点击后如何主动通知并停止正在运行的任务？这很简单，可以使用 channel + select 的方式来解决。

使用 channel + select
go
package main

import (
    "fmt"
    "time"
)

func main() {
    stop := make(chan bool)

    go func() {
        for {
            select {
            case <-stop:
                fmt.Println("got the stop channel")
                return
            default:
                fmt.Println("still working")
                time.Sleep(1 * time.Second)
            }
        }
    }()

    time.Sleep(5 * time.Second)
    fmt.Println("stop the goroutine")
    stop <- true
    time.Sleep(5 * time.Second)
}
从上面可以看到，通过 select + channel 可以快速解决这个问题：只要在任意地方向 stop channel 发送一个 true 值，就可以停止后台正在执行的任务。

虽然使用 channel 能解决这个问题，但现在出现一个新的问题：如果后台启动了无数个 goroutine，甚至 goroutine 内部又启动了新的 goroutine，结构就会变得非常复杂。例如下面这种情况：

cancel
这时就无法再用简单的 channel 方式来统一管理了，而需要使用今天的核心主题：context。

认识 context
从上图可以看出，我们建立了三个 worker 节点来处理不同的任务。因此，在主程序最开始会声明一个根 context：context.Background()，然后每个 worker 节点分别创建自己的子 context。这样做的主要目的是：当关闭某个 context 时，可以直接取消该 worker 内部所有正在运行的任务。

我们用前面的例子进行改写：

go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    go func() {
        for {
            select {
            case <-ctx.Done():
                fmt.Println("got the stop channel")
                return
            default:
                fmt.Println("still working")
                time.Sleep(1 * time.Second)
            }
        }
    }()

    time.Sleep(5 * time.Second)
    fmt.Println("stop the goroutine")
    cancel()
    time.Sleep(5 * time.Second)
}
可以看到，这里只是把原来的 channel 替换成了 context，其他逻辑完全不变。我们使用了 context.WithCancel，其声明方式如下：

go
ctx, cancel := context.WithCancel(context.Background())
这样做的意义在于：每个 worker 都有独立的 cancel 函数，开发者可以在任意位置调用 cancel()，来决定停止哪一个 worker。这样就可以通过 context 实现对多个 goroutine 的统一取消控制。

再看一个实际例子：

go
package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    ctx, cancel := context.WithCancel(context.Background())

    go worker(ctx, "node01")
    go worker(ctx, "node02")
    go worker(ctx, "node03")

    time.Sleep(5 * time.Second)
    fmt.Println("stop the goroutine")
    cancel()
    time.Sleep(5 * time.Second)
}

func worker(ctx context.Context, name string) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println(name, "got the stop channel")
            return
        default:
            fmt.Println(name, "still working")
            time.Sleep(1 * time.Second)
        }
    }
}
上面的例子中，通过一个 context 就可以一次性停止多个 worker。

关键在于：如何声明 context，以及在什么时机调用 cancel()。我个人通常会将 context 与优雅关闭（graceful shutdown）结合使用，用来取消正在运行的任务，或关闭数据库连接等资源。

总结
WaitGroup 适用于“等待所有任务完成”的场景。
channel + select 适用于简单的任务取消。
context 更适合在复杂、嵌套的并发结构中，进行统一的生命周期管理和取消控制。
context 不仅能传递取消信号，还能传递超时、截止时间和请求范围的数据，是构建可维护、高并发 Go 程序的重要工具。
