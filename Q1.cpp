//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.

#include <iostream>
#include <vector>
#include <conio.h>
using namespace std;

bool mul3(const int i) { return ( i % 3 == 0 ); }
bool mul4(const int i) { return ( i % 5 == 0 ); }
void getMuls(vector<int> &x, const int bound) {
	for (int i = 1; i < bound; i++) {
		if (mul3(i)) x.push_back(i);
		else if (mul4(i)) x.push_back(i);
	}
}

int findSum(const vector<int>& x) {
	int sum = 0;
	for (int i = 0; i < x.size(); i++) 
		sum += x[i]; 
	return sum;
}

int main() {
	vector<int> arr;
	getMuls(arr, 1000);
	int sum = findSum(arr);
	cout << sum;
}