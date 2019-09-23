    #include <iostream>
    #include <string>
    #include "sehp.h"
    using namespace std;
    int main(int argc,char *argv[])
    {
        sehp Delete(argc,argv);
        string ID = Delete.GetFrom("ID");
        string UserName = Delete.GetCookies("Keyt");
        string SEPHWorkFolder = Delete.GetSEPHWorkFolder();
        string Command = SEPHWorkFolder+"\\Executable\\UserDefined\\DeleteAsset\\main.exe -username " + UserName + " -id "+ID;
        system(Command.c_str());
        cout<<endl;
        exit(0);
        return 0;
    }
