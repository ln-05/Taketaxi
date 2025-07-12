<template>
  <div class="order-page">
    <!-- 路线信息 -->
    <div class="route-info">
      <div class="location-item start">
        <el-icon><Location /></el-icon>
        <div class="location-text">
          <div class="label">起点</div>
          <div class="address">{{ startLocation }}</div>
        </div>
      </div>
      <div class="divider"></div>
      <div class="location-item end">
        <el-icon><LocationFilled /></el-icon>
        <div class="location-text">
          <div class="label">终点</div>
          <div class="address">{{ endLocation }}</div>
        </div>
      </div>
    </div>

    <!-- 车型选择 -->
    <div class="vehicle-selection">
      <div class="section-title">选择车型</div>
      <div class="vehicle-list">
        <div
          v-for="vehicle in vehicles"
          :key="vehicle.type"
          class="vehicle-item"
          :class="{ active: selectedVehicle === vehicle.type }"
          @click="selectVehicle(vehicle.type)"
        >
          <img :src="vehicle.icon" :alt="vehicle.name" class="vehicle-icon">
          <div class="vehicle-info">
            <div class="vehicle-name">{{ vehicle.name }}</div>
            <div class="vehicle-desc">{{ vehicle.description }}</div>
            <div class="vehicle-price">¥{{ vehicle.price }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- 支付方式 -->
    <div class="payment-method">
      <div class="section-title">支付方式</div>
      <div class="payment-list">
        <div
          v-for="method in paymentMethods"
          :key="method.type"
          class="payment-item"
          :class="{ active: selectedPayment === method.type }"
          @click="selectPayment(method.type)"
        >
          <img :src="method.icon" :alt="method.name" class="payment-icon">
          <div class="payment-name">{{ method.name }}</div>
          <el-icon v-if="selectedPayment === method.type"><Check /></el-icon>
        </div>
      </div>
    </div>

    <!-- 优惠券 -->
    <div class="coupon-section" @click="showCouponList">
      <div class="section-title">优惠券</div>
      <div class="coupon-info">
        <span>{{ selectedCoupon ? `已选择：${selectedCoupon.name}` : '暂无可用优惠券' }}</span>
        <el-icon><ArrowRight /></el-icon>
      </div>
    </div>

    <!-- 底部确认栏 -->
    <div class="bottom-bar">
      <div class="price-info">
        <div class="total-price">¥{{ totalPrice }}</div>
        <div class="price-detail" @click="showPriceDetail">
          明细 <el-icon><ArrowRight /></el-icon>
        </div>
      </div>
      <el-button type="primary" class="confirm-button" @click="confirmOrder">
        确认用车
      </el-button>
    </div>

    <!-- 价格明细弹窗 -->
    <el-dialog
      v-model="priceDetailVisible"
      title="费用明细"
      width="90%"
      custom-class="price-detail-dialog"
    >
      <div class="price-detail-list">
        <div class="price-item">
          <span>起步价</span>
          <span>¥{{ basePrice }}</span>
        </div>
        <div class="price-item">
          <span>里程费</span>
          <span>¥{{ distancePrice }}</span>
        </div>
        <div class="price-item">
          <span>时长费</span>
          <span>¥{{ timePrice }}</span>
        </div>
        <div class="price-item discount" v-if="discountAmount > 0">
          <span>优惠券</span>
          <span>-¥{{ discountAmount }}</span>
        </div>
        <div class="price-item total">
          <span>总计</span>
          <span>¥{{ totalPrice }}</span>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import {
  Location,
  LocationFilled,
  Check,
  ArrowRight
} from '@element-plus/icons-vue'

const router = useRouter()

// 路线信息
const startLocation = ref('上海市浦东新区陆家嘴')
const endLocation = ref('上海市黄浦区外滩')

// 车型选择
const vehicles = ref([
  {
    type: 'economy',
    name: '经济型',
    description: '舒适代步 省心出行',
    price: '23.00',
    icon: '/icons/economy.png'
  },
  {
    type: 'comfort',
    name: '舒适型',
    description: '舒适空间 品质出行',
    price: '32.00',
    icon: '/icons/comfort.png'
  },
  {
    type: 'luxury',
    name: '豪华型',
    description: '高端享受 尊贵体验',
    price: '45.00',
    icon: '/icons/luxury.png'
  }
])
const selectedVehicle = ref('economy')

// 支付方式
const paymentMethods = ref([
  {
    type: 'balance',
    name: '余额支付',
    icon: '/icons/balance.png'
  },
  {
    type: 'wechat',
    name: '微信支付',
    icon: '/icons/wechat.png'
  },
  {
    type: 'alipay',
    name: '支付宝',
    icon: '/icons/alipay.png'
  }
])
const selectedPayment = ref('balance')

interface Coupon {
  id: number
  name: string
  type: 'discount' | 'amount'  // discount: 折扣券, amount: 满减券
  value: number               // 折扣券：折扣率（如 85 表示 85折）, 满减券：减免金额
  minAmount: number          // 最低使用金额
  expireDate: string        // 过期时间
}

// 优惠券
const selectedCoupon = ref<Coupon | null>(null)

// 价格计算
const basePrice = ref(10)
const distancePrice = ref(13)
const timePrice = ref(5)
const discountAmount = ref(5)

const totalPrice = computed(() => {
  const total = basePrice.value + distancePrice.value + timePrice.value
  return (total - discountAmount.value).toFixed(2)
})

// 价格明细弹窗
const priceDetailVisible = ref(false)

// 方法
const selectVehicle = (type: string) => {
  selectedVehicle.value = type
}

const selectPayment = (type: string) => {
  selectedPayment.value = type
}

const showCouponList = () => {
  ElMessage.info('优惠券功能开发中')
}

const showPriceDetail = () => {
  priceDetailVisible.value = true
}

const confirmOrder = async () => {
  try {
    ElMessage.success('订单创建成功')
    router.push('/order/detail')
  } catch (error) {
    console.error('创建订单失败:', error)
    ElMessage.error('创建订单失败，请重试')
  }
}
</script>

<style lang="scss" scoped>
.order-page {
  min-height: 100vh;
  background: #f5f7fa;
  padding: 16px;
  padding-bottom: 80px;
}

.route-info {
  background: #fff;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 16px;

  .location-item {
    display: flex;
    align-items: center;
    gap: 12px;

    .el-icon {
      font-size: 20px;
      color: #1890ff;
    }

    .location-text {
      flex: 1;

      .label {
        font-size: 12px;
        color: #666;
      }

      .address {
        font-size: 14px;
        color: #333;
        margin-top: 4px;
      }
    }
  }

  .divider {
    width: 1px;
    height: 24px;
    background: #e8e8e8;
    margin: 8px 0 8px 10px;
  }
}

.vehicle-selection {
  background: #fff;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 16px;

  .section-title {
    font-size: 16px;
    font-weight: 500;
    margin-bottom: 12px;
  }

  .vehicle-list {
    display: flex;
    flex-direction: column;
    gap: 12px;
  }

  .vehicle-item {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 12px;
    border: 1px solid #e8e8e8;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.3s;

    &.active {
      border-color: #1890ff;
      background: #e6f7ff;
    }

    .vehicle-icon {
      width: 48px;
      height: 48px;
    }

    .vehicle-info {
      flex: 1;

      .vehicle-name {
        font-size: 16px;
        font-weight: 500;
      }

      .vehicle-desc {
        font-size: 12px;
        color: #666;
        margin-top: 4px;
      }

      .vehicle-price {
        font-size: 16px;
        color: #f5222d;
        margin-top: 4px;
      }
    }
  }
}

.payment-method {
  background: #fff;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 16px;

  .section-title {
    font-size: 16px;
    font-weight: 500;
    margin-bottom: 12px;
  }

  .payment-list {
    display: flex;
    gap: 12px;
  }

  .payment-item {
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 8px;
    padding: 12px;
    border: 1px solid #e8e8e8;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.3s;

    &.active {
      border-color: #1890ff;
      background: #e6f7ff;
    }

    .payment-icon {
      width: 32px;
      height: 32px;
    }

    .payment-name {
      font-size: 14px;
    }

    .el-icon {
      color: #1890ff;
    }
  }
}

.coupon-section {
  background: #fff;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 16px;
  cursor: pointer;

  .section-title {
    font-size: 16px;
    font-weight: 500;
    margin-bottom: 12px;
  }

  .coupon-info {
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: #666;

    .el-icon {
      color: #999;
    }
  }
}

.bottom-bar {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background: #fff;
  padding: 12px 16px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  box-shadow: 0 -2px 8px rgba(0, 0, 0, 0.05);

  .price-info {
    .total-price {
      font-size: 24px;
      font-weight: 500;
      color: #f5222d;
    }

    .price-detail {
      font-size: 12px;
      color: #666;
      display: flex;
      align-items: center;
      gap: 4px;
      cursor: pointer;
    }
  }

  .confirm-button {
    width: 120px;
    height: 40px;
    border-radius: 20px;
  }
}

.price-detail-dialog {
  .price-detail-list {
    .price-item {
      display: flex;
      justify-content: space-between;
      margin-bottom: 12px;
      font-size: 14px;

      &.discount {
        color: #52c41a;
      }

      &.total {
        font-size: 16px;
        font-weight: 500;
        margin-top: 16px;
        padding-top: 16px;
        border-top: 1px solid #f0f0f0;
      }
    }
  }
}
</style>
