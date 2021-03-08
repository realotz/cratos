import React, {useEffect, useState} from 'react';
import ProForm, {
  ModalForm,
  ProFormText,
  ProFormSelect, ProFormList,
} from '@ant-design/pro-form';
import {ListTags} from "@/services/namespace_service.pb";
import {Space} from "antd";
// @ts-ignore
import ProCard from '@ant-design/pro-card';
// @ts-ignore
import {UnControlled as CodeMirror} from 'react-codemirror2'
import yaml from "js-yaml"
import 'codemirror/lib/codemirror.css';
import 'codemirror/lib/codemirror.js'
import 'codemirror/mode/yaml/yaml';
import 'codemirror/theme/material.css';

//ctrl+空格代码提示补全
import 'codemirror/addon/hint/show-hint.css';
import 'codemirror/addon/hint/show-hint';
import 'codemirror/addon/hint/anyword-hint.js';
//代码高亮
import 'codemirror/addon/selection/active-line';
//折叠代码
import 'codemirror/addon/fold/foldgutter.css';
import 'codemirror/addon/fold/foldcode.js';
import 'codemirror/addon/fold/foldgutter.js';
import 'codemirror/addon/fold/brace-fold.js';
import 'codemirror/addon/fold/comment-fold.js';
import 'codemirror/addon/edit/closebrackets';

export type FormValueType = {
  target?: string;
  template?: string;
  type?: string;
  time?: string;
  frequency?: string;
} & CratosApiV1Gateway.Gateway;

// form props 类型定义
export type UpdateFormProps = {
  onCancel: (flag?: boolean, formVals?: FormValueType) => void;
  onSubmit: (values: FormValueType) => Promise<boolean | void>;
  updateModalVisible: boolean;
  values: CratosApiV1Gateway.Gateway|undefined;
};

const UpdateForm: React.FC<UpdateFormProps> = (props) => {
  const [yamlEdit, handleYamlEdit] = useState<boolean>(false);
  const [values, handleValues] = useState<CratosApiV1Gateway.Gateway>();
  useEffect(() => {
    if(props.values && props.values.kind){
      handleValues(props.values)
    }else{
      handleValues({
        kind:"Gateway",
        apiVersion:"networking.istio.io/v1alpha3",
        metadata:{
          name:"",
          namespace:""
        },
        spec:{
          servers:[
            {
              hosts:[""],
              port:{
                name:"",
                number:80,
                protocol:"",
              }
            }
          ]
        }
      })
    }
  }, [props.updateModalVisible]);
  // 表单编辑
  const FormEdit = (
    <>
      {/*基础信息*/}
      <ProForm.Group>
        <ProFormSelect
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
          width="md"
          name="namespace"
          label="命名空间"
          rules={[
            {
              required: true,
              message: '请选择命名空间!',
            },
          ]}
        />
        <ProFormText width="md" name="name" label="网关名称" placeholder="请输入网关名称" bordered
                     rules={[
                       {
                         required: true,
                         message: '请输入网关名称!',
                       },
                     ]}
        />
      </ProForm.Group>

      {/*selector选择器*/}
      <ProFormList
        name="selector"
        initialValue={[
          {
            useMode: 'chapter',
          },
        ]}
        creatorButtonProps={{
          creatorButtonText: '增加选择器',
          type: "primary"
        }}
        creatorRecord={{
          useMode: 'none',
        }}
        label={"selector选择器"}
      >
        <Space style={{display: 'flex', width: "100%"}} align="baseline">
          <ProFormText width={340} name="key" placeholder="key"
                       rules={[
                         {
                           required: true,
                           message: '至少添加一个选择器!',
                         },
                       ]}
          /> = <ProFormText width={340} name="value" placeholder="value" rules={[
          {
            required: true,
            message: '至少添加一个选择器!',
          },
        ]}/>
        </Space>
      </ProFormList>

      {/*服务器配置*/}
      <ProFormList
        name="server"
        initialValue={[{useMode: 'chapter'}]}
        creatorButtonProps={{
          creatorButtonText: '增加server',
          type: "primary"
        }}
        creatorRecord={{
          useMode: 'none',
        }}
        label={"服务器配置"}
        itemRender={({listDom, action}, {}) => {
          return (
            <ProCard
              bordered
              extra={action}
              style={{width: 750}}
              size="small"
            >
              {listDom}
            </ProCard>
          );
        }}
      >
        <Space style={{display: 'flex', width: "100%"}} align="baseline">
          <ProFormText name="port" width={250} label={"端口号："} placeholder="端口号" rules={[
            {
              required: true,
              message: '请输入端口号!',
            }]}/>
          <ProFormText name="portName" width={250} label={"端口名称："} placeholder="端口名称" rules={[
            {
              required: true,
              message: '请输入端口名称!',
            }]}/>
          <ProFormSelect name="端口协议" width={180} label={"端口协议："} options={[
            {label: "HTTP", value: "HTTP"},
            {label: "HTTPS", value: "HTTPS"},
            {label: "HTTP2", value: "HTTP2"},
            {label: "GRPC", value: "GRPC"},
            {label: "MONGO", value: "MONGO"},
            {label: "TCP", value: "TCP"},
            {label: "TLS", value: "TLS"},
          ]} rules={[
            {
              required: true,
              message: '请选择端口协议!',
            }]}/>
        </Space>
        <ProFormList name="hosts" label="hosts" creatorButtonProps={{
          creatorButtonText: '增加host',
        }}>
          <Space style={{display: 'flex', marginBottom: 8, width: "100%"}} align="baseline">
            <ProFormText width={680} placeholder="输入host地址 例如：*" rules={[
              {
                required: true,
                message: '请输入host地址!',
              }]}/>
          </Space>
        </ProFormList>
      </ProFormList>

      {/*注释与标签*/}
      <Space style={{display: 'flex', width: "100%"}} align="baseline">
        <ProFormList
          name="tabs"
          creatorButtonProps={{
            creatorButtonText: '增加标签',
            type: "primary"
          }}
          creatorRecord={{
            useMode: 'none',
          }}
          label={"标签"}
        >
          <Space style={{display: 'flex', width: "100%"}} align="baseline">
            <ProFormText width={150} name="key" placeholder="key"/> = <ProFormText width={150} name="value"
                                                                                   placeholder="value"/>
          </Space>
        </ProFormList>
        <ProFormList
          name="test"
          creatorButtonProps={{
            creatorButtonText: '增加注释',
            type: "primary"
          }}
          creatorRecord={{
            useMode: 'none',
          }}
          label={"注释"}
        >
          <Space style={{display: 'flex', width: "100%"}} align="baseline">
            <ProFormText width={150} name="key" placeholder="key"/> = <ProFormText width={150} name="value"
                                                                                   placeholder="value"/>
          </Space>
        </ProFormList>
      </Space>
    </>
  )

  return (
    <ModalForm<FormValueType>
      title={(<><span>新建网关 </span>{(yamlEdit ? (<a onClick={() => handleYamlEdit(false)}>YAML</a>) : (
        <a onClick={() => handleYamlEdit(true)}>表单</a>))}</>)}
      visible={props.updateModalVisible}
      modalProps={{
        onCancel: () => props.onCancel(),
      }}
      onFinish={props.onSubmit}
    >
      {yamlEdit ? (
        <CodeMirror
          value={yaml.dump(values)}
          options={{
            autoCloseBrackets: true, //键入时将自动关闭()[]{}''""
            lineNumbers: true, // 显示行号
            theme: 'material', // 设置主题
            mode: {
              name: 'text/x-yaml',
            },

            // (以下三行)设置支持代码折叠
            lineWrapping:true,
            foldGutter:true,
            gutters:['CodeMirror-linenumbers','CodeMirror-foldgutter'],
          }}
          onChange={(editor, data, value) => {
          }}
        />
      ) : FormEdit}
    </ModalForm>
  );
};

export default UpdateForm;
