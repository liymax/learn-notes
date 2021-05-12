http://clab-mtc.huawei.com/#!/devices

(() => {
  //屏蔽切屏限制
  window.onblur=null
  window.onblur=function(){console.debug(1);}

  //解除快捷键操作屏蔽
  window.onkeyup = window.onkeydown = window.onKeyPress = document.onkeyup = document.onkeydown = document.onKeyPress = document.body.onkeyup = document.body.onkeydown = document.body.onKeyPress = onkeyup = onkeydown = onKeyPress = null;

  //解除复制粘贴限制
  window.oncopy = window.onpaste = document.oncopy = document.onpaste = document.body.oncopy = document.body.onpaste = oncopy = onpaste = null;
})()

git config --global http.sslVerify false

code font: DejaVu/Droid Sans Mono, Cascadia Code

git config --global alias.st status
git config --global alias.ca "commit -a -m"

git config --global alias.isse "pull isse dev"

git config --global alias.dev "pull origin dev"

wsl --shutdown

C:\MongoDB\Server\3.4\bin\mongod.exe --config D:\Tools\mongodb\mongo.config --install --serviceName "MongoDB"

net start mongodb

db.createUser({user:'test',pwd:'tiger',roles:[{ role: "root", db: "admin" }]})

db.createUser({user:"test1",pwd:"test1",roles:[{role: 'readWrite', db: 'yapi'},{role: 'dbOwner', db: 'yapi'}]})

db.createUser({user:"test2",pwd:"test2",roles:[{role: 'readWrite', db: 'test'},{role: 'dbOwner', db: 'test'}]})



https://skidrowcodex.co/878-gamepc-last-epoch.html
