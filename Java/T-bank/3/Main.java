import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Main {

    public static void main(String[] args) throws IOException {

        try (
                BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        ) {
            int n = Integer.parseInt(br.readLine());
            int[][] mat = new int[n][n];
            int[] sumsRows = new int[n];
            int[] sumsCols = new int[n];

            for (int i = 0; i < n; i++) {

                String[] nums = br.readLine().split(" ");
                int k = 0;
                for (String num : nums) {
                    mat[i][k++] = Integer.parseInt(num);
                }
            }

            for (int i = 0; i < n; i++) {
                for (int j = 0; j < n; j++) {
                    sumsRows[i] += mat[i][j];
                    sumsCols[j] += mat[i][j];
                }
            }
            int counter = 0;
            for (int i = 0; i < n; i++) {
                for (int j = 0; j < n; j++) {
                    if (Math.abs(sumsRows[i] - sumsCols[j]) <= mat[i][j]) {
                        counter++;
                    }
                }
            }
            System.out.println(counter);
            int i = 9;
        }
    }
}

