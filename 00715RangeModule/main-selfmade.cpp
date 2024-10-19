#include<iostream>
#include<map>

class RangeModule {
private:
    std::map<int, bool> m_trackedNums;
public:
    RangeModule() {
    }

    void setNumber(int n, bool tracked) {
        m_trackedNums[n] = tracked;
    }

    bool isTracked(int n) {
        if (m_trackedNums.find(n) != m_trackedNums.end()) {
            return m_trackedNums[n];
        }
        return false;
    }
    
    void addRange(int left, int right) {
        for (int i = left; i < right; i++)
        {
            this->setNumber(i, true);
        }
    }
    
    bool queryRange(int left, int right) {
        for (int i = left; i < right; i++)
        {
            if(!this->isTracked(i)){
                return false;
            }
        }
        return true;
    }
    
    void removeRange(int left, int right) {
        for (int i = left; i < right; i++)
        {
            this->setNumber(i, false);
        }
    }
};

/**
 * Your RangeModule object will be instantiated and called as such:
 * RangeModule* obj = new RangeModule();
 * obj->addRange(left,right);
 * bool param_2 = obj->queryRange(left,right);
 * obj->removeRange(left,right);
 */

int main() {
    RangeModule* obj = new RangeModule();
    obj->addRange(10, 20);
    obj->removeRange(14, 16);
    bool param_2 = obj->queryRange(10, 14);
    std::cout << param_2 << std::endl;
    param_2 = obj->queryRange(13, 15);
    std::cout << param_2 << std::endl;
    param_2 = obj->queryRange(16, 17);
    std::cout << param_2 << std::endl;
    return 0;
}