#include <iostream>

using namespace std;

int main(){
    string a, b;

    cin>>a>>b;
    for (int i = 0; i < a.size(); i++){
        string temp = a;
        temp.erase(i, 1);
        if (temp == b){
            cout<<"Tentu saja bisa!\n";
            return 0;
        }
    }
    cout<<"Wah, tidak bisa :(\n";

    return 0;
}