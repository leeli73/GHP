import java.util.*;

public class test2java
{
    public static void main(String args[])
    {
        int[] arr = new int[5000];
        CHPXLib2Java test = new CHPXLib2Java(args,"METHOD2UPPER");
        System.out.println("你的浏览器信息为:"+test.UserAgrnt+"<br>");
        System.out.println("你访问的URL为:"+test.Url+"<br>");
        System.out.println("你的访问方式为:"+test.Method+"<br>");
        if(test.isSet("From", "chpx"))
        {
            System.out.println("检测到表单中存在chpx,对应的值为:"+test.Froms.get("chpx")+"<br>");
        }
        else
        {
            System.out.println("检测到表单中不存在chpx<br>");
        }
        System.out.println("表单测试");
        System.out.println("<form action=\"demo2java.chpx\" method=\"POST\">");
        System.out.println("chpx的值:<br><input type=\"text\" name=\"chpx\" value=\"Test\"><br>");
        System.out.println("<input type=\"submit\" value=\"提交\">");
        System.out.println("</form>");
        System.out.println("<a href=\"demo2java.chpx2java.txt\">查看源码</a><br>");
        System.out.println("<a href=\"demo2java.chpx?type=bubble\">冒泡排序</a><br>");
        System.out.println("<a href=\"demo2java.chpx?type=fast\">快速排序</a><br>");
        System.out.println("性能测试:默认使用快排排序5000个随机数(真正原生Java的性能)输入Get参数type=bubble切换为冒泡排序,删除参数或type=fast为快速排序<br>");
        Random rand = new Random();
        System.out.println("原始数据:<br>");
        for(int i=0;i<5000;i++)
        {
            if(i%10 == 0)
            {
                System.out.print("<br>");;
            }
            arr[i] = rand.nextInt(10000);
            System.out.print(arr[i]);
            System.out.print(",");
        }
        long startTime=System.currentTimeMillis();
        if(test.isSet("From","type"))
        {
            if(test.Froms.get("type").equals(new String("fast")))
            {
                quickSort(arr, 0, 4999);
            }
            else if(test.Froms.get("type").equals(new String("bubble")))
            {
                BubbleSort(arr);
            }
            else
            {
                System.out.println("<script>alert(\"未知的排序方式!\")</script>");
            }
        }
        else
        {
            quickSort(arr, 0, 4999);
        }
        System.out.println("排序后数据:<br>");
        for(int i=0;i<5000;i++)
        {
            if(i%10 == 0)
            {
                System.out.print("<br>");;
            }
            System.out.print(arr[i]);
            System.out.print(",");
        }
        long endTime=System.currentTimeMillis();
        System.out.println("程序运行时长:" + (endTime-startTime) + "ms");
    }
    private static void quickSort(int[] a, int low, int high) {
        if( low > high) {
            return;
        }
        int i = low;
        int j = high;
        int key = a[ low ];
        while( i< j) {
            while(i<j && a[j] > key){
                j--;
            }
            while( i<j && a[i] <= key) {
                i++;
            }
            if(i<j) {
                int p = a[i];
                a[i] = a[j];
                a[j] = p;
            }
        }
        int p = a[i];
        a[i] = a[low];
        a[low] = p;
        quickSort(a, low, i-1 );
        quickSort(a, i+1, high);
    }
    private static void BubbleSort(int[] a)
    {
        for(int i=0;i<a.length-1;i++){
            for(int j=0;j<a.length-1-i;j++){
                if(a[j]>a[j+1]){
                    int temp=a[j];
                    a[j]=a[j+1];
                    a[j+1]=temp;
                }
            }
        }
    }
}
class CHPXLib2Java
{
    public String UserAgrnt;
	public String Header;
	public String Method;
	public String Host;
	public String Url;
	public String ContentType;
    public Map Froms = new HashMap<String,String>();
    public Map Cookies = new HashMap<String,String>();
    public chpxlib2java(String argv[],String Style)
    {
        ContentType = new String("text/html");
        int Sign = 0;
        int fromflag = 1;
		int coookieflag = 1;
        for(int i=0;i<argv.length;i++)
        {
            if (i == 1)
			{
				UserAgrnt = new String(argv[1]);
			}
			else if (i == 2)
			{
				Header = new String(argv[2]);
			}
			else if (i == 3)
			{
				String temp = new String(argv[3]);;
				if (Style == "METHOD2UPPER")
				{
					temp.toUpperCase();
				}
				else if(Style == "METHOD2LOWER")
				{
					temp.toLowerCase();
				}
				Method = temp;
			}
			else if (i == 4)
			{
				Host = new String(argv[4]);
			}
			else if (i == 5)
			{
				Url = new String(argv[5]);
			}
			else
			{
                if(argv[i].equals(new String("#from")))
                {
                    Sign = 1;
                    continue;
                }
                else if(argv[i].equals(new String("from#")))
                {
                    Sign = 2;
                    continue;
                }
                else if(argv[i].equals(new String("from#")))
                {
                    Sign = 3;
                    continue;
                }
                else if(argv[i].equals(new String("from#")))
                {
                    Sign = 4;
                    continue;
                }

                if(Sign == 1)
				{
					if (fromflag > 0)
					{
						fromflag *= -1;
					}
					else
					{
                        Froms.put(argv[i-1], argv[i]);
						i++;
					}
				} 
                else if(Sign == 3)
				{
					
					if (coookieflag > 0)
					{
						coookieflag *= -1;
					}
					else
					{
						Cookies.put(argv[i-1], argv[i]);
						i++;
					}
				}
			}
        }
    }
    public boolean isSet(String Type,String Key)
    {
        if(Type.equals(new String("From").toUpperCase()))
        {
            return From.containsKey(Key);
        }
        else if(Type.equals(new String("Cookies").toUpperCase()))
        {
            return Cookies.containsKey(Key);
        }
        else
        {
            return false;
        }
    }
    public void AddCookies(String Key,String Value)
    {
        System.out.println("#AddCookies:"+Key+"-"+Value);
    }
    public void SetContentType(String ContentType)
    {
        System.out.println("#ContentType:"+ContentType);
    }
}