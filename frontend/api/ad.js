import {
    request
} from './request'

module.exports = {
    ad: (data) => request('/ad', 'POST', data),
}