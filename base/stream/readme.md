Java 8 中的 Stream 是对集合（Collection）对象功能的增强，它专注于对集合对象进行各种非常便利、高效的聚合操作（aggregate operation），
或者大批量数据操作 (bulk data operation)。Stream API 借助于同样新出现的 Lambda 表达式，极大的提高编程效率和程序可读性。
同时它提供串行和并行两种模式进行汇聚操作，并发模式能够充分利用多核处理器的优势，使用 fork/join 并行方式来拆分任务和加速处理过程。


## 什么是聚合操作


## operation type
- Intermediate
map(mapToInt, flatMap...), filter,distinct,sorted,peek,limit,skip,parallel,sequential, unordered

- Terminal
forEach, forEachOrdered, toArray, reduce,collect, min,max,count,anyMatch,allMatch, noneMatch,findFirst,findAny,iterator

- Short-circuiting
anyMatch, allMatch, noneMatch, findFirst,findAny, limit


