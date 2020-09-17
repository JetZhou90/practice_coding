#include <iostream>

extern int count;

void write_extern(void)
{
	std::cout << "Count is " << count << std::endl;
}

//int main()
//{
//	cout << "Hello, World! My C++ " << endl;
//	return 0;
//}