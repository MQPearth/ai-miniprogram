// pages/naming.js

const util = require('../../utils/util.js')
const $api = require('../../api/translate.js')

Page({
    data: {
        result: '',
        content: '',
        list: ['中文', '英语', '日语', '俄语', '韩语', '希腊语', '斯洛文尼亚语', '捷克语', '波兰语', '越南语', '葡萄牙语', '德语', '泰语', '瑞典语', '芬兰语', '保加利亚语', '法语', '意大利语', '阿拉伯语', '匈牙利语', '罗马尼亚语', '丹麦语', '爱沙尼亚语'],
        mainSelected: 1
    },
    bindPickerChange: function (e) {
        this.setData({
            mainSelected: e.detail.value
        })
    },
    naming() {
        let token = wx.getStorageSync('token')
        if (!token) {
            wx.showToast({
                title: '即将授权登录',
                icon: 'loading',
                duration: 2000
            })
            util.login()
            // 刷新当前页
            this.onLoad()
            return
        }
        if (!this.data.content) {
            wx.showToast({
                title: '请填写翻译内容',
                icon: 'none'
            })
            return
        }
        wx.showLoading({
            title: "翻译中",
            mask: true
        });
        $api.translate({
            content: this.data.content,
            target: this.data.list[this.data.mainSelected]
        }).then(res => {
            wx.hideLoading()
            while (res.startsWith('\n')) {
                res = res.substr(1, res.length - 1)
            }
            this.setData({
                result: res
            })
        })

    },


    bindConetent: function (inputEle) {
        this.setData({
            content: inputEle.detail.value
        })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {

    },

    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady: function () {

    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow: function () {

    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide: function () {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload: function () {

    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh: function () {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom: function () {

    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage: function () {

    }
})