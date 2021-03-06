// @ts-ignore
/* eslint-disable */
// Code generated by protoc-gen-ts-umi. DO NOT EDIT.
import {request} from 'umi';

const APIService = '/api';
// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.

type Options = {
  [key: string]: any
}

/** List  gateway列表 /api */
export async function List(params: CratosApiV1Gateway.ListOption, options?: Options) {
	return request<CratosApiV1Gateway.GatewayList>(APIService + '/cratos.api.v1.Gateway/List', {
    	method: 'GET',
    	params: {...params},
    	...(options || {}),
	});
}

/** Get  获取gateway /api */
export async function Get(params: CratosApiV1Gateway.GetKind, options?: Options) {
	return request<CratosApiV1Gateway.Gateway>(APIService + '/cratos.api.v1.Gateway/Get', {
    	method: 'GET',
    	params: {...params},
    	...(options || {}),
	});
}

/** Create  创建 Gateway /api */
export async function Create(params: CratosApiV1Gateway.Gateway, options?: Options) {
	return request<CratosApiV1Gateway.Response>(APIService + '/cratos.api.v1.Gateway/Create', {
    	method: 'POST',
		headers: {'Content-Type': 'application/json',},
    	data: {...params},
    	...(options || {}),
	});
}

/** Update  更新 Gateway /api */
export async function Update(params: CratosApiV1Gateway.Gateway, options?: Options) {
	return request<CratosApiV1Gateway.Response>(APIService + '/cratos.api.v1.Gateway/Update', {
    	method: 'POST',
		headers: {'Content-Type': 'application/json',},
    	data: {...params},
    	...(options || {}),
	});
}

/** Delete  删除gateway /api */
export async function Delete(params: CratosApiV1Gateway.DeleteKind, options?: Options) {
	return request<CratosApiV1Gateway.Response>(APIService + '/cratos.api.v1.Gateway/Delete', {
    	method: 'DELETE',
    	params: {...params},
    	...(options || {}),
	});
}

