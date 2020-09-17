#include <iostream>
#include <cmath>
#include <ctime>
#include <cstdlib>
#include <cstring>

using namespace std;

// ��������
int count;
// defineԤ������
#define LENGTH 10   
#define WIDTH  5
#define NEWLINE '\n'

/*
extern �洢�������ṩһ��ȫ�ֱ��������ã�ȫ�ֱ��������еĳ����ļ����ǿɼ��ġ�
����ʹ�� 'extern' ʱ�������޷���ʼ���ı�������ѱ�����ָ��һ��֮ǰ������Ĵ洢λ�á�
extern ����������һ���ļ�������һ��ȫ�ֱ���������
*/
extern void write_extern();
extern int a, b;
extern int c;
extern float f;

// ����һ���ṹ������ Books 
struct Books
{
	char  title[50];
	char  author[50];
	char  subject[100];
	int   book_id;
};

typedef struct {
	char  title[50];
	char  author[50];
	char  subject[100];
	int   book_id;
}eBooks;

/*
 ��Ĺ��캯�� 
 ���캯��������ΪĳЩ��Ա�������ó�ʼֵ��
 �����������
 ��������������ÿ��ɾ���������Ķ���ʱִ��,
 ���������������򣨱���ر��ļ����ͷ��ڴ�ȣ�ǰ�ͷ���Դ��
*/ 

class Box {

private:
	double length;   // ����
	double breadth;  // ���
	double height;   // �߶�
public:
	static int boxNum; //��̬��Ա����
	double volume;
	virtual double getVolume2(void);
	void setLength(double len);
	void setBreadth(double bre);
	void setHeight(double hei);
	int compare(Box box) {
		/*
		this ָ��
		ÿһ��������ͨ�� this ָ���������Լ��ĵ�ַ
		��Ԫ����û�� this ָ�룬��Ϊ��Ԫ������ĳ�Ա��
		ֻ�г�Ա�������� this ָ�롣
		*/
		return this->getVolume() > box.getVolume();
	}
	double getVolume(void) {
		volume = length * breadth * height;
		return volume;
	}

	void setVolume(double volume) {
		this->volume = volume;

	}
	// ���ظ�������� - ��
	Box operator- ()
	{
		length = -length;
		breadth = -breadth;
		height = -height;

		return Box(length, breadth, height);
	}
	// �����������
	Box operator+(const Box& b) {
		Box box;
		box.length = this->length + b.length;
		box.height = this->height + b.height;
		box.breadth = this->breadth + b.breadth;
		return box;
	}
	Box();
	Box(double length, double breadth, double height);
	~Box();
	/*
	�����Ԫ�����Ƕ��������ⲿ������Ȩ�����������˽�У�private����Ա�ͱ�����protected����Ա
	������Ԫ������ԭ��������Ķ����г��ֹ���������Ԫ���������ǳ�Ա����
	��Ԫ������һ���������ú�������Ϊ��Ԫ��������ԪҲ������һ���࣬���౻��Ϊ��Ԫ��.
	*/
	friend void printLength(Box box);
};

void printLength(Box box) {
	cout << "Length of box : " << box.length << endl;
}
Box::Box(double length, double breadth, double height) {
	boxNum++;
	length = length;
	breadth = breadth;
	height = height;
	cout << "init Box Volume: " << getVolume() << endl;
}
Box::Box(void) {
	boxNum++;
	length = 3.;
	breadth = 4;
	height = 2;
	cout << "init Box Volume: " << getVolume() << endl;
}
Box::~Box(void) {
	cout << "Object is being deleted" << endl;
}
double Box::getVolume2(void) {
	return length * breadth * height;
}
void Box::setLength(double len)
{
	length = len;
}
void Box::setBreadth(double bre)
{
	breadth = bre;
}
void Box::setHeight(double hei)
{
	height = hei;
}
int Box::boxNum = 0;

class redBox : public Box {
	public:
	string getColor() {
		return "RED";
	}
};

class Cost {
public:
	int getCost(int volume) {
		return volume * 10;
	}
};

class RectRedBox : public Box, public Cost {
	public:
		
		string getShape() {
			return "����";
		}
};


void MutiExtendTest() {
	
	RectRedBox crb;
	crb.setLength(3.0);
	crb.setBreadth(3.0);
	crb.setHeight(3.0);
	RectRedBox crb2;
	crb2.setLength(5.0);
	crb2.setBreadth(3.0);
	crb2.setHeight(3.0);



	RectRedBox *ptrBox;
	ptrBox = &crb;
	cout << "Volume of Box1: " << ptrBox->getVolume() << endl;
	if (crb.compare(crb2)) {
		cout << "Box2 is smaller than Box1" << endl;
	}
	else {
		cout << "Box2 is equal to or larger than Box1" << endl;
	}
	printLength(crb);
	cout << "���: " << crb.getVolume() << endl;
	cout << "��״: " << crb.getShape() << endl;
	cout << "����: " << crb.getCost(crb.getVolume()) << endl;
	cout << "total boxes " << Box::boxNum << endl;
	
}


void ExtendTest() {
	redBox rb;
	rb.setLength(3.0);
	rb.setBreadth(3.0);
	rb.setHeight(3.0);
	cout << "���: " << rb.getVolume() << endl;
	cout << "��ɫ: " << rb.getColor() << endl;
}

void ClassTest() {
	Box Box1;        // ���� Box1������Ϊ Box
	Box Box2;        // ���� Box2������Ϊ Box
	double volume = 0.0;     // ���ڴ洢���

	// box 1 ����
	Box1.setLength(6.0);
	Box1.setBreadth(7.0);
	Box1.setHeight(5.0);

	// box 2 ����
	Box2.setLength(10.0);
	Box2.setBreadth(12.0);
	Box2.setHeight(13.0);

	// box 1 �����
	volume = Box1.getVolume();
	cout << "Box1 �������" << volume << endl;

	// box 2 �����
	volume = Box2.getVolume();
	cout << "Box2 �������" << volume << endl;
	
}

void SructureTest() {
	Books b1;
	eBooks b2;

	strcpy_s(b1.title , "C++");
	strcpy_s(b1.author, "Runoob");
	strcpy_s(b1.subject, "�������");
	b1.book_id = 11123;

	strcpy_s(b2.title, "CSS �̳�");
	strcpy_s(b2.author, "Runoob");
	strcpy_s(b2.subject, "ǰ�˼���");
	b2.book_id = 12346;

	cout << "��һ������� : " << b1.title << endl;
	cout << "��һ�������� : " << b1.author << endl;
	cout << "��һ������Ŀ : " << b1.subject << endl;
	cout << "��һ���� ID : "  <<  b1.book_id << endl;

	// ��� Book2 ��Ϣ
	cout << "�ڶ�������� : " << b2.title << endl;
	cout << "�ڶ��������� : " << b2.author << endl;
	cout << "�ڶ�������Ŀ : " << b2.subject << endl;
	cout << "�ڶ����� ID : "  <<  b2.book_id << endl;

}


int func()
{
	return 0;
}

void funcB(void) 
{
	static int i = 5;
	i++;
	cout << i;
	cout << ::count << std::endl;
}

int max(int num1, int num2) {
	return num1 > num2 ? num1: num2;
}

void MathCalculateExample() {
	short  s = 10;
	int    i = -1000;
	long   l = 100000;
	float  f = 230.47;
	double d = 200.374;
	// ��ѧ����
	cout << "sin(d) :" << sin(d) << endl;
	cout << "abs(i)  :" << abs(i) << endl;
	cout << "floor(d) :" << floor(d) << endl;
	cout << "sqrt(f) :" << sqrt(f) << endl;
	cout << "pow( d, 2) :" << pow(d, 2) << endl;

}

void RandomNum() {
	int i, j;
	// ��������
	srand((unsigned)time(NULL));
	/* ���� 10 ������� */
	for (i = 0; i < 10; i++)
	{
		// ����ʵ�ʵ������
		j = rand();
		cout << "������� " << j << endl;
	}
}

void PointArray() {
	char  var[3] = { 'a', 'b', 'c' };
	char *ptr;
	ptr = var;
	for (int i = 0; i < 3; i++)
	{
		cout << "Address of var[" << i << "] = ";
		cout << (int*)ptr << endl;
		cout << "Value of var[" << i << "] = ";
		cout << *(ptr+i) << endl;
		// �ƶ�����һ��λ��
	}
}

void funC() {
	int a = 21;
	int c;
	c = a > 0 ? a : 0;
	c = a;
	cout << "Line 1 - =  �����ʵ����c ��ֵ = : " << c << endl;
	c += a;
	cout << "Line 2 - += �����ʵ����c ��ֵ = : " << c << endl;
	c -= a;
	cout << "Line 3 - -= �����ʵ����c ��ֵ = : " << c << endl;
	c *= a;
	cout << "Line 4 - *= �����ʵ����c ��ֵ = : " << c << endl;
	c /= a;
	cout << "Line 5 - /= �����ʵ����c ��ֵ = : " << c << endl;
	c = 200;
	c %= a;
	cout << "Line 6 - %= �����ʵ����c ��ֵ = : " << c << endl;
	c <<= 2;
	cout << "Line 7 - <<= �����ʵ����c ��ֵ = : " << c << endl;
	c >>= 2;
	cout << "Line 8 - >>= �����ʵ����c ��ֵ = : " << c << endl;
	c &= 2;
	cout << "Line 9 - &= �����ʵ����c ��ֵ = : " << c << endl;
	c ^= 2;
	cout << "Line 10 - ^= �����ʵ����c ��ֵ = : " << c << endl;
	c |= 2;
	cout << "Line 11 - |= �����ʵ����c ��ֵ = : " << c << endl;
}

double getAverage(int arr[], int size) {
	int sum = 0;
	for (int i = 0; i < size; i++) {
		sum += arr[i];
	}
	return double(sum) / size;
}

void AverageTest() {
	int balance[5] = { 1000, 2, 3, 17, 50 };
	double avg = getAverage(balance, 5);
	cout << "ƽ��ֵ�ǣ�" << avg << endl;
}

int * getRandomArr() {
	static int  r[10];

	// ��������
	srand((unsigned)time(NULL));
	for (int i = 0; i < 10; ++i)
	{
		r[i] = rand();
		cout << r[i] << endl;
	}
	return r;
}

void RandomArrTest() {
	int *p;
	p = getRandomArr();
	for (int i = 0; i < 10; i++)
	{
		cout << "*(p + " << i << ") : ";
		cout << *(p + i) << endl;
	}
}

void Point2Point() {
	int  var;
	int* ptr;
	int** pptr;
	var = 3000;
	// ��ȡ var �ĵ�ַ
	ptr = &var;
	// ʹ������� & ��ȡ ptr �ĵ�ַ
	pptr = &ptr;
	// ʹ�� pptr ��ȡֵ
	cout << "var ֵΪ :" << var << endl;
	cout << "*ptr ֵΪ:" << *ptr << endl;
	cout << "**pptr ֵΪ:" << **pptr << endl;
}

void ReferenceTest() {
	int    i;
	double d;
	// �������ñ���
	int& r = i;
	double& s = d;
	i = 5;
	cout << "Value of i : " << i << endl;
	cout << "Value of i reference : " << r << endl;
	d = 11.7;
	cout << "Value of d : " << d << endl;
	cout << "Value of d reference : " << s << endl;

}

void swap(int& x, int& y) {
	int temp;
	temp = x; /* �����ַ x ��ֵ */
	x = y;    /* �� y ��ֵ�� x */
	y = temp; /* �� x ��ֵ�� y  */
}

void SwapTest() {
	int a = 100;
	int b = 200;

	cout << "����ǰ��a ��ֵ��" << a << endl;
	cout << "����ǰ��b ��ֵ��" << b << endl;

	/* ���ú���������ֵ */
	swap(a, b);

	cout << "������a ��ֵ��" << a << endl;
	cout << "������b ��ֵ��" << b << endl;
}