//http://www.axios-js.com/zh-cn/docs/
import axios from 'axios';
import storageService from '../service/storageService.js';
const service = axios.create({
    // baseURL: 'http://localhost:1016/api/',
    baseURL: process.env.VUE_APP_BASE_URL,
    timeout: 1000 * 5, //5秒
});
//config/dev.env.js
console.log(process.env.VUE_APP_BASE_URL)
service.interceptors.request.use(function (config) {
    // Do something before request is sent
    Object.assign(config.headers, { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` })
    return config;
}, function (error) {
    // Do something with request error
    return Promise.reject(error);
});
export default service;