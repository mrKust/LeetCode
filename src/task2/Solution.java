package task2;

public class Solution {

    public static void main(String[] args) {

        ListNode head = new ListNode(1, new ListNode(2, new ListNode(2)));
        System.out.println(isPalindrome(head));
    }

    public static boolean isPalindrome(ListNode head) {

        StringBuilder tmp = new StringBuilder();

        tmp.append(head.val);
        while (head.next != null) {
            tmp.append(head.next.val);
            head = head.next;
        }

        if ( (tmp.length() % 2) == 0 ) {
            String part1 = tmp.substring(0, tmp.length() / 2);
            String part2 = new StringBuilder(tmp.substring((tmp.length() / 2),
                    tmp.length())).reverse().toString();
            if (part1.equals(part2))
                return true;
        } else {
            String part1 = tmp.substring(0, (tmp.length() / 2) + 1);
            String part2 = new StringBuilder(tmp.substring((tmp.length() / 2),
                    tmp.length())).reverse().toString();
            if (part1.equals(part2))
                return true;

        }

        return false;
    }
}
