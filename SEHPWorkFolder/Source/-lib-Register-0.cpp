    #include <iostream>
    #include <string>
    #include "sehp.h"
    using namespace std;
    int main(int argc,char *argv[])
    {
        sehp Login(argc,argv);
        string UserName = Login.GetFrom("UserName");
        string PassWord = Login.GetFrom("PassWord");
        string Phone = Login.GetFrom("Phone");
        string Email = Login.GetFrom("Email");
        string SEPHWorkFolder = Login.GetSEPHWorkFolder();
        string Command = SEPHWorkFolder+"\\Executable\\UserDefined\\Register\\main.exe -username " + UserName+" -passwd "+PassWord+" -phone "+Phone+" -email " + Email;
        system(Command.c_str());
        cout<<endl;
        exit(0);
        return 0;
    }
