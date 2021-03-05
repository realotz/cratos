import {Button, message, Tooltip} from 'antd';
import React from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { QuestionCircleOutlined } from '@ant-design/icons';
import ProTable, { ActionType, TableDropdown } from '@ant-design/pro-table'
import type { ProColumns } from '@ant-design/pro-table';
import {GetGatewayList,UpdateGateway} from '@/services/mesh_service.pb'
import UpdateForm from './components/UpdateForm';
import { useState } from 'react';
import { useRef } from 'react';


/**
 * 更新节点
 *
 * @param fields
 */
const handleUpdate = async (fields: CratosApiV1.Gateway) => {
  const hide = message.loading('正在配置');
  try {
    await UpdateGateway(fields);
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
  const [createModalVisible, handleModalVisible] = useState<boolean>(false);
  /** 分布更新窗口的弹窗 */
  const [updateModalVisible, handleUpdateModalVisible] = useState<boolean>(false);

  const [showDetail, setShowDetail] = useState<boolean>(false);

  const actionRef = useRef<ActionType>();
  const [currentRow, setCurrentRow] = useState<API.RuleListItem>();
  const [selectedRowsState, setSelectedRows] = useState<API.RuleListItem[]>([]);
  
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
      search:false,
      render: (_: React.ReactNode,row) => <span>{row.metadata?.creationTimestamp}</span>,
    },
    {
      title: '操作',
      width: '164px',
      key: 'option',
      valueType: 'option',
      render: () => [
        <a key="link">编辑</a>,
        <a key="link2">复制</a>,
        <a key="link3">删除</a>,
        // <TableDropdown
        //   key="actionGroup"
        //   menus={[
        //     { key: 'copy', name: '复制' },
        //     { key: 'delete', name: '删除' },
        //   ]}
        // />,
      ],
    },
  ];
  return (
    <PageContainer>
      <ProTable<CratosApiV1.Gateway>
        columns={columns}
        request={async (params, sorter, filter) => {
          // 表单搜索项会从 params 传入，传递给后端接口。
          let offset = 0
          let pageSize = params?.pageSize?params?.pageSize:10
          if (params?.current && params?.current>=1){
            offset = (params?.current-1)*pageSize
          }
          const res = await GetGatewayList({
            limit:params.pageSize,
            Offset:offset
          })
          return Promise.resolve({
            data: res.list,
            success: res.total,
          });
        }}
        rowKey={(record) => {
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
        values={currentRow || {}}
      />
    </PageContainer>
  );
};

export default GatewayPage;
