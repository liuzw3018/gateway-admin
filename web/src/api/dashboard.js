import request from '@/utils/request'

export function flowStat(data) {
  return request({
    url: '/api/dashboard/flowStat',
    method: 'get',
    params: data
  })
}

export function panelGroupData(data) {
  return request({
    url: '/api/dashboard/panelGroupData',
    method: 'get',
    params: data
  })
}

export function serviceStat(data) {
  return request({
    url: '/api/dashboard/serviceStat',
    method: 'get',
    params: data
  })
}
