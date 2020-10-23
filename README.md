# 数据结构与算法养成记


## 索引

  - [常用库](#常用库) 
  - [常用技巧](#常用技巧)
  - [数据结构与算法](#数据结构与算法)
  - [面试注意点](#面试注意点)

## 常用库

- go 通过切片模拟栈和队列
    ```go
        // 创建栈
        stack:=make([]int,0)
        // push压入
        stack=append(stack,10)
        // pop弹出
        v:=stack[len(stack)-1]
        stack=stack[:len(stack)-1]
        // 检查栈空
        len(stack)==0
  
  
        // 创建队列
        queue:=make([]int,0)
        // enqueue入队
        queue=append(queue,10)
        // dequeue出队
        v:=queue[0]
        queue=queue[1:]
        // 长度0为空
        len(queue)==0
    ```
    -  参数传递，只能修改，不能新增或者删除原始数据
    - 默认 s=[0:len(s)],取下限不取上限，数学表示为：[)
    
- 字典
    ```go
        // 创建
        m:=make(map[string]int)
        // 设置kv
        m["hello"]=1
        // 删除k
        delete(m,"hello")
        // 遍历
        for k,v:=range m{
        	println(k,v)
        }
    ```
    - map键需要可比较，不能为slice、map、function
    - map值都有默认值，可以直接操作默认值，如：m[age]++值由0变为1
    - 比较两个map需要遍历，其中的kv是否相同，因为有默认值关系，所以需要检查val和ok两个值

- 标准库
    - sort
    ```go
      // int排序
      sort.Ints([]int{})
      // 字符串排序
      sort.Strings([]string{})
      // 自定义排序
      sort.Slice(s,func(i,j int)bool{return s[i]<s[j]})
    ```
    - math
    ```go
      // int32 最大最小值
      math.MaxInt32 // 实际值：1<<31-1
      math.MinInt32 // 实际值：-1<<31
      // int64 最大最小值（int默认是int64）
      math.MaxInt64
      math.MinInt64
    ```
    - copy
    ```go
      // 删除a[i]，可以用 copy 将i+1到末尾的值覆盖到i,然后末尾-1
      copy(a[i:],a[i+1:])
      a=a[:len(a)-1]
      
      // make创建长度，则通过索引赋值
      a:=make([]int,n)
      a[n]=x
      // make长度为0，则通过append()赋值
      a:=make([]int,0)
      a=append(a,x)
    ```

## 常用技巧

- 类型转换
    ```go
    // byte转数字
    s="12345"  // s[0] 类型是byte
    num:=int(s[0]-'0') // 1
    str:=string(s[0]) // "1"
    b:=byte(num+'0') // '1'
    fmt.Printf("%d%s%c\n", num, str, b) // 111
    
    // 字符串转数字
    num,_:=strconv.Atoi()
    str:=strconv.Itoa()
    ```

## 数据结构与算法

> 数据结构是一种数据的表现形式，如链表、二叉树、栈、队列等都是内存中的一段数据表现的形式。
> 算法是一种通用的解决问题的模板或思路，大部分数据结构都有一套通用的算法模板，所以掌握这些通用的算法模板即可解决各种算法问题。

- 回溯法
    - 
    - 通过不停的选择，撤销选择，来穷尽所有可能性，最后将满足条件的结果返回。
    - 穷尽所有可能性，算法模板如下：
    ```go
      result = []
      func backtrack(选择列表,路径):
          if 满足结束条件:
              result.add(路径)
              return
          for 选择 in 选择列表:
              做选择
              backtrack(选择列表,路径)
              撤销选择
    ```
    - 典型应用
    > 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
    ```go
      func subsets(nums []int) [][]int {
          // 保存最终结果
          result := make([][]int, 0)
          // 保存中间结果
          list := make([]int, 0)
          backtrack(nums, 0, list, &result)
          return result
      }
      
      // nums 给定的集合
      // pos 下次添加到集合中的元素位置索引
      // list 临时结果集合(每次需要复制保存)
      // result 最终结果
      func backtrack(nums []int, pos int, list []int, result *[][]int) {
          // 把临时结果复制出来保存到最终结果
          ans := make([]int, len(list))
          copy(ans, list)
          *result = append(*result, ans)
          // 选择、处理结果、再撤销选择
          for i := pos; i < len(nums); i++ {
              list = append(list, nums[i])
              backtrack(nums, i+1, list, result)
              list = list[0 : len(list)-1]
          }
      }
    ```


## 面试注意点

- 快速定位到题目的知识点，找到知识点的**通用模板**，可能需要根据题目**特殊情况做特殊处理**。
- 先去朝一个解决问题的方向！**先抛出可行解**，而不是最优解！先解决，再优化！
- 代码的风格要统一，熟悉各类语言的代码规范。
    - 命名尽量简洁明了，尽量不用数字命名：i1、node1、a1、b2
- 常见错误总结
    - 访问下标时，不能访问越界。
    - 空值nil问题 run time error
