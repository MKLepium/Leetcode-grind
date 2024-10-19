#include <mutex>
#include <condition_variable>
#include <functional>

using namespace std;


class H2O {
public:
    mutex m_lock;
    condition_variable m_cv;
    int turn;

    H2O() {
        
    }

    void hydrogen(function<void()> releaseHydrogen) {
        unique_lock<mutex> lock(m_lock);
        while(turn==2) m_cv.wait(lock);
        // releaseHydrogen() outputs "H". Do not change or remove this line.
        releaseHydrogen();
        turn++;
        m_cv.notify_all();
    }

    void oxygen(function<void()> releaseOxygen) {
        unique_lock<mutex> lock(m_lock);
        while(turn<2) m_cv.wait(lock);
        // releaseOxygen() outputs "O". Do not change or remove this line.
        releaseOxygen();
        turn = 0;
        m_cv.notify_all();
    }
};