import java.util.function.IntConsumer;

class FizzBuzz {
    private int n;
    private int i;

    public FizzBuzz(int n) {
        this.n = n;
        this.i = 1;
    }

    // printFizz.run() outputs "fizz".
    public void fizz(Runnable printFizz) throws InterruptedException {
        // lock1.lock();
        while (i <= n) {
            if (i % 3 == 0 && i % 15 != 0) {

                printFizz.run();
                i++;
                // printFizz.wait();
            }
        }
        // printFizz.notifyAll();
        // lock1.unlock();
    }

    // printBuzz.run() outputs "buzz".
    public void buzz(Runnable printBuzz) throws InterruptedException {
        // lock2.lock();
        while (i <= n) {
            if (i % 5 == 0 && i % 15 != 0) {
                printBuzz.run();
                i++;
                // printBuzz.wait();
            }
        }

        // printBuzz.notifyAll();
        // lock2.unlock();
    }

    // printFizzBuzz.run() outputs "fizzbuzz".
    public void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {
        while (i <= n) {
            if (i % 15 == 0) {
                printFizzBuzz.run();

                i++;
            }
        }
        // printFizzBuzz.notifyAll();
        // lock3.unlock();
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void number(IntConsumer printNumber) throws InterruptedException {
        while (i <= n) {
            if (i % 3 != 0 && i % 5 != 0) {
                printNumber.accept(i);
                i++;
            }
        }
        // printNumber.notifyAll();
        // lock4.unlock();
    }
}