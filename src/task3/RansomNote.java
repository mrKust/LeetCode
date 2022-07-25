package task3;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

public class RansomNote {

    public static void main(String[] args) {

        System.out.println(canConstruct("aa", "aab"));
    }

    public static boolean canConstruct(String ransomNote, String magazine) {

        List<Character> text = ransomNote.chars().mapToObj(c -> (char) c).collect(Collectors.toList());
        List<Character> leftLetters = magazine.chars().mapToObj(c -> (char) c).collect(Collectors.toList());
        List<Character> textCopy = new ArrayList<>(text);


        for (Character c: text) {
            if (leftLetters.contains(c)) {
                leftLetters.remove(c);
                textCopy.remove(c);
            } else return false;
        }

        return true;
    }
}
