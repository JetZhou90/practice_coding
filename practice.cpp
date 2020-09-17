#include <iostream>
#include <cmath>
#include <ctime>
#include <cstdlib>
#include <cstring>

using namespace std;

// 变量声明
int count;
// define预处理器
#define LENGTH 10   
#define WIDTH  5
#define NEWLINE '\n'

/*
extern 存储类用于提供一个全局变量的引用，全局变量对所有的程序文件都是可见的。
当您使用 'extern' 时，对于无法初始化的变量，会把变量名指向一个之前定义过的存储位置。
extern 是用来在另一个文件中声明一个全局变量或函数。
*/
extern void write_extern();
extern int a, b;
extern int c;
extern float f;

// 声明一个结构体类型 Books 
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
 类的构造函数 
 构造函数可用于为某些成员变量设置初始值。
 类的析构函数
 析构函数它会在每次删除所创建的对象时执行,
 有助于在跳出程序（比如关闭文件、释放内存等）前释放资源。
*/ 

class Box {

private:
	double length;   // 长度
	double breadth;  // 宽度
	double height;   // 高度
public:
	static int boxNum; //静态成员数据
	double volume;
	virtual double getVolume2(void);
	void setLength(double len);
	void setBreadth(double bre);
	void setHeight(double hei);
	int compare(Box box) {
		/*
		this 指针
		每一个对象都能通过 this 指针来访问自己的地址
		友元函数没有 this 指针，因为友元不是类的成员。
		只有成员函数才有 this 指针。
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
	// 重载负运算符（ - ）
	Box operator- ()
	{
		length = -length;
		breadth = -breadth;
		height = -height;

		return Box(length, breadth, height);
	}
	// 重载相加运算
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
	类的友元函数是定义在类外部，但有权访问类的所有私有（private）成员和保护（protected）成员
	尽管友元函数的原型有在类的定义中出现过，但是友元函数并不是成员函数
	友元可以是一个函数，该函数被称为友元函数；友元也可以是一个类，该类被称为友元类.
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
			return "方形";
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
	cout << "体积: " << crb.getVolume() << endl;
	cout << "形状: " << crb.getShape() << endl;
	cout << "花费: " << crb.getCost(crb.getVolume()) << endl;
	cout << "total boxes " << Box::boxNum << endl;
	
}


void ExtendTest() {
	redBox rb;
	rb.setLength(3.0);
	rb.setBreadth(3.0);
	rb.setHeight(3.0);
	cout << "体积: " << rb.getVolume() << endl;
	cout << "颜色: " << rb.getColor() << endl;
}

void ClassTest() {
	Box Box1;        // 声明 Box1，类型为 Box
	Box Box2;        // 声明 Box2，类型为 Box
	double volume = 0.0;     // 用于存储体积

	// box 1 详述
	Box1.setLength(6.0);
	Box1.setBreadth(7.0);
	Box1.setHeight(5.0);

	// box 2 详述
	Box2.setLength(10.0);
	Box2.setBreadth(12.0);
	Box2.setHeight(13.0);

	// box 1 的体积
	volume = Box1.getVolume();
	cout << "Box1 的体积：" << volume << endl;

	// box 2 的体积
	volume = Box2.getVolume();
	cout << "Box2 的体积：" << volume << endl;
	
}

void SructureTest() {
	Books b1;
	eBooks b2;

	strcpy_s(b1.title , "C++");
	strcpy_s(b1.author, "Runoob");
	strcpy_s(b1.subject, "编程语言");
	b1.book_id = 11123;

	strcpy_s(b2.title, "CSS 教程");
	strcpy_s(b2.author, "Runoob");
	strcpy_s(b2.subject, "前端技术");
	b2.book_id = 12346;

	cout << "第一本书标题 : " << b1.title << endl;
	cout << "第一本书作者 : " << b1.author << endl;
	cout << "第一本书类目 : " << b1.subject << endl;
	cout << "第一本书 ID : "  <<  b1.book_id << endl;

	// 输出 Book2 信息
	cout << "第二本书标题 : " << b2.title << endl;
	cout << "第二本书作者 : " << b2.author << endl;
	cout << "第二本书类目 : " << b2.subject << endl;
	cout << "第二本书 ID : "  <<  b2.book_id << endl;

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
	// 数学运算
	cout << "sin(d) :" << sin(d) << endl;
	cout << "abs(i)  :" << abs(i) << endl;
	cout << "floor(d) :" << floor(d) << endl;
	cout << "sqrt(f) :" << sqrt(f) << endl;
	cout << "pow( d, 2) :" << pow(d, 2) << endl;

}

void RandomNum() {
	int i, j;
	// 设置种子
	srand((unsigned)time(NULL));
	/* 生成 10 个随机数 */
	for (i = 0; i < 10; i++)
	{
		// 生成实际的随机数
		j = rand();
		cout << "随机数： " << j << endl;
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
		// 移动到下一个位置
	}
}

void funC() {
	int a = 21;
	int c;
	c = a > 0 ? a : 0;
	c = a;
	cout << "Line 1 - =  运算符实例，c 的值 = : " << c << endl;
	c += a;
	cout << "Line 2 - += 运算符实例，c 的值 = : " << c << endl;
	c -= a;
	cout << "Line 3 - -= 运算符实例，c 的值 = : " << c << endl;
	c *= a;
	cout << "Line 4 - *= 运算符实例，c 的值 = : " << c << endl;
	c /= a;
	cout << "Line 5 - /= 运算符实例，c 的值 = : " << c << endl;
	c = 200;
	c %= a;
	cout << "Line 6 - %= 运算符实例，c 的值 = : " << c << endl;
	c <<= 2;
	cout << "Line 7 - <<= 运算符实例，c 的值 = : " << c << endl;
	c >>= 2;
	cout << "Line 8 - >>= 运算符实例，c 的值 = : " << c << endl;
	c &= 2;
	cout << "Line 9 - &= 运算符实例，c 的值 = : " << c << endl;
	c ^= 2;
	cout << "Line 10 - ^= 运算符实例，c 的值 = : " << c << endl;
	c |= 2;
	cout << "Line 11 - |= 运算符实例，c 的值 = : " << c << endl;
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
	cout << "平均值是：" << avg << endl;
}

int * getRandomArr() {
	static int  r[10];

	// 设置种子
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
	// 获取 var 的地址
	ptr = &var;
	// 使用运算符 & 获取 ptr 的地址
	pptr = &ptr;
	// 使用 pptr 获取值
	cout << "var 值为 :" << var << endl;
	cout << "*ptr 值为:" << *ptr << endl;
	cout << "**pptr 值为:" << **pptr << endl;
}

void ReferenceTest() {
	int    i;
	double d;
	// 声明引用变量
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
	temp = x; /* 保存地址 x 的值 */
	x = y;    /* 把 y 赋值给 x */
	y = temp; /* 把 x 赋值给 y  */
}

void SwapTest() {
	int a = 100;
	int b = 200;

	cout << "交换前，a 的值：" << a << endl;
	cout << "交换前，b 的值：" << b << endl;

	/* 调用函数来交换值 */
	swap(a, b);

	cout << "交换后，a 的值：" << a << endl;
	cout << "交换后，b 的值：" << b << endl;
}