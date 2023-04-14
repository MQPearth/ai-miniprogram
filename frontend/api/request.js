import {
    baseUrl
} from './host.js'

import {
    login
} from '../utils/util.js'

module.exports = {
    request: function (url, methodType, data) {
        let fullUrl = `${baseUrl}${url}`
        let token = wx.getStorageSync('token') ? wx.getStorageSync('token') : ''
        return new Promise((resolve, reject) => {
            wx.request({
                url: fullUrl,
                method: methodType,
                timeout: 120000,
                data,
                header: {
                    'content-type': 'application/json',
                    'token': token,
                },
                success(res) {
                    if (res.data.code == 0) {
                        resolve(res.data.data)
                        wx.hideLoading()
                    } else if (res.data.code == 2) {
                        login()
                        reject(res.data.message)
                    } else {
                        wx.hideLoading()
                        wx.showToast({
                            title: res.data.msg,
                            icon: 'none'
                        })
                        reject(res.data.message)
                    }
                },
                fail() {
                    wx.showToast({
                        title: '接口请求错误',
                        icon: 'none'
                    })
                    reject('接口请求错误')
                }
            })
        })
    }
}