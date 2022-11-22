import time
import threading
from queue import Queue
import random
 
# producer thread
def producer(queue):
    print('running producer')
    for i in range(1000):
        # generate random value between 0 and 1
        val = random.random()
        # block for that fraction of a second
        time.sleep(val)
        # add value to the queue
        queue.put(val)
    queue.put(None)
    print('completed running producer')
 
# consumer thread  
def consumer(queue):
    print('running consumer')
    while True:
        # get a value from the queue
        val = queue.get()
        # if val  is None, break from the infinite loop and this thread will be terminated
        if val is None:
            break
        # print the received value from the shared queue
        print(f'Received value {val}')
    print('completed running consumer')

if __name__=="__main__":
    # shared queue between consumer and producer threads
    queue = Queue()
    # start the consumer thread
    consumer = threading.Thread(target=consumer, args=(queue,))
    consumer.start()
    # start the producer thread
    producer = threading.Thread(target=producer, args=(queue,))
    producer.start()
    # join threads when done 
    producer.join()
    consumer.join()