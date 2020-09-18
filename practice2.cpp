#include <fstream>
#include <iostream>
#include <csignal>
#include <windows.h>


using namespace std;

/*
��ȡ & д��ʵ��
*/

void IOtest() {
	char data[100];

	// ��дģʽ���ļ�
	ofstream outfile;
	outfile.open("afile.dat");
	cout << "Writing to the file" << endl;
	cout << "Enter your name: ";
	cin.getline(data, 100);

	// ���ļ�д���û����������
	outfile << data << endl;

	cout << "Enter your age: ";
	cin >> data;
	cin.ignore();

	// �ٴ����ļ�д���û����������
	outfile << data << endl;

	// �رմ򿪵��ļ�
	outfile.close();

	// �Զ�ģʽ���ļ�
	ifstream infile;
	infile.open("afile.dat");

	cout << "Reading from the file" << endl;
	infile >> data;

	// ����Ļ��д������
	cout << data << endl;

	// �ٴδ��ļ���ȡ���ݣ�����ʾ��
	infile >> data;
	cout << data << endl;

	// �رմ򿪵��ļ�
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
��̬�ڴ�
*/

void memoryTest() {
	double* pvalue = NULL; // ��ʼ��Ϊ null ��ָ��
	if (!(pvalue = new double))
	{
		cout << "Error: out of memory." << endl;
		exit(1);
	}
	pvalue = new double;   // Ϊ���������ڴ�
	*pvalue = 29494.99;     // �ڷ���ĵ�ַ�洢ֵ
	cout << "Value of pvalue : " << *pvalue << endl;
	delete pvalue;         // �ͷ��ڴ�

	/*
	����Ķ�̬�ڴ����
	*/
	int m = 4;
	int n = 8;
	int** array;
	// �ٶ������һά����Ϊ m�� �ڶ�ά����Ϊ n
	// ��̬����ռ�
	array = new int* [m];
	for (int i = 0; i < m; i++)
	{
		array[i] = new int[n];
		for (int j = 0; j < n; j++) {
			array[i][j] = i * j;
			cout << array[i][j] << "\n";
		}
	}

	//�ͷ�
	for (int i = 0; i < m; i++)
	{
		delete[] array[i];
	}
	delete[] array;

}

/*
# �� ## �����
*/
#define MKSTR( x ) #x // # �������� replacement-text ����ת��Ϊ���������������ַ�����
#define concat(a, b) a ## b //


/*
C++ �е�Ԥ�����
__LINE__	����ڳ������ʱ������ǰ�кš�
__FILE__	����ڳ������ʱ������ǰ�ļ�����
__DATE__	������һ����ʽΪ month/day/year ���ַ���������ʾ��Դ�ļ�ת��ΪĿ���������ڡ�
__TIME__	������һ����ʽΪ hour:minute:second ���ַ���������ʾ���򱻱����ʱ�䡣
*/

void PracticeTest() {
	cout << MKSTR(HELLO C++) << endl;
	int xy = 100;
	cout << concat(x, y)<<endl; // concat(x, y) ��x,yת����xy
	cout << "Value of __LINE__ : " << __LINE__ << endl;
	cout << "Value of __FILE__ : " << __FILE__ << endl;
	cout << "Value of __DATE__ : " << __DATE__ << endl;
	cout << "Value of __TIME__ : " << __TIME__ << endl;
}

void signalHandler(int signum)
{
	cout << "Interrupt signal (" << signum << ") received.\n";
	// �����ر�
	// ��ֹ����  
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