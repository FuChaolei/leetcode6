#include <bits/stdc++.h>
using namespace std;
class AbstrcatFruit
{
public:
    virtual void ShowName() = 0;
};
class Apple : public AbstractFruit
{
    virtual void ShowName()
    {
        cout << "i'm apple" << endl;
    }
};
class Banana : public AbstractFruit
{
    virtual void ShowName()
    {
        cout << "i'm banana" << endl;
    }
};
class Factory
{
public:
    static AbstrctFruit *Createfruit(string fruit)
    {
        if (fruit == "apple")
        {
            return new Apple;
        }
        else if (fruit == "Banana")
        {
            return new Banana;
        }
        return nullptr;
    }
};
int main()
{
    Factory *factory = new Factory;
    AbstrctFruit *fruit = factory->Createfruit("apple");
    fruit->ShowName();
    delete fruit;
    delete factory;
    fruit = nullptr;
    factory = nullptr;
    return 0;
}
