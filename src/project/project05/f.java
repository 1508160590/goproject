public class Fibonacci {

    public static void main(String[] args) {
        long startTime = System.currentTimeMillis(); // 获取开始时间
        for (int i = 1; i < 50; i++) {
            f(i);  // 调用 Fibonacci 函数
        }
        // 计算并输出程序执行时间
        long duration = System.currentTimeMillis() - startTime;
        System.out.println("Execution time in milliseconds: " + duration);
    }

    // Fibonacci 函数
    public static int f(int n) {
        if (n == 1) {
            return 1;
        } else if (n == 2) {
            return 2;
        } else {
            return f(n - 1) + f(n - 2);  // 递归计算 Fibonacci 数列
        }
    }
}
