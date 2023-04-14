// pages/naming.js

const util = require('../../utils/util.js')
const $api = require('../../api/ad.js')

Page({
    data: {
        result: '',
        content: ''
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
                title: '请填写产品内容',
                icon: 'none'
            })
            return
        }

        if (this.data.content.length > 50) {
            wx.showToast({
                title: '产品内容不能大于50个字符',
                icon: 'none'
            })
            return
        }
        wx.showLoading({
            title: "生成中",
            mask: true
        });
        $api.ad({
            content: this.data.content,
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