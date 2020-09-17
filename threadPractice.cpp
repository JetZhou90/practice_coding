#include <iostream>
#include <thread>
using namespace std;

thread::id main_thread_id = this_thread::get_id();
#define NUM_THREADS 5

// 线程的运行函数
void say_hello()
{
    cout << "Hello Runoob！\n" << endl;
    if (main_thread_id == this_thread::get_id())
        cout << "this is from main thread\n" ;
    else
        cout << "this is not the main thread\n";
}

void pause_thread(int n) {
    this_thread::sleep_for(chrono::seconds(n));
    cout << "pause of " << n << " seconds ended\n";
}


void threadTest() {
    thread t(say_hello);
    cout << t.hardware_concurrency() << endl;
    cout << "native_handle " << t.native_handle() << endl;
    t.join(); //链接线程， 主线程等待子线程运行结束后方可执行下一步
    thread a(say_hello);
    a.detach(); // 分离线程，独立于主线程并发执行，主线程后续代码段无需等待
    thread threads[5];
    cout << "Spawning 5 threads...\n";
    for (int i = 0; i < 5; i++)
        threads[i] = thread(pause_thread, i + 1);
    cout << "Done spawning threads. Now waiting for them to join:\n";
    for (auto& thread : threads)
        thread.join();
    cout << "All threads joined\n";
}


