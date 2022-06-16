## 解法一： 暴力循环所有子字符串
- 注意边界条件
- 单个字符也是回文
- 判断方法： s[left_end] == s[right_end], 由两边向中心
- 终止条件

## 解法二： 中心扩散法
![img.png](img.png)
- 判断方法： s[left_end] == s[right_end], 由中心向两边扩展

## 解法三： 动态规划法


## 解法四： Manacher
