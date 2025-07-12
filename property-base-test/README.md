# property-base-test

使用 [Roman Numberals Kata](https://codingdojo.org/kata/RomanNumerals/) 這個題目來展示如何在 TDD 開發流使用 property base test 來讓測試更穩固。

## 講述順序

1. 先用原本 TDD 的方式來逐步拆解任務
2. 介紹 property base test
3. 使用 property base test 來讓測試更加穩固
4. 分析測試的邊界條件

## [Roman Numberals Kata](https://codingdojo.org/kata/RomanNumerals/) 

羅馬數字表達式是一種表達數值的方式

舉例來說

| 數值 | 羅馬數字表達式 |
| ---- | -------------- |
| 1    | I              |
| 2    | II             |
| 3    | III            |
| 4    | IV             |
| 5    | V              |
| 6    | VI             |
| 7    | VII            |
| 8    | VIII           |
| 9    | IX             |
| 10   | X              |
| 50   | L              |
| 100  | C              |
| 500  | D              |
| 1000 | M              |

MCMLXXXIV is 1984 

這邊我們需要實做兩個功能

1. ConvertToRoman : 從一般數值換成羅馬數字表達式
2. ConvertToArabic : 從羅馬數字表達式一般數值換成


