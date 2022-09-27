#include <bits/stdc++.h>
using namespace std;
#define ll long long

int main(){

    int n, arr[] = {1, 2, 5, 10, 20, 50, 100, 200, 500, 1000};

    cin>>n;
    for (int i = 9; i >= 0; i--){
        if(n >= arr[i]){
            cout<<arr[i] << " " << n/arr[i] << "\n";
            n %= arr[i];
        }
    }
    return 0;
}