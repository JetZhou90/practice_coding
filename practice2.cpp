#include <fstream>
#include <iostream>
#include <csignal>
#ifdef WIN32
#include <windows.h>
#else
#include <unistd.h>
#endif

using namespace std;

/*
读取 & 写入实例
*/

void IOtest() {
	char data[100];

	// 以写模式打开文件
	ofstream outfile;
	outfile.open("afile.dat");
	cout << "Writing to the file" << endl;
	cout << "Enter your name: ";
	cin.getline(data, 100);

	// 向文件写入用户输入的数据
	outfile << data << endl;

	cout << "Enter your age: ";
	cin >> data;
	cin.ignore();

	// 再次向文件写入用户输入的数据
	outfile << data << endl;

	// 关闭打开的文件
	outfile.close();

	// 以读模式打开文件
	ifstream infile;
	infile.open("afile.dat");

	cout << "Reading from the file" << endl;
	infile >> data;

	// 在屏幕上写入数据
	cout << data << endl;

	// 再次从文件读取数据，并显示它
	infile >> data;
	cout << data << endl;

	// 关闭打开的文件
	infile.close();
}

double division(int a, int b) {
	if (b == 0) {
		throw "Division by 0 condition!";
	}
	return (a / b);
}

void try_catch_test() {
	int x = 50;
	int y = 2;
	double z = 0;
	try {
		z = division(x, y);
		cout << z << endl;
	}
	catch (const char* msg) {
		cerr << msg << endl;
	}

}

/*
动态内存
*/

void memoryTest() {
	double* pvalue = NULL; // 初始化为 null 的指针
	if (!(pvalue = new double))
	{
		cout << "Error: out of memory." << endl;
		exit(1);
	}
	pvalue = new double;   // 为变量请求内存
	*pvalue = 29494.99;     // 在分配的地址存储值
	cout << "Value of pvalue : " << *pvalue << endl;
	delete pvalue;         // 释放内存

	/*
	数组的动态内存分配
	*/
	int m = 4;
	int n = 8;
	int** array;
	// 假定数组第一维长度为 m， 第二维长度为 n
	// 动态分配空间
	array = new int* [m];
	for (int i = 0; i < m; i++)
	{
		array[i] = new int[n];
		for (int j = 0; j < n; j++) {
			array[i][j] = i * j;
			cout << array[i][j] << "\n";
		}
	}

	//释放
	for (int i = 0; i < m; i++)
	{
		delete[] array[i];
	}
	delete[] array;

}

/*
# 和 ## 运算符
*/
#define MKSTR( x ) #x // # 运算符会把 replacement-text 令牌转换为用引号引起来的字符串。
#define concat(a, b) a ## b //


/*
C++ 中的预定义宏
__LINE__	这会在程序编译时包含当前行号。
__FILE__	这会在程序编译时包含当前文件名。
__DATE__	这会包含一个形式为 month/day/year 的字符串，它表示把源文件转换为目标代码的日期。
__TIME__	这会包含一个形式为 hour:minute:second 的字符串，它表示程序被编译的时间。
*/

void PracticeTest() {
	cout << MKSTR(HELLO C++) << endl;
	int xy = 100;
	cout << concat(x, y)<<endl; // concat(x, y) 将x,y转换成xy
	cout << "Value of __LINE__ : " << __LINE__ << endl;
	cout << "Value of __FILE__ : " << __FILE__ << endl;
	cout << "Value of __DATE__ : " << __DATE__ << endl;
	cout << "Value of __TIME__ : " << __TIME__ << endl;
}

void signalHandler(int signum)
{
	cout << "Interrupt signal (" << signum << ") received.\n";
	// 清理并关闭
	// 终止程序  
	exit(signum);
}

void signalTest() {
	int i = 0;
	signal(SIGINT, signalHandler);

	while (++i) {
		cout << "Going to sleep...." << endl;
		if (i == 100) {
			raise(SIGINT);
		}
		Sleep(5);
	}

}