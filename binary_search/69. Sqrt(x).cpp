class Solution {
public:
    int mySqrt(int x) {
        int mid, left = 0;
        int right = x;
        while(left <= right){
            mid = (right - left)/2 + left;
            if(pow(mid, 2) <= x){
                left = mid + 1;
            }else{
                right = mid - 1;
            }
        }
        return (left - 1);
    }
};