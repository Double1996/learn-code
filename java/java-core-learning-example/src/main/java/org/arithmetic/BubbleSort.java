package org.arithmetic;

public class BubbleSort {

    /*  冒泡排序
        两两相邻比较
        然后排序
     */
    public void sort(int[] array) {
        // 创建一个数组
        for (int i = 1; i < array.length; i++) {
            for (int j = 0; j < array.length; j++) {
                if (array[j] > array[j + 1]) {
                    int temp = array[j];
                    array[j] = array[j + 1];
                    array[j + 1] = temp;
                }
            }
        }
        for (int i : array) {
            System.out.println(">" + i);
        }
    }

    public static void main(String[] args) {
        int[] array = {63, 4, 24, 1, 3, 15};
        BubbleSort sorter = new BubbleSort();
        sorter.sort(array);
    }
}
