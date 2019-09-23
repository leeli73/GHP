                    #include <iostream>
                    #include <string>
                    #include "sehp.h"
                    using namespace std;
                    int main(int argc,char *argv[])
                    {
                        sehp MessageList(argc,argv);
                        string Keyt = MessageList.GetCookies("Keyt");
                        string SEPHWorkFolder = MessageList.GetSEPHWorkFolder();
                        string Command = SEPHWorkFolder+"\\Executable\\UserDefined\\GetMessageList\\main.exe -username "+Keyt;
                        system(Command.c_str());
                        cout<<endl;
                        exit(0);
                        return 0;
                    }
