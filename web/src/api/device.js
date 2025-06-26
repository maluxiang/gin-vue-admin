import request from '@/utils/request'; // 假设项目中已封装好的请求工具，`gin-vue-admin` 一般有此封装

// 获取设备列表
export function getDeviceList(params) {
    return request({
        url: '/device/list', // 后端接口地址，需与 gin 后端路由对应
        method: 'get',
        params,
    });
}

// 获取设备告警列表
export function getDeviceAlarmList(params) {
    return request({
        url: '/device/alarm/list',
        method: 'get',
        params,
    });
}