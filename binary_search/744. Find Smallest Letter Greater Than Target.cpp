class Solution {
public:
    char nextGreatestLetter(vector<char>& letters, char target) {
        int mid, left = 0;
        int right = letters.size() - 1;

        if(target >= letters.back()){
            return letters[0];
        }

        while(left <= right){
            mid = (right - left)/2 + left;
            if(letters[mid] <= target){
                left = mid + 1;
            }else{
                right = mid - 1;
            }
        }
        return letters[left];
    }
};