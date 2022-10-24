import request from '@/utils/request'

export function appList(data) {
  return request({
    url: '/api/app/list',
    method: 'get',
    params: data
  })
}

export function appCreate(data) {
  return request({
    url: '/api/app/add',
    method: 'post',
    params: data
  })
}

export function appUpdate(data) {
  return request({
    url: '/api/app/update',
    method: 'post',
    params: data
  })
}

export function appDetail(data) {
  return request({
    url: '/api/app/detail',
    method: 'get',
    params: data
  })
}

export function appDelete(data) {
  return request({
    url: '/api/app/delete',
    method: 'get',
    params: data
  })
}

export function appStat(data) {
  return request({
    url: '/api/app/stat',
    method: 'get',
    params: data
  })
}
