// 太慢了!!!
// int hammingDistance(int x, int y) {
//     int result = 0;
//     while(x > 0 || y > 0) {
//         if( (x&1) != (y&1) ) { // 要加上(x&1)
//             result++;
//         }
//         x = x>>1;
//         y = y>>1;
//     }
//     return result;
// }

int hammingDistance(int x, int y) {
    int result = 0;
    for(int i = 0; i < 31; i++) {
        if( ((x>>i)&1) ^ ((y>>i)&1) ) { // 位运算, 执行速度快
            result++;
        }
    }
    return result;
}
