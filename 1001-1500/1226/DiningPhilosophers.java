import java.util.HashMap;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

class DiningPhilosophers {
    private ConcurrentHashMap<Integer, Lock> lockMap;

    public DiningPhilosophers() {
        lockMap = new ConcurrentHashMap<>();
        lockMap.put(0, new ReentrantLock());
        lockMap.put(1, new ReentrantLock());
        lockMap.put(2, new ReentrantLock());
        lockMap.put(3, new ReentrantLock());
        lockMap.put(4, new ReentrantLock());
    }

    // call the run() method of any runnable to execute its code
    public void wantsToEat(int philosopher,
            Runnable pickLeftFork,
            Runnable pickRightFork,
            Runnable eat,
            Runnable putLeftFork,
            Runnable putRightFork) throws InterruptedException {
        int leftChop = philosopher;
        int rightChop = (philosopher + 1) % 5;
        lockMap.get(leftChop).lock();
        lockMap.get(rightChop).lock();
        pickLeftFork.run();
        pickRightFork.run();
        eat.run();
        putRightFork.run();
        putLeftFork.run();
        lockMap.get(leftChop).unlock();
        lockMap.get(rightChop).unlock();
    }
}