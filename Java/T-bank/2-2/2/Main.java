import java.io.IOException;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) throws IOException {
        int infected = 1;

        try (
                Scanner scanner = new Scanner(System.in);
        ) {
            int minutes = scanner.nextInt();
            if (minutes != 1) {
                infected = (minutes - 1) * 4;
            }
            System.out.println(infected);
        }
    }
}
