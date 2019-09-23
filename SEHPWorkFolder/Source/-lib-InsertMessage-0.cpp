    #include <iostream>
    #include <string>
    #include "sehp.h"
    using namespace std;
    int main(int argc,char *argv[])
    {
        sehp Login(argc,argv);
        string MessageFrom = Login.GetFrom("MessageFrom");
        string MessageTo = Login.GetFrom("MessageTo");
        string MessageDate = Login.GetFrom("MessageDate");
        string SEPHWorkFolder = Login.GetSEPHWorkFolder();
        string Command = SEPHWorkFolder+"\\Executable\\UserDefined\\InsertMessage\\main.exe -MessageFrom " + MessageFrom + " -MessageTo "+MessageTo+" -MessageDate "+MessageDate;
        system(Command.c_str());
        cout<<endl;
        exit(0);
        return 0;
    }
