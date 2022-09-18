#include <iostream>
#include <string.h>
using namespace std;

int main(){
    int n;

    cin>>n;
    while (n--){
        int temp = 0;

        string s;
        cin>>s;
        
        if (s[0]=='Y' || s[0]=='y')
        {
            temp++;
        }
        if (s[1]=='E' || s[1]=='e')
        {
            temp++;
        }
        if (s[2]=='S' || s[2]=='s')
        {
            temp++;
        }
        if(temp == 3){
            cout<<"YES"<<endl;
        } else{
            cout<<"NO"<<endl;
        }
    }

    return 0;
}