#include <bits/stdc++.h>

using namespace std;

int main(){
    int t;
    int a,b,c;

    cin>>t;
    while(t--){
        cin>>a>>b>>c;
        if (a+b==c || a+c==b || b+c==a){
            cout<<"YES"<<endl;
        } else{
            cout<<"NO"<<endl;
        }
    }


    return 0;
}
