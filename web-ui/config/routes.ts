export default [
  {
    path: '/user',
    layout: false,
    routes: [
      {
        path: '/user',
        routes: [
          {
            name: 'login',
            path: '/user/login',
            component: './User/login',
          },
        ],
      },
    ],
  },
  {
    name: 'dashboard.analysis',
    icon: 'smile',
    path: '/dashboard/analysis',
    component: './Dashboard/Analysis',
  },
  {
    path: '/mesh',
    name: 'mesh',
    icon: 'crown',
    routes: [
      {
        path: '/mesh/networking',
        name: 'networking',
        routes: [
          {
            path: '/mesh/networking/gateway',
            name: 'gateway',
            component: './Mesh/Gateway',
          },
          {
            path: '/mesh/networking/virtual-service',
            name: 'virtual-service',
            component: './Mesh/VirtualService',
          },
          {
            path: '/mesh/networking/destination-rule',
            name: 'destination-rule',
            component: './Mesh/DestinationRule',
          },
          {
            path: '/mesh/networking/envoy-filter',
            name: 'envoy-filter',
            component: './Mesh/EnvoyFilter',
          },
          {
            path: '/mesh/networking/sidecar',
            name: 'sidecar',
            icon: 'smile',
            component: './Mesh/Sidecar',
          },
          {
            path: '/mesh/networking/service-entry',
            name: 'service-entry',
            component: './Mesh/ServiceEntry',
          },
        ],
      },
      {
        path: '/mesh/security',
        name: 'security',
        icon: 'crown',
        routes: [
          {
            path: '/mesh/security/request-authentication',
            name: 'request-authentication',
            component: './Mesh/RequestAuthentication',
          },
          {
            path: '/mesh/security/peer-authentication',
            name: 'peer-authentication',
            component: './Mesh/PeerAuthentication',
          },
          {
            path: '/mesh/security/policy',
            name: 'policy',
            component: './Mesh/Policy',
          },
          {
            path: '/mesh/security/authorizatio-policy',
            name: 'authorizatio-policy',
            component: './Mesh/AuthorizatioPolic',
          },
        ],
      },
    ],
  },
  {
    name: '标准列表',
    icon: 'smile',
    path: '/listbasiclist',
    component: './ListBasicList',
  },
  {
    name: 'list.table-list',
    icon: 'table',
    path: '/list',
    component: './TableList',
  },
  {
    path: '/',
    redirect: '/dashboard/analysis',
  },
  {
    component: './404',
  },
];
