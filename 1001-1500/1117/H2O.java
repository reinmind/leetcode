import java.util.concurrent.Semaphore;

class H2O {
    Semaphore s1,s2;
    int h,o;
    public H2O() {
        s1 = new Semaphore(2);
        s2 = new Semaphore(1);
    }
    public void consume(){
        if (h == 2 && o == 1){
            h = 0;
            o = 0;
            s1.release();
            s1.release();
            s2.release();
        }
    }
    public void hydrogen(Runnable releaseHydrogen) throws InterruptedException {
		s1.acquire();
        // releaseHydrogen.run() outputs "H". Do not change or remove this line.
        releaseHydrogen.run();
        h++;
        consume();
    }

    public void oxygen(Runnable releaseOxygen) throws InterruptedException {
        s2.acquire();
        // releaseOxygen.run() outputs "O". Do not change or remove this line.
		releaseOxygen.run();
        o++;
        consume();
    }
}