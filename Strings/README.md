# 205. Isomorphic Strings

Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.


Example 1:
```
Input: s = "egg", t = "add"
Output: true
```
Example 2:
```
Input: s = "foo", t = "bar"
Output: false
```
Example 3:
```
Input: s = "paper", t = "title"
Output: true
```

# 168. Excel Sheet Column Title

Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.

For example:
```
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...
```


Example 1:

Input: columnNumber = 1
Output: "A"

Example 2:

Input: columnNumber = 28
Output: "AB"

Example 3:

Input: columnNumber = 701
Output: "ZY"


# 171. Excel Sheet Column Number

Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.

For example:

```js
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28 
...
```


Example 1:

Input: columnTitle = "A"
Output: 1


Example 2:

Input: columnTitle = "AB"
Output: 28


Example 3:

Input: columnTitle = "ZY"
Output: 701