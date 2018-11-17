#include<stdbool.h>
// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

// 二分查找思想
int firstBadVersion(int n) {
    int low = 1;
    int high = n;
    int mid;
    
    while(low < high) {
        // mid = (low+high)/2; // 这里有个坑, 数据溢出, 报了超时错误
        mid = low + (high-low)/2;
        if(isBadVersion(mid)) {
            high = mid;
        }else {
            low = mid+1;
        }
    }
    return low;
}
