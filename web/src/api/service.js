import request from '@/utils/request'

export function serviceList(data) {
  return request({
    url: '/api/service/list',
    method: 'get',
    params: data
  })
}

export function serviceDelete(data) {
  return request({
    url: '/api/service/delete',
    method: 'get',
    params: data
  })
}

export function addHttpService(data) {
  return request({
    url: '/api/service/http/add',
    method: 'post',
    params: data
  })
}

export function updateHttpService(data) {
  return request({
    url: '/api/service/http/update',
    method: 'post',
    params: data
  })
}

export function getServiceDetail(data) {
  return request({
    url: '/api/service/detail',
    method: 'get',
    params: data
  })
}

export function addTCPService(data) {
  return request({
    url: '/api/service/tcp/add',
    method: 'post',
    params: data
  })
}

export function updateTCPService(data) {
  return request({
    url: '/api/service/tcp/update',
    method: 'post',
    params: data
  })
}

export function addGRPCService(data) {
  return request({
    url: '/api/service/grpc/add',
    method: 'post',
    params: data
  })
}

export function updateGRPCService(data) {
  return request({
    url: '/api/service/grpc/update',
    method: 'post',
    params: data
  })
}

export function serviceStat(data) {
  return request({
    url: '/api/service/stat',
    method: 'get',
    params: data
  })
}
