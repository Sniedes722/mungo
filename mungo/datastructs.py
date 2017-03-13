import ctypes
import random

lib = ctypes.cdll.LoadLibrary('./mungo/build/mungo.so')

def Array(nums):
	return (ctypes.c_int * len(nums))(*nums)
	
pyarr = random.sample(range(100), 10)
lib.average(Array(pyarr))
