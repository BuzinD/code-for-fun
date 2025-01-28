import java.io.*;
import java.nio.file.Paths;
import java.util.Scanner;

public class Main {

//    private static final String source = "data/dim_sum_delivery_validation_input.txt";
    private static final String source = "data/dim_sum_delivery_input.txt";

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
        long rows = scaner.nextLong();
        System.out.println("rows " + rows);
        long columns = scaner.nextLong();
        System.out.println("columns " + columns);
        long aliceSteps = scaner.nextLong();
        long bobSteps = scaner.nextLong();
        boolean possible = rows > columns;
        writer.printf("Case #%d: %s\n", caseNumber, possible ? "YES" : "NO");
    }
}