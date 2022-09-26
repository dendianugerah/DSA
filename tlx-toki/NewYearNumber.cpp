#include <bits/stdc++.h>
using namespace std;
#define ll long long

int main(){
    ios::sync_with_stdio(false);
    cin.tie(0);

    int n;
    scanf("%d", &n);
    if (n == 2 || n == 3 || n == 5){
        cout<<"YES";
    } else{
        cout<<"NO";
    }
    return 0;
}