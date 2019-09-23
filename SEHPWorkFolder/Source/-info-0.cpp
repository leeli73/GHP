      #include <iostream>
      #include <string>
      #include "sehp.h"
      using namespace std;
      int main(int argc,char *argv[])
      {
          sehp Admin(argc,argv);
          string Keyt = Admin.GetCookies("Keyt");
          string UserName = Admin.GetCookies("UserName");
          if(UserName.compare("admin"))
          {
              cout<<endl;
              exit(0);
              return 0;
          }
          string SEPHWorkFolder = Admin.GetSEPHWorkFolder();
          cout<< "<div class=\"py-5\">" <<endl;
          cout<< "<div class=\"container\">" <<endl;
          cout<< "<div class=\"row\">" <<endl;
          cout<< "<div class=\"col-md-12\">" <<endl;
          cout<< "<center><h1>用户管理</h1></center>" <<endl;
          cout<< "<ul class=\"list-group\">" <<endl;
          string Command = SEPHWorkFolder+"\\Executable\\UserDefined\\GetAllUser\\main.exe";
          system(Command.c_str());
          cout<< "</ul></div></div></div></div>" <<endl;

          cout<< "<div class=\"py-5\">" <<endl;
          cout<< "<div class=\"container\">" <<endl;
          cout<< "<div class=\"row\">" <<endl;
          cout<< "<div class=\"col-md-12\">" <<endl;
          cout<< "<center><h1>资源管理</h1></center>" <<endl;
          cout<< "<ul class=\"list-group\">" <<endl;
          Command = SEPHWorkFolder+"\\Executable\\UserDefined\\GetAllAsset\\main.exe";
          system(Command.c_str());
          cout<< "</ul></div></div></div></div>" <<endl;

          cout<< "<div class=\"py-5\">" <<endl;
          cout<< "<div class=\"container\">" <<endl;
          cout<< "<div class=\"row\">" <<endl;
          cout<< "<div class=\"col-md-12\">" <<endl;
          cout<< "<center><h1>消息管理</h1></center>" <<endl;
          cout<< "<ul class=\"list-group\">" <<endl;
          Command = SEPHWorkFolder+"\\Executable\\UserDefined\\GetAllMessage\\main.exe";
          system(Command.c_str());
          cout<< "</ul></div></div></div></div>" <<endl;
          exit(0);
          return 0;
      }
