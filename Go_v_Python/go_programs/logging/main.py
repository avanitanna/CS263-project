import random
import threading 
import time 
import logging
 
# logging function
def log(p):
    logging.info(f'running thread')
    for i in range(100):
        # generate random value between 0 and 1
        val = random.random()
        # log values
        logging.debug(f'Received value {val}.')
        # block for that fraction of a second
        time.sleep(val)
        # Creates warning and stops if below a certain value p
        if val < p:
            logging.warning(f'Stop! We have reached a value less than {p}')
            break
    logging.info(f'completed running thread')

# thread function in python  
def threadFunction():
    list_threads = []
    threads = 10
    for j in range(0,threads):
        t = threading.Thread(target = log, args=(0.2,))
        list_threads.append(t)
        t.start()   
        
    for j in range(0,threads):
        list_threads[j].join()

if __name__ == "__main__":

    # configure log handler 
    h = logging.FileHandler('out.log')
    h.setLevel(logging.DEBUG)
    h.setFormatter(logging.Formatter('[%(levelname)s] %(name)s: [%(threadName)s] %(message)s'))
    # set the new log handler
    logger = logging.getLogger()
    logger.setLevel(logging.DEBUG)
    logger.addHandler(h) 

    threadFunction()