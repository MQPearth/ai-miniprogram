import {
    request
} from './request'

module.exports = {
    getNameList: (data) => request('/naming/getNameList', 'GET', data),
}