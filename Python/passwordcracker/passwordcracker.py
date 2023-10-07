# Simple Brute Force Password Cracker
import string
from itertools import product
import time

def bruteforce(pw):
    attempts = 0
    chars = string.printable
    for i in range(1, len(pw)+1):
        # giving cartesian product
        for guess in product(chars, repeat=i):
            attempts += 1
            guess = ''.join(guess)
            if guess == pw:
                return (attempts, guess)
    return (attempts, None)

start_time = time.time()
pw = input("Input the password to crack: ")
attempts, guess = bruteforce(pw)
print(f"Password cracked in {attempts} attempts. The password is {guess}.") if guess else print(f"Password not cracked after {attempts} attempts.")
print(f"Time taken to complete execution: {time.time() - start_time}")