#include <bits/stdc++.h>
#define ll long long

using namespace std;

int main(){
    ll temp = 0;
    int n;
    scanf("%d", &n);
    
    for (int i = 1; i < n; i++){
        if (i % 3 == 0 || i % 5 == 0){
            temp += i;
        }
    }
    printf("%d", temp);

    return 0;
}