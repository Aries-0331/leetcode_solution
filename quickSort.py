# 快速排序，平均时间复杂度 O(nlogn)，最糟时间复杂度O(n^2)

def quickSort(list):
    if len(list) < 2:
        return list
    else:
        pivot = list[0]
        less = [i for i in list[1:] if i < pivot]
        greater = [i for i in list[1:] if i >= pivot]
        return quickSort(less) + [pivot] + quickSort(greater)


list = [5, 2, 7, 4, 8, 4, 1]
print(quickSort(list))
