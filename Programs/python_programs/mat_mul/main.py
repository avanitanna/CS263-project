import threading
import numpy as np

# default num of threads
threads = 10 
# default mat dim
mat_dim = 500
mat_A = []
mat_B = []
# res matrix to store mat mul of A and B 
res = np.zeros((mat_dim,mat_dim))

# matrix multiplication function 
def matMul(start,end):
    for i in range(start,end):
        for j in range(mat_dim):
            for k in range(mat_dim):
                res[i][j] += int(mat_A[i][k] * mat_B[k][j])  

# random generation of matrices
def generateMat():
    global mat_A
    global mat_B
    mat_A = np.random.randint(1, size=(mat_dim, mat_dim))
    mat_B = np.random.randint(1, size=(mat_dim, mat_dim))

# we can change the num of threads or dim of matrix    
def userInput():
    global mat_dim
    global threads
    mat_dim = int(input("Please enter matrix dim n: "))
    threads = int(input("Please enter number of threads: "))

# threading function in python  
def threadFunction():
    global threads
    list_threads = []

    for j in range(0,threads):
        t = threading.Thread(target = matMul, args=(int((mat_dim/threads)*j),int((mat_dim/threads)*(j+1))))
        list_threads.append(t)
        t.start()   
        
    for j in range(0,threads):
        list_threads[j].join()

if __name__=="__main__":
    generateMat()
    threadFunction()