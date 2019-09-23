#pragma once
#include <string>
#include <iostream>
#include <stdio.h>
#include <WINSOCK2.H>
#include <vector>
using namespace std;
#pragma comment(lib,"WS2_32.lib")
class sehp
{
private:
	string SendData(string data);
	void SplitString(const string& s, vector<string>& v, const string& c);
public:
	string ID;
	sehp();//DEBUG
	sehp(int argc, char ** argv);
	~sehp();
	string GetFrom(string key);
	string GetCookies(string key);
	string GetUserAgent();
	string GetMethod();
	string GetHost();
	string GetPath();
	string SetCookies(string key, string value);
	string GetFile(string key);
	string SaveFile(string key, string name, string path);
	string GetSEPHWorkFolder();
};

