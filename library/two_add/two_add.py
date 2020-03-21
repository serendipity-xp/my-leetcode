#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2020/3/21 下午3:52
# @Author  : serendipity-xp
# @FileName: two_add.py
# @Software: PyCharm
# @GitHub  ：https://github.com/serendipity-xp


# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        dummy_head = ListNode(0)
        cur_l1, cur_l2 = l1, l2
        cur_node = dummy_head
        carry = 0  # 表示进位
        while cur_l1 is not None or cur_l2 is not None:
            x = (cur_l1.val if cur_l1 else 0)
            y = (cur_l2.val if cur_l2 else 0)
            sum_val = x + y + carry
            carry = sum_val / 10
            cur_node.next = ListNode(sum_val % 10)
            cur_node = cur_node.next
            cur_l1 = (cur_l1.next if cur_l1 else None)
            cur_l2 = (cur_l2.next if cur_l2 else None)
        if carry > 0:
            cur_node.next = ListNode(carry)
        return dummy_head.next


if __name__ == '__main__':
    l1 = ListNode(2)
    l1.next = ListNode(4)
    l1.next.next = ListNode(3)
    l2 = ListNode(5)
    l2.next = ListNode(6)
    l2.next.next = ListNode(4)
    Solution().addTwoNumbers(l1, l2)


