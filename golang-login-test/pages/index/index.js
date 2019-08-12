//index.js
//获取应用实例
const app = getApp();

Page({
  data: {
    motto: "Hello World",
    userInfo: {},
    hasUserInfo: false,
    canIUse: wx.canIUse("button.open-type.getUserInfo")
  },
  //事件处理函数
  bindViewTap: function() {
    wx.navigateTo({
      url: "../logs/logs"
    });
  },
  wxLogin: function() {
    console.log("wx login ...");
    wx.login({
      success: res => {
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
        console.log(res);
        wx.request({
          url: "http://localhost:8080/login/wx",
          success: res => console.log(res),
          data: {
            code: res.code
          }
        });
      }
    });
  },
  qywxLogin: function() {
    console.log("qywx login ...");
    wx.qy.login({
      success: res => {
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
        console.log(res);
        wx.request({
          url: "http://localhost:8080/login/qywx",
          success: res => console.log(res),
          data: {
            code: res.code
          }
        });
      }
    });
  },

  qywxtpLogin: function () {
    console.log("qywxtp login ...");
    wx.qy.login({
      success: res => {
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
        console.log(res);
        wx.request({
          url: "http://localhost:8080/login/qywxtp",
          success: res => console.log(res),
          data: {
            code: res.code
          }
        });
      }
    });
  },

  
});
