
#include "sehp.h"


sehp::sehp()
{
	ID = "test";
}
sehp::sehp(int argc, char ** argv)
{
	ID = string(argv[1]);
}
sehp::~sehp()
{
	SendData("Delete|" + ID);
}
string sehp::SendData(string data)
{
	WSADATA wsd;
	SOCKET  s;
	if (WSAStartup(MAKEWORD(2, 2), &wsd) != 0)
	{
		return "WSAStartup failed !";
	}
	s = socket(AF_INET, SOCK_DGRAM, 0);
	if (s == INVALID_SOCKET)
	{
		WSACleanup();
		return "socket() failed, Error Code:" + WSAGetLastError();
	}
	char        buf[1024];
	SOCKADDR_IN servAddr;
	SOCKET      sockClient = socket(AF_INET, SOCK_DGRAM, 0);
	int         nRet;
	int size = data.length();
	ZeroMemory(buf, size);
	strcpy(buf, data.c_str());
	servAddr.sin_family = AF_INET;
	servAddr.sin_addr.S_un.S_addr = inet_addr("127.0.0.1");
	servAddr.sin_port = htons(9981);
	int nServAddLen = sizeof(servAddr);
	if (sendto(sockClient, buf, size, 0, (sockaddr *)&servAddr, nServAddLen) == SOCKET_ERROR)
	{
		closesocket(s);
		WSACleanup();
		return "recvfrom() failed:" + WSAGetLastError();
	}
	ZeroMemory(buf, 1024);
	nRet = recvfrom(sockClient, buf, 1024, 0, (sockaddr *)&servAddr, &nServAddLen);
	if (SOCKET_ERROR == nRet)
	{
		closesocket(s);
		WSACleanup();
		return "recvfrom failed !";
	}
	closesocket(s);
	WSACleanup();
	return string(buf);
}
void sehp::SplitString(const string& s, vector<string>& v, const string& c)
{
	string::size_type pos1, pos2;
	pos2 = s.find(c);
	pos1 = 0;
	while (string::npos != pos2)
	{
		v.push_back(s.substr(pos1, pos2 - pos1));
		pos1 = pos2 + c.size();
		pos2 = s.find(c, pos1);
	}
	if (pos1 != s.length())
		v.push_back(s.substr(pos1));
}
string sehp::GetFrom(string key)
{
	vector<string> v;
	SplitString(SendData("GetFrom|" + ID + "|" + key), v, "|");
	return v[1];
}
string sehp::GetCookies(string key)
{
	vector<string> v;
	SplitString(SendData("GetCookies|" + ID + "|" + key), v, "|");
	return v[1];
}
string sehp::GetUserAgent()
{
	vector<string> v;
	SplitString(SendData("GetUserAgent|" + ID), v, "|");
	return v[1];
}
string sehp::GetMethod()
{
	vector<string> v;
	SplitString(SendData("GetMethod|" + ID), v, "|");
	return v[1];
}
string sehp::GetHost()
{
	vector<string> v;
	SplitString(SendData("GetHost|" + ID), v, "|");
	return v[1];
}
string sehp::GetPath()
{
	vector<string> v;
	SplitString(SendData("GetPath|" + ID), v, "|");
	return v[1];
}
string sehp::GetFile(string key)
{
	vector<string> v;
	SplitString(SendData("GetFile|" + ID +"|" + key), v, "|");
	return v[1];
}
string sehp::SaveFile(string key,string name,string path)
{
	vector<string> v;
	SplitString(SendData("SaveFile|" + ID+"|"+key+"|"+name+"|"+path), v, "|");
	return v[1];
}
string sehp::GetSEPHWorkFolder()
{
	vector<string> v;
	SplitString(SendData("GetSEPHWorkFolder|" + ID), v, "|");
	return v[1];
}