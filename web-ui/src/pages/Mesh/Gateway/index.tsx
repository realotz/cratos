import {Button, Tooltip} from 'antd';
import React from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { QuestionCircleOutlined } from '@ant-design/icons';
import ProTable, { TableDropdown } from '@ant-design/pro-table'
import type { ProColumns } from '@ant-design/pro-table';
import {GetGatewayList} from '@/services/mesh_service.pb'


const GatewayPage: React.FC = () => {

  const columns: ProColumns<CratosApiV1.Gateway>[] = [
    {
      title: '应用名称',
      dataIndex: 'metadata.name',
      render: (_: React.ReactNode,row) => <a>{row.metadata?.name}</a>,
    },
    {
      title: '命名空间',
      dataIndex: 'metadata.namespace',
      render: (_: React.ReactNode,row) => <a>{row.metadata?.namespace}</a>,
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
      render: (_: React.ReactNode,row) => <span>{row.metadata?.creationTimestamp}</span>,
    },
    {
      title: '操作',
      width: '164px',
      key: 'option',
      valueType: 'option',
      render: () => [
        <a key="link">链路</a>,
        <a key="link2">报警</a>,
        <a key="link3">监控</a>,
        <TableDropdown
          key="actionGroup"
          menus={[
            { key: 'copy', name: '复制' },
            { key: 'delete', name: '删除' },
          ]}
        />,
      ],
    },
  ];
  return (
    <PageContainer>
      <ProTable<CratosApiV1.Gateway>
        columns={columns}
        request={async (params, sorter, filter) => {
          // 表单搜索项会从 params 传入，传递给后端接口。
          console.log(params, sorter, filter);
          const res = await GetGatewayList({
            limit:params.pageSize
          })
          return Promise.resolve({
            data: res.list,
            success: true,
          });
        }}
        rowKey={(record) => {
          return record.metadata.name
        }}
        pagination={{
          showQuickJumper: true,
        }}
        search={{
          filterType: 'light',
        }}
        dateFormatter="string"
        headerTitle="网关列表"
        toolBarRender={() => [
          <Button key="show">查看日志</Button>,
          <Button type="primary" key="primary">
            创建应用
          </Button>,
        ]}
      />
    </PageContainer>
  );
};

export default GatewayPage;
