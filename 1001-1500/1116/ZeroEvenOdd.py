import threading
class ZeroEvenOdd(object):
    def __init__(self, n):
        self.n = n
        self.s1 = threading.Semaphore(1)
        self.s2 = threading.Semaphore(0)
        self.s3 = threading.Semaphore(0)
        
        
	# printNumber(x) outputs "x", where x is an integer.
    def zero(self, printNumber):
        for i in range(1,self.n+1,1):
            self.s1.acquire()
            printNumber(0)
            if i % 2 == 0:
                self.s2.release()
            else:
                self.s3.release()
        """
        :type printNumber: method
        :rtype: void
        """
        

    def even(self, printNumber):
        for i in range (2,self.n+1,2) :
            self.s2.acquire()
            printNumber(i)
            self.s1.release()
        """
        :type printNumber: method
        :rtype: void
        """
        
        
    def odd(self, printNumber):
        for i in range (1,self.n+1,2):
            self.s3.acquire()
            printNumber(i)
            self.s1.release()
        """
        :type printNumber: method
        :rtype: void
        """