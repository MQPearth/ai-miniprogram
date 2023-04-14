// index.js
// 获取应用实例
const app = getApp()

const $api = require('../../api/user')

Page({
    data: {
        nickName: '',
        avatarUrl: ''
    },

    onLoad() {

    },
    getUserProfile(e) {
        wx.getUserProfile({
            desc: '展示用户信息',
            success: (res) => {
                console.log(res)
                let userRes = res
                wx.login({
                    desc: '展示用户信息',
                    success: (res) => {
                        console.log(res)
                        $api.login({
                            "appid": "wxcad3cb8297a1390c",
                            "code": res.code,
                            "nickName": userRes.userInfo.nickName,
                            "avatarUrl": userRes.userInfo.avatarUrl,
                        }).then(ret => {
                            console.log(ret)
                            this.setData({
                                nickName: ret.nickName,
                                avatarUrl: ret.avatarUrl
                            })
                        })
                    }
                })
            }
        })
    },
})