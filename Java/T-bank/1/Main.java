import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scaner = new Scanner(System.in);
        int lastNumber = scaner.nextInt();

        long res;
        int col = lastNumber - 100 + 1;
        res = (long) col * (100 + lastNumber) / 2;
        System.out.println(res);
    }
}