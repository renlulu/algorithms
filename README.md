# Algorithm (We are seeking AC solution for leetcode)

# Table of Content

- [Overview](#overview)
- [Divide and Conquer](#divide-and-conquer)
- [Dynamic Programming](#dynamic-programming)
- [Greedy](#greedy)
- [Breadth First Search](#breadth-first-search)
- [Depth First Search](#depth-first-search)
- [HashMap](#hashmap)
- [Algorithm in practise](#algorithm-in-practise)

# Overview

The table blow summarizes categories of algorithm we will focus on, before you go, we recommend two open courses by MIT:

- [MIT6.006 Introduction to Algorithms](https://www.youtube.com/playlist?list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
- [MIT6.046 Design And Analysis of Algorithms](https://www.youtube.com/playlist?list=PLUl4u3cNGP6317WaSNfmCvGym2ucw3oGp)

Cheer, thanks MIT for providing such good material for us.

| Algorithm | Description | Sample problem |
|--|--| --|
| Divide and Conquer |   | [0004.Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) |
| Dynamic Programming |   | []() |
| Greedy |  | []() |
| Breadth First Search | | [0101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/) |
| Depth First Search | | []() |
| Hash Map | | []() |


## Divide and Conquer

Divide and conquer is an algorithm design paradigm. A divide-and-conquer algorithm recursively breaks down a problem into two or more sub-problems of the same or related type, until these become simple enough to be solved directly. The solutions to the sub-problems are then combined to give a solution to the original problem. --wikipedia

| Problem | Leetcode Link | Solution | Description |
|--|--|--|--|
| 0023. Merge k sorted lists | [0023. Merge k sorted lists](https://leetcode.com/problems/merge-k-sorted-lists/) | [click me](0023_merge_k_sorted_lists.go)| |
| 0023. Median of Two Sorted Arrays | [0004.Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) | [click me](0004_median_of_two_sorted_arrays.go)| sub procedure of merge sort |

## Dynamic Programming

Dynamic programming is both a mathematical optimization method and a computer programming method. The method was developed by Richard Bellman in the 1950s and has found applications in numerous fields, from aerospace engineering to economics. --wikipedia

| Problem | Leetcode Link | Solution | Description |
|--|--|--|--|

## Greedy

A greedy algorithm is any algorithm that follows the problem-solving heuristic of making the locally optimal choice at each stage.[1] In many problems, a greedy strategy does not usually produce an optimal solution, but nonetheless, a greedy heuristic may yield locally optimal solutions that approximate a globally optimal solution in a reasonable amount of time. --wikipedia

| Problem | Leetcode Link | Solution | Description |
|--|--|--|--|

## Breadth First Search

Breadth-first search (BFS) is an algorithm for traversing or searching tree or graph data structures. It starts at the tree root (or some arbitrary node of a graph, sometimes referred to as a 'search key'[1]), and explores all of the neighbor nodes at the present depth prior to moving on to the nodes at the next depth level. --wikipedia

| Problem | Leetcode Link | Solution | Description |
|--|--|--|--|
| 0101. Symmetric tree | [0023. Symmetric tree](https://leetcode.com/problems/symmetric-tree/) | [click me](0101_symmetric_tree.go)| carefully handle if node is nil |

## Depth First Search

Depth-first search (DFS) is an algorithm for traversing or searching tree or graph data structures. The algorithm starts at the root node (selecting some arbitrary node as the root node in the case of a graph) and explores as far as possible along each branch before backtracking. --wikipedia

| Problem | Leetcode Link | Solution | Description |
|--|--|--|--|

## HashMap

A hash table (hash map) is a data structure that implements an associative array abstract data type, a structure that can map keys to values. A hash table uses a hash function to compute an index, also called a hash code, into an array of buckets or slots, from which the desired value can be found. During lookup, the key is hashed and the resulting hash indicates where the corresponding value is stored. --wikipedia


| Problem | Leetcode Link | Solution | Description |
|--|--|--|--|
| 0003. Longest Substring Without Repeating Characters | [0003. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | [click me](0003_longest_substring_without_repeating_characters.go)| |

## Algorithm in practise

Some useful algorithm and implement in real world.

- [Merkle Patricia Trie](https://eth.wiki/en/fundamentals/patricia-tree)

