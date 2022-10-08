def binary_search(list, item):
    if len(list) == 1 and list[0] == item:
        return list[0]
    half = int(len(list)/2)
    if list[half] == item:
        return list[half]
    elif item > list[half]:
        return binary_search(list[half:], item)
    else:
        return binary_search(list[:half], item)


list = [1, 3, 5, 7, 9]
print(binary_search(list, 3))
