package task4;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class FizzBuzz {

    public static void main(String[] args) {

        int n = 3;
        System.out.println(Arrays.toString(fizzBuzz(n).toArray()));
    }

    public static List<String> fizzBuzz(int n) {

        List<String> result = new ArrayList<>();

        for (int i = 1; i <= n; i++) {

            if (( (i % 3) == 0 ) && ( (i % 5) == 0 )) {
                result.add("FizzBuzz");

            } else if ( (i % 3) == 0 ) {
                result.add("Fizz");

            } else if ( (i % 5) == 0 ) {
                result.add("Buzz");

            } else {
                result.add(Integer.toString(i));
            }
        }

        return result;
    }
}
