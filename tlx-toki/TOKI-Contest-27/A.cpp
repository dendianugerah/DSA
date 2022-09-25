#include <bits/stdc++.h>
using namespace std;

int main(){
    cin.tie(0);
    int N;
    vector<int> A(N);
    //int A[N];

    cin>>N;
    for (int i = 0; i < N; i++){
        cin>>A[i];
    }
    for (int i = 0 ; i < N; i++){
        if (A[i] % 2 == 0){
            cout<<"YA"<<endl;
            return 0;
        } 
    }

    cout<<"TIDAK"<<endl;

    return 0;
}