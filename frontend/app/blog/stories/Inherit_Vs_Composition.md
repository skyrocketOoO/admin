20240725
Software

作者：Tyrion Flores
链接：https://www.zhihu.com/question/511958588/answer/3137705299
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

因为组合优于继承。我们有一只鸭子和一只鸡，他们工作得很好。我们发现鸭子和鸡有很多重复的地方，他们都会飞，都有两只脚两个翅膀，都会唧唧或者嘎嘎叫。于是我们抽象出鸟这个父类，鸭子和鸡都继承了鸟这个父类， 当我们想要在飞的时候额外做点什么，只需要修改鸟就好了，代码得到了缩减，维护起来看似方便了。鸟工作得也很好。我们业务不断扩展，企鹅出现了。 它不会飞，但是会游泳。 鸟的工作出了问题，于是我们把飞行这个功能被下沉到了会飞的鸟类，企鹅继承自一个不会飞的鸟类。接下来橡皮鸭子出现了，人们对于它究竟是不是鸟有了争议。开始浪费时间大量的讨论什么是鸟，鸟该做些什么。（分割线）但我们的生活中没有鸟（请注意这句话），鸟是一个抽象， 我们生活中有鸡，有鸭。我们觉得他们有一些相同的地方，于是把拥有这些相同点的东西叫做鸟，但永远不知道下一个遇见的，能不能算鸟， 鸟的定义要不要修改。这就是继承不适用的原因，让我们看看组合会怎么样。我们找到了鸡和鸭的共同点， 会飞，两只脚，两个翅膀，会叫。 这些东西加上其他的特质『组合』成了鸡或鸭。 会飞这个能力就能提出来，使用在每一个需要飞行能力的地方。 当我遇到企鹅，就不用拿飞行来『组合』它。飞行，不应该是鸡或鸭从父类继承的能力，而应该是『飞行能力』组合成了鸡鸭的一部分。

这是因为抽象能力不佳，陷入自己的逻辑误区，鸟原本就是一个符号，代表特定功能的组合，企鹅不会飞为什么要继承鸟，这不是继承的原罪，jdk的代码就没有这样的坏味道，只能说在工程演进的过程中，继承需要被及时重构，但是很难实现，造成最终滥用，越来越四不像

就是继承的原罪，继承是本末倒置了事物和事物的特点的关系，并不是事物派生于事物的理型，而是事物共有的特点总结出理型，继承通过事物继承于基类(理型)就是一种本末倒置

继承的原则是：当一个东西它确实是某一个东西的时候才集成它(is-a)。企鹅是不是鸟？正方形是不是矩形？这都算是经典问题了，但你实际开发中会浪费大量时间去考虑到底是不是is-a，还会判断错误。

哪来的这么多理由，就算组合确实优于继承，你提供组合的同时也提供继承，让程序员自己来选择，而不是程序语言设计者用自己的口味替程序员选择！

实际上就是语言设计者偷懒，不想提供这么多特性，不给程序员自己选择。例如 Java，潜台词就是：你们这一群渣渣，一个指针都用不好学不会，我不提供（go 语言就提供了指针）；操作符重载估计你们也学不会，通通用函数吧；多继承你们学不会，那就单继承吧。。。等等等等。

rust 则是：每次你们都容易搞出内存泄漏，我来替你们规定严格的生命周期卡一卡你们；继承你们也学不会用不好（Java 为什么允许继承？）则禁止继承；还有一点，多线程这一群渣渣程序员也搞不定，那我就规定严格的生命周期一并把多线程容易出错的部分也帮你们选择了。

而 C/C++ 则是另一个极端：你们都已经是程序员了，你们想要的功能都通通给，功能越多越好，想要的概念全部都提供，至于你们用这些组合出问题，你们自己搞定。甚至一些语法糖也越来越多（例如 C++17 的结构化绑定就是明显的语法糖）。

实际工程中我看到很多人在企鹅的飞行那放了个空实现让代码能跑起来，然后跟测试讲反正企鹅不能飞，你让他飞也飞不起来

复杂性不会消失只会转移，这种情况下，使用继承确实会出现父类有上百万接口的情况，但组合的情况下，你只需要实现你需要的最少接口——而且如果可以复用，你就不需要实现，也不用担心多继承，菱形继承等复杂情况

面向对象编程语言的问题在于，它总是附带着所有它需要的隐含环境。你想要一个香蕉，但得到的却是一个大猩猩拿着香蕉，而其还有整个丛林。” — Joe Armstrong（Erlang语言发明人）