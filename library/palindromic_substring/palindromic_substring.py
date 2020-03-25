#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2020/3/23 下午8:14
# @Author  : serendipity-xp
# @FileName: palindromic_substring.py
# @Software: PyCharm
# @GitHub  ：https://github.com/serendipity-xp


class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        # Transform S into T.
        # For example, S = "abba", T = "^#a#b#b#a#$".
        # ^ and $ signs are sentinels appended to each end to avoid bounds checking
        T = '#'.join('^{}$'.format(s))
        n = len(T)
        P = [0] * n
        C = R = 0
        for i in range(1, n - 1):
            P[i] = (R > i) and min(R - i, P[2 * C - i])  # equals to i' = C - (i-C)
            # Attempt to expand palindrome centered at i
            while T[i + 1 + P[i]] == T[i - 1 - P[i]]:
                P[i] += 1

            # If palindrome centered at i expand past R,
            # adjust center based on expanded palindrome.
            if i + P[i] > R:
                C, R = i, i + P[i]

        # Find the maximum element in P.
        maxLen, centerIndex = max((n, i) for i, n in enumerate(P))
        return s[(centerIndex - maxLen) // 2: (centerIndex + maxLen) // 2]

    def longest_palindrome_dp(self, s):
        """
             :type s: str
             :rtype: str
             """
        # 1. create a 2D array, size is len(s) * len(s)
        dp = [[False] * len(s) for _ in range(len(s))]

        # 2. initial lcs related char index
        lcs_start_index = 0
        lcs_end_index = 0

        # 3. dp algo
        """
            a b a
            0 1 2
        a 0 T X X
        b 1 F T X
        a 2 T F T
        """
        for i in range(len(s)):
            # what we need is only the left bottom part
            start = i
            end = i
            while start >= 0:
                # case1. if sub-string is 'a'
                if start == end:
                    dp[start][end] = True
                # case2. if sub-string is 'ab'
                # We need this case because start + 1 may larger than end - 1 if using case3 directly
                elif start + 1 == end:
                    dp[start][end] = s[start] == s[end]
                # case3. if sub-string is 'aba' 'abac' ..etc, i.e. len(sub) >= 3
                else:
                    dp[start][end] = dp[start + 1][end - 1] and (s[start] == s[end])

                # if dp[start][end] is palidromic, check is it longer than current solution
                if dp[start][end] and (end - start + 1) > (lcs_end_index - lcs_start_index + 1):
                    lcs_start_index = start
                    lcs_end_index = end

                start = start - 1

        return s[lcs_start_index:lcs_end_index + 1]


if __name__ == '__main__':
    test_input = 'babadc'
    solution_obj = Solution()
    result = solution_obj.longestPalindrome(test_input)
    assert result == 'aba' or result == 'bab'
    result = solution_obj.longest_palindrome_dp(test_input)
    print result
    assert result == 'aba' or result == 'bab'

