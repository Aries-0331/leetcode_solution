class Solution {
public:
    bool isPerfectSquare(int num) {
        int mid, left = 0;
        int right = num;
        while(left <= right){
            mid = (right - left)/2 + left;
            if(pow(mid ,2) == num){
                return true;
            }else if(pow(mid, 2) < num){
                left = mid + 1;
            }else{
                right = mid - 1;
            }
        }
        return false;
    }
};