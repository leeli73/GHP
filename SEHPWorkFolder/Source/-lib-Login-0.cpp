    #include <iostream>
    #include <string>
    #include "sehp.h"
    using namespace std;
    int execmd(const char* cmd, char* result) {
        char buffer[128];
        FILE* pipe = _popen(cmd, "r");
        if (!pipe)
            return 0;
        while (!feof(pipe)) {
            if (fgets(buffer, 128, pipe)) {
                strcat(result, buffer);
            }
        }
        _pclose(pipe);
        return 1;
    }
    int main(int argc,char *argv[])
    {
        sehp Login(argc,argv);
        string UserName = Login.GetFrom("UserName");
        string PassWord = Login.GetFrom("PassWord");
        string SEPHWorkFolder = Login.GetSEPHWorkFolder();
        string Command = SEPHWorkFolder+"\\Executable\\UserDefined\\Login\\main.exe -username " + UserName + " -passwd " + PassWord;
        system(Command.c_str());
        cout<<endl;
        exit(0);
        return 0;
    }
