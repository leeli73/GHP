#pragma once
#include "iostream"
#include "string"
#include "map"
#include "vector"
#include <algorithm>
#define METHOD2UPPER 1
#define METHOD2LOWER 2
#define COOKIES 1
#define FROM 2
using namespace std;
class Request2CPP
{
public:
	string UserAgrnt;
	string Header;
	string Method;
	string Host;
	string Url;
	string ContentType;
	map<string, string> Froms;
	map<string, string> Cookies;
	Request2CPP(int argc, char ** argv,int style=1)
	{
		int sign = 0;
		int fromflag = 1;
		int coookieflag = 1;
		ContentType = string("text/html");
		for (int i = 1; i < argc; i++)
		{
			if (i == 1)
			{
				UserAgrnt = string(argv[1]);
			}
			else if (i == 2)
			{
				Header = string(argv[2]);
			}
			else if (i == 3)
			{
				string temp = string(argv[3]);;
				if (style == METHOD2UPPER)
				{
					transform(temp.begin(), temp.end(), temp.begin(), ::toupper);
				}
				else if(style == METHOD2LOWER)
				{
					transform(temp.begin(), temp.end(), temp.begin(), ::tolower);
				}
				Method = temp;
			}
			else if (i == 4)
			{
				Host = string(argv[4]);
			}
			else if (i == 5)
			{
				Url = string(argv[5]);
			}
			else
			{
				if (string(argv[i]) == string("#from")) 
				{
					sign = 1;
					continue;
				}
				else if(string(argv[i]) == string("from#"))
				{
					sign = 2; //停止接收from
					continue;
				}
				else if (string(argv[i]) == string("#cookies"))
				{
					sign = 3;
					continue;
				}
				else if (string(argv[i]) == string("cookies#"))
				{
					sign = 4; //停止接收cookies
					continue;
				}

				if (sign == 1) //开始接收from
				{
					if (fromflag > 0)
					{
						fromflag *= -1;
					}
					else
					{
						Froms.insert(pair<string, string>(string(argv[i-1]), string(argv[i])));
						i++;
					}
				}
				else if (sign == 3)
				{
					
					if (coookieflag > 0)
					{
						coookieflag *= -1;
					}
					else
					{
						Cookies.insert(pair<string, string>(string(argv[i - 1]), string(argv[i])));
						i++;
					}
				}
			}
		}
	}
	void AddCookies(string key, string value) 
	{
		cout << "#AddCookies:" << key << "-" << value << endl;
	}
	void SetContentType(string ContentType)
	{
		cout << "#ContentType:" << ContentType << endl;
	}
	bool isSet(int type,string key)
	{
		if (type == COOKIES)
		{
			if (Cookies.find(key) != Cookies.end())
			{
				return true;
			}
			else
			{
				return false;
			}
		}
		else if (type == FROM)
		{
			if (Froms.find(key) != Froms.end())
			{
				return true;
			}
			else
			{
				return false;
			}
		}
	}
};