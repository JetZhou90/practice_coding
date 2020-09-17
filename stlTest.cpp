#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;



void vectorOperate() {
	vector<int> vec;
	int i;
	cout << "vector size = " << vec.size() << endl;
	for (i = 0; i < 5; i++)
		vec.push_back(i); // 添加值
	cout << "extended vector size = " << vec.size() << endl;
	// 访问向量中的 5 个值
	for (i = 0; i < 5; i++)
		cout << "value of vec [" << i << "] = " << vec[i] << endl;
	// 使用迭代器 iterator 访问值
	vector<int>::iterator v = vec.begin();
	while (v != vec.end()) {
		cout << "value of v = " << *v << endl;
		v++;
	}
}

vector<int> pick_vector_with_biggest_fifth_element(vector<int> left, vector<int> right)
{
	if (left[5] < right[5])
	{
		return(right);
	}
	// else
	return left;
}

int vector_demo(void)
{
	cout << "vector demo" << endl;
	vector<int> left(7);
	vector<int> right(7);
	left[5] = 7;
	right[5] = 8;
	cout << left[5] << endl;
	cout << right[5] << endl;
	vector<int> biggest(pick_vector_with_biggest_fifth_element(left, right));
	cout << "the biggest 5th element is "<<biggest[5] << endl;

	return 0;
}

void sortVector() {
	vector<int> vec;
	int num = 0;
	int tmp;
	while (num <7) {
		cin >> tmp;
		vec.push_back(tmp);
		num++;
	}
	cout << "Sorted: " << endl;
	sort(vec.begin(), vec.end());
	int i = 0;
	for (i = 0; i < vec.size(); i++) {
		cout << vec[i] << endl;;
	}
}


void vectorTest() {
	
	sortVector();

}