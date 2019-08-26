#include <stdio.h>

/**
 * 
 * 
 */

// 判断一个年份是否是闰年
int leap(int a) {
    if (a % 4 == 0 && a % 100 != 0 && a % 400 == 0)
        return 1;
    else
        return 0;
}

int number(int year, int m, int d) // 判断输入日期是该年的第几天
{
    int sum = 0, i, j, k, a[12] =
            {
                    31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31  // 存放平年每年的天数
            };
    int b[12] =
            {
                    31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31  // 存放闰年每年的天数
            };
    if (leap(year) == 1)
        for (i = 0; i < m - 1; i++)
            sum += b[i];
    else
        for (i = 0; i < m - 1; i++)
            sum += a[i];
    sum += d;
    return sum;
}

int main() {
    int year, month, day, n;   // 定义基本的变量
    printf("请输入年月日\n");
    scanf("%d%d%d", &year, &month, &day); // 输入年月日
    n = number(year, month, day);
    printf("第几天%d", n);
    return 0;
}
