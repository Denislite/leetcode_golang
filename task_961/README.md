# 961. N-Repeated Element in Size 2N Array

---

You are given an integer array nums with the following properties:

nums.length == 2 * n. <br>
nums contains n + 1 unique elements. <br>
Exactly one element of nums is repeated n times. <br>
Return the element that is repeated n times.

#### Example 1:
```
Input: nums = [1,2,3,3]
Output: 3
```
#### Example 2:
```
Input: nums = [2,1,2,5,3,2]
Output: 2
```

---
### Constraints:

- 2 <= n <= 5000
- nums.length == 2 * n
- 0 <= nums[i] <= 10ˆ4 
- nums contains n + 1 unique elements and one of them is repeated exactly n times.