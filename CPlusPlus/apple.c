#include <stdio.h>

int main(int argc, char const *argv[]) {
    int n = 2, day = 0;    /* 定义基本类型 */
    float monkey = 0, ave; /* 定义单精度基本变量 */
    while (n < 100) {
        monkey += 0.8 * n;
        ++day;
        n *= 2;
    }
    ave = monkey / day;
    printf("The result of %.6f ", ave);
    return 0;
}
