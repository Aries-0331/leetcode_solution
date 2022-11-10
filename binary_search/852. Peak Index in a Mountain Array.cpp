class Solution {
public:
    int peakIndexInMountainArray(vector<int>& arr) {
        int mid, left = 0;
        int right = arr.size() - 1;
        while(left != right){
            mid = (right - left)/2 + left;
            if(arr[mid] <= arr[mid+1]){
                left = mid + 1;
            }else{
                right = mid;
            }
        }
        return left;
    }
};