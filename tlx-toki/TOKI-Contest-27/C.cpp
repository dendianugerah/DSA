#include <bits/stdc++.h>
using namespace std;
#define ll long long


int main(){
    ll n, m;
    cin>>n>>m;

    cout<<n*m - max(n, m);

    return 0;
}