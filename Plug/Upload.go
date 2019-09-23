package main

import (
    "gitee.com/johng/gf/g"
    "gitee.com/johng/gf/g/os/gfile"
    "gitee.com/johng/gf/g/net/ghttp"
    "os/exec"
    "fmt"
)
// 执行文件上传处理，上传到系统临时目录 /tmp
func Upload(r *ghttp.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    if _,ok := r.Form["UserName"];ok{
        if f, h, e := r.FormFile("upload-file"); e == nil {
            defer f.Close()
            name   := gfile.Basename(h.Filename)
            buffer := make([]byte, h.Size)
            f.Read(buffer)
            UserName,_ := r.Form["UserName"]
            gfile.PutBinContents("D:\\CHP\\Code\\Version-Public-1\\WWW\\asset\\" + UserName[0] + "_" +name, buffer)
            var isPublic string
            isPublic = "0"
            if _,ok := r.Form["isPublic"];ok{
                isPublic = "1"
            }
            params := []string{"-name",name,"-username",UserName[0],"-public",isPublic}
            cmd := exec.Command("D:\\CHP\\Code\\Version-Public-1\\SEHPWorkFolder\\Executable\\UserDefined\\NewAsset\\main.exe", params...)
            fmt.Println(cmd.Args)
            cmd.Start()
            cmd.Wait()
            r.Response.Header().Set("content-type","text/html;charset=utf-8")
            r.Response.Write("文件:"+name + "上传成功...<script>alert('"+name+"上传成功');window.location = \"upload/show\";parent.location.reload();</script>")
        } else {
            r.Response.Write(e.Error())
        }
    } else {
        r.Response.Write("非法请求")
    }
}

// 展示文件上传页面
func UploadShow(r *ghttp.Request) {
    /*r.Response.Write(`
    <html>
    <head>
        <title>上传文件</title>
    </head>
        <body>
            <form enctype="multipart/form-data" action="/upload" method="post">
                <input type="file" name="upload-file" />
                <input type="submit" value="upload" />
            </form>
        </body>
    </html>
    `)*/
    r.Response.Write(`<!DOCTYPE html>
    <html>
    
    <head>
      <meta charset="utf-8">
      <meta name="viewport" content="width=device-width, initial-scale=1">
      <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css" type="text/css">
      <link rel="stylesheet" href="https://static.pingendo.com/bootstrap/bootstrap-4.1.3.css">
    </head>
    
    <body>
      <form class="form-inline w-100" enctype="multipart/form-data" action="/upload" method="post" id="fileform" width="100%">
        <div class="input-group" width="100%">
          <input type="file" multiple="multiple" class="form-control" placeholder="请选择文件" name="upload-file" id="upload-file" width="70%">
          <div class="input-group-append" width="30%"><button class="btn btn-primary" type="submit" id="uploadb">点击上传</button></div>
        </div>
        <div class="form-check">
            <label class="form-check-label">
                <input class="form-check-input" type="checkbox" name="isPublic">是否公开
            </label>
        </div>
      </form>
      <br>
      <h3>关闭此模态框上传仍在进行...</h3>
      <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
      <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.3/umd/popper.min.js" integrity="sha384-ZMP7rVo3mIykV+2+9J3UJ46jBk0WLaUAdn689aCwoqbBJiSnjAK/l8WvCWPIPm49" crossorigin="anonymous"></script>
      <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/js/bootstrap.min.js" integrity="sha384-ChfqqxuZUCnJSK3+MXmPNIyE6ZbWh2IMqE241rYiqJxyMiZ6OW/JmZQ5stwEULTy" crossorigin="anonymous"></script>
      <script src="https://cdn.bootcss.com/jquery.form/4.2.2/jquery.form.min.js"></script>
      <script>
      var search = window.location.search;
      var UserName = getSearchString('UserName', search); //结果：18
      document.getElementById("fileform").action = "/upload?UserName="+UserName;
      function getSearchString(key, Url) {
          var str = Url;
          str = str.substring(1, str.length); // 获取URL中?之后的字符（去掉第一位的问号）
          // 以&分隔字符串，获得类似name=xiaoli这样的元素数组
          var arr = str.split("&");
          var obj = new Object();
      
          // 将每一个数组元素以=分隔并赋给obj对象 
          for (var i = 0; i < arr.length; i++) {
              var tmp_arr = arr[i].split("=");
              obj[decodeURIComponent(tmp_arr[0])] = decodeURIComponent(tmp_arr[1]);
          }
          return obj[key];
      }
      </script>
    </body>
    </html>`)
}

func main() {
    s := g.Server()
    s.BindHandler("/upload",      Upload)
    s.BindHandler("/upload/show", UploadShow)
    s.SetPort(8199)
    s.Run()
}