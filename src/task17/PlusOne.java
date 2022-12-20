package task17;

import java.util.Arrays;

public class PlusOne {

    public static void main(String[] args) {
        int[] test = new int[]{8, 9, 9, 9};
        test = plusOne(test);
        System.out.println(Arrays.toString(test));
    }

    public static int[] plusOne(int[] digits) {

        int previousDigitTrans = 0;
        for (int i = digits.length - 1; i > 0; i--) {
            int digitResult = digits[i];
            if (i == digits.length - 1)
                digitResult++;
                else digitResult += previousDigitTrans;

            if (digitResult < 10) {
                digits[i] = digitResult;
                previousDigitTrans = 0;
            } else {
                previousDigitTrans = 1;
                digits[i] = digitResult % 10;
            }
        }

        if ( (digits[0] == 9) && ( (previousDigitTrans == 1) || (digits.length == 1)) ) {
            int[] extendedDigits = new int[digits.length + 1];
            for (int i = extendedDigits.length - 1; i > 1; i--) {
                extendedDigits[i] = digits[i - 1];
            }
            extendedDigits[1] = 0;
            extendedDigits[0] = 1;
            return extendedDigits;
        } else if ( (digits[0] != 9) && ( previousDigitTrans == 1) ) {
            digits[0] = digits[0] + previousDigitTrans;
        }

        if ( digits.length == 1 ) {
            digits[0]++;
        }

        return digits;
    }
}
