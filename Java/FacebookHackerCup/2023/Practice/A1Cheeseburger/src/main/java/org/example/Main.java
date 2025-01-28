package org.example;

import java.io.*;
import java.nio.file.Paths;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) throws IOException {
        long start = System.currentTimeMillis();
        try (Scanner scanner = new Scanner(Paths.get(args[0]));
             PrintWriter writer = getWriter(args)) {
            int t = scanner.nextInt();

            for (int i = 1; i <= t; i++) {
                long testCaseStart = System.currentTimeMillis();
                solve(i, scanner, writer);
                System.out.printf("Case #%02d: %d ms%n", i, System.currentTimeMillis() - testCaseStart);
            }
        }
        System.out.printf("%nTotal: %d ms%n", System.currentTimeMillis() - start);
    }

    private static PrintWriter getWriter(String[] args) throws FileNotFoundException {
        return new PrintWriter(new OutputStreamWriter(new FileOutputStream(args[0].replace("input", "output"))));
    }

    private static void solve(int caseNumber, Scanner scaner, PrintWriter writer) {
        int s = scaner.nextInt();
        int d = scaner.nextInt();
        int k = scaner.nextInt();

        int patty = s + 2 * d;
        int buns = s * 2 + d * 2;

        int needBuns = 1 + k;
        int needPatty = k;
        boolean possible = patty >= needPatty && buns >= needBuns;

        writer.printf("Case #%d: %s\n", caseNumber, possible ? "YES" : "NO");
    }
}