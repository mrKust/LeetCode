package task16;

public class LengthOfLastWord {

    public static void main(String[] args) {
        String test = new String("Hello World");
        System.out.println(lengthOfLastWord(test));
    }

    public static int lengthOfLastWord(String s) {

        String[] words = s.split(" ");

        return words[words.length - 1].length();
    }
}
