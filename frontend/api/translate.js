import {
    request
} from './request'

module.exports = {
    translate: (data) => request('/translate', 'POST', data),
}