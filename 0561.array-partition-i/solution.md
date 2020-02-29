### 复杂度

执行用时 : 52 ms, 99.36%
内存消耗 : 6.7 MB, 6.17%


### 解题思路

感谢[❤❤❤花花酱的视频解说❤❤❤](https://zxi.mytechroad.com/blog/hashtable/leetcode-561-array-partition-i/)。我这里只是用Golang实现了他的C++思路。


这道题可以用简单排序方法将输入排序，然后从第一位开始，隔一个数取出求和就能得到答案。
比如：
```
[1,4,3,2]  排序后 [1,2,3,4]

然后取[1,3] 求和得出 4
```
使用标准库的排序算法一般是O(nlogn),时间复杂度高。

花花酱提供一个线性复杂度的解法。私用的是*记数排序*的思想。这种排序方法达到了O(n),但是额外的记数统计表增加了空间复杂度。因为这道题的输入值范围是[-10000, 10000],所以可以使用记数排序。

例如
对于输入 [5,5,3,2,4,2]

得到的输入表
```
输入   [0, 2, 1, 1, 2]
位置   [0, 1, 2, 3, 4]
```
之后我们从左到右循环输入表:
* 如果在位置`i`上这个数统计为0，我们跳过。
* 如果不为0，而且是第一次见到，则把他加入答案。这里需要有一个标记是否是第一次见到。统计之后将这个`i`位置上的数递减。

### 示例
*输入规模为正整数[0,4]*，题目中可以为复数，所以需要整体像右边shift一下。

```
input   [0, 1, 1, 1, 1]
count   [0, 1, 2, 3, 4]
```

初始化
ans = 0
n = 0
maxValue = 4
first = true

1.  n = 0, count[0] = 0 => n++
2.  n = 1, count[1] = 1 , first = true => ans += 1, first = false, count[1]--
3.  n = 1, count[1] = 0 => n++
4.  n = 2, count[2] = 1, first = false => first = true, count[2]-- 
5.  n =2, count[2] = 0 => n++
6.  n = 3, count[3] = 1, first = true => ans += 3, first=false, count[3]--
7.  n = 4, count[4] = 1, first=false => first = true, count[4]--
   
推出循环， ans = 4 返回。
这里的原理就是利用记数统计，如果是第一次见到这个数，就加到答案。如果不是，就忽略。因为记数排序已经排好了



### 代码

```golang
func arrayPairSum(nums []int) int {

	// The maximum of inputs' range.
	const maxValue = 10000
	// The count table. Since the range is [-10000,10000],
	// we will shift every value right 10000.
	count := make([]int, 2*maxValue+1)

	// Do counting.
	for _, v := range nums {
		count[v+maxValue]++
	}

	ans := 0
	n := -maxValue
	first := true

	for n <= maxValue {
		if count[n+maxValue] == 0{
			n++
			continue
		}
		if first {
			ans += n
			first = false
		} else {
			first = true
		}
		count[n+maxValue]--
	}

	return ans
}

```