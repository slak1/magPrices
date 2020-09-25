

package main
import (
        "fmt"
        //"math"
          )

func przez100(x int) float64{
  wyn:=float64(x)/100
  return wyn
  }
func main(){
     a:=0; b:=0; c:=0; x:=0; y:=0
   for{
      a=a+1
      if a>10000 {a=0;break}
    for{
       b=b+1
       if b>10000 {b=0;break}
     for {
       c=c+1;
       if c>10000 {y=y+c;c=0;break}
       suma:=a+b+c; iloczyn:=a*b*c; 
       if suma*10000<iloczyn{y=y+c;c=0;break}
       if suma*10000==iloczyn {
           x=x+1
           fmt.Println("RESULT nr.",x,"-> ",przez100(a),"*",przez100(b),"*",przez100(c)," = ",przez100(suma));
         }
     }
    }
   }
   fmt.Println("of actions performed {loop c} =",y)
 }
