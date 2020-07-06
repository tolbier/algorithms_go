# Multiplier

This multiplier uses the KAratsuba Multiply
n = len(x & y)
x = (10^n) * a  + b
y = (10^n) * c  + d

result = (10^n) * ac + (10^(n/2))*((a+b)(c+d) - ac - bd) + bd
 