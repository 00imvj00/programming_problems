# programming_problems

Some of tricky programming problems and solutions.

## Array Problems

### Problem1

- Given an unsorted array of integers, find **a pair** with given sum in it.
  - Input: [8,7,2,5,3,1]
  - Sum: 10
  - OutPut:
    - (8, 2)
    - (7, 3)

### Problem2

- Given an unsorted array of integers, print **all sub-arrays** with sum 0.
- This problem can also be written as find all sub-arrays whose sum is given sum like 10.
- Remember that, in this case **sub-array has to be consicutive**. **Re-arrannging elements are not allowed**.
  - Input: [4,2,-3,-1,0,4]
  - Sum: 0
  - OutPut:
    - [-3,-1,0,4]
    - [0]

### Problem3

- Given an unsorted binary array, Sort it in linear time and constant space.
  - Input: [1,0,1,0,1,0,0,1]
  - OutPut: [0,0,0,0,1,1,1,1]

### Problem4

- Given an unsorted array, find total occurence count.

  - Input: [1,2,4,7,2,3,6,3,4,5,7,3,3,4,5]
  - OutPut: {1: 1, 2: 2, 3: 4, 4: 3, 7: 2,6: 1, 5: 2}

### Problem5 (Impl Pending)

- Given an array of integers, find largest sub-array formed by consecutive integers. The sub-array should contain all the distinct values.
  - Input: [2,1,2,1,4,3,1,0]
  - OutPut: [0,2,1,4,3]

### Problem6 (Impl Pending)

- Given an binary array containing 0s and 1s, Find maximum length sub-array having equal number of 0s and 1s.
  - Input: [0,0,1,0,1,0,0]
  - Output: [0,1,0,1]

### Problem7

- Given an array containing only 0s, 1s and 2s, Sort this array in linear time O(n) with constant space O(1).
- Fun fact, This is also known as **_Dutch National Flag Problem_**.
  - Input: [0,1,2,2,1,0,0,2,0,1,1,0]
  - Output: [0,0,0,0,0,1,1,1,1,2,2,2]

### Problem8 (Impl Pending)

- **Inplace merge two sorted arrays**:
  - Given two sorted arrays X[] &amp; Y[] of size m and n each,merge elements of X[] with array Y[] by maintaining the sorted order i.e. fill X[] with first m smallest elements and fill Y[] with remaining elements. The conversion should be done in-place without using any other data structure.
  - Input:
    - X = [ 1, 4, 7, 8, 10 ]
    - Y = [ 2, 3, 9 ]
  - Output:
    - X = [ 1, 2, 3, 4, 7 ]
    - Y = [ 8, 9, 10 ]

### Problem9 (Impl Pending)

- Given two sorted arrays X[] and Y[] of size m and n each where m &gt;= n and X[] has exactly n vacant cells, merge elements of Y[] in their correct position in array X[] i.e. merge (X, Y) by keeping the sorted order.
  - Input:
    - X = [ 0, 2, 0, 3, 0, 5, 6, 0, 0]
    - Y = [ 1, 8, 9, 10, 15 ]
  - Output:
    - X = [ 1, 2, 3, 5, 6, 8, 9, 10, 15 ]

### Problem10 (Impl Pending)

- **Find index of 0 to be replaced to get maximum length sequence of continuous ones**
  - Given a binary array, find the index of 0 to replaced with 1 to get maximum length sequence of continuous ones.We can efficiently solve this problem in linear time and constant space.
  - Input: [ 0, 0, 1, 0, 1, 1, 1, 0, 1, 1 ]
  - Output: The index to be replaced is 7.

### Problem11 (Impl Pending)

- Find maximum product of two integer in given array.
  - Input: [-10,-3,5,6,-2]
  - OutPut: [(-10,-3),(5,6)]

### Problem12

- Shuffle a given array of elements.
- FunFact: This is also known as **Fisher-Yates Shuffle**.
  - Given an array of elements, in-place shuffle it.The algorithm should produce un-biased permutation.

### Problem13

- Re-arrange array with alternate high and low elements.
- Given an array of integers, re-arrange this array in such a way that every second element in this array is greater than it's left and right element.
- **Assume no duplicate elements present in given array**
  Input: [1,2,3,4,5,6,7]
  OutPut: [1,3,2,5,4,7,6]

### Problem14

- Find equilibrium index of given array.

### Problem15

- Find majority element in given array.
- FunFact: This problem is also known as **Boyer-Moore Majority vote Algorithm**
  - Given an array of integer containing duplicates, Return the majority element in an array if present.
  - **TIP**: Majority element usually appears more than n/2 times in array, where n is the size of the array.
  - Input: [2,8,7,2,2,5,2,3,1,2,2]
  - OutPut: 2

### Problem16

- Move all 0s in given array at the end of the array.
  - Input: [6,0,8,2,3,0,4,0,1]
  - OutPut: [6,8,2,3,4,1,0,0,0]

### Problem17

- Given an array of integer, replace each element of the array with product of every other element in the array without using division operator.
  - Input: [1,2,3,4,5]
  - Output: [120,60,40,30,24]

### Problem18

- Find longest bitonic sub-array in an array.
- The lingest bitonic sub-array problem is to find a subarray of a given sequence in which the sub-array's elements are first sorted in increasing order, then in decreasing order, and the sub-array is as long as possible. Strictly ascending or descending sub-arrays are also accepted.
  - Input: [3,5,8,4,5,9,10,8,5,3,4]
  - Output: [4,5,9,10,8,5,3]

### Problem19

- Geiven an array of integers, find the maximum difference between two elements in the array such that smaller element appears before the larger element.
  - Input: [2,7,9,5,1,3,5]
  - OutPut: Max Difference : 7, Pair: (2,9)

### Problem20

- **Kadane's Algorithm**
- Given an array of integers, find contignous sub-array within it which has the largest sum.
  - Input: [-2,1,-3,4,-1,2,1,-5,4]
  - Output:
    - Sum: 6
    - sub-array: [4,-1,2,1]
