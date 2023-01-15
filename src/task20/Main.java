package task20;

public class Main {

    private static int num = 0;

    public static void main(String[] args) {
        System.out.println(climbStairs(44));
    }
    
    public static int climbStairs(int n) {

        helper(n);

        return Main.num;
    }

    public static void helper(int n) {
        if (n == 0)
            Main.num++;
        else if (n >= 2) {
            helper(n - 2);
            helper(n - 1);
        } else {
            helper(n - 1);
        }
    }

}
