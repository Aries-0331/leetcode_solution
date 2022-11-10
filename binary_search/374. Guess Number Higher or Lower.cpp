/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * int guess(int num);
 */

class Solution {
public:
    int guessNumber(int n) {
        int num, left = 1;
        while(n > left){
            num = (n - left) / 2 + left;
            if(guess(num) <= 0){
                n = num;
            }else{
                left = num + 1;
            }
        }

        return left;
    }
};