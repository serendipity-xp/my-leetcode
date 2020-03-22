#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2020/3/22 下午1:23
# @Author  : serendipity-xp
# @FileName: median_num.py
# @Software: PyCharm
# @GitHub  ：https://github.com/serendipity-xp

class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        sum_list = list(nums1 + nums2)
        sum_list.sort()
        length = len(sum_list)
        result = float(sum_list[length / 2])
        if length % 2 == 0:
            result = float(sum_list[length / 2] + sum_list[length / 2 - 1]) / 2
        return result


if __name__ == '__main__':
    list1 = [3, 1, 5]
    list2 = [2]
    solution_obj = Solution()
    print solution_obj.findMedianSortedArrays(list1, list2)

