# Simple CLI Binary To Decimal & Vice Versa

def menu():
    print("Binary To Decimal & Decimal to Binary")
    print("1. Decimal To Binary")
    print("2. Binary To Decimal")
    print("3. Exit")

def decimalToBinary():
    try:
        decimal = int(input("Enter a decimal number:"))
        binary = bin(decimal).replace("0b", "")
        print(f"Converted Binary Value: {binary}")
    except ValueError:
        print("Please Enter a Valid Decimal")

def binaryToDecimal():
    try:
        conversion = int(input("Enter a binary number:"), 2)
        print(f"Converted Decimal Value: {conversion}")
    except ValueError:
        print("Please Enter a Valid Binary")

while 1:
    menu()
    try:
        choice = int(input("Enter your choice: "))
        if choice > 3:
            print("Please Enter a Valid Choice: ")
        if choice == 1:
            decimalToBinary()
        if choice == 2:
            binaryToDecimal()
        if choice == 3:
            break
    except ValueError:
        print("Please Enter a Valid Choice")