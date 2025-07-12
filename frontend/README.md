# Helloworld 前端项目

基于 Vue 3 + TypeScript + Vite + Element Plus 构建的现代化前端管理系统。

## 技术栈

- **框架**: Vue 3 (Composition API)
- **语言**: TypeScript
- **构建工具**: Vite
- **UI框架**: Element Plus
- **状态管理**: Pinia
- **路由**: Vue Router 4
- **HTTP客户端**: Axios
- **图表**: ECharts
- **代码规范**: ESLint + Prettier

## 项目结构

```
frontend/
├── public/                 # 静态资源
├── src/
│   ├── api/               # API接口
│   │   ├── request.ts     # HTTP请求封装
│   │   ├── user.ts        # 用户相关API
│   │   ├── order.ts       # 订单相关API
│   │   └── wechat.ts      # 微信相关API
│   ├── assets/            # 资源文件
│   ├── components/        # 通用组件
│   ├── router/            # 路由配置
│   │   └── index.ts
│   ├── store/             # 状态管理
│   │   ├── index.ts       # 导出文件
│   │   ├── user.ts        # 用户状态
│   │   ├── theme.ts       # 主题状态
│   │   └── order.ts       # 订单状态
│   ├── styles/            # 样式文件
│   │   └── index.scss     # 全局样式
│   ├── types/             # 类型定义
│   │   ├── user.ts        # 用户类型
│   │   └── order.ts       # 订单类型
│   ├── utils/             # 工具函数
│   ├── views/             # 页面组件
│   │   ├── auth/          # 认证页面
│   │   │   └── Login.vue  # 登录页
│   │   ├── dashboard/     # 仪表盘
│   │   │   └── Dashboard.vue
│   │   ├── user/          # 用户管理
│   │   ├── order/         # 订单管理
│   │   ├── wechat/        # 微信集成
│   │   ├── settings/      # 系统设置
│   │   └── error/         # 错误页面
│   │       └── NotFound.vue
│   ├── App.vue            # 根组件
│   └── main.ts            # 入口文件
├── .env                   # 环境变量
├── .env.development       # 开发环境变量
├── .env.production        # 生产环境变量
├── index.html             # HTML模板
├── package.json           # 依赖配置
├── tsconfig.json          # TypeScript配置
├── vite.config.ts         # Vite配置
└── README.md              # 项目说明
```

## 功能特性
### 🎯 核心功能
- 用户登录认证（短信验证码）
- 微信扫码登录
- 用户管理
- 订单管理
- 数据统计图表
- 系统设置

### 🎨 界面特性
- 响应式设计，支持多端适配
- 暗黑模式支持
- 国际化支持（中英文）
- 现代化UI设计
- 流畅的交互体验

### 🛠 开发特性
- TypeScript 类型安全
- 组件化开发
- 状态管理（Pinia）
- 路由守卫
- HTTP请求拦截
- 代码规范检查

## 开发指南

### 环境要求
- Node.js >= 16
- pnpm >= 7 (推荐) 或 npm >= 8

### 安装依赖
```bash
# 使用 pnpm (推荐)
pnpm install

# 或使用 npm
npm install
```

### 开发运行
```bash
# 启动开发服务器
pnpm dev

# 或
npm run dev
```

访问 http://localhost:3000

### 构建部署
```bash
# 构建生产版本
pnpm build

# 或
npm run build

# 预览构建结果
pnpm preview

# 或
npm run preview
```

### 代码检查
```bash
# ESLint检查
pnpm lint

# 或
npm run lint

# TypeScript类型检查
pnpm type-check

# 或
npm run type-check

# 代码格式化
pnpm format

# 或
npm run format
```

## API接口

### 基础配置
- 开发环境: http://localhost:8000
- 生产环境: 需要配置实际的API地址

### 接口列表
- `POST /v1/api/sendSms` - 发送短信验证码
- `POST /v1/api/login` - 用户登录
- `POST /v1/api/real` - 实名认证
- `POST /v1/api/placeOrder` - 下单
- `GET /v1/api/wechat/check` - 微信签名验证
- `GET /v1/api/wechat/login` - 生成微信登录二维码
- `GET /v1/api/wechat/callback` - 微信登录回调

## 部署说明

### Nginx配置示例
```nginx
server {
    listen 80;
    server_name your-domain.com;
    root /path/to/frontend/dist;
    index index.html;

    # 前端路由支持
    location / {
        try_files $uri $uri/ /index.html;
    }

    # API代理
    location /api/ {
        proxy_pass http://localhost:8000/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # 静态资源缓存
    location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg)$ {
        expires 1y;
        add_header Cache-Control "public, immutable";
    }
}
```

### Docker部署
```dockerfile
FROM node:18-alpine as builder

WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production

COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

## 注意事项

1. **环境变量**: 修改 `.env` 文件中的配置以适应实际环境
2. **API地址**: 确保后端API服务正常运行
3. **路由模式**: 使用 history 模式需要服务器支持
4. **浏览器兼容**: 支持现代浏览器，IE需要额外配置
5. **HTTPS**: 生产环境建议使用HTTPS

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 开启 Pull Request

## 许可证

MIT License 
