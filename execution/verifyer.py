import sys
import random

def extract_numbers(wallet_address):
    # Extract numbers from the wallet address
    numbers = [int(char) for char in wallet_address if char.isdigit()]
    return numbers

def main(wallet_address):
    numbers = extract_numbers(wallet_address)
    total_sum = 0
    
    for num in numbers:
        if num % 2 == 0:
            value = random.randint(1, 10000)
        else:
            value = random.randint(-10000, -1)
        
        total_sum += value
        print(f"Number: {num}, Value: {value}")
    
    print("Total Sum:", total_sum)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <wallet_address>")
    else:
        wallet_address = sys.argv[1]
        main(wallet_address)