package memory;
public class Mem{
    public static void main(String[] args) throws InterruptedException {
        System.out.println(Integer.MAX_VALUE);
        byte[] ca = new byte[Integer.MAX_VALUE / 2];
        System.out.println(ca.length);
        String s = new String();
        Thread.sleep(10000);
    }
}