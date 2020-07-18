#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2020/3/21 下午8:31
# @Author  : serendipity-xp
# @FileName: longest_substring.py
# @Software: PyCharm
# @GitHub  ：https://github.com/xuliangpan


class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        longest_substring = []
        max_length = 0
        for character in s:
            if character in longest_substring:
                longest_substring = longest_substring[longest_substring.index(character)+1:]
            longest_substring.append(character)
            max_length = max(max_length, len(longest_substring))
        return max_length


if __name__ == '__main__':
    solution_obj = Solution()
    test_str = 'pwwkew'
    print solution_obj.lengthOfLongestSubstring(test_str)
