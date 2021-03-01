import React from 'react';
import { GithubOutlined } from '@ant-design/icons';
import { DefaultFooter } from '@ant-design/pro-layout';

export default () => (
  <DefaultFooter
    copyright="2020 cratos ui for antd"
    links={[
      {
        key: 'antd',
        title: 'antd',
        href: 'https://ant.design',
        blankTarget: true,
      },
      {
        key: 'github',
        title: <GithubOutlined />,
        href: 'https://github.com/realotz/cratos',
        blankTarget: true,
      },
    ]}
  />
);
