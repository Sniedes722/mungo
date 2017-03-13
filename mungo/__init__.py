from ctypes import cdll

lib = cdll.LoadLibrary('./mungo/build/mungo.so')

lib.

class MungoArray:
    
    def __init__(self):
        self.mgArray = []
        
    def addArray(self,num1,num2):
        self.arraySum = lib.add(num1,num2)
        return self.arraySum