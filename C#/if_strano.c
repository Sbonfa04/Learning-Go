#include <stdio.h>

int main(){
    int x,y=0,z=3;
    scanf("%d",&x);
    y = x>3? z++ : z--; //if (x>3) y=z and z++; else y=z and z--; 
    printf("%d,%d,%d \n",x,y,z);
    return 0;
}