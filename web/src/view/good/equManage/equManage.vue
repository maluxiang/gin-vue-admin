<template>
  <div class="device-management-container p-4 md:p-6 max-w-7xl mx-auto">
    <!-- 设备统计卡片 -->
    <div class="stats-card bg-white rounded-xl shadow-sm border border-gray-100 mb-6 overflow-hidden">
      <div class="stats-header bg-gradient-to-r from-blue-500 to-blue-600 px-6 py-4">
        <div class="flex items-center">
          <i class="fa fa-bar-chart text-white text-xl mr-3"></i>
          <h3 class="text-white font-semibold text-lg">设备统计</h3>
        </div>
      </div>

      <div class="stats-body p-6">
        <!-- 加载状态 -->
        <div v-if="statsLoading" class="flex justify-center items-center py-8">
          <el-loading v-loading="true" text="加载中..."></el-loading>
        </div>

        <!-- 错误状态 -->
        <div v-else-if="statsError" class="flex flex-col items-center justify-center py-8 text-center">
          <i class="fa fa-exclamation-circle text-red-500 text-3xl mb-3"></i>
          <p class="text-gray-600 mb-2">获取统计数据失败</p>
          <el-button size="small" type="primary" @click="retryFetchStats">重试</el-button>
        </div>

        <!-- 正常状态 -->
        <div v-else class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <!-- 全部设备 -->
          <div
              class="stat-item p-4 rounded-lg border border-gray-100 hover:border-blue-200 transition-all cursor-pointer"
              :class="{ 'bg-blue-50 border-blue-200': activeFilter === 'all' }"
              @click="filterDevices('all')"
          >
            <div class="text-gray-500 text-sm font-medium mb-1">全部设备</div>
            <div class="text-2xl font-bold text-gray-800">{{ deviceStats.total || 0 }}</div>
          </div>

          <!-- 在线设备 -->
          <div
              class="stat-item p-4 rounded-lg border border-gray-100 hover:border-green-200 transition-all cursor-pointer"
              :class="{ 'bg-green-50 border-green-200': activeFilter === 'online' }"
              @click="filterDevices('online')"
          >
            <div class="text-gray-500 text-sm font-medium mb-1">在线</div>
            <div class="text-2xl font-bold text-green-600">{{ deviceStats.online || 0 }}</div>
          </div>

          <!-- 离线设备 -->
          <div
              class="stat-item p-4 rounded-lg border border-gray-100 hover:border-red-200 transition-all cursor-pointer"
              :class="{ 'bg-red-50 border-red-200': activeFilter === 'offline' }"
              @click="filterDevices('offline')"
          >
            <div class="text-gray-500 text-sm font-medium mb-1">离线</div>
            <div class="text-2xl font-bold text-red-600">{{ deviceStats.offline || 0 }}</div>
          </div>

          <!-- 未启用设备 -->
          <div
              class="stat-item p-4 rounded-lg border border-gray-100 hover:border-gray-300 transition-all cursor-pointer"
              :class="{ 'bg-gray-50 border-gray-300': activeFilter === 'disabled' }"
              @click="filterDevices('disabled')"
          >
            <div class="text-gray-500 text-sm font-medium mb-1">未启用</div>
            <div class="text-2xl font-bold text-gray-600">{{ deviceStats.disabled || 0 }}</div>
          </div>
        </div>

        <div class="mt-4 text-sm text-gray-500">
          数据更新时间: {{ deviceStats.updateTime || 'N/A' }}
        </div>
      </div>
    </div>

    <!-- 搜索框 -->
    <div class="search-box bg-white rounded-xl shadow-sm border border-gray-100 p-4 mb-6">
      <el-form
          ref="elSearchFormRef"
          :inline="true"
          :model="searchInfo"
          class="demo-form-inline"
          @keyup.enter="onSubmit"
      >
        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button
              link
              type="primary"
              icon="arrow-down"
              @click="showAllQuery = true"
              v-if="!showAllQuery"
          >
            展开
          </el-button>
          <el-button
              link
              type="primary"
              icon="arrow-up"
              @click="showAllQuery = false"
              v-else
          >
            收起
          </el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 设备表格 -->
    <div class="device-table bg-white rounded-xl shadow-sm border border-gray-100 overflow-hidden">
      <!-- 表格操作按钮 -->
      <div class="p-4 border-b border-gray-100 flex flex-wrap items-center justify-between gap-2">
        <el-button type="primary" icon="plus" @click="openDialog()">新增</el-button>
        <el-button
            icon="delete"
            style="margin-left: 0;"
            :disabled="!multipleSelection.length"
            @click="onDelete"
        >
          删除
        </el-button>
      </div>

      <!-- 表格内容 -->
      <div class="overflow-x-auto">
        <!-- 加载状态 -->
        <div v-if="tableLoading" class="flex justify-center items-center py-16">
          <el-loading v-loading="true" text="加载中..."></el-loading>
        </div>

        <!-- 错误状态 -->
        <div v-else-if="tableError" class="flex flex-col items-center justify-center py-16 text-center">
          <i class="fa fa-exclamation-circle text-red-500 text-3xl mb-3"></i>
          <p class="text-gray-600 mb-2">获取表格数据失败</p>
          <el-button size="small" type="primary" @click="retryFetchTable">重试</el-button>
        </div>

        <!-- 空数据状态 -->
        <div v-else-if="tableData.length === 0" class="flex flex-col items-center justify-center py-16 text-center">
          <i class="fa fa-inbox text-gray-300 text-5xl mb-4"></i>
          <p class="text-gray-500 mb-2">暂无数据</p>
          <p class="text-gray-400 text-sm">请尝试其他筛选条件或添加数据</p>
        </div>

        <!-- 正常状态 -->
        <div v-else>
          <el-table
              ref="multipleTable"
              style="width: 100%"
              tooltip-effect="dark"
              :data="tableData"
              row-key="id"
              @selection-change="handleSelectionChange"
              border
          >
            <el-table-column type="selection" width="55" />
            <el-table-column align="left" label="设备id" prop="equId" width="200" />
            <el-table-column align="left" label="设备名称" prop="equName" width="200" />
            <el-table-column align="left" label="设备分类" prop="equClass" width="200" />
            <el-table-column align="left" label="设备状态" prop="equStatus" width="200">
              <template #default="scope">
                <span
                    class="status-indicator inline-block w-2 h-2 rounded-full mr-2"
                    :class="getStatusClass(scope.row.equStatus)"
                ></span>
                {{ scope.row.equStatus }}
              </template>
            </el-table-column>
            <el-table-column align="left" label="备注说明" prop="remark" width="200" />
            <el-table-column
                align="left"
                label="注册时间"
                prop="createdAt"
                width="180"
            >
              <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
            </el-table-column>
            <el-table-column
                align="left"
                label="操作"
                fixed="right"
                :min-width="appStore.operateMinWith"
            >
              <template #default="scope">
                <el-button
                    type="primary"
                    link
                    class="table-button"
                    @click="getDetails(scope.row)"
                >
                  <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看
                </el-button>
                <el-button
                    type="primary"
                    link
                    icon="edit"
                    class="table-button"
                    @click="updateEquManageFunc(scope.row)"
                >
                  编辑
                </el-button>
                <el-button
                    type="primary"
                    link
                    icon="delete"
                    @click="deleteRow(scope.row)"
                >
                  删除
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>

      <!-- 分页 -->
      <div v-if="tableData.length > 0" class="p-4 border-t border-gray-100 flex justify-between items-center">
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
    </div>

    <!-- 新增/编辑对话框 -->
    <el-drawer
        destroy-on-close
        :size="appStore.drawerSize"
        v-model="dialogFormVisible"
        :show-close="false"
        :before-close="closeDialog"
    >
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg font-medium">{{ type === 'create' ? '新增设备' : '编辑设备' }}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">
              确 定
            </el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form
          :model="formData"
          label-position="top"
          ref="elFormRef"
          :rules="rule"
          label-width="80px"
          class="mt-4"
      >
        <el-form-item label="设备ID:" prop="equId">
          <el-input
              v-model="formData.equId"
              :clearable="true"
              placeholder="请输入设备ID"
              class="w-full"
          />
        </el-form-item>
        <el-form-item label="设备名称:" prop="equName">
          <el-input
              v-model="formData.equName"
              :clearable="true"
              placeholder="请输入设备名称"
              class="w-full"
          />
        </el-form-item>
        <el-form-item label="设备分类:" prop="equClass">
          <el-input
              v-model="formData.equClass"
              :clearable="true"
              placeholder="请输入设备分类"
              class="w-full"
          />
        </el-form-item>
        <el-form-item label="设备状态:" prop="equStatus">
          <el-select v-model="formData.equStatus" placeholder="请选择设备状态">
            <el-option label="在线" value="在线"></el-option>
            <el-option label="离线" value="离线"></el-option>
            <el-option label="未启用" value="未启用"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="备注说明:" prop="remark">
          <el-input
              v-model="formData.remark"
              :clearable="true"
              placeholder="请输入备注说明"
              class="w-full"
          />
        </el-form-item>
      </el-form>
    </el-drawer>

    <!-- 查看详情对话框 -->
    <el-drawer
        destroy-on-close
        :size="appStore.drawerSize"
        v-model="detailShow"
        :show-close="true"
        :before-close="closeDetailShow"
        title="查看设备详情"
    >
      <el-descriptions :column="1" border class="mt-4">
        <el-descriptions-item label="ID" span="2">{{ detailFrom.id }}</el-descriptions-item>
        <el-descriptions-item label="设备ID" span="2">{{ detailFrom.equId }}</el-descriptions-item>
        <el-descriptions-item label="设备名称" span="2">{{ detailFrom.equName }}</el-descriptions-item>
        <el-descriptions-item label="设备分类" span="2">{{ detailFrom.equClass }}</el-descriptions-item>
        <el-descriptions-item label="设备状态" span="2">
          <span
              class="status-indicator inline-block w-2 h-2 rounded-full mr-2"
              :class="getStatusClass(detailFrom.equStatus)"
          ></span>
          {{ detailFrom.equStatus }}
        </el-descriptions-item>
        <el-descriptions-item label="备注说明" span="2">{{ detailFrom.remark }}</el-descriptions-item>
        <el-descriptions-item label="注册时间" span="2">{{ detailFrom.createdAt }}</el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createEquManage,
  deleteEquManage,
  deleteEquManageByIds,
  updateEquManage,
  findEquManage,
  getEquManageList,
  counts
} from '@/api/good/equManage'
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, onMounted } from 'vue'
import { useAppStore } from '@/pinia'
import { InfoFilled } from '@element-plus/icons-vue'

// 定义响应式数据来存储设备统计信息
const deviceStats = ref({
  total: 0,
  online: 0,
  offline: 0,
  disabled: 0,
  updateTime: ''
})

// 状态管理
const statsLoading = ref(false)
const statsError = ref(false)
const tableLoading = ref(false)
const tableError = ref(false)

// 定义获取设备统计数据的函数
const fetchDeviceStats = async () => {
  statsLoading.value = true
  statsError.value = false

  try {
    // 模拟网络请求延迟
    await new Promise(resolve => setTimeout(resolve, 500))

    const response = await counts()

    // 检查API响应状态
    if (response.code === 0) {
      deviceStats.value = response.data
    } else {
      throw new Error(response.message || '获取统计数据失败')
    }
  } catch (error) {
    console.error('获取设备统计数据时出错:', error)
    statsError.value = true
    ElMessage.error('获取统计数据失败，请稍后重试')
  } finally {
    statsLoading.value = false
  }
}

// 重试获取统计数据
const retryFetchStats = () => {
  fetchDeviceStats()
}

defineOptions({
  name: 'EquManage'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  equId: '',
  equName: '',
  equClass: '',
  equStatus: '',
  remark: ''
})

// 验证规则
const rule = reactive({
  equId: [
    { required: true, message: '请输入设备ID', trigger: 'blur' }
  ],
  equName: [
    { required: true, message: '请输入设备名称', trigger: 'blur' }
  ],
  equStatus: [
    { required: true, message: '请选择设备状态', trigger: 'change' }
  ]
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 筛选状态，默认 'all'，可选 'online' 'offline' 'disabled'
const activeFilter = ref('all')

// 处理筛选
const filterDevices = (filterType) => {
  // 点击相同筛选状态时不刷新，提升性能
  if (activeFilter.value === filterType) return

  activeFilter.value = filterType // 更新筛选状态
  page.value = 1 // 重置页码为1
  getTableData() // 重新获取数据
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  filterDevices('all') // 重置筛选状态为全部
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1 // 搜索时重置页码为1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 根据筛选状态获取表格数据
const getTableData = async () => {
  tableLoading.value = true
  tableError.value = false

  try {
    // 模拟网络请求延迟
    await new Promise(resolve => setTimeout(resolve, 500))

    let params = {
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    }

    // 根据筛选状态添加设备状态参数
    if (activeFilter.value === 'online') {
      params.equStatus = '在线'
    } else if (activeFilter.value === 'offline') {
      params.equStatus = '离线'
    } else if (activeFilter.value === 'disabled') {
      params.equStatus = '未启用'
    } else {
      // 全部设备，不传递equStatus参数或清空它
      delete params.equStatus
    }

    const table = await getEquManageList(params)

    // 检查API响应状态
    if (table.code === 0) {
      tableData.value = table.data.list || []
      total.value = table.data.total || 0
      page.value = table.data.page || 1
      pageSize.value = table.data.pageSize || 10
    } else {
      throw new Error(table.message || '获取表格数据失败')
    }
  } catch (error) {
    console.error('获取表格数据时出错:', error)
    tableError.value = true
    ElMessage.error('获取表格数据失败，请稍后重试')
  } finally {
    tableLoading.value = false
  }
}

// 重试获取表格数据
const retryFetchTable = () => {
  getTableData()
}

// 初始化获取表格数据
onMounted(() => {
  // 同时获取统计数据和表格数据
  Promise.all([
    fetchDeviceStats(),
    getTableData()
  ]).catch(error => {
    console.error('初始化数据获取失败:', error)
  })
})

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {}
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除该设备吗?', '删除确认', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteEquManageFunc(row)
  })
}

// 多选删除
const onDelete = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要删除的设备')
    return
  }

  ElMessageBox.confirm(
      `确定要删除选中的 ${multipleSelection.value.length} 个设备吗?`,
      '批量删除确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
  ).then(async () => {
    try {
      const ids = multipleSelection.value.map(item => item.id)
      const res = await deleteEquManageByIds({ ids })

      if (res.code === 0) {
        ElMessage.success('删除成功')

        // 如果当前页只有一个设备且被删除，跳转到上一页
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }

        // 刷新表格和统计数据
        await Promise.all([
          getTableData(),
          fetchDeviceStats()
        ])
      } else {
        throw new Error(res.message || '删除失败')
      }
    } catch (error) {
      console.error('删除设备时出错:', error)
      ElMessage.error('删除失败，请稍后重试')
    }
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateEquManageFunc = async (row) => {
  try {
    const res = await findEquManage({ id: row.id })
    if (res.code === 0) {
      formData.value = { ...res.data } // 深拷贝数据
      type.value = 'update'
      dialogFormVisible.value = true
    } else {
      throw new Error(res.message || '获取设备信息失败')
    }
  } catch (error) {
    console.error('获取设备信息时出错:', error)
    ElMessage.error('获取设备信息失败，请稍后重试')
  }
}

// 删除行
const deleteEquManageFunc = async (row) => {
  try {
    const res = await deleteEquManage({ id: row.id })
    if (res.code === 0) {
      ElMessage.success('删除成功')

      // 如果当前页只有一个设备且被删除，跳转到上一页
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }

      // 刷新表格和统计数据
      await Promise.all([
        getTableData(),
        fetchDeviceStats()
      ])
    } else {
      throw new Error(res.message || '删除失败')
    }
  } catch (error) {
    console.error('删除设备时出错:', error)
    ElMessage.error('删除失败，请稍后重试')
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  // 重置表单数据
  formData.value = {
    equId: '',
    equName: '',
    equClass: '',
    equStatus: '',
    remark: ''
  }
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  // 重置表单数据
  formData.value = {
    equId: '',
    equName: '',
    equClass: '',
    equStatus: '',
    remark: ''
  }
}

// 弹窗确定
const enterDialog = async () => {
  btnLoading.value = true
  elFormRef.value?.validate(async (valid) => {
    if (!valid) {
      btnLoading.value = false
      return
    }

    try {
      let res
      switch (type.value) {
        case 'create':
          res = await createEquManage(formData.value)
          break
        case 'update':
          res = await updateEquManage(formData.value)
          break
        default:
          res = await createEquManage(formData.value)
          break
      }

      btnLoading.value = false

      if (res.code === 0) {
        ElMessage.success(`${type.value === 'create' ? '创建' : '更新'}成功`)
        closeDialog()

        // 刷新表格和统计数据
        await Promise.all([
          getTableData(),
          fetchDeviceStats()
        ])
      } else {
        throw new Error(res.message || `${type.value === 'create' ? '创建' : '更新'}失败`)
      }
    } catch (error) {
      btnLoading.value = false
      console.error(`${type.value === 'create' ? '创建' : '更新'}设备时出错:`, error)
      ElMessage.error(`${type.value === 'create' ? '创建' : '更新'}失败，请稍后重试`)
    }
  })
}

const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async (row) => {
  try {
    const res = await findEquManage({ id: row.id })
    if (res.code === 0) {
      detailFrom.value = { ...res.data } // 深拷贝数据
      openDetailShow()
    } else {
      throw new Error(res.message || '获取设备详情失败')
    }
  } catch (error) {
    console.error('获取设备详情时出错:', error)
    ElMessage.error('获取设备详情失败，请稍后重试')
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}

// 获取状态样式
const getStatusClass = (status) => {
  switch (status) {
    case '在线':
      return 'bg-green-500'
    case '离线':
      return 'bg-red-500'
    case '未启用':
      return 'bg-gray-500'
    default:
      return ''
  }
}
</script>

<style scoped>
/* 设备状态指示器样式 */
.status-indicator {
  display: inline-block;
  width: 10px;
  height: 10px;
  border-radius: 50%;
  margin-right: 6px;
  vertical-align: middle;
}

/* 统计卡片样式 */
.stats-card {
  transition: all 0.3s ease;
}

.stats-card:hover {
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05);
}

/* 统计项样式 */
.stat-item {
  transition: all 0.2s ease;
}

/* 表单样式优化 */
.el-form-item {
  margin-bottom: 16px;
}

/* 表格操作按钮样式 */
.table-button {
  padding: 0 8px;
  height: 28px;
  line-height: 28px;
}

/* 分页样式 */
.gva-pagination {
  padding: 12px 0;
}
</style>