#include <iostream>
using namespace std;

int main(){
    int n, m, c;
    int totalm = 0;
    int totalc = 0;

    cin>>n;
    for (int i = 0; i < n; i++){
        cin>>m>>c;
        if (m > c){
            totalm++;
        } else if (c > m){
            totalc++;
        } else{
            continue;
        }
    }

    if (totalc > totalm){
        cout<<"Chris"<<endl;
    } else if (totalm > totalc){
        cout<<"Mishka"<<endl;
    } else{
        cout<<"Friendship is magic!^^"<<endl;
    }
    

    return 0;
}