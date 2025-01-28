import java.io.IOException;
import java.util.ArrayDeque;
import java.util.Queue;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) throws IOException {
        long infected = 0;

        try (
                Scanner scanner = new Scanner(System.in);
        ) {
            int minutes = scanner.nextInt();
            int n = minutes;
            int curMin = 1;

            int[][] grid = new int[n][n];

            for (int i = 0; i < n; i++) {
                for (int j = 0; j < n; j++) {
                    grid[i][j] = 0;
                }
            }
            grid[0][0] = -1;
            //0 healthy
            //-1 start ill
            //-3 has imun
            int[][] directions = new int[][]{{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

            Queue<int[]> queue = new ArrayDeque<>();

            queue.add(new int[]{0, 0});

            while (!queue.isEmpty()) {
                int size = queue.size();

                for (int k = 0; k < size; k++) {
                    int[] current = queue.remove();
                    int cR = current[0];
                    int cC = current[1];

                    //если клетка с имунитетом или здоровая, ничего не делаем. тут будут все больные
//                    if (grid[cR][cC] == -3 || grid[cR][cC] == 0) {
//                        continue;
//                    }
                    for (int[] dir : directions) {
                        int dirR = cR + dir[0];
                        int dirC = cC + dir[1];

                        if (dirR < 0 || dirR >= n || dirC < 0 || dirC >= n) {
                            continue;
                        }

                        if (grid[dirR][dirC] == -3) {
                            continue;
                        }

                        /*if (grid[dirR][dirC] >= curMin -1 && grid[dirR][dirC] < curMin) {
                            grid[dirR][dirC] = -3;
                        }*/

                        queue.add(new int[]{dirR, dirC});
                        grid[dirR][dirC] = curMin;
                    }

//                    for (int i = 0; i < n; i++) {
//                        for (int j = 0; j < n; j++) {
//
//                            if (grid[i][j] > 0 && grid[i][j] < curMin) {
//                                grid[i][j] = -3;
//                            }
//                        }
//                    }
                    grid[cR][cC] = -3;
                }
                curMin++;
            }
            for (int i = 0; i < n; i++) {
                for (int j = 0; j < n; j++) {
                    if (grid[i][j] == minutes) {
                        infected++;
                    }

                }
            }
            System.out.println(infected);
        }

    }

}
