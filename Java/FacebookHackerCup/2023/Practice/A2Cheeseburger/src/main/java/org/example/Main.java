package org.example;

import java.io.*;
import java.nio.file.Paths;
import java.util.Scanner;

public class Main {
    //private static final String inputPath = "data/cheeseburger_corollary_2_validation_input.txt";

    private static final String inputPath = "data/cheeseburger_corollary_2_input.txt";
    public static void main(String[] args) throws IOException {
        long start = System.currentTimeMillis();
        try (Scanner scanner = new Scanner(Paths.get(inputPath));
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
        return new PrintWriter(new OutputStreamWriter(new FileOutputStream(inputPath.replace("input", "output"))));
    }

    private static void solve(int caseNumber, Scanner scaner, PrintWriter writer) {
        long cs = scaner.nextLong();
        long cd = scaner.nextLong();
        long m = scaner.nextLong();

        long patty = 0;
        long buns = 0;
        long canMake = 0;
        long canMake2 = 0;
        long maxS = m / cs;
        long maxD = m / cd;

        long remAfterS = m - maxS * cs;
        long remAfterD = m - maxD * cd;

        long appendD = remAfterS / cd;
        long appendS = remAfterD / cs;

        //1 S
        patty = maxS + 2 * appendD;
        buns = maxS * 2 + appendD * 2;
        canMake = buns > 0 && patty > 0 ? Math.min((buns - 1), patty) : 0;

        //2 D
        patty = maxD * 2 + appendS;
        buns = maxD * 2 + appendS * 2;

        canMake2 =  buns > 0 && patty > 0? Math.min((buns - 1), patty) : 0;

        long res = Math.max(canMake2, canMake);

        writer.printf("Case #%d: %d\n", caseNumber, res);
    }
}