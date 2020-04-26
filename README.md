运行过程
1、Go env GOPATH
2、创建目录和文件
3、Go build 
4、go run hello.go
5、Go install

一些特性
1. 保留但大幅度简化指针
Go语言保留着C中值和指针的区别，但是对于指针繁琐用法进行了大量的简化，引入引用的概念。所以在Go语言中，你几乎不用担心会因为直接操作内寸而引起各式各样的错误。
2. 多参数返回
还记得在C里面为了回馈多个参数，不得不开辟几段指针传到目标函数中让其操作么？在Go里面这是完全不必要的。而且多参数的支持让Go无需使用繁琐的exceptions体系，一个函数可以返回期待的返回值加上error，调用函数后立刻处理错误信息，清晰明了。
3. Array，slice，map等内置基本数据结构
如果你习惯了Python中简洁的list和dict操作，在Go语言中，你不会感到孤单。一切都是那么熟悉，而且更加高效。如果你是C++程序员，你会发现你又找到了STL的vector 和 map这对朋友。
4. Interface
Go语言最让人赞叹不易的特性，就是interface的设计。任何数据结构，只要实现了interface所定义的函数，自动就implement了这个interface，没有像Java那样冗长的class申明，提供了灵活太多的设计度和OO抽象度，让你的代码也非常干净。千万不要以为你习惯了Java那种一条一条加implements的方式，感觉还行，等接口的设计越来越复杂的时候，无数Bug正在后面等着你。
同时，正因为如此，Go语言的interface可以用来表示任何generic的东西，比如一个空的interface，可以是string可以是int，可以是任何数据类型，因为这些数据类型都不需要实现任何函数，自然就满足空interface的定义了。加上Go语言的type assertion，可以提供一般动态语言才有的duck typing特性， 而仍然能在compile中捕捉明显的错误。
5. OO
Go语言本质上不是面向对象语言，它还是过程化的。但是，在Go语言中， 你可以很轻易的做大部分你在别的OO语言中能做的事，用更简单清晰的逻辑。是的，在这里，不需要class，仍然可以继承，仍然可以多态，但是速度却快得多。因为本质上，OO在Go语言中，就是普通的struct操作。
6. Goroutine
这个几乎算是Go语言的招牌特性之一了，我也不想多提。如果你完全不了解Goroutine，那么你只需要知道，这玩意是超级轻量级的类似线程的东西，但通过它，你不需要复杂的线程操作锁操作，不需要care调度，就能玩转基本的并行程序。在Go语言里，触发一个routine和erlang spawn一样简单。基本上要掌握Go语言，以Goroutine和channel为核心的内存模型是必须要懂的。不过请放心，真的非常简单。
7. 更多现代的特性
和C比较，Go语言完全就是一门现代化语言，原生支持的Unicode, garbage collection, Closures(是的，和functional programming language类似), function是first class object，等等等等。
