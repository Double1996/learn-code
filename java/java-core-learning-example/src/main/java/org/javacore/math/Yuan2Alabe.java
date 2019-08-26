package org.javacore.math;

/**
 * 这是一个阿拉伯数字转转换成大写数字的算法
 * Created by double2 on 2017/7/3.
 */
public class Yuan2Alabe {
    private static final char[] data = new char[]{
        '零', '贰', '叁', '肆', '伍', '陆', '柒', '捌', '玖'
    };
    private static final char[] units = new char[]{
        '圆', '佰', '仟', '万', '亿'
    };

    private static String convert(int money) {
        StringBuffer stringBuffer = new StringBuffer();
        int unit = 0;
        while (money != 0) {
            stringBuffer.insert(0, units[unit++]);
            int number = money % 10;
            stringBuffer.insert(0, data[number]);
            money /= 10;
        }
        return stringBuffer.toString();
    }

    // Todo: 只能支持int字符
    public static void main(String[] args) {
        System.out.println(convert(16111));

    }
}
