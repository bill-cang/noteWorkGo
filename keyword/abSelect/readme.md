本文基于 go 1.14

select 允许在一个goroutine中管理多个channel。但是，当所有channel同时就绪的时候，go需要在其中选择一个执行。
go还需要处理没有channel就绪的情况，我们先从就绪的channel开始。

Order
select 不会按照任何规则或者优先级选择到达的channel。go标准库在每次访问的时候，
都会将他们顺序打乱，也就是说不能保证任何顺序。