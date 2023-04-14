const $api = require('../api/user')

export function login() {
    wx.showLoading({
        title: "发起登录",
        mask: true
    });
    wx.getUserProfile({
        desc: '登录授权',
        success: (res) => {
            let userRes = res
            wx.login({
                desc: '展示用户信息',
                success: async (res) => {
                    let data = await $api.login({
                        "appid": "xxxx",
                        "code": res.code,
                        "nickName": userRes.userInfo.nickName,
                        "avatarUrl": userRes.userInfo.avatarUrl,
                    })
                    wx.setStorageSync('token', data.token)
                    wx.showToast({
                        title: '登录成功',
                        icon: 'success'
                    })
                }
            })
        },
        fail: (res) => {
            wx.hideLoading()
        }
    })
}