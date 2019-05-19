//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?

#include <iostream>
#include <vector>
#include <conio.h>
#include <math.h>
using namespace std;

#define VAL 600851475143

bool isFactor(long long int val, const int divi) { return val % divi == 0; }

bool isPrime(const int val) {
	for (int i = 2; i < sqrt(val) + 1; i++) {
		if (val % i == 0) return false;
	}
	return true;
}

int main() {
	int largest = 2;
	auto x = VAL;
	auto divi = 2;
	while (x > 1) {
		if (isFactor(x, divi) && isPrime(divi)) {
			if (divi > largest) largest = divi;
			x = x / divi;
			divi = 2;
		}
		else {
			divi++;
		}
	}
	cout << "The largest prime factor of " << VAL << " is " << largest;
	return 0;
}