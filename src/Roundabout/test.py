from sympy.ntheory import factorint
from itertools import product

# The number provided by the user
number = 3066844496547985532785966973086993824

# Factorizing the number using sympy's factorint function
factors = factorint(number)

# Showing the factors
print(factors)
