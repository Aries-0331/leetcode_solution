def max(list):
    if len(list) == 2:
        return list[0] if list[0] > list[1] else list[1]
    submax = max(list[1:])
    return list[0] if list[0] > submax else submax


list = [1, 2, 3, 4, 5, 6]
print(max(list))
