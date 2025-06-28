<template>
  <div class="gva-container">
    <!-- 搜索及操作区域 -->
    <div class="gva-header">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="gva-search-form" @keyup.enter="onSubmit">
        <el-form-item label="产品名称">
          <el-input v-model="searchInfo.productName" placeholder="请输入产品名称" clearable />
        </el-form-item>
        <el-form-item label="所属品类">
          <el-select v-model="searchInfo.category" placeholder="全部" clearable>
            <el-option
                v-for="item in categoryOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="goToDetail(item.id)">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery = true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery = false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
      <div class="gva-header-actions">
        <el-button type="primary" icon="plus" @click="openDialog()">新增</el-button>
        <el-button type="primary" icon="delete" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
      </div>
    </div>

    <!-- 卡片列表区域 -->
    <div class="gva-cards-wrapper">
      <div
          v-for="(item, index) in tableData"
          :key="index"
          class="gva-device-card"
      >
        <!-- 新增：标题+数据 横向容器 -->
        <div class="card-content">
          <!-- 左侧内容：图标、标题、代码 -->
          <div class="card-left">
            <div class="card-icon">
              <el-image
                  :src="getIconUrl(item.pictrue)"
                  fit="cover"
                  style="width: 60px; height: 60px"
              />
            </div>
            <div class="card-title">
              <h3 class="device-name">{{ item.name }}</h3>
              <p class="device-code">{{ item.code || 'AIOTxxxx' }}</p>
            </div>
          </div>
          <!-- 右侧内容：数据信息 -->
          <div class="card-right">
            <div class="info-item">
              <span class="label">设备数量：</span>
              <span class="value">{{ item.count }}</span>
            </div>
            <div class="info-item">
              <span class="label">发布状态：</span>
              <span class="value">
                <span
                    class="status-indicator"
                    :class="getStatusClass(item.status)"
                ></span>
                {{ item.status }}
              </span>
            </div>
            <div class="info-item">
              <span class="label">设备类型：</span>
              <span class="value">{{ item.equStatus }}</span>
            </div>
          </div>
        </div>
        <!-- 底部操作按钮 -->
        <div class="card-actions">
          <el-button
              type="primary"
              icon="search"
              @click="getDetails(item)"
          >查看</el-button>
          <el-button
              type="primary"
              icon="edit"
              @click="updateGoodAccountFunc(item)"
          >编辑</el-button>
          <el-button
              type="primary"
              icon="delete"
              @click="deleteRow(item)"
          >删除</el-button>
          <el-button
              type="primary"
              icon="more"
              @click="handleMoreActions(item)"
          >更多</el-button>
        </div>
      </div>
    </div>

    <!-- 分页区域 -->
    <div class="gva-pagination">
      <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
      />
    </div>

    <!-- 新增/编辑抽屉 -->
    <el-drawer
        destroy-on-close
        :size="appStore.drawerSize"
        v-model="dialogFormVisible"
        :show-close="false"
        :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '新增' : '编辑' }}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>
      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="图片字段:" prop="pictrue">
          <SelectImage
              v-model="formData.pictrue"
              file-type="image"
          />
        </el-form-item>
        <el-form-item label="设备名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入设备名称" />
        </el-form-item>
        <el-form-item label="设备数量:" prop="count">
          <el-input v-model.number="formData.count" :clearable="true" placeholder="请输入设备数量" />
        </el-form-item>
        <el-form-item label="发布状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true" placeholder="请输入发布状态" />
        </el-form-item>
        <el-form-item label="设备类型:" prop="equStatus">
          <el-input v-model="formData.equStatus" :clearable="true" placeholder="请输入设备类型" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <!-- 详情抽屉 -->
    <el-drawer
        destroy-on-close
        :size="appStore.drawerSize"
        v-model="detailShow"
        :show-close="true"
        :before-close="closeDetailShow"
        title="查看"
    >
      <el-descriptions :column="1" border>
        <el-descriptions-item label="ID">
          {{ detailFrom.id }}
        </el-descriptions-item>
        <el-descriptions-item label="图片">
          <el-image
              :src="getIconUrl(detailFrom.pictrue)"
              fit="cover"
              style="width: 60px; height: 60px"
          />
        </el-descriptions-item>
        <el-descriptions-item label="设备名称">
          {{ detailFrom.name }}
        </el-descriptions-item>
        <el-descriptions-item label="设备数量">
          {{ detailFrom.count }}
        </el-descriptions-item>
        <el-descriptions-item label="发布状态">
          {{ detailFrom.status }}
        </el-descriptions-item>
        <el-descriptions-item label="设备类型">
          {{ detailFrom.equStatus }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
// 保持原有逻辑代码不变（JS部分无需修改）
import {
  createGoodAccount,
  deleteGoodAccount,
  deleteGoodAccountByIds,
  updateGoodAccount,
  findGoodAccount,
  getGoodAccountList
} from '@/api/good/goodAccount'
import { getUrl } from '@/utils/image'
import SelectImage from '@/components/selectImage/selectImage.vue'
import { returnArrImg } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, computed } from 'vue'
import { useAppStore } from '@/pinia'
import { InfoFilled } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router';

const router = useRouter();

const getDetails =  (row) => {
  // 新逻辑：跳转到详情页面
  router.push({
    name: 'addressdateil',
    params: { id: row.id }
  });
};

defineOptions({
  name: 'GoodAccount'
})

const btnLoading = ref(false)
const appStore = useAppStore()
const showAllQuery = ref(false)
const formData = ref({
  pictrue: '',
  name: '',
  count: undefined,
  status: '',
  equStatus: ''
})
const rule = reactive({})
const elFormRef = ref()
const elSearchFormRef = ref()

// 表格控制
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const getTableData = async () => {
  const table = await getGoodAccountList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData()

// 字典/下拉选项
const categoryOptions = ref([
  { label: '全部', value: '' },
  { label: '直联设备', value: '直联设备' },
  { label: '网关设备', value: '网关设备' }
])


// 设备图标映射（可根据实际类型扩展）
const getIconUrl = computed(() => (pictrue) => {
  return getUrl(pictrue)
})

// 多选相关
const multipleSelection = ref([])
// 删除
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteGoodAccountFunc(row)
  })
}

const onDelete = async () => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const ids = multipleSelection.value.map(item => item.id)
    if (ids.length === 0) {
      ElMessage.warning('请选择要删除的数据')
      return
    }
    const res = await deleteGoodAccountByIds({ ids })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      if (tableData.value.length === ids.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 弹窗相关
const type = ref('')
const dialogFormVisible = ref(false)

const updateGoodAccountFunc = async (row) => {
  const res = await findGoodAccount({ id: row.id })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}

const deleteGoodAccountFunc = async (row) => {
  const res = await deleteGoodAccount({ id: row.id })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    pictrue: '',
    name: '',
    count: undefined,
    status: '',
    equStatus: ''
  }
}

const enterDialog = async () => {
  btnLoading.value = true
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return (btnLoading.value = false)
    let res
    if (type.value === 'create') {
      res = await createGoodAccount(formData.value)
    } else {
      res = await updateGoodAccount(formData.value)
    }
    btnLoading.value = false
    if (res.code === 0) {
      ElMessage.success('创建/更改成功')
      closeDialog()
      getTableData()
    }
  })
}

// 详情弹窗
const detailFrom = ref({})
const detailShow = ref(false)

const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}

// 状态样式
const getStatusClass = (status) => {
  return status === '已发布' ? 'status-online' : 'status-offline'
}

// 更多操作（示例，需补充实际逻辑）
const handleMoreActions = (item) => {
  console.log('更多操作', item)
  ElMessage.info('更多操作功能待实现')
}
</script>

<style scoped>
/* 全局容器 */
.gva-container {
  width: 100%;
  min-height: calc(100vh - 64px);
  padding: 20px;
  box-sizing: border-box;
  background-color: #f5f7fa;
}

/* 头部搜索+操作 */
.gva-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #fff;
  padding: 16px 24px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  margin-bottom: 24px;
}

.gva-search-form {
  display: flex;
  align-items: center;
  gap: 16px;
}

.gva-header-actions {
  display: flex;
  gap: 12px;
}

/* 卡片列表 */
.gva-cards-wrapper {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(323px, 1fr));
  grid-gap: 20px;
}

/* 卡片容器：改为垂直方向布局，控制内部元素排列 */
.gva-device-card {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.03);
  padding: 20px;
  display: flex;
  flex-direction: column; /* 垂直排列子元素 */
  transition: all 0.2s ease;
}

.gva-device-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  transform: translateY(-2px);
}

/* 标题+数据 容器：横向排列 */
.card-content {
  display: flex;
  justify-content: space-between; /* 标题居左、数据居右 */
  margin-bottom: 16px; /* 与按钮保持间距 */
}

/* 卡片左侧：图标、标题、代码 */
.card-left {
  display: flex;
  flex-direction: column;
  align-items: flex-start; /* 标题、图标居左 */
}

.card-icon {
  margin-bottom: 12px;
}

.device-name {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin: 0 0 4px 0;
}

.device-code {
  font-size: 14px;
  color: #999;
  margin: 0;
}

/* 卡片右侧：数据信息 */
.card-right {
  display: flex;
  flex-direction: column;
  align-items: flex-end; /* 数据信息居右 */
}

.info-item {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  margin-bottom: 8px;
  font-size: 14px;
  color: #666;
}

.info-item .value {
  color: #333;
  font-weight: 500;
}

/* 底部操作按钮：自动顶到底部 */
.card-actions {
  margin-top: auto; /* 关键：把按钮推到最底部 */
  display: flex;
  gap: 8px; /* 按钮之间间距 */
  justify-content: flex-start; /* 按钮居左，可根据需求改 flex-end */
}

/* 状态标签 */
.status-indicator {
  display: inline-block;
  width: 10px;
  height: 10px;
  border-radius: 50%;
  margin-right: 6px;
  vertical-align: middle;
}

.status-online {
  background-color: #10b981; /* 绿色：在线 */
}

.status-offline {
  background-color: #ef4444; /* 红色：离线 */
}

.status-disabled {
  background-color: #9ca3af; /* 灰色：未启用 */
}
</style>