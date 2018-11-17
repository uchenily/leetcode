int missingNumber(int* nums, int numsSize) {
    int result = 0;
    for (int i = 0; i < numsSize; i++) {
        result = result ^ nums[i];
        result ^= i;
    }
    result ^= numsSize;
    return result;
}
