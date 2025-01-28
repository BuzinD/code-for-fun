import java.io.*;
import java.nio.file.Paths;
import java.util.Scanner;

public class Main {

    //private static final String source = "data/two_apples_a_day_input.txt";
    private static final String source = "data/two_apples_a_day_validation_input.txt";

    public static void main(String[] args) throws IOException {
        long start = System.currentTimeMillis();
        try (Scanner scanner = new Scanner(Paths.get(source));
             PrintWriter writer = getWriter()) {
            int t = scanner.nextInt();

            for (int i = 1; i <= t; i++) {
                long testCaseStart = System.currentTimeMillis();
                solve(i, scanner, writer);
                System.out.printf("Case #%02d: %d ms%n", i, System.currentTimeMillis() - testCaseStart);
            }
        }
        System.out.printf("%nTotal: %d ms%n", System.currentTimeMillis() - start);
    }

    private static PrintWriter getWriter() throws FileNotFoundException {
        return new PrintWriter(new OutputStreamWriter(new FileOutputStream(source.replace("input", "output"))));
    }

    private static void solve(int caseNumber, Scanner scaner, PrintWriter writer) {
        int days = scaner.nextInt();
        int n = 2 * days - 1;
        int[] weights = new int[n];
        int[][] sums = new int[n][n];
        for (int i = 0; i < n; i++) {
            weights[i] = scaner.nextInt();
        }
        int minRepitedSum = Integer.MIN_VALUE;

        for (int i = 0; i < n; i++) {
            for (int j = i; j < n; j++) {
                 sums[i][j] = weights[i] + weights[j];
            }
        }
//        for (int i = n - 1; i >= 0; i--) {
//            for (int j = i; j >= 0 / 2; j++) {
//                sums[i][j] = weights[i] + weights[j];
//            }
//        }
        for (int i = 0; i < n; i++) {
            for (int j = i; j < n; j++) {
                sums[i][j] = weights[i] + weights[j];
            }
        }
        writer.printf("Case #%d: %d\n", caseNumber, result != Long.MIN_VALUE ? result : -1);
    }
}