#include <iostream>
using namespace std;
//1721

int main(){
    int n;
    char a, b, c, d;

    cin>>n;
    int arr[n];
    for (int i = 0; i < n; i++){
        cin>>a>>b>>c>>d;
        int count =1;
        if (a!=b)
        {
            count++;
        }
        if (b!=c && a!=c){
            count++;
        }
        if (a!=d && b!=d && c!=d){
            count++;
        }
        arr[i] = count-1;
    }
    for(int i = 0; i < n; i++){
        cout<<arr[i];
        if(i<n-1){
            cout<<" "<<endl;
        }
    }
    return 0;
}