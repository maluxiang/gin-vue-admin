
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
            <el-form-item label="设备id" prop="equId">
  <el-input v-model="searchInfo.equId" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="设备名称" prop="equName">
  <el-input v-model="searchInfo.equName" placeholder="搜索条件" />
</el-form-item>
            
            <el-form-item label="告警时间" prop="alarmTime">
  <template #label>
    <span>
      alarmTime字段
      <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
        <el-icon><QuestionFilled /></el-icon>
      </el-tooltip>
    </span>
  </template>
<el-date-picker class="w-[380px]" v-model="searchInfo.alarmTimeRange" type="datetimerange" range-separator="至"  start-placeholder="开始时间" end-placeholder="结束时间"></el-date-picker></el-form-item>
            

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
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="序号" prop="id" width="120" />

            <el-table-column align="left" label="设备id" prop="equId" width="120" />

            <el-table-column align="left" label="设备名称" prop="equName" width="120" />

            <el-table-column align="left" label="告警类型" prop="alarmStatus" width="120" />


            <el-table-column align="left" label="告警时间" prop="alarmTime" width="180">
   <template #default="scope">{{ formatDate(scope.row.alarmTime) }}</template>
</el-table-column>
            <el-table-column align="left" label="告警状态" prop="cpuStatus" width="120" >
          <template #default="scope">
            <span class="status-indicator" :class="getStatusClass(scope.row.cpuStatus)"></span>
            {{ scope.row.cpuStatus }}
          </template>
          </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateAlarmManageFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
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
            <el-form-item label="equId字段:" prop="equId">
    <el-input v-model="formData.equId" :clearable="true" placeholder="请输入equId字段" />
</el-form-item>
            <el-form-item label="equName字段:" prop="equName">
    <el-input v-model="formData.equName" :clearable="true" placeholder="请输入equName字段" />
</el-form-item>
            <el-form-item label="alarmStatus字段:" prop="alarmStatus">
    <el-input v-model="formData.alarmStatus" :clearable="true" placeholder="请输入alarmStatus字段" />
</el-form-item>
            <el-form-item label="cpuStatus字段:" prop="cpuStatus">
    <el-input v-model="formData.cpuStatus" :clearable="true" placeholder="请输入cpuStatus字段" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="id字段">
    {{ detailFrom.id }}
</el-descriptions-item>
                    <el-descriptions-item label="equId字段">
    {{ detailFrom.equId }}
</el-descriptions-item>
                    <el-descriptions-item label="equName字段">
    {{ detailFrom.equName }}
</el-descriptions-item>
                    <el-descriptions-item label="alarmStatus字段">
    {{ detailFrom.alarmStatus }}
</el-descriptions-item>
                    <el-descriptions-item label="alarmTime字段">
    {{ detailFrom.alarmTime }}
</el-descriptions-item>
                    <el-descriptions-item label="cpuStatus字段">
    {{ detailFrom.cpuStatus }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createAlarmManage,
  deleteAlarmManage,
  deleteAlarmManageByIds,
  updateAlarmManage,
  findAlarmManage,
  getAlarmManageList
} from '@/api/good/alarmManage'

// 全量引入格式化工具 请按需保留
import {  formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'AlarmManage'
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
            alarmStatus: '',
            cpuStatus: '',
        })



// 验证规则
const rule = reactive({
})

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
  const table = await getAlarmManageList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


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
            deleteAlarmManageFunc(row)
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
      const res = await deleteAlarmManageByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateAlarmManageFunc = async(row) => {
    const res = await findAlarmManage({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAlarmManageFunc = async (row) => {
    const res = await deleteAlarmManage({ id: row.id })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        equId: '',
        equName: '',
        alarmStatus: '',
        cpuStatus: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createAlarmManage(formData.value)
                  break
                case 'update':
                  res = await updateAlarmManage(formData.value)
                  break
                default:
                  res = await createAlarmManage(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
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
  // 打开弹窗
  const res = await findAlarmManage({ id: row.id })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}

const getStatusClass = (status) => {
  switch (status) {
    case '已处理':
      return 'status-online';
    case '未处理':
      return 'status-offline';
  }
};

</script>

<style scoped>
.status-indicator {
  display: inline-block;
  width: 10px;
  height: 10px;
  /* border-radius: 50%; !* 圆形，如果要方形则删除此行 *!*/
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