<template>
  <div class="didi-container">
    <!-- 首页内容 -->
    <div v-if="activeTab === 'home'">
      <!-- header浮层 -->
      <div class="didi-header-float">
        <div class="didi-header-left">
          <div class="didi-header-location">
            <span class="didi-header-city">{{ cityName }}</span>
            <span class="didi-header-time">{{ currentTime }}</span>
          </div>
        </div>
        <div class="didi-header-btns">
          <button class="didi-header-btn">乘车码</button>
          <button class="didi-header-btn">扫一扫</button>
          <button class="didi-header-btn icon-only"><img class="didi-header-icon" src="https://img.icons8.com/ios-glyphs/30/4caf50/near-me.png" /></button>
          <button class="didi-header-btn icon-only"><img class="didi-header-icon" src="https://img.icons8.com/ios-glyphs/30/ff8200/settings.png" /></button>
        </div>
      </div>
      <!-- 地图区域 -->
      <div class="didi-map-card">
        <div id="amap-container" class="didi-map"></div>
      </div>
      <!-- 优惠券Banner -->
      <div class="didi-banner-card">
        <img class="didi-banner-icon" src="https://img.icons8.com/color/48/000000/sun--v1.png" />
        <span class="didi-banner-text">毕业季包车游 领券减66元</span>
        <button class="didi-banner-btn">收下</button>
      </div>
      <!-- 行程提示与目的地输入 -->
      <div class="didi-trip-card">
        <div class="didi-trip-row">
          <span class="didi-trip-label">当前地址</span>
          <span class="didi-trip-place">{{ currentAddress }}</span>
          <span class="didi-trip-arrow">›</span>
        </div>
        <div class="didi-trip-divider"></div>
        <div class="didi-dest-row">
          <span class="didi-dot"></span>
          <span class="didi-dest-q">您想去哪儿？</span>
        </div>
        <input
          class="didi-dest-input"
          type="text"
          v-model="destination"
          placeholder="请输入目的地地址"
        />
        <div class="didi-order-types didi-order-types-small">
          <span class="didi-order-type">预约</span>
          <span class="didi-order-type">帮人叫车</span>
          <span class="didi-order-type">接送机</span>
        </div>
      </div>
      <!-- 服务宫格 -->
      <div class="didi-services-card">
        <div class="didi-service" v-for="item in services" :key="item.text">
          <img :src="item.icon" class="didi-service-icon" />
          <span class="didi-service-text">{{ item.text }}</span>
        </div>
      </div>
      <!-- 广告位 -->
      <div class="didi-ads-row">
        <div class="didi-ad-card orange">
          <div class="didi-ad-title">5折嗨惠</div>
          <div class="didi-ad-desc">最高立减15元</div>
          <button class="didi-ad-btn">立即领</button>
        </div>
        <div class="didi-ad-card blue">
          <div class="didi-ad-title">抽无门槛券</div>
          <div class="didi-ad-desc">新客专享</div>
        </div>
      </div>
    </div>
    <!-- 车主服务页面 -->
    <div v-if="activeTab === 'owner'" class="owner-container">
      <!-- 顶部tab栏 -->
      <div class="owner-header-row">
        <div class="owner-header-tabs">
          <span :class="['owner-header-tab', ownerTab === 'owner' ? 'active' : '']" @click="ownerTab = 'owner'">车主服务</span>
          <span :class="['owner-header-tab', ownerTab === 'shunfeng' ? 'active' : '']" @click="ownerTab = 'shunfeng'">顺风车</span>
          <span class="owner-reward">50元奖</span>
        </div>
        <button class="owner-center-btn">个人中心</button>
      </div>
      <!-- 车主服务内容 -->
      <div v-if="ownerTab === 'owner'">
        <div class="owner-addcar-banner">
          <button class="owner-addcar-btn">+ 添加爱车</button>
          <span class="owner-gift">得 <b>100</b> 元大礼包</span>
        </div>
        <div class="owner-grid owner-grid1">
          <div v-for="item in ownerGrid1" :key="item.text" class="owner-grid-item">
            <img :src="item.icon" class="owner-grid-icon" />
            <span class="owner-grid-text">{{ item.text }}</span>
          </div>
        </div>
        <div class="owner-grid owner-grid2">
          <div v-for="item in ownerGrid2" :key="item.text" class="owner-grid-item">
            <img :src="item.icon" class="owner-grid-icon" />
            <span class="owner-grid-text">{{ item.text }}</span>
          </div>
        </div>
        <div class="owner-cards-row">
          <div class="owner-card orange">
            <div class="owner-card-title">洗车新客专享</div>
            <div class="owner-card-desc">首单5折 最高立减15元</div>
            <button class="owner-card-btn">立即领</button>
          </div>
          <div class="owner-card blue">
            <div class="owner-card-title">养车大礼包</div>
            <div class="owner-card-desc">￥0.1 滴滴专享福利</div>
          </div>
          <div class="owner-card green">
            <div class="owner-card-title">成为滴滴司机</div>
            <div class="owner-card-desc">领最高813元新人奖</div>
          </div>
        </div>
        <div class="owner-fuel-section">
          <div class="owner-fuel-tabs">
            <span class="active">优惠加油</span>
            <span>充电站</span>
            <span>特惠洗车</span>
            <span>看车选车</span>
          </div>
          <div class="owner-fuel-info">
            <div class="owner-fuel-title">金益一站</div>
            <div class="owner-fuel-desc">浦东新区惠南镇沪南公路9735号 | 7.1km</div>
            <div class="owner-fuel-price">
              <span class="price">￥6.31</span>
              <span class="member">会员价</span>
              <span class="origin">滴滴价￥6.65</span>
              <button class="owner-fuel-btn">去加油</button>
            </div>
          </div>
        </div>
      </div>
      <!-- 顺风车内容 -->
      <div v-if="ownerTab === 'shunfeng'" class="sf-container">
        <div class="sf-header-row">
          <span class="sf-title">顺风车</span>
          <span class="sf-time">21:20</span>
        </div>
        <div class="sf-banner">
          <img class="sf-banner-img" src="/static/logo.png" alt="顺风车banner" />
          <div class="sf-banner-text">
            <span class="sf-banner-main">成为车主 最高得50元!!</span>
            <span class="sf-banner-sub">再送最高60元加油充电券</span>
          </div>
        </div>
        <div class="sf-advantage-row">
          <div class="sf-advantage-item">
            <img src="https://img.icons8.com/color/48/000000/briefcase.png" class="sf-advantage-icon" />
            <div class="sf-advantage-title">顺路订单多</div>
          </div>
          <div class="sf-advantage-item">
            <img src="https://img.icons8.com/color/48/000000/discount--v1.png" class="sf-advantage-icon" />
            <div class="sf-advantage-title">低信息服务费</div>
          </div>
          <div class="sf-advantage-item">
            <img src="https://img.icons8.com/color/48/000000/ok.png" class="sf-advantage-icon" />
            <div class="sf-advantage-title">合规共享出行</div>
          </div>
        </div>
        <div class="sf-faq-section">
          <div class="sf-faq-title">常见问题</div>
          <div class="sf-faq-list">
            <div class="sf-faq-item">
              <div class="sf-faq-q">Q：注册车主需要哪些资料？</div>
              <div class="sf-faq-a">A：上传真实有效的<em>驾驶证、行驶证</em>。证件不在身边可通过相册上传。</div>
            </div>
            <div class="sf-faq-item">
              <div class="sf-faq-q">Q：注册顺风车需要网约车证吗？</div>
              <div class="sf-faq-a">A：不需要，根据交通运输部规定，顺风车属于<em>非营利性的合乘共享出行，无需网约车证</em>。</div>
            </div>
          </div>
        </div>
        <button class="sf-join-btn">立即成为车主</button>
      </div>
    </div>
    <!-- 订票页面 -->
    <div v-if="activeTab === 'order'" class="ticket-container">
      <!-- 顶部tab -->
      <div class="ticket-tabs">
        <span :class="['ticket-tab', ticketTab === 'plane' ? 'active' : '']" @click="ticketTab = 'plane'">机票</span>
        <span :class="['ticket-tab', ticketTab === 'train' ? 'active' : '']" @click="ticketTab = 'train'">火车票</span>
      </div>
      <!-- 查询区 -->
      <div class="ticket-search-card">
        <div class="ticket-city-row">
          <span class="ticket-city">北京</span>
          <span class="ticket-switch"><img src='https://img.icons8.com/ios-glyphs/24/ff8200/sort-right.png' style='vertical-align:middle;'/></span>
          <span class="ticket-city">上海</span>
        </div>
        <div class="ticket-date-row">
          <span class="ticket-date">07月02日 <span class="ticket-date-tip">明天</span></span>
        </div>
        <div class="ticket-options-row">
          <label><input type="checkbox" /> 学生票</label>
          <label><input type="checkbox" /> 高铁动车</label>
          <span class="ticket-all-orders">全部订单 ></span>
        </div>
        <button class="ticket-search-btn">查询预订</button>
      </div>
      <!-- 服务宫格 -->
      <div class="ticket-services">
        <div class="ticket-service-item">
          <img src="https://img.icons8.com/color/48/000000/rent.png" />
          <span>滴滴租车</span>
        </div>
        <div class="ticket-service-item">
          <img src="https://img.icons8.com/color/48/000000/airport.png" />
          <span>接送机</span>
        </div>
        <div class="ticket-service-item">
          <img src="https://img.icons8.com/color/48/000000/clock--v1.png" />
          <span>预约用车</span>
        </div>
        <div class="ticket-service-item">
          <img src="https://img.icons8.com/color/48/000000/beach.png" />
          <span>旅游</span>
        </div>
      </div>
      <!-- 优惠券区 -->
      <div class="ticket-coupons-row">
        <div class="ticket-coupons-col">
          <div class="ticket-coupons-title">滴滴好券 <span class="ticket-coupons-more">更多好券</span></div>
          <div class="ticket-coupon-list">
            <div class="ticket-coupon-item"><b>5元</b> 火车票立减券 <span class="ticket-coupon-get">立即领取</span></div>
            <div class="ticket-coupon-item"><b>1元</b> 火车票立减券 <span class="ticket-coupon-get">立即领取</span></div>
            <div class="ticket-coupon-item"><b>1天</b> 租车免租券 <span class="ticket-coupon-get">立即领取</span></div>
          </div>
        </div>
        <div class="ticket-coupons-side">
          <div class="ticket-side-card">
            <div class="ticket-side-title">火车通勤 一单回本</div>
            <div class="ticket-side-desc">最高立省18元 <span class="ticket-side-link">去看看 ></span></div>
          </div>
          <div class="ticket-side-card">
            <div class="ticket-side-title">邀约助力 三券齐发</div>
            <div class="ticket-side-desc">打车+单车+火车票 <span class="ticket-side-link">去看看 ></span></div>
          </div>
        </div>
      </div>
      <!-- 底部说明 -->
      <div class="ticket-bottom-desc">
        滴滴火车票 · 无忧便捷行<br />
        <span class="ticket-bottom-tips">售后服务有保障 · 便捷购票无套路</span>
      </div>
    </div>
    <!-- 送货页面 -->
    <div v-if="activeTab === 'delivery'" class="delivery-container">
      <!-- 顶部 header -->
      <div class="delivery-header-row">
        <span class="delivery-title">滴滴送货</span>
        <span class="delivery-city">{{ cityName }}</span>
        <span class="delivery-time">{{ currentTime }}</span>
      </div>
      <!-- 服务类型tab -->
      <div class="delivery-tabs">
        <span class="delivery-tab active">送货</span>
        <span class="delivery-tab">快送</span>
        <span class="delivery-tab">搬家</span>
        <span class="delivery-tab">宠物专送</span>
        <span class="delivery-tab">零担拼货</span>
      </div>
      <!-- 车型选择 -->
      <div class="delivery-car-types">
        <span class="delivery-car-type">跑腿</span>
        <span class="delivery-car-type">四轮小件</span>
        <span class="delivery-car-type">微面</span>
        <span class="delivery-car-type active">小面</span>
        <span class="delivery-car-type">中面</span>
        <span class="delivery-car-type">依维柯</span>
      </div>
      <div class="delivery-car-sizes">
        <span class="delivery-car-size">小货</span>
        <span class="delivery-car-size">4米2</span>
        <span class="delivery-car-size">5米2</span>
        <span class="delivery-car-size">6米8</span>
        <span class="delivery-car-size">7米6</span>
        <span class="delivery-car-size">9米6</span>
        <span class="delivery-car-size active">全部</span>
      </div>
      <!-- 车辆图片和参数 -->
      <div class="delivery-car-info-vertical">
        <img class="delivery-car-img-vertical" src="https://img.icons8.com/color/96/000000/delivery.png" alt="小面" />
        <div class="delivery-car-title">小面</div>
        <div class="delivery-car-desc-vertical">
          <div>载方 <b>2.4~4立方</b> 载重 <b>0.5~0.8吨</b></div>
          <div>尺寸 <b>1.8~2.4米</b>长 <b>1.1~1.4米</b>高</div>
        </div>
        <div class="delivery-car-tags-vertical">
          <span class="delivery-car-tag">新能源</span>
          <span class="delivery-car-tag">拆座</span>
          <span class="delivery-car-tag">含支撑架</span>
        </div>
      </div>
      <!-- 地址输入 -->
      <div class="delivery-address-card">
        <div class="delivery-address-row">
          <span class="delivery-dot green"></span>
          <span class="delivery-address-label">{{ currentAddress }}</span>
          <span class="delivery-add-waypoint">+加途经点</span>
        </div>
        <div class="delivery-address-row">
          <span class="delivery-dot orange"></span>
          <input class="delivery-address-input" type="text" v-model="destination" placeholder="输入收货地" />
        </div>
      </div>
      <!-- 快捷功能区 -->
      <div class="delivery-quick-row">
        <div class="delivery-quick-item">
          <img src="https://img.icons8.com/ios-filled/50/4caf50/order-history.png" />
          <span>订单</span>
        </div>
        <div class="delivery-quick-item">
          <img src="https://img.icons8.com/fluency/48/hand-truck.png" />
          <span>装卸搬抬</span>
        </div>
        <div class="delivery-quick-item">
          <img src="https://img.icons8.com/ios-filled/50/4caf50/wallet-app.png" />
          <span>邀用户返现</span>
        </div>
        <div class="delivery-quick-item">
          <img src="https://img.icons8.com/ios-filled/50/ff8200/discount--v1.png" />
          <span>省钱中心</span>
        </div>
      </div>
      <!-- 底部logo -->
      <div class="delivery-bottom-logo">
        <img src="https://img.icons8.com/color/48/ff8200/delivery.png" />
        <span>滴滴送货</span>
      </div>
    </div>
    <!-- 个人中心页面 -->
    <div v-if="activeTab === 'mine'" class="mine-container">
      <!-- 顶部信息栏 -->
      <div class="mine-header-row">
        <span class="mine-time">{{ currentTime }}</span>
        <div class="mine-header-actions">
          <span class="mine-header-action">消息</span>
          <span class="mine-header-action">设置</span>
        </div>
      </div>
      <div class="mine-profile-row">
        <img class="mine-avatar" src="https://img.icons8.com/color/48/000000/user-male-circle--v2.png" />
        <div class="mine-profile-info">
          <div class="mine-phone">{{ maskedPhone }}</div>
          <div class="mine-profile-links">
            <span>里程</span> <span class="mine-link-arrow">›</span>
            <span>周报</span>
          </div>
        </div>
        <button class="logout-btn" @click="logout">退出登录</button>
      </div>
      <!-- 会员卡片 -->
      <div class="mine-vip-card">
        <span class="mine-vip-level">V1会员</span>
        <span class="mine-vip-desc">行程意外险 <b style='color:#ff8200'>5</b> 次</span>
        <button class="mine-vip-btn">查看</button>
      </div>
      <!-- 订单快捷入口 -->
      <div class="mine-order-row">
        <div class="mine-order-item">
          <img src="https://img.icons8.com/ios-filled/32/000000/todo-list.png" />
          <span>全部订单</span>
        </div>
        <div class="mine-order-item">
          <img src="https://img.icons8.com/ios-filled/32/000000/clock--v1.png" />
          <span>待出发</span>
        </div>
        <div class="mine-order-item">
          <img src="https://img.icons8.com/ios-filled/32/000000/paid.png" />
          <span>待支付</span>
        </div>
        <div class="mine-order-item">
          <img src="https://img.icons8.com/ios-filled/32/000000/ticket.png" />
          <span>开发票</span>
        </div>
        <div class="mine-order-item">
          <img src="https://img.icons8.com/ios-filled/32/000000/online-support.png" />
          <span>客服</span>
        </div>
      </div>
      <!-- 钱包区 -->
      <div class="mine-wallet-row">
        <div class="mine-wallet-item">
          <img src="https://img.icons8.com/ios-filled/32/000000/wallet-app.png" />
          <span>我的钱包</span>
        </div>
        <div class="mine-wallet-info">
          <span>6张优惠卡券</span>
          <span>18.8元单单返现</span>
          <span>1个金融福利</span>
          <span>9.68万可借</span>
        </div>
      </div>
      <!-- 功能区 -->
      <div class="mine-func-row">
        <div class="mine-func-item">
          <img src="https://img.icons8.com/ios-filled/32/ff8200/steering-wheel.png" />
          <span>成为司机</span>
        </div>
        <div class="mine-func-item">
          <img src="https://img.icons8.com/ios-filled/32/ff8200/family.png" />
          <span>家庭账户</span>
        </div>
        <div class="mine-func-item">
          <img src="https://img.icons8.com/ios-filled/32/ff8200/briefcase.png" />
          <span>企业权益</span>
        </div>
        <div class="mine-func-item">
          <img src="https://img.icons8.com/ios-filled/32/ff8200/medal2.png" />
          <span>出行成就</span>
        </div>
        <div class="mine-func-item">
          <img src="https://img.icons8.com/ios-filled/32/ff8200/coin-wallet.png" />
          <span>充值中心</span>
        </div>
        <div class="mine-func-item">
          <img src="https://img.icons8.com/ios-filled/32/ff8200/microphone.png" />
          <span>公众评议</span>
        </div>
        <div class="mine-func-item">
          <img src="https://img.icons8.com/ios-filled/32/ff8200/ok.png" />
          <span>评价有奖</span>
        </div>
        <div class="mine-func-item">
          <img src="https://img.icons8.com/ios-filled/32/ff8200/leaf.png" />
          <span>绿色出行</span>
        </div>
        <div class="mine-func-item">
          <img src="https://img.icons8.com/ios-filled/32/ff8200/t-shirt.png" />
          <span>个性皮肤</span>
        </div>
        <div class="mine-func-item">
          <img src="https://img.icons8.com/ios-filled/32/ff8200/alipay.png" />
          <span>滴滴支付</span>
        </div>
      </div>
      <!-- 福利横滑区 -->
      <div class="mine-benefit-row">
        <div class="mine-benefit-item">
          <img src="https://img.icons8.com/color/48/000000/sale.png" />
          <span>省钱套餐</span>
        </div>
        <div class="mine-benefit-item">
          <img src="https://img.icons8.com/color/48/000000/gift-card.png" />
          <span>天天神券</span>
        </div>
        <div class="mine-benefit-item">
          <img src="https://img.icons8.com/color/48/000000/money-bag.png" />
          <span>邀请赚钱</span>
        </div>
        <div class="mine-benefit-item">
          <img src="https://img.icons8.com/color/48/000000/orange.png" />
          <span>福利群</span>
        </div>
        <div class="mine-benefit-item">
          <img src="https://img.icons8.com/color/48/000000/free-shipping.png" />
          <span>邀请领券</span>
        </div>
      </div>
      <!-- 底部banner区 -->
      <div class="mine-banner-row">
        <div class="mine-banner-item orange">
          <div class="mine-banner-title">宠物打车立减10元</div>
          <div class="mine-banner-desc">夏日萌宠趣出行</div>
          <button class="mine-banner-btn">点我抢</button>
        </div>
        <div class="mine-banner-item blue">
          <div class="mine-banner-title">车主服务</div>
          <div class="mine-banner-desc">有车的人就用TA</div>
          <button class="mine-banner-btn">立即领</button>
        </div>
      </div>
      <!-- 底部logo -->
      <div class="mine-bottom-logo">
        <img src="https://img.icons8.com/color/48/ff8200/delivery.png" />
        <span>滴滴出行</span>
      </div>
    </div>
    <!-- tabbar 不变 -->
    <div class="didi-tabbar">
      <div class="didi-tabbar-item" v-for="tab in tabbar" :key="tab.text" @click="activeTab = tab.key">
        <img :src="tab.icon" class="didi-tabbar-icon" />
        <span class="didi-tabbar-text" :class="{active: activeTab === tab.key}">{{ tab.text }}</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      services: [
        { icon: 'https://img.icons8.com/color/48/000000/car--v2.png', text: '打车' },
        { icon: 'https://img.icons8.com/color/48/000000/good-pincode.png', text: '顺风车' },
        { icon: 'https://img.icons8.com/color/48/000000/driver.png', text: '代驾' },
        { icon: 'https://img.icons8.com/color/48/000000/sedan.png', text: '城际拼车' },
        { icon: 'https://img.icons8.com/color/48/000000/bus.png', text: '公交地铁' },
        { icon: 'https://img.icons8.com/color/48/000000/discount.png', text: '特价拼车' },
        { icon: 'https://img.icons8.com/color/48/000000/delivery--v2.png', text: '快送跑腿' },
        { icon: 'https://img.icons8.com/color/48/000000/money.png', text: '借钱' },
        { icon: 'https://img.icons8.com/color/48/000000/gas-pump.png', text: '优惠加油' },
        { icon: 'https://img.icons8.com/color/48/000000/taxi.png', text: '出租车' },
        { icon: 'https://img.icons8.com/color/48/000000/bicycle.png', text: '青桔骑行' },
        { icon: 'https://img.icons8.com/color/48/000000/delivery.png', text: '送货' },
        { icon: 'https://img.icons8.com/color/48/000000/car-rental.png', text: '滴滴租车' },
        { icon: 'https://img.icons8.com/color/48/000000/bus.png', text: '搬家' },
        { icon: 'https://img.icons8.com/color/48/000000/more.png', text: '更多服务' }
      ],
      tabbar: [
        { icon: 'https://img.icons8.com/color/48/000000/car--v2.png', text: '首页', key: 'home' },
        { icon: 'https://img.icons8.com/color/48/000000/smiling.png', text: '车主', key: 'owner' },
        { icon: 'https://img.icons8.com/color/48/000000/ticket--v1.png', text: '订票', key: 'order' },
        { icon: 'https://img.icons8.com/color/48/000000/delivery.png', text: '送货', key: 'delivery' },
        { icon: 'https://img.icons8.com/color/48/000000/user.png', text: '我的', key: 'mine' }
      ],
      canNext: true,
      currentAddress: '定位中...',
      activeTab: 'home',
      ownerTab: 'owner',
      ownerGrid1: [
        { icon: 'https://img.icons8.com/ios-filled/50/000000/to-do.png', text: '车主订单' },
        { icon: 'https://img.icons8.com/ios-filled/50/000000/discount--v1.png', text: '优惠券' },
        { icon: 'https://img.icons8.com/ios-filled/50/000000/wallet-app.png', text: '我的钱包' },
        { icon: 'https://img.icons8.com/ios-filled/50/000000/crown.png', text: '车主权益' }
      ],
      ownerGrid2: [
        { icon: 'https://img.icons8.com/color/48/000000/gas-pump.png', text: '优惠加油' },
        { icon: 'https://img.icons8.com/color/48/000000/flash-on.png', text: '扫码充电' },
        { icon: 'https://img.icons8.com/color/48/000000/driver.png', text: '叫代驾' },
        { icon: 'https://img.icons8.com/color/48/000000/car-maintenance.png', text: '滴滴保养' },
        { icon: 'https://img.icons8.com/color/48/000000/car-wash.png', text: '特惠洗车' },
        { icon: 'https://img.icons8.com/color/48/000000/money.png', text: '车主用钱' },
        { icon: 'https://img.icons8.com/color/48/000000/worker-male.png', text: '成为司机' },
        { icon: 'https://img.icons8.com/color/48/000000/sell.png', text: '高价卖车' },
        { icon: 'https://img.icons8.com/color/48/000000/insurance.png', text: '滴滴保险' },
        { icon: 'https://img.icons8.com/color/48/000000/more.png', text: '更多' }
      ],
      ticketTab: 'train',
      destination: '',
      cityName: '定位中',
      currentTime: '',
      userPhone: '18512342848', // 示例，实际应由登录赋值
    };
  },
  computed: {
    maskedPhone() {
      const phone = this.userPhone || '';
      if (phone.length === 11) {
        return phone.substr(0, 3) + '****' + phone.substr(7, 4);
      }
      return phone;
    }
  },
  mounted() {
    this.loadBMapScript(() => {
      this.initBMapWithBrowserLocation();
    });
    this.updateTime();
    setInterval(this.updateTime, 1000);
    // 读取本地存储手机号
    const phone = localStorage.getItem('userPhone');
    if (phone) {
      this.userPhone = phone;
    }
  },
  methods: {
    loadBMapScript(callback) {
      if (window.BMapGL) {
        callback();
        return;
      }
      // 直接加载 getscript，不要加载 api 入口脚本，避免 document.write 报错
      const script = document.createElement('script');
      script.type = 'text/javascript';
      script.src = 'https://api.map.baidu.com/getscript?type=webgl&v=1.0&ak=Kdx19DyjGxp4balzBctaBkKAdNIP6zI2&services=&t=' + Date.now();
      script.onload = callback;
      document.head.appendChild(script);
    },
    initBMapWithBrowserLocation() {
      this.$nextTick(() => {
        if (!window.BMapGL) {
          alert('百度地图API加载失败，请检查网络！');
          return;
        }
        const map = new window.BMapGL.Map('amap-container');
        map.enableScrollWheelZoom(true);
        // 默认中心点（上海）
        let point = new window.BMapGL.Point(121.47, 31.23);
        map.centerAndZoom(point, 16);
        const self = this;
        // 优先用浏览器原生定位
        if (navigator.geolocation) {
          navigator.geolocation.getCurrentPosition(
            (pos) => {
              const lng = pos.coords.longitude;
              const lat = pos.coords.latitude;
              // 用百度坐标转换
              const convertor = new window.BMapGL.Convertor();
              convertor.translate(
                [new window.BMapGL.Point(lng, lat)],
                1, // 1: WGS-84 to BD-09
                5,
                (data) => {
                  if (data.status === 0) {
                    const bdPoint = data.points[0];
                    map.centerAndZoom(bdPoint, 16);
                    const marker = new window.BMapGL.Marker(bdPoint);
                    map.addOverlay(marker);
                    // 逆地理编码获取地址
                    const geocoder = new window.BMapGL.Geocoder();
                    geocoder.getLocation(bdPoint, (result) => {
                      if (result && result.address) {
                        self.currentAddress = result.address;
                        self.setCityName(result.address);
                      } else {
                        self.currentAddress = '未知位置';
                        self.setCityName('');
                      }
                    });
                  } else {
                    // 坐标转换失败，降级用百度API定位
                    self.useBaiduGeolocation(map);
                  }
                }
              );
            },
            (err) => {
              // 浏览器定位失败，降级用百度API定位
              self.useBaiduGeolocation(map);
            },
            { enableHighAccuracy: true, timeout: 5000 }
          );
        } else {
          // 不支持浏览器定位，直接用百度API定位
          self.useBaiduGeolocation(map);
        }
      });
    },
    useBaiduGeolocation(map) {
      const self = this;
      const geolocation = new window.BMapGL.Geolocation();
      geolocation.getCurrentPosition(function(r){
        console.log('百度定位返回：', r);
        if(this.getStatus() == window.BMAP_STATUS_SUCCESS){
          map.centerAndZoom(r.point, 16);
          const marker = new window.BMapGL.Marker(r.point);
          map.addOverlay(marker);
          // 逆地理编码获取地址
          const geocoder = new window.BMapGL.Geocoder();
          geocoder.getLocation(r.point, (result) => {
            if (result && result.address) {
              self.currentAddress = result.address;
              self.setCityName(result.address);
            } else {
              self.currentAddress = '未知位置';
              self.setCityName('');
            }
          });
        }
        else {
          alert('定位失败：' + this.getStatus());
        }
      }, {enableHighAccuracy: true});
    },
    updateTime() {
      // 获取北京时间
      const now = new Date();
      // 转为东八区
      const utc = now.getTime() + now.getTimezoneOffset() * 60000;
      const beijing = new Date(utc + 8 * 3600000);
      const h = beijing.getHours().toString().padStart(2, '0');
      const m = beijing.getMinutes().toString().padStart(2, '0');
      this.currentTime = `${h}:${m}`;
    },
    setCityName(address) {
      // 只提取省级（如"台湾省"、"江苏省"、"上海市"等）
      if (!address) {
        this.cityName = '定位中';
        return;
      }
      // 匹配"xx省"或"xx市"或"xx自治区"或"xx特别行政区"
      const match = address.match(/([\u4e00-\u9fa5]+?(省|市|自治区|特别行政区))/);
      if (match && match[1]) {
        this.cityName = match[1];
      } else {
        // 没有省级时，取第一个空格或逗号前的内容
        this.cityName = address.split(/[ ,，]/)[0] || address;
      }
    },
    nextStep() {
      console.log('点击了下一步');
    },
    logout() {
      localStorage.removeItem('userPhone');
      this.userPhone = '';
      uni.redirectTo({ url: '/pages/index/index' });
    },
  }
}
</script>

<style scoped>
body, html {
  background: #f7f8fa;
}
.didi-container {
  background: linear-gradient(135deg, #fff7e6 0%, #f7f8fa 100%);
  min-height: 100vh;
  max-width: 480px;
  margin: 0 auto;
  padding: 0 0 80px 0;
  box-sizing: border-box;
  position: relative;
  font-family: 'PingFang SC', 'Helvetica Neue', Arial, 'Microsoft YaHei', sans-serif;
}
.didi-header-float {
  position: static;
  margin: 18px auto 0 auto;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 8px;
  height: 54px;
  background: rgba(255,255,255,0.95);
  border-radius: 18px;
  box-shadow: 0 4px 18px rgba(0,0,0,0.06);
  gap: 12px;
  max-width: 480px;
  width: calc(100% - 8px);
}
.didi-header-left {
  display: flex;
  align-items: center;
  gap: 6px;
  flex-shrink: 0;
}
.didi-header-location {
  display: flex;
  align-items: baseline;
  gap: 6px;
  min-width: 0;
}
.didi-header-city {
  font-size: 18px;
  color: #222;
  font-weight: 500;
  line-height: 1;
  letter-spacing: 1px;
  max-width: 150px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.didi-header-time {
  font-size: 16px;
  color: #888;
  font-weight: 500;
  margin-left: 2px;
  letter-spacing: 1px;
}
.didi-header-btns {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-shrink: 0;
}
.didi-header-btn {
  background: #fff;
  border: none;
  border-radius: 18px;
  min-width: 48px;
  padding: 0 12px;
  font-size: 13px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  white-space: nowrap;
  box-sizing: border-box;
  box-shadow: 0 2px 8px rgba(0,0,0,0.08);
  transition: box-shadow 0.2s, transform 0.2s;
}
.didi-header-btn.icon-only {
  width: 36px;
  padding: 0;
  justify-content: center;
}
.didi-header-btn:active {
  background: #f7f8fa;
  transform: scale(0.97);
}
.didi-header-icon {
  width: 20px;
  height: 20px;
}
.didi-map-card {
  position: relative;
  margin: 18px 16px 18px 16px;
  padding: 0;
}
.didi-map {
  width: 100%;
  height: 260px;
  border-radius: 16px;
  overflow: hidden;
  background: #f5f5f5;
  position: relative;
  z-index: 1;
}
.didi-banner-card {
  background: linear-gradient(90deg, #fff7e6 0%, #fff 100%);
  border-radius: 16px;
  box-shadow: 0 2px 12px rgba(255,130,0,0.10);
  margin: 0 16px 16px 16px;
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px 18px;
  animation: fadeIn 1s;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(-10px);}
  to { opacity: 1; transform: translateY(0);}
}

.didi-banner-icon {
  width: 24px;
  height: 24px;
}

.didi-banner-text {
  font-size: 15px;
  color: #ff8200;
  flex: 1;
  font-weight: 600;
}

.didi-banner-btn {
  background: #FFF0E1;
  color: #ff8200;
  border: none;
  border-radius: 12px;
  padding: 0 14px;
  font-size: 14px;
  height: 28px;
  line-height: 28px;
  font-weight: 600;
  transition: background 0.2s, transform 0.2s;
}
.didi-banner-btn:active {
  background: #ffd6a0;
  transform: scale(0.96);
}

.didi-trip-card {
  background: #fff;
  border-radius: 20px;
  box-shadow: 0 4px 24px rgba(255,130,0,0.06);
  margin: 0 16px 18px 16px;
  padding: 18px 20px 20px 20px;
  display: flex;
  flex-direction: column;
  gap: 0;
}

.didi-trip-row {
  display: flex;
  align-items: center;
  font-size: 15px;
  font-weight: 500;
  margin-bottom: 0;
}

.didi-trip-label {
  background: #e8f5e9;
  color: #4caf50;
  border-radius: 6px;
  font-size: 13px;
  font-weight: 600;
  padding: 3px 10px;
  margin-right: 8px;
}

.didi-trip-place {
  color: #222;
  font-size: 15px;
  font-weight: 500;
  flex: 1;
}

.didi-trip-arrow {
  color: #bbb;
  font-size: 18px;
  margin-left: 6px;
}

.didi-trip-divider {
  height: 1px;
  background: #f2f3f5;
  margin: 14px 0 14px 0;
  width: 100%;
}

.didi-dest-row {
  display: flex;
  align-items: center;
  margin-bottom: 8px;
}

.didi-dot {
  width: 8px;
  height: 8px;
  background: #ff8200;
  border-radius: 50%;
  margin-right: 8px;
}

.didi-dest-q {
  font-size: 17px;
  color: #222;
  font-weight: bold;
}

.didi-order-types {
  display: flex;
  align-items: center;
  gap: 18px;
  padding-left: 8px;
  margin-bottom: 18px;
}

.didi-order-type {
  font-size: 13px;
  color: #888;
  background: #f7f8fa;
  border-radius: 8px;
  padding: 2px 10px;
  font-weight: 500;
}

.didi-dest-input {
  width: 100%;
  height: 44px;
  border: 1px solid #eee;
  border-radius: 12px;
  padding: 0 14px;
  font-size: 16px;
  margin-bottom: 10px;
  box-sizing: border-box;
  outline: none;
  background: #f7f8fa;
}
.didi-dest-input:focus {
  border-color: #ff8200;
  background: #fff;
}

.didi-next-btn {
  width: 100%;
  background: linear-gradient(90deg, #ff8200 0%, #ffb300 100%);
  color: #fff;
  border: none;
  border-radius: 24px;
  font-size: 19px;
  font-weight: bold;
  height: 50px;
  margin-top: 8px;
  box-shadow: 0 2px 12px rgba(255,130,0,0.12);
  transition: background 0.2s, transform 0.2s;
}
.didi-next-btn:active {
  background: #ff8200;
  transform: scale(0.98);
}
.didi-next-btn:disabled {
  background: #ffd6a0;
  color: #fff;
}

.didi-services-card {
  display: flex;
  flex-wrap: wrap;
  background: #fff;
  border-radius: 20px;
  padding: 12px 0 0 0;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
  margin: 0 16px 18px 16px;
  justify-content: space-between;
}

.didi-service {
  width: 20%;
  min-width: 70px;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 20px;
  transition: transform 0.2s;
}
.didi-service:hover {
  transform: translateY(-4px) scale(1.05);
}

.didi-service-icon {
  width: 38px;
  height: 38px;
  margin-bottom: 8px;
}

.didi-service-text {
  font-size: 15px;
  color: #333;
  font-weight: 600;
}

.didi-ads-row {
  display: flex;
  gap: 12px;
  margin: 0 16px 18px 16px;
}

.didi-ad-card {
  flex: 1;
  background: #fff;
  border-radius: 18px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
  padding: 14px 16px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
}
.didi-ad-card.orange {
  background: linear-gradient(90deg, #fff7e6 0%, #fff 100%);
}
.didi-ad-card.blue {
  background: linear-gradient(90deg, #e6f7ff 0%, #fff 100%);
}
.didi-ad-title {
  font-size: 16px;
  font-weight: bold;
  color: #ff8200;
  margin-bottom: 4px;
}
.didi-ad-desc {
  font-size: 14px;
  color: #222;
  margin-bottom: 8px;
}
.didi-ad-btn {
  background: #ff8200;
  color: #fff;
  border: none;
  border-radius: 14px;
  font-size: 14px;
  padding: 4px 16px;
  font-weight: bold;
  transition: background 0.2s, transform 0.2s;
}
.didi-ad-btn:active {
  background: #ffb300;
  transform: scale(0.97);
}

.didi-tabbar {
  position: fixed;
  left: 0;
  bottom: 0;
  width: 100vw;
  max-width: 480px;
  height: 56px;
  background: #fff;
  display: flex;
  align-items: center;
  justify-content: space-around;
  box-shadow: 0 -2px 16px rgba(255,130,0,0.06);
  z-index: 10;
  border-radius: 18px 18px 0 0;
  padding-bottom: env(safe-area-inset-bottom);
}

.didi-tabbar-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  flex: 1;
  transition: background 0.2s;
  border-radius: 12px;
  padding: 2px 0;
}
.didi-tabbar-item:active {
  background: #fff7e6;
}

.didi-tabbar-icon {
  width: 24px;
  height: 24px;
  margin-bottom: 2px;
}

.didi-tabbar-text {
  font-size: 12px;
  color: #333;
  transition: color 0.2s;
}
.didi-tabbar-text.active {
  color: #ff8200;
  font-weight: bold;
  text-shadow: 0 1px 4px #fff0e1;
}

/* 车主服务页面样式优化 */
.owner-container {
  background: #f7f8fa;
  min-height: 100vh;
  max-width: 480px;
  margin: 0 auto;
  padding: 0 0 80px 0;
  box-sizing: border-box;
  position: relative;
}
.owner-header-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin: 18px 16px 0 16px;
  padding: 0 18px;
  height: 54px;
  background: rgba(255,255,255,0.95);
  border-radius: 18px;
  box-shadow: 0 4px 18px rgba(0,0,0,0.06);
}
.owner-header-tabs {
  display: flex;
  align-items: center;
  gap: 18px;
}
.owner-header-tab {
  font-size: 20px;
  font-weight: 600;
  color: #888;
  cursor: pointer;
  padding: 0 6px 6px 6px;
  border-bottom: 2px solid transparent;
  transition: color 0.2s, border-bottom 0.2s;
}
.owner-header-tab.active {
  color: #222;
  border-bottom: 3px solid #ff8200;
}
.owner-center-btn {
  background: none;
  border: none;
  color: #888;
  font-size: 15px;
  font-weight: 500;
}
.owner-addcar-banner {
  background: linear-gradient(90deg, #fff7e6 0%, #fff 100%);
  border-radius: 16px;
  box-shadow: 0 2px 12px rgba(255,130,0,0.10);
  margin: 16px 16px 16px 16px;
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px 18px;
}
.owner-addcar-btn {
  background: #ff8200;
  color: #fff;
  border: none;
  border-radius: 12px;
  padding: 0 14px;
  font-size: 14px;
  height: 32px;
  line-height: 32px;
  font-weight: 600;
  transition: background 0.2s, transform 0.2s;
}
.owner-addcar-btn:active {
  background: #ffb300;
  transform: scale(0.97);
}
.owner-gift {
  font-size: 15px;
  color: #ff8200;
  flex: 1;
  font-weight: 600;
}
.owner-grid {
  display: flex;
  flex-wrap: wrap;
  background: #fff;
  border-radius: 20px;
  padding: 12px 0 0 0;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
  margin: 0 16px 18px 16px;
  justify-content: space-between;
}
.owner-grid-item {
  width: 25%;
  min-width: 70px;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 20px;
  transition: transform 0.2s;
}
.owner-grid-item:hover {
  transform: translateY(-4px) scale(1.05);
}
.owner-grid-icon {
  width: 38px;
  height: 38px;
  margin-bottom: 8px;
}
.owner-grid-text {
  font-size: 15px;
  color: #333;
  font-weight: 600;
}
.owner-cards-row {
  display: flex;
  gap: 12px;
  margin: 0 16px 18px 16px;
}
.owner-card {
  flex: 1;
  background: #fff;
  border-radius: 18px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
  padding: 14px 16px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
}
.owner-card.orange {
  background: linear-gradient(90deg, #fff7e6 0%, #fff 100%);
}
.owner-card.blue {
  background: linear-gradient(90deg, #e6f7ff 0%, #fff 100%);
}
.owner-card.green {
  background: linear-gradient(90deg, #e6ffe6 0%, #fff 100%);
}
.owner-card-title {
  font-size: 16px;
  font-weight: bold;
  color: #ff8200;
  margin-bottom: 4px;
}
.owner-card-desc {
  font-size: 14px;
  color: #222;
  margin-bottom: 8px;
}
.owner-card-btn {
  background: #ff8200;
  color: #fff;
  border: none;
  border-radius: 14px;
  font-size: 14px;
  padding: 4px 16px;
  font-weight: bold;
  transition: background 0.2s, transform 0.2s;
}
.owner-card-btn:active {
  background: #ffb300;
  transform: scale(0.97);
}
.owner-fuel-section {
  background: #fff;
  border-radius: 18px;
  box-shadow: 0 2px 12px rgba(255,130,0,0.06);
  margin: 0 16px 18px 16px;
  padding: 16px;
}
.owner-fuel-tabs {
  display: flex;
  gap: 18px;
  font-size: 15px;
  margin-bottom: 10px;
}
.owner-fuel-tabs span {
  color: #888;
  cursor: pointer;
}
.owner-fuel-tabs .active {
  color: #ff8200;
  font-weight: bold;
  border-bottom: 2px solid #ff8200;
}
.owner-fuel-info {
  margin-top: 8px;
}
.owner-fuel-title {
  font-size: 16px;
  font-weight: bold;
  color: #222;
}
.owner-fuel-desc {
  font-size: 13px;
  color: #888;
  margin-bottom: 6px;
}
.owner-fuel-price {
  display: flex;
  align-items: center;
  gap: 8px;
}
.owner-fuel-price .price {
  color: #ff8200;
  font-size: 20px;
  font-weight: bold;
}
.owner-fuel-price .member {
  background: #fff0e1;
  color: #ff8200;
  border-radius: 6px;
  font-size: 12px;
  padding: 2px 8px;
  margin-left: 4px;
}
.owner-fuel-price .origin {
  color: #bbb;
  font-size: 13px;
  text-decoration: line-through;
}
.owner-fuel-btn {
  background: #ff8200;
  color: #fff;
  border: none;
  border-radius: 12px;
  font-size: 14px;
  padding: 4px 16px;
  font-weight: bold;
  margin-left: 12px;
  transition: background 0.2s, transform 0.2s;
}
.owner-fuel-btn:active {
  background: #ffb300;
  transform: scale(0.97);
}

/* 顺风车页面样式 */
.sf-container {
  background: #f7f8fa;
  border-radius: 18px;
  margin: 0 0 16px 0;
  padding: 0 0 24px 0;
}
.sf-header-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 18px 18px 0 18px;
}
.sf-title {
  font-size: 22px;
  font-weight: bold;
  color: #222;
}
.sf-time {
  font-size: 16px;
  color: #888;
}
.sf-banner {
  background: linear-gradient(90deg, #e6ffe6 0%, #fff 100%);
  border-radius: 16px;
  margin: 16px 16px 0 16px;
  padding: 18px 18px 0 18px;
  display: flex;
  align-items: center;
  gap: 16px;
  position: relative;
}
.sf-banner-img {
  width: 120px;
  height: 80px;
  border-radius: 12px;
  object-fit: cover;
}
.sf-banner-text {
  display: flex;
  flex-direction: column;
  gap: 6px;
}
.sf-banner-main {
  font-size: 18px;
  color: #222;
  font-weight: bold;
}
.sf-banner-sub {
  font-size: 15px;
  color: #ff8200;
}
.sf-advantage-row {
  display: flex;
  justify-content: space-around;
  margin: 18px 0 0 0;
  padding: 0 16px;
}
.sf-advantage-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
}
.sf-advantage-icon {
  width: 38px;
  height: 38px;
}
.sf-advantage-title {
  font-size: 15px;
  color: #333;
  font-weight: 600;
}
.sf-faq-section {
  background: #fff;
  border-radius: 16px;
  margin: 18px 16px 0 16px;
  padding: 16px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
}
.sf-faq-title {
  font-size: 16px;
  font-weight: bold;
  color: #222;
  margin-bottom: 10px;
}
.sf-faq-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.sf-faq-item {
  font-size: 14px;
  color: #333;
}
.sf-faq-q {
  font-weight: bold;
  margin-bottom: 2px;
}
.sf-faq-a em {
  color: #ff8200;
  font-style: normal;
  font-weight: bold;
}
.sf-join-btn {
  width: calc(100% - 32px);
  margin: 24px 16px 0 16px;
  background: linear-gradient(90deg, #4be281 0%, #00c853 100%);
  color: #fff;
  border: none;
  border-radius: 24px;
  font-size: 20px;
  font-weight: bold;
  height: 50px;
  box-shadow: 0 2px 12px rgba(76,226,129,0.12);
  transition: background 0.2s, transform 0.2s;
}
.sf-join-btn:active {
  background: #00c853;
  transform: scale(0.98);
}

/* 订票页面样式 */
.ticket-container {
  background: #f7f8fa;
  min-height: 100vh;
  max-width: 480px;
  margin: 0 auto;
  padding: 0 0 80px 0;
  box-sizing: border-box;
  position: relative;
}
.ticket-tabs {
  display: flex;
  align-items: center;
  background: #fff;
  border-radius: 18px 18px 0 0;
  margin: 0 0 0 0;
  padding: 0 0 0 18px;
  gap: 24px;
  height: 48px;
}
.ticket-tab {
  font-size: 18px;
  color: #888;
  font-weight: 500;
  padding: 0 8px 8px 8px;
  border-bottom: 2px solid transparent;
  cursor: pointer;
  transition: color 0.2s, border-bottom 0.2s;
}
.ticket-tab.active {
  color: #ff8200;
  font-weight: bold;
  border-bottom: 3px solid #ff8200;
}
.ticket-search-card {
  background: #fff;
  border-radius: 18px;
  margin: 0 16px 16px 16px;
  padding: 18px 18px 12px 18px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
}
.ticket-city-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 22px;
  font-weight: bold;
  color: #222;
  margin-bottom: 10px;
}
.ticket-city {
  flex: 1;
  text-align: center;
}
.ticket-switch {
  flex: 0 0 40px;
  text-align: center;
}
.ticket-date-row {
  font-size: 18px;
  color: #222;
  margin-bottom: 10px;
}
.ticket-date-tip {
  font-size: 14px;
  color: #888;
  margin-left: 6px;
}
.ticket-options-row {
  display: flex;
  align-items: center;
  gap: 18px;
  font-size: 15px;
  color: #888;
  margin-bottom: 12px;
}
.ticket-all-orders {
  margin-left: auto;
  color: #888;
  font-size: 14px;
  cursor: pointer;
}
.ticket-search-btn {
  width: 100%;
  background: linear-gradient(90deg, #ff8200 0%, #ffb300 100%);
  color: #fff;
  border: none;
  border-radius: 24px;
  font-size: 20px;
  font-weight: bold;
  height: 48px;
  margin-top: 8px;
  box-shadow: 0 2px 12px rgba(255,130,0,0.12);
  transition: background 0.2s, transform 0.2s;
}
.ticket-search-btn:active {
  background: #ff8200;
  transform: scale(0.98);
}
.ticket-services {
  display: flex;
  background: #fff;
  border-radius: 16px;
  margin: 0 16px 16px 16px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
  padding: 12px 0;
  justify-content: space-around;
}
.ticket-service-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 15px;
  color: #333;
  font-weight: 600;
}
.ticket-service-item img {
  width: 38px;
  height: 38px;
  margin-bottom: 6px;
}
.ticket-coupons-row {
  display: flex;
  gap: 12px;
  margin: 0 16px 16px 16px;
}
.ticket-coupons-col {
  flex: 1.2;
  background: #fff;
  border-radius: 16px;
  padding: 14px 10px 10px 14px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
}
.ticket-coupons-title {
  font-size: 15px;
  font-weight: bold;
  color: #ff8200;
  margin-bottom: 8px;
}
.ticket-coupons-more {
  color: #888;
  font-size: 13px;
  float: right;
  font-weight: normal;
}
.ticket-coupon-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.ticket-coupon-item {
  background: #fff7e6;
  border-radius: 8px;
  padding: 6px 10px;
  font-size: 15px;
  color: #ff8200;
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
}
.ticket-coupon-item b {
  font-size: 18px;
  margin-right: 4px;
}
.ticket-coupon-get {
  color: #ff8200;
  font-size: 13px;
  margin-left: auto;
}
.ticket-coupons-side {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.ticket-side-card {
  background: #fff;
  border-radius: 16px;
  padding: 10px 12px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
}
.ticket-side-title {
  font-size: 15px;
  font-weight: bold;
  color: #ff8200;
  margin-bottom: 4px;
}
.ticket-side-desc {
  font-size: 14px;
  color: #222;
}
.ticket-side-link {
  color: #ff8200;
  font-size: 13px;
  margin-left: 6px;
}
.ticket-bottom-desc {
  text-align: center;
  color: #888;
  font-size: 14px;
  margin: 18px 0 0 0;
}
.ticket-bottom-tips {
  color: #bbb;
  font-size: 13px;
}
.didi-order-types-small {
  margin-top: 4px;
  margin-bottom: 10px;
}
.didi-order-types-small .didi-order-type {
  font-size: 12px;
  padding: 2px 8px;
}
/* 送货页面样式 */
.delivery-container {
  background: #f7f8fa;
  min-height: 100vh;
  max-width: 480px;
  margin: 0 auto;
  padding: 0 0 80px 0;
  box-sizing: border-box;
  position: relative;
}
.delivery-header-row {
  display: flex;
  align-items: center;
  gap: 12px;
  margin: 18px 16px 0 16px;
  padding: 0 8px;
  height: 54px;
  background: rgba(255,255,255,0.95);
  border-radius: 18px;
  box-shadow: 0 4px 18px rgba(0,0,0,0.06);
  font-size: 18px;
  font-weight: 500;
}
.delivery-title {
  color: #ff8200;
  font-weight: bold;
  font-size: 20px;
}
.delivery-city {
  color: #222;
  font-size: 18px;
  font-weight: 500;
}
.delivery-time {
  color: #888;
  font-size: 16px;
  margin-left: auto;
}
.delivery-tabs {
  display: flex;
  align-items: center;
  gap: 18px;
  background: #fff;
  border-radius: 18px 18px 0 0;
  margin: 0 0 0 0;
  padding: 0 0 0 18px;
  height: 48px;
  font-size: 16px;
  font-weight: 500;
}
.delivery-tab {
  color: #888;
  padding: 0 8px 8px 8px;
  border-bottom: 2px solid transparent;
  cursor: pointer;
  transition: color 0.2s, border-bottom 0.2s;
}
.delivery-tab.active {
  color: #ff8200;
  font-weight: bold;
  border-bottom: 3px solid #ff8200;
}
.delivery-car-types, .delivery-car-sizes {
  display: flex;
  align-items: center;
  gap: 10px;
  margin: 0 16px 0 16px;
  padding: 10px 0 0 0;
  font-size: 15px;
  flex-wrap: wrap;
}
.delivery-car-type, .delivery-car-size {
  color: #888;
  background: #f7f8fa;
  border-radius: 8px;
  padding: 2px 10px;
  font-weight: 500;
  cursor: pointer;
  margin-bottom: 6px;
}
.delivery-car-type.active, .delivery-car-size.active {
  color: #fff;
  background: #ff8200;
}
.delivery-car-info-vertical {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: #fff;
  border-radius: 16px;
  margin: 12px 16px 0 16px;
  padding: 16px 12px 12px 12px;
  box-shadow: 0 2px 12px rgba(255,130,0,0.10);
  gap: 6px;
}
.delivery-car-img-vertical {
  width: 90px;
  height: 60px;
  object-fit: contain;
  margin-bottom: 4px;
}
.delivery-car-title {
  font-size: 18px;
  color: #222;
  font-weight: bold;
  margin-bottom: 2px;
}
.delivery-car-desc-vertical {
  font-size: 14px;
  color: #333;
  text-align: center;
  margin-bottom: 2px;
}
.delivery-car-tags-vertical {
  display: flex;
  gap: 8px;
  margin-top: 2px;
  flex-wrap: wrap;
  justify-content: center;
}
.delivery-address-card {
  background: #fff;
  border-radius: 16px;
  margin: 12px 16px 0 16px;
  padding: 14px 18px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
}
.delivery-address-row {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 10px;
}
.delivery-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  display: inline-block;
}
.delivery-dot.green {
  background: #4caf50;
}
.delivery-dot.orange {
  background: #ff8200;
}
.delivery-address-label {
  color: #222;
  font-size: 15px;
  flex: 1;
}
.delivery-add-waypoint {
  color: #888;
  font-size: 13px;
  cursor: pointer;
}
.delivery-address-input {
  width: 100%;
  height: 36px;
  border: 1px solid #eee;
  border-radius: 8px;
  padding: 0 10px;
  font-size: 15px;
  background: #f7f8fa;
  outline: none;
}
.delivery-address-input:focus {
  border-color: #ff8200;
  background: #fff;
}
.delivery-quick-row {
  display: flex;
  justify-content: space-around;
  margin: 18px 16px 0 16px;
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
  padding: 10px 0;
}
.delivery-quick-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 14px;
  color: #333;
  font-weight: 600;
  gap: 4px;
}
.delivery-quick-item img {
  width: 28px;
  height: 28px;
}
.delivery-bottom-logo {
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 24px 0 0 0;
  gap: 8px;
  color: #ff8200;
  font-size: 18px;
  font-weight: bold;
}
.delivery-bottom-logo img {
  width: 28px;
  height: 28px;
}
/* 个人中心页面样式 */
.mine-container {
  background: #f7f8fa;
  min-height: 100vh;
  max-width: 480px;
  margin: 0 auto;
  padding: 0 0 80px 0;
  box-sizing: border-box;
  position: relative;
}
.mine-header-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin: 0 0 0 0;
  padding: 18px 18px 0 18px;
}
.mine-time {
  font-size: 20px;
  color: #222;
  font-weight: bold;
}
.mine-header-actions {
  display: flex;
  gap: 18px;
  font-size: 15px;
  color: #888;
}
.mine-header-action {
  cursor: pointer;
}
.mine-profile-row {
  display: flex;
  align-items: center;
  gap: 14px;
  margin: 8px 18px 0 18px;
  background: #fff;
  border-radius: 16px;
  padding: 12px 18px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
}
.mine-avatar {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: #f7f8fa;
}
.mine-profile-info {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
}
.mine-phone {
  font-size: 18px;
  color: #222;
  font-weight: bold;
}
.mine-profile-links {
  font-size: 14px;
  color: #888;
  display: flex;
  gap: 8px;
  align-items: center;
}
.mine-link-arrow {
  color: #bbb;
  font-size: 16px;
}
.mine-vip-card {
  background: #e6f2ff;
  border-radius: 14px;
  margin: 14px 18px 0 18px;
  padding: 12px 18px;
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 15px;
  color: #222;
  position: relative;
}
.mine-vip-level {
  color: #3b7cff;
  font-weight: bold;
  font-size: 18px;
}
.mine-vip-desc {
  flex: 1;
  color: #888;
  font-size: 15px;
}
.mine-vip-btn {
  background: #fff;
  color: #3b7cff;
  border: 1px solid #3b7cff;
  border-radius: 10px;
  font-size: 14px;
  padding: 2px 16px;
  font-weight: 600;
  cursor: pointer;
}
.mine-order-row {
  display: flex;
  justify-content: space-around;
  background: #fff;
  border-radius: 16px;
  margin: 14px 18px 0 18px;
  padding: 12px 0;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
}
.mine-order-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 13px;
  color: #333;
  font-weight: 600;
  gap: 4px;
}
.mine-order-item img {
  width: 28px;
  height: 28px;
}
.mine-wallet-row {
  display: flex;
  align-items: center;
  background: #fff;
  border-radius: 16px;
  margin: 14px 18px 0 18px;
  padding: 12px 18px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
  gap: 18px;
}
.mine-wallet-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 14px;
  color: #333;
  font-weight: 600;
  gap: 4px;
}
.mine-wallet-item img {
  width: 28px;
  height: 28px;
}
.mine-wallet-info {
  display: flex;
  flex-wrap: wrap;
  gap: 12px 18px;
  font-size: 13px;
  color: #888;
}
.mine-func-row {
  display: flex;
  flex-wrap: wrap;
  background: #fff;
  border-radius: 16px;
  margin: 14px 18px 0 18px;
  padding: 12px 0 0 0;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
  justify-content: space-between;
}
.mine-func-item {
  width: 20%;
  min-width: 70px;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 18px;
  font-size: 13px;
  color: #333;
  font-weight: 600;
  gap: 4px;
}
.mine-func-item img {
  width: 28px;
  height: 28px;
}
.mine-benefit-row {
  display: flex;
  background: #fff;
  border-radius: 16px;
  margin: 14px 18px 0 18px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
  padding: 10px 0;
  justify-content: space-around;
  overflow-x: auto;
}
.mine-benefit-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 13px;
  color: #333;
  font-weight: 600;
  gap: 4px;
  min-width: 70px;
}
.mine-benefit-item img {
  width: 32px;
  height: 32px;
}
.mine-banner-row {
  display: flex;
  gap: 12px;
  margin: 18px 18px 0 18px;
}
.mine-banner-item {
  flex: 1;
  background: #fff;
  border-radius: 18px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.04);
  padding: 14px 16px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
}
.mine-banner-item.orange {
  background: linear-gradient(90deg, #fff7e6 0%, #fff 100%);
}
.mine-banner-item.blue {
  background: linear-gradient(90deg, #e6f7ff 0%, #fff 100%);
}
.mine-banner-title {
  font-size: 16px;
  font-weight: bold;
  color: #ff8200;
  margin-bottom: 4px;
}
.mine-banner-desc {
  font-size: 14px;
  color: #222;
  margin-bottom: 8px;
}
.mine-banner-btn {
  background: #ff8200;
  color: #fff;
  border: none;
  border-radius: 14px;
  font-size: 14px;
  padding: 4px 16px;
  font-weight: bold;
  transition: background 0.2s, transform 0.2s;
  cursor: pointer;
}
.mine-banner-btn:active {
  background: #ffb300;
  transform: scale(0.97);
}
.mine-bottom-logo {
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 24px 0 0 0;
  gap: 8px;
  color: #ff8200;
  font-size: 18px;
  font-weight: bold;
}
.mine-bottom-logo img {
  width: 28px;
  height: 28px;
}
.logout-btn {
  margin-left: 20px;
  padding: 8px 18px;
  background: #ff4d4f;
  color: #fff;
  border: none;
  border-radius: 16px;
  font-size: 16px;
  cursor: pointer;
  box-shadow: 0 2px 8px rgba(255,77,79,0.08);
  transition: background 0.2s;
}
.logout-btn:hover {
  background: #d9363e;
}
</style> 