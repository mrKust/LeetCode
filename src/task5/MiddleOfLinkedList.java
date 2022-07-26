package task5;

public class MiddleOfLinkedList {

    public static void main(String[] args) {

        ListNode head = new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4,
                new ListNode(5, new ListNode(6))))));
        ListNode tmp = middleNode(head);
        while (tmp.next != null) {
            System.out.print(tmp.val + " ");
            tmp = tmp.next;
        }
        System.out.print(tmp.val);
    }

    public static ListNode middleNode(ListNode head) {

        ListNode forward = head;
        ListNode behind = head;

        while (forward.next != null) {
            if (forward.next.next != null) {
                forward = forward.next.next;
            } else forward = forward.next;
            behind = behind.next;
        }

        return behind;
    }
}
