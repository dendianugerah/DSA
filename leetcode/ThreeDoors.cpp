#include <bits/stdc++.h>
#define ll long long


using namespace std;

int main(){
    int tt;

    cin>>tt;
    while (tt--){
        ll x;
        cin>>x;

        ll p[4];
        cin>>p[1]>>p[2]>>p[3];

        if(p[x] != 0 && p[p[x]] != 0) cout<<"YES"<<endl;
        else cout<<"NO"<<endl;
    }
}