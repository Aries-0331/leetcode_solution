#include <iostream>
#include <vector>

int arrSum(int arr[], int length)
{
    if (arr[length - 1] == 0)
    {
        return 0;
    }
    else
    {
        return arr[length - 1] + arrSum(arr, length - 1);
    }
}

int main()
{
    int array[6] = {2, 4, 6, 8, 10, 12};
    std::cout << "sum = " << arrSum(array, 6) << std::endl;
    return 0;
}