#include <bits/stdc++.h>
using namespace std;
#define ll long long


int main(){
    ll n;
    scanf("%lld", &n);

    ll k = 1;
    while (k <= n){
        k *= 2;
    }
    printf("%lld", k - n);



    return 0;
}