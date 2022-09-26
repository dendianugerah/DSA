#include <bits/stdc++.h>
using namespace std;
#define ll long long

string s;
int temp;
int main(){
    ios::sync_with_stdio(false);
    cin.tie(0);
    
    string Ex = "halo dunia";
    string X = "HALO DUNIA";

    getline(cin, s);
    for(int i = 0; i < s.size(); i++){
        if(s[i] == Ex[i] || s[i] == X[i]){
            temp++;
        }
    }
    printf("%d", temp);
    


    return 0;
}