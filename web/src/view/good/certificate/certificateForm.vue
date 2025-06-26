
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="certificateName字段:" prop="certificateName">
    <el-input v-model="formData.certificateName" :clearable="true" placeholder="请输入certificateName字段" />
</el-form-item>
        <el-form-item label="versionNumber字段:" prop="versionNumber">
    <el-input v-model="formData.versionNumber" :clearable="true" placeholder="请输入versionNumber字段" />
</el-form-item>
        <el-form-item label="certificateInfo字段:" prop="certificateInfo">
    <el-input v-model="formData.certificateInfo" :clearable="true" placeholder="请输入certificateInfo字段" />
</el-form-item>
        <el-form-item label="certificateStatus字段:" prop="certificateStatus">
    <el-input v-model="formData.certificateStatus" :clearable="true" placeholder="请输入certificateStatus字段" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createCertificate,
  updateCertificate,
  findCertificate
} from '@/api/good/certificate'

defineOptions({
    name: 'CertificateForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            certificateName: '',
            versionNumber: '',
            certificateInfo: '',
            certificateStatus: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCertificate({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createCertificate(formData.value)
               break
             case 'update':
               res = await updateCertificate(formData.value)
               break
             default:
               res = await createCertificate(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
