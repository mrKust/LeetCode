package task29;

public class Main {

    public static void main(String[] args) {
        String one = "leetcode";
        String two = "leeto";
        System.out.println(strStr(one, two));
    }

    public static int strStr(String haystack, String needle) {
        int indexStart = -1;

        if (haystack.length() < needle.length())
            return -1;

        for (int i = 0; i <= haystack.length() - needle.length(); i++) {
            int tmp = 0;
            while ( (tmp < needle.length()) && (haystack.charAt(i + tmp) == needle.charAt(tmp)) ) {
                tmp++;
                if (tmp == needle.length())
                    return i;
            }
        }

        return indexStart;
    }
}
