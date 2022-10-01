#include <bits/stdc++.h>
using namespace std;
#define ll long long

int main(){
    int tt;
    int temp = 0;
    ll a, b, c, d;

    cin>>tt;
    while (tt--){
        cin>>a>>b>>c>>d;
        cout<<(a < b) + (a < c) + (a < d)<<"\n";
    }
}