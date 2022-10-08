def arrSum(list):
    if list == []:
        return 0
    return list[0] + arrSum(list[1:])


list = [1, 2, 3, 4]
print(arrSum(list))
