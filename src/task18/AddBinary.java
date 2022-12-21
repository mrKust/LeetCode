package task18;

import java.util.Arrays;

public class AddBinary {
    public static void main(String[] args) {

        String a = "11";
        String b = "1";
        System.out.println(addBinary(a, b));
    }

    public static String addBinary(String a, String b) {

        int aIndex = a.length();
        int bIndex = b.length();

        if (aIndex < bIndex)
            return addBinary(b, a);

        int maxLength = Math.max(aIndex, bIndex);
        StringBuilder sb = new StringBuilder();
        int transfer = 0;
        int j = bIndex - 1;

        for (int i = maxLength - 1; i > -1; --i) {
            if (a.charAt(i) == '1')
                ++transfer;
            if (j > -1 && b.charAt(j--) == '1')
                ++transfer;

            if (transfer % 2 == 1)
                sb.append('1');
            else sb.append('0');

            transfer /= 2;
        }

        if (transfer == 1) sb.append('1');
        sb.reverse();

        return sb.toString();
    }
}
