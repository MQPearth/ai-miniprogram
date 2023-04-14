import {
    request
} from './request'

module.exports = {
    save: (data) => request('/feedback/save', 'POST', data),
    view: () => request('/feedback/view', 'GET'),
}