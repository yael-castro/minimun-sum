# Minimun sum

Given an array of integers, perform some number *k* of operations.
Each operation consists of removing an element from the array, dividing it by 2,
and inserting the ceiling of that result back into the array.

*Minimize the sum of the elements in the final array*

###### Example
nums = [10, 20, 7]<br>
k = 4

```
Pick    Pick/2  Ceiling     Result
7       3.5     4           [10, 20, 4]
10      5       5           [5 , 20, 4]
20      10      10          [5 , 10, 4]
10      5       5           [5 , 5 , 4]
```
The sum of the final array is *5 + 5 + 4 = 14*, and that sum is minimal.

###### Instructions
Complete the function *minSum* in the *main.go* file

###### Check the solution
```shell
go test -v
```

###### Check solution performance
```shell
go test -bench . -run ^$
```

###### Solution performance
```
goos: linux
goarch: amd64
pkg: github.com/yael-castro/minimun-sum
cpu: AMD Ryzen 7 3700U with Radeon Vega Mobile Gfx
BenchmarkMinSum-8              2        1011610204 ns/op
```
```
goos: linux
goarch: amd64
pkg: github.com/yael-castro/minimun-sum
cpu: AMD Ryzen 7 3700U with Radeon Vega Mobile Gfx
BenchmarkMinSum-8              1        1009105070 ns/op
```

###### Notes
The three variable values that affect the execution time of the
function are *num*, *k*, and the range of values in *num*