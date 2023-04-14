// pages/naming.js

const util = require('../../utils/util.js')
const $api = require('../../api/naming')

Page({
    data: {
        result: '',
        gender: '男',
        surname: '',
        length: '3',
        qty: '5',
        list: ['正常', '诗经', '国学', '中药材'],
        selected: 0,
    },
    bindPickerChange: function (e) {
        this.setData({
            selected: e.detail.value
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
        if (!this.data.gender) {
            wx.showToast({
                title: '请填写性别',
                icon: 'none'
            })
            return
        }

        if (!this.data.surname) {
            wx.showToast({
                title: '请填写姓氏',
                icon: 'none'
            })
            return
        }

        if (!this.data.length) {
            wx.showToast({
                title: '请填写姓名长度',
                icon: 'none'
            })
            return
        }

        if (!this.data.qty) {
            wx.showToast({
                title: '请填写生成数量',
                icon: 'none'
            })
            return
        }
        let lengthValue = Number(this.data.length)
        if (lengthValue == NaN) {
            wx.showToast({
                title: '请正确填写姓名长度',
                icon: 'none'
            })
            return
        }

        if (lengthValue <= 1 || lengthValue > 4) {
            wx.showToast({
                title: '姓名长度必须大于1 小于等于4',
                icon: 'none'
            })
            return
        }

        let qtyValue = Number(this.data.qty)

        if (qtyValue == NaN) {
            wx.showToast({
                title: '请正确填写生成数量',
                icon: 'none'
            })
            return
        }

        if (qtyValue <= 0 || qtyValue > 10) {
            wx.showToast({
                title: '生成数量必须大于0 小于等于10',
                icon: 'none'
            })
            return
        }
        wx.showLoading({
            title: "生成中",
            mask: true
        });
        $api.getNameList({
            gender: this.data.gender,
            surname: this.data.surname,
            length: this.data.length,
            qty: this.data.qty,
            style: this.data.list[this.data.selected]
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

    bindGender: function (inputEle) {
        this.setData({
            gender: inputEle.detail.value
        })
    },

    bindSurname: function (inputEle) {
        this.setData({
            surname: inputEle.detail.value
        })
    },


    bindQty: function (inputEle) {
        this.setData({
            qty: inputEle.detail.value
        })
    },

    bindLength: function (inputEle) {
        this.setData({
            length: inputEle.detail.value
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