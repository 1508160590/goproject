public class cm {
    public static void main(String[] args) {
        long startTime = System.currentTimeMillis(); // 获取开始时间

        for (int i = 0; i < 100000; i++) {
            fibonacci(i);
        }

        long endTime = System.currentTimeMillis(); // 获取结束时间
        long duration = endTime - startTime; // 计算执行时间
        System.out.println("Execution time in milliseconds: " + duration);
    }

    private static long fibonacci(int n) {
        if (n < 1) {
            return n;
        }

        long[] fib = new long[n + 1];
        fib[0] = 0;
        fib[1] = 1;

        for (int i = 2; i <= n; i++) {
            fib[i] = fib[i - 1] + fib[i - 2];
        }

        return fib[n];
    }
}
