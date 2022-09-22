#include <bits/stdc++.h>
using namespace std;

string ToLower(string s){
    for(int i = 0; i < s.size(); i++){
        if(s[i] >= 'A' && s[i] <= 'Z'){
            s[i] += 32;
        }
    }
    return s;

};

int main(){
    cout<<ToLower("IkanHiuAA");

    return 0;
}