import java.util.concurrent.Semaphore;
import java.util.concurrent.locks.ReentrantLock;
import java.util.function.IntConsumer;

import javax.swing.plaf.basic.BasicColorChooserUI.PropertyHandler;

class ZeroEvenOdd {
    private int n;
    // private int i;
    // private int pre;
    private Semaphore s1;
    private Semaphore s2;
    private Semaphore s3;

    // private ReentrantLock lock;

    public ZeroEvenOdd(int n) {
        this.n = n;
        s1 = new Semaphore(1);
        s2 = new Semaphore(0);
        s3 = new Semaphore(0);
        // lock = new ReentrantLock();
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void zero(IntConsumer printNumber) throws InterruptedException {
        // lock.lock();
        for (int i = 1; i <= n; i++) {
            s1.acquire();
            printNumber.accept(0);
            if(i % 2 == 0){
                s2.release();
            }else{
                s3.release();
            }
        }
        // lock.unlock();
    }

    public void even(IntConsumer printNumber) throws InterruptedException {
        for (int i = 2; i <= n; i += 2) {
            s2.acquire();
            printNumber.accept(i);
            s1.release();
        }
    }

    public void odd(IntConsumer printNumber) throws InterruptedException {
        for (int i = 1; i <= n; i += 2) {
            s3.acquire();
            printNumber.accept(i);
            s1.release();
        }
    }
}