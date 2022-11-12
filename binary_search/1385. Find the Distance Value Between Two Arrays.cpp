class Solution {
public:
    bool binarySearch(vector<int>& arr, int low, int high){
        int mid, left = 0;
        int right = arr.size() - 1;
        while(left <= right){
            mid = (right - left)/2 + left;
            if(arr[mid] >= low && arr[mid] <= high){
                return true;
            }else if(arr[mid] < low){
                left = mid + 1;
            }else{
                right = mid - 1;
            }
        }
        return false;
    }

    int findTheDistanceValue(vector<int>& arr1, vector<int>& arr2, int d) {
        int val = 0;
        sort(arr2.begin(), arr2.end());
        for(int i : arr1){
            int low = i - d;
            int high = i + d;
            if(!binarySearch(arr2, low, high)){
                val++;
            }
        }
        return val;
    }
};