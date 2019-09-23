                    #include <iostream>
                    #include <string>
                    #include "sehp.h"
                    using namespace std;
                    int main(int argc,char *argv[])
                    {
                        sehp UserList(argc,argv);
                        string Keyt = UserList.GetCookies("Keyt");
                        string SEPHWorkFolder = UserList.GetSEPHWorkFolder();
                        string Command = SEPHWorkFolder+"\\Executable\\UserDefined\\GetSelectUser\\main.exe -username "+Keyt;
                        system(Command.c_str());
                        cout<<endl;
                        exit(0);
                        return 0;
                    }
