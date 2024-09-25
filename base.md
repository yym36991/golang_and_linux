基本语法学习
# 1、new和make的区别
<details>
  <summary>查看/隐藏代码</summary>
  ```go
  入参和出参不同，new接收一个形参，返回一个类型的指针，make 用来构造slice,map,chan类型，等引用类型。比如slice包括一个数据结构：返回底层数组的指针， 长度，和容量
  使用场景不同
</details>

# 2、切片和数组的区别？
<details>
  <summary> 查看/隐藏</summary>
  切片是引用，数组是底层存储结构，
  切片会动态扩容，如果A切片有引用B，则引用的B不会随着扩容而变化，但是A变成了新扩容的新地址
</details>

# 3、for range 底层地址发生改变吗？
<details>
  <summary> 查看/隐藏</summary>
  不会，这个遍历是值遍历，每次都新创建当前元素的副本，新建元素副本的地址会变化，但是原来slice不会变化。
</details>
