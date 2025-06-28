<template>
  <div class="product-detail-container">
    <!-- 页面头部 -->
    <div class="detail-header">
      <h2>产品详情</h2>
      <div class="action-buttons">
        <el-button @click="returnList()">返回</el-button>
      </div>
    </div>

    <!-- 主选项卡：产品信息、物模型、告警设置 -->
    <el-tabs v-model="activeMainTab" class="main-tabs">
      <!-- 1. 产品信息详情 -->
      <el-tab-pane label="产品信息" name="product-info">
        <br>
        <el-button type="primary">编辑</el-button>
        <el-descriptions :column="2" border class="info-descriptions">
          <el-descriptions-item label="产品ID">{{ productDetailFrom.id }}</el-descriptions-item>
          <el-descriptions-item label="产品名称">{{ productDetailFrom.name }}</el-descriptions-item>
          <el-descriptions-item label="产品型号">{{ productDetailFrom.model }}</el-descriptions-item>
          <el-descriptions-item label="产品状态">{{ productDetailFrom.num }}</el-descriptions-item>
          <el-descriptions-item label="创建时间">{{ productDetailFrom.status }}</el-descriptions-item>
          <el-descriptions-item label="更新时间">{{ productDetailFrom.kind }}</el-descriptions-item>
          <el-descriptions-item label="所属项目">{{ productDetailFrom.createdAt }}</el-descriptions-item>
          <el-descriptions-item label="制造商">{{ productDetailFrom.updatedAt }}</el-descriptions-item>
        </el-descriptions>
      </el-tab-pane>

      <!-- 2. 物模型（嵌套选项卡） -->
      <el-tab-pane label="物模型" name="thing-model">
        <el-tabs v-model="activeModelTab" class="nested-tabs">

          <el-tab-pane label="属性" name="properties">
            <div class="gva-search-box">
              <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
                <el-form-item label="属性名称" prop="attributeName">
                  <el-input v-model="searchInfo.attributeName" placeholder="搜索条件" />
                </el-form-item>

                <template v-if="showAllQuery">
                  <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
                </template>

                <el-form-item>
                  <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
                  <el-button icon="refresh" @click="onReset">重置</el-button>
                  <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
                  <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
                </el-form-item>
              </el-form>
            </div>
            <div class="gva-table-box">
              <div class="gva-btn-list">
                <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
                <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
              </div>
              <el-table
                  ref="multipleTable"
                  style="width: 100%"
                  tooltip-effect="dark"
                  :data="tableData"
                  row-key="id"
                  @selection-change="handleSelectionChange"
              >
<!--                <el-table-column type="selection" width="55" />-->
<!--                <el-table-column align="left" label="序号" prop="id" width="150" />-->
<!--                <el-table-column align="left" label="属性标识" prop="attributeId" width="150" />-->
<!--                <el-table-column align="left" label="属性名称" prop="attributeName" width="150" />-->
<!--                <el-table-column align="left" label="数据类型" prop="dataType" width="150" />-->
<!--                <el-table-column align="left" label="属性值来源" prop="attributeFrom" width="150" />-->
<!--                <el-table-column align="left" label="是否只读" prop="read" width="150" />-->
<!--                <el-table-column align="left" label="其他说明" prop="description" width="150" />-->
<!--                <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">-->
<!--                  <template #default="scope">-->
<!--                    &lt;!&ndash;                    <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>&ndash;&gt;-->
<!--                    <el-button  type="primary" link icon="edit" class="table-button" @click="updateDeviceManageModelAttributeFunc(scope.row)">编辑</el-button>-->
<!--                    <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>-->
<!--                  </template>-->
<!--                </el-table-column>-->
              </el-table>
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
            </div>
            <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
              <template #header>
                <div class="flex justify-between items-center">
                  <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                  <div>
                    <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                    <el-button @click="closeDialog">取 消</el-button>
                  </div>
                </div>
              </template>

              <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
                <el-form-item label="属性标识:" prop="attributeId">
                  <el-input v-model="formData.attributeId" :clearable="true" placeholder="请输入属性标识" />
                </el-form-item>
                <el-form-item label="属性名称:" prop="attributeName">
                  <el-input v-model="formData.attributeName" :clearable="true" placeholder="请输入属性名称" />
                </el-form-item>
                <el-form-item label="数据类型:" prop="dataType">
                  <el-input v-model="formData.dataType" :clearable="true" placeholder="请输入数据类型" />
                </el-form-item>
                <el-form-item label="属性值来源:" prop="attributeFrom">
                  <el-input v-model="formData.attributeFrom" :clearable="true" placeholder="请输入属性值来源" />
                </el-form-item>
                <el-form-item label="是否只读:" prop="read">
                  <el-input v-model="formData.read" :clearable="true" placeholder="请输入是否只读" />
                </el-form-item>
                <el-form-item label="其他说明:" prop="description">
                  <el-input v-model="formData.description" :clearable="true" placeholder="请输入其他说明" />
                </el-form-item>
              </el-form>
            </el-drawer>

            <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
              <el-descriptions :column="1" border>
                <el-descriptions-item label="序号">
                  {{ detailFrom.id }}
                </el-descriptions-item>
                <el-descriptions-item label="产品管理ID">
                  {{ detailFrom.deviceManageId }}
                </el-descriptions-item>
                <el-descriptions-item label="属性标识">
                  {{ detailFrom.attributeId }}
                </el-descriptions-item>
                <el-descriptions-item label="属性名称">
                  {{ detailFrom.attributeName }}
                </el-descriptions-item>
                <el-descriptions-item label="数据类型">
                  {{ detailFrom.dataType }}
                </el-descriptions-item>
                <el-descriptions-item label="属性值来源">
                  {{ detailFrom.attributeFrom }}
                </el-descriptions-item>
                <el-descriptions-item label="是否只读">
                  {{ detailFrom.read }}
                </el-descriptions-item>
                <el-descriptions-item label="其他说明">
                  {{ detailFrom.description }}
                </el-descriptions-item>
                <el-descriptions-item label="创建时间">
                  {{ detailFrom.createdAt }}
                </el-descriptions-item>
              </el-descriptions>
            </el-drawer>
          </el-tab-pane>

          <el-tab-pane label="功能" name="functions">
            <div class="model-list">
              <el-table :data="[]" stripe style="width: 100%">
                <el-table-column prop="name" label="功能名称" width="180"></el-table-column>
                <el-table-column prop="type" label="功能类型" width="120"></el-table-column>
                <el-table-column prop="inputParams" label="输入参数"></el-table-column>
                <el-table-column prop="outputParams" label="输出参数"></el-table-column>
              </el-table>
            </div>
          </el-tab-pane>

          <el-tab-pane label="事件" name="events">
            <div class="model-list">
              <el-table :data="[]" stripe style="width: 100%">
                <el-table-column prop="name" label="事件名称" width="180"></el-table-column>
                <el-table-column prop="type" label="事件类型" width="120"></el-table-column>
                <el-table-column prop="params" label="参数"></el-table-column>
                <el-table-column prop="description" label="描述"></el-table-column>
              </el-table>
            </div>
          </el-tab-pane>

          <el-tab-pane label="标签定义" name="tags">
            <div class="model-list">
              <el-table :data="[]" stripe style="width: 100%">
                <el-table-column prop="name" label="标签名称" width="180"></el-table-column>
                <el-table-column prop="value" label="标签值" width="120"></el-table-column>
                <el-table-column prop="description" label="描述"></el-table-column>
              </el-table>
            </div>
          </el-tab-pane>
        </el-tabs>
      </el-tab-pane>

      <!-- 3. 告警设置（嵌套选项卡） -->
      <el-tab-pane label="告警设置" name="alarm-setting">
        <el-tabs v-model="activeAlarmTab" class="nested-tabs">

          <el-tab-pane label="告警规则" name="alarm-rules">
            <div class="alarm-setting">
              <el-table :data="[]" stripe style="width: 100%">
                <el-table-column prop="ruleName" label="规则名称" width="180"></el-table-column>
                <el-table-column prop="monitorParam" label="监控参数" width="180"></el-table-column>
                <el-table-column prop="condition" label="触发条件" width="200"></el-table-column>
                <el-table-column prop="level" label="告警级别" width="100"></el-table-column>
                <el-table-column prop="status" label="状态" width="100"></el-table-column>
              </el-table>
            </div>
          </el-tab-pane>

          <el-tab-pane label="告警记录" name="alarm-logs">
            <div class="alarm-logs">
              <el-table :data="[]" stripe style="width: 100%">
                <el-table-column prop="logTime" label="告警时间" width="180"></el-table-column>
                <el-table-column prop="deviceName" label="设备名称" width="180"></el-table-column>
                <el-table-column prop="alarmInfo" label="告警信息"></el-table-column>
                <el-table-column prop="level" label="告警级别" width="100"></el-table-column>
                <el-table-column prop="status" label="处理状态" width="120"></el-table-column>
              </el-table>
            </div>
          </el-tab-pane>

        </el-tabs>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup>
import {
  createGoodAccount,
  deleteGoodAccount,
  deleteGoodAccountByIds,
  findGoodAccount, getGoodAccountList, updateGoodAccount
} from '@/api/good/goodAccount'
import {

} from '@/api/good/goodAccount'
import {onMounted, ref} from 'vue';
import { useRoute } from 'vue-router';
import { useRouter } from 'vue-router';
// 其他导入语句
import { ElMessage, ElMessageBox } from 'element-plus';
import { useAppStore } from "@/pinia";
import { InfoFilled } from '@element-plus/icons-vue';
// 全量引入格式化工具 请按需保留

// 主选项卡激活状态
const activeMainTab = ref('product-info');
// 物模型子选项卡激活状态
const activeModelTab = ref('properties');
// 告警设置子选项卡激活状态
const activeAlarmTab = ref('alarm-rules');

// 返回列表页面
const route = useRoute();
const router = useRouter();
const returnList = () => {
  router.push({
    name: "deviceManage"
  });
};

// 产品详情
const productDetailFrom = ref({})
const getDeviceManageId = async (row) => {
  const res = await findGoodAccount({
    id: route.params.id
  })
  productDetailFrom.value = res.data
}
onMounted(() => {
  getDeviceManageId();
});

// 物模型-属性
// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()
// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)
// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  attributeId: '',
  attributeName: '',
  dataType: '',
  attributeFrom: '',
  read: '',
  description: '',
})
// 验证规则
const rule = ref({})
const elFormRef = ref()
const elSearchFormRef = ref()
// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}
// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
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
// 查询
const getTableData = async() => {
  try {
    const table = await getGoodAccountList({
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value,
      deviceId: route.params.id // 添加设备ID参数
    })
    if (table.code === 0) {
      tableData.value = table.data.list
      total.value = table.data.total
      page.value = table.data.page
      pageSize.value = table.data.pageSize
    } else {
      ElMessage.error(table.message || '获取属性列表失败')
    }
  } catch (error) {
    ElMessage.error('网络错误，请稍后重试')
  }
}
onMounted(() => {
  getTableData()
})
// ============== 表格控制部分结束 ===============
// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}
// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteDeviceManageModelAttributeFunc(row)
  })
}
// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const ids = []
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要删除的数据'
      })
      return
    }
    multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.id)
    })
    try {
      const res = await deleteGoodAccountByIds({
        ids
      })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      } else {
        ElMessage.error(res.message || '删除失败')
      }
    } catch (error) {
      ElMessage.error('网络错误，请稍后重试')
    }
  })
}
// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')
// 更新行
const updateDeviceManageModelAttributeFunc = async(row) => {
  try {
    const res = await findGoodAccount({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
      formData.value = res.data
      dialogFormVisible.value = true
    } else {
      ElMessage.error(res.message || '获取属性详情失败')
    }
  } catch (error) {
    ElMessage.error('网络错误，请稍后重试')
  }
}
// 删除行
const deleteDeviceManageModelAttributeFunc = async (row) => {
  try {
    const res = await deleteGoodAccount({ id: row.id })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    } else {
      ElMessage.error(res.message || '删除失败')
    }
  } catch (error) {
    ElMessage.error('网络错误，请稍后重试')
  }
}
// 弹窗控制标记
const dialogFormVisible = ref(false)
// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
  // 清空表单
  formData.value = {
    deviceId: route.params.id, // 设置设备ID
    attributeId: '',
    attributeName: '',
    dataType: '',
    attributeFrom: '',
    read: '',
    description: '',
  }
}
// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    attributeId: '',
    attributeName: '',
    dataType: '',
    attributeFrom: '',
    read: '',
    description: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  btnLoading.value = true
  elFormRef.value?.validate( async (valid) => {
    if (!valid) return btnLoading.value = false
    let res
    try {
      switch (type.value) {
        case 'create':
          res = await createGoodAccount(formData.value)
          break
        case 'update':
          res = await updateGoodAccount(formData.value)
          break
        default:
          res = await createGoodAccount(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
        closeDialog()
        getTableData()
      } else {
        ElMessage.error(res.message || '操作失败')
      }
    } catch (error) {
      ElMessage.error('网络错误，请稍后重试')
    } finally {
      btnLoading.value = false
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
const getDetailsAttribute = async (row) => {
  try {
    const res = await findDeviceManageModelAttribute({ id: row.id })
    if (res.code === 0) {
      detailFrom.value = res.data
      openDetailShow()
    } else {
      ElMessage.error(res.message || '获取属性详情失败')
    }
  } catch (error) {
    ElMessage.error('网络错误，请稍后重试')
  }
}
// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style scoped>
.product-detail-container {
  padding: 20px;
  background-color: #fff;
  border-radius: 4px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.detail-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
  padding-bottom: 10px;
  border-bottom: 1px solid #ebeef5;
}

.action-buttons {
  display: flex;
  gap: 10px;
}

.main-tabs {
  margin-top: 10px;
}

.nested-tabs {
  margin-top: 10px;
}

.info-descriptions {
  margin-top: 10px;
}

.model-list, .alarm-setting, .alarm-logs {
  margin-top: 10px;
}
</style>