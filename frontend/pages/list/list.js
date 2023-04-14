// logs.js
const util = require('../../utils/util.js')
const $api = require('../../api/feedback')

import {
    version
} from '../../api/host.js'

Page({
    data: {
        logs: []
    },
    onLoad() {},
    onShow() {
        $api.view().then(res => {
            let localVersion = version.split('.')
            let remoteVersion = res.split('.')
            let needUpdate = false
            for (let i = 0; i < localVersion.length; i++) {
                if (localVersion[i] < remoteVersion[i]) {
                    needUpdate = true
                    break
                }
            }
            if (needUpdate) {
                if (wx.canIUse('getUpdateManager')) {
                    const updateManager = wx.getUpdateManager()
                    updateManager.onCheckForUpdate(function (res) {
                        if (res.hasUpdate) {
                            wx.showLoading();
                            updateManager.onUpdateReady(function () {
                                wx.hideLoading()
                                updateManager.applyUpdate()
                            })
                            updateManager.onUpdateFailed(function () {
                                wx.hideLoading()
                                wx.showModal({
                                    title: '已经有新版本了',
                                    content: '新版本已经上线啦，请您删除当前小程序，重新搜索打开',
                                })
                            })
                        }
                    })
                }
            }

        })

    },
    naming() {
        wx.navigateTo({
            url: '../naming/naming'
        })
    },
    translate() {
        wx.navigateTo({
            url: '../translate/translate'
        })
    },
    ad() {
        wx.navigateTo({
            url: '../ad/ad'
        })
    },
    more() {
        wx.showToast({
            title: '更多功能敬请期待',
            icon: 'none'
        })
    },

    feedback() {
        wx.showModal({
            title: '向开发者反馈',
            content: '',
            editable: true,
            showCancel: true,
            cancelText: '取消',
            confirmText: '确定',
            success: function (res) {


                if (!res.content) {
                    wx.showToast({
                        title: '请填写反馈内容',
                        icon: 'none'
                    })
                    return
                }

                if (res.confirm) {
                    $api.save({
                        content: res.content
                    }).then(res => {
                        wx.showToast({
                            title: '反馈成功'
                        })
                    })
                }
            }
        })
    },

    dev() {
        wx.showToast({
            title: '功能加急开发中, 敬请期待',
            icon: 'none'
        })
    }
})