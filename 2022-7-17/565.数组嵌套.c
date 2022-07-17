#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#define MAX(a, b) ((a) > (b) ? (a) : (b))

int array_nesting(int* nums, int numsSize){
    int ans = 0;
    bool *vis = (bool *)malloc(sizeof(bool) * numsSize);
    for (int i = 0; i < numsSize; ++i) {
        int cnt = 0;
        memset(vis, 0, sizeof(bool) * numsSize);
        while (!vis[i]) {
            vis[i] = true;
            i = nums[i];
            ++cnt;
        }
        ans = MAX(ans, cnt);
    }
    free(vis);
    return ans;
}

int main()
{
    int ret;
    int V[7] = {5, 4, 0, 3, 1, 6, 2};

    ret = array_nesting(V, 7);
    printf("ret is %d\n", ret);

    return 0;
}