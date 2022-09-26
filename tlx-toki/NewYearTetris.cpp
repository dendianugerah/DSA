#include <bits/stdc++.h>
using namespace std;
#define ll long long

int main(){
 ios::sync_with_stdio(false);
 cin.tie(0);

    int n;
    scanf("%d", &n);
    if(n % 2 == 1){
        cout<<(n*n) - 1;
    } else{
        cout<<(n*n);
    }

    return 0;
}