<template>
  <div class="login-container">
    <!-- 背景等装饰可自行用 CSS 处理，这里先放核心表单 -->
    <div class="login-box">
      <h2>IoT物联设备管理平台</h2>
      <el-form :model="loginForm" label-width="80px">
        <el-form-item label="账号">
          <el-input v-model="loginForm.username"></el-input>
        </el-form-item>
        <el-form-item label="密码">
          <el-input type="password" v-model="loginForm.password"></el-input>
        </el-form-item>
        <el-form-item label="企业编号">
          <el-input v-model="loginForm.enterpriseCode"></el-input>
        </el-form-item>
        <el-form-item label="验证码">
          <el-input v-model="loginForm.captcha" style="width: 60%"></el-input>
          <div class="captcha-img" @click="refreshCaptcha">
            <!-- 实际开发要替换成真实验证码图片渲染，这里先模拟 -->
            <img :src="captchaUrl" alt="验证码">
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin">登录</el-button>
          <el-button @click="handleForget">忘记密码</el-button>
        </el-form-item>
        <el-checkbox v-model="loginForm.rememberMe">记住密码</el-checkbox>
      </el-form>
    </div>
  </div>
</template>

<script>

</script>

<style scoped>

</style>