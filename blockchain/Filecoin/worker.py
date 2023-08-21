import sys
import random
import os
import pandas as pd
import os
import locale

locale.setlocale(locale.LC_ALL, 'C')

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
    print("Total Sum:", total_sum)
    # Initialize document object
    #os.makedirs("outputs", exist_ok=True)
    #data = {'id': [1], 'value': total_sum}
    #df = pd.DataFrame(data)
    #df.to_csv('outputs/out.csv')
    #print("CSV saved as:")

if __name__ == "__main__":
    wallet_address = sys.argv[1]
    main(wallet_address)
