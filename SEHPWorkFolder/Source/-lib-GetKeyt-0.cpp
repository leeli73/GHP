    #include <iostream>
    #include <string>
    #include "sehp.h"
    using namespace std;
    int main(int argc,char *argv[])
    {
        sehp Login(argc,argv);
        string UserName = Login.GetFrom("UserName");
        string SEPHWorkFolder = Login.GetSEPHWorkFolder();
        string Command = SEPHWorkFolder+"\\Executable\\UserDefined\\GetKeyt\\main.exe -username " + UserName;
        system(Command.c_str());
        cout<<endl;
        exit(0);
        return 0;
    }
