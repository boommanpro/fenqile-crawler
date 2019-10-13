# 分期乐乐疯抢爬虫

## 最终效果

每日定时(如早上9点)收取乐疯抢的抢购信息(10点,16点,22点)
收取方式为Server酱的微信公众号
内容是图片


## why make this project

这个乐疯抢时而有些比较划算的电子产品,遂编写爬虫方便查看。

## theroy

乐疯抢手机地址:https://hui.m.fenqile.com/

使用selenium访问目标地址,然后进行屏幕截图

使用腾讯云的cos进行存储

然后通过Server酱进行推送


## To do List

1.Server酱推送内容编写  Markdown格式

2.golang定时任务编写

3.相关配置抽取

4.相关日志操作统计编写

5.腾讯cos代码优化

6.selenium爬取代码优化

7.项目坑点总结。
