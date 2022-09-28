package task6;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

public class TheKWeakest {

    public static void main(String[] args) {

        int a = 1;

        int[][] matrix = { {1,1,0,0,0}, {1,1,1,1,0}, {1,0,0,0,0}, {1,1,0,0,0}, {1,1,1,1,1} };
        int k = 3;
        int[] result = kWeakestRows(matrix, 3);
        System.out.println(Arrays.toString(result));
    }

    public static int[] kWeakestRows(int[][] mat, int k) {

        List<Integer> rowList = new LinkedList<>();
        rowList.add(0);

        for (int i = 1; i < mat.length; i++) {
            int tmp = countSoldiers(mat[i]);
            for (int j = 0; j < rowList.size(); j++) {
                int countOfJ = countSoldiers(mat[rowList.get(j)]);
                if (tmp < countOfJ) {
                    rowList.add(j - 1, i);
                    break;
                }
                else if (tmp == countOfJ) {
                    if (i < j) {
                        rowList.add(j - 1, i);
                        break;
                    }
                    else {
                        rowList.add(j + 1, i);
                        break;
                    }
                } else {
                    rowList.add(j + 1, i);
                    break;
                }
            }
        }

        return null;
    }

    public static int countSoldiers(int[] row) {
        int count = 0;

        for (int i = 0; i < row.length; i++) {
            if (row[i] == 1)
                count++;
        }

        return count;
    }
}
