def binary_search(list, left, right, item):
    if left > right:
        return None
    mid = (left + right) // 2
    if list[mid] == item:
        return item
    elif item > list[mid]:
        return binary_search(list, mid + 1, right, item)
    else:
        return binary_search(list, left, mid - 1, item)


list = [1, 3, 5, 7, 9]
print(binary_search(list, 0, 4, 5))
