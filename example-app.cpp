#include <opencv2/core/core.hpp>
#include <opencv2/highgui/highgui.hpp>
#include <iostream>
#include <torch/script.h>




int torchTest() {
	torch::Tensor tensor = torch::rand({ 2, 3 });
	std::cout << tensor << std::endl;
	
	//cv::Mat img = cv::imread("D:\\Projects\\seal_validation\\cmb_seal\\circle.png");
	//// ����һ����Ϊ "ͼƬ"����    
	//cv::namedWindow("opencv-c++");
	//// �ڴ�������ʾͼƬ   
	//cv::imshow("opencv-c++", img);
	//// �ȴ�6000 ms�󴰿��Զ��ر�    
	//cv::waitKey(6000);
	return 0;

}
