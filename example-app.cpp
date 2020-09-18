#include <opencv2/core/core.hpp>
#include <opencv2/highgui/highgui.hpp>
#include <iostream>
#include <torch/script.h>




int torchTest() {
	torch::Tensor tensor = torch::rand({ 2, 3 });
	std::cout << tensor << std::endl;
	
	//cv::Mat img = cv::imread("D:\\Projects\\seal_validation\\cmb_seal\\circle.png");
	//// 创建一个名为 "图片"窗口    
	//cv::namedWindow("opencv-c++");
	//// 在窗口中显示图片   
	//cv::imshow("opencv-c++", img);
	//// 等待6000 ms后窗口自动关闭    
	//cv::waitKey(6000);
	return 0;

}
