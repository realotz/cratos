import {Button, message} from 'antd';
import React from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import ProTable, { ActionType} from '@ant-design/pro-table'
import type { ProColumns } from '@ant-design/pro-table';
import {List,Update} from '@/services/gateway_service.pb'
import UpdateForm from './components/UpdateForm';
import { useState } from 'react';
import { useRef } from 'react';
import {ListTags} from "@/services/namespace_service.pb";
import {ProFormSelect} from "@ant-design/pro-form";


/**
 * 更新节点
 *
 * @param fields
 */
const handleUpdate = async (fields: CratosApiV1Gateway.Gateway) => {
  const hide = message.loading('正在配置');
  try {
    await Update(fields);
    hide();

    message.success('配置成功');
    return true;
  } catch (error) {
    hide();
    message.error('配置失败请重试！');
    return false;
  }
};

const GatewayPage: React.FC = () => {
  /** 新建窗口的弹窗 */
  /** 分布更新窗口的弹窗 */
  const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);

  const actionRef = useRef<ActionType>();
  const [currentRow, setCurrentRow] = useState<CratosApiV1Gateway.Gateway>();

  const columns: ProColumns<CratosApiV1Gateway.Gateway>[] = [
    {
      title: '应用名称',
      dataIndex: 'metadata.name',
      render: (_: React.ReactNode,row) => <a>{row.metadata?.name}</a>,
    },
    {
      title: '命名空间',
      dataIndex: 'metadata.namespace',
      render: (_: React.ReactNode,row) => <a>{row.metadata?.namespace}</a>,
      renderFormItem:(item)=>{
        return <ProFormSelect
          showSearch
          request={async (params: any, props: any) => {
            const res = await ListTags({name: params.keyWords, limit: 10})
            var listTags = []
            if (res.name) {
              for (let i = 0; i < res?.name.length; i++) {
                listTags.push({
                  label: res.name[i],
                  value: res.name[i],
                  key: res.name[i],
                })
              }
            }
            return Promise.resolve(listTags)
          }}
          label={false}
          name="namespace"
        />
      }
    },
    {
      title: '主机名',
      render: (_: React.ReactNode,row) => <div>{row.spec?.servers?.map(function (params,index) {
        return <span key={index}>{params.hosts?.toLocaleString()}</span>
      })}</div>,
    },
    {
      title: '端口',
      render: (_: React.ReactNode,row) => <div>{row.spec?.servers?.map(function (params,index) {
        return <span key={index}>{params.port?.protocol} :{params.port?.number}</span>
      })}</div>,
    },
    {
      title: "创建时间",
      dataIndex: 'metadata.creationTimestamp',
      valueType: 'date',
      search:false,
      render: (_: React.ReactNode,row) => <span>{row.metadata?.creationTimestamp}</span>,
    },
    {
      title: '操作',
      width: '164px',
      key: 'option',
      valueType: 'option',
      render: (_: React.ReactNode,row) => [
        <a key="link" onClick={()=>{setCurrentRow(row);handleUpdateModalVisible(true)}}>编辑</a>,
        <a key="link2">复制</a>,
        <a key="link3">删除</a>,
      ],
    },
  ];
  return (
    <PageContainer>
      <ProTable<CratosApiV1Gateway.Gateway>
        columns={columns}
        request={async (params, sorter, filter) => {
          // 表单搜索项会从 params 传入，传递给后端接口。
          let offset = 0
          let pageSize = params?.pageSize?params?.pageSize:10
          if (params?.current && params?.current>=1){
            offset = (params?.current-1)*pageSize
          }
          const res = await List({
            limit:params.pageSize,
            Offset:offset
          })
          return Promise.resolve({
            data: res.list,
            success: res.total,
          });
        }}
        rowKey={(record:any) => {
          return record.metadata.name
        }}
        search={{
          filterType: 'light',
        }}
        dateFormatter="string"
        headerTitle="网关列表"
        toolBarRender={() => [
          <Button onClick={() => {handleUpdateModalVisible(true);}} type="primary" key="primary">
            创建网关
          </Button>,
        ]}
      />
      <UpdateForm
        onSubmit={async (value) => {
          const success = await handleUpdate(value);
          if (success) {
            handleUpdateModalVisible(false);
            setCurrentRow(undefined);
            if (actionRef.current) {
              actionRef.current.reload();
            }
          }
        }}
        onCancel={() => {
          handleUpdateModalVisible(false);
          setCurrentRow(undefined);
        }}
        updateModalVisible={updateModalVisible}
        values={currentRow}
      />
    </PageContainer>
  );
};

export default GatewayPage;
