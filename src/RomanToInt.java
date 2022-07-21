public class RomanToInt {
    public static void main(String[] args) {
        String s = "MCMXCIV";

        System.out.println(romanToInt(s));
    }

    public static int romanToInt(String s) {

        char[] symbols = s.toCharArray();
        int result = 0;

        int startIndex = 0;
        int finishIndex = symbols.length;

        while (startIndex < finishIndex) {
            if (startIndex == (finishIndex - 1) ) {
                result += getValue(symbols[startIndex]);
                startIndex++;
                continue;
            }

            if ( (symbols[startIndex] == 'I') &&
            ( (symbols[startIndex + 1] == 'V') || (symbols[startIndex + 1] == 'X') ) ) {
                result += (getValue(symbols[startIndex + 1]) - getValue(symbols[startIndex]));
                startIndex += 2;
                continue;
            }

            if ( (symbols[startIndex] == 'X') &&
                    ( (symbols[startIndex + 1] == 'L') || (symbols[startIndex + 1] == 'C') ) ) {
                result += (getValue(symbols[startIndex + 1]) - getValue(symbols[startIndex]));
                startIndex += 2;
                continue;
            }

            if ( (symbols[startIndex] == 'C') &&
                    ( (symbols[startIndex + 1] == 'D') || (symbols[startIndex + 1] == 'M') ) ) {
                result += (getValue(symbols[startIndex + 1]) - getValue(symbols[startIndex]));
                startIndex += 2;
                continue;
            }

            result += getValue(symbols[startIndex]);
            startIndex++;
        }

        return result;
    }

    public static int getValue(char symbol) {

        switch (symbol) {
            case 'I':
                return 1;
            case 'V':
                return 5;
            case 'X':
                return 10;
            case 'L':
                return 50;
            case 'C':
                return 100;
            case 'D':
                return 500;
            case 'M':
                return 1000;
            default:
                break;
        }

        return 0;
    }
}