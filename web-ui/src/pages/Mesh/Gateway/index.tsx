import { Button, Result } from 'antd';
import React from 'react';
import { history } from 'umi';

const GatewayPage: React.FC = () => (
  <Result
    status="404"
    title="200"
    subTitle="devops"
    extra={
      <Button type="primary" onClick={() => history.push('/')}>
        Back Home
      </Button>
    }
  />
);

export default GatewayPage;
