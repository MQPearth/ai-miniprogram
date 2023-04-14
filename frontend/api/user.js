import {
    request
} from './request'

module.exports = {

    login: (data) => request('/wechat/user/login', 'POST', data),
}