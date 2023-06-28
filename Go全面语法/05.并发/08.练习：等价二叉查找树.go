//1. 实现 Walk 函数。
//2. 测试 Walk 函数。
//函数 tree.New(k) 用于构造一个随机结构的已排序二叉查找树，它保存了值 k, 2k, 3k, ..., 10k。
//创建一个新的信道 ch 并且对其进行步进：
//go Walk(tree.New(1), ch)
//然后从信道中读取并打印 10 个值。应当是数字 1, 2, 3, ..., 10。
//3. 用 Walk 实现 Same 函数来检测 t1 和 t2 是否存储了相同的值。
//4. 测试 Same 函数。
//Same(tree.New(1), tree.New(1)) 应当返回 true，而 Same(tree.New(1), tree.New(2)) 应当返回 false。