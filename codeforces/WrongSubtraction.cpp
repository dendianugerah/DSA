#include <iostream>
#include <algorithm>

using namespace std;

int main(){
    int n, k;
    cin>>n>>k;

    while (k)
    {
        if (n%10 == 0){
            n /= 10;
            k--;
        } else{
            n -= 1;
            k--;
        }
    }
    cout<<n<<endl;

    return 0;
}