<template>
  <div class="product-category">
    <h1>产品分类管理</h1>
    <!-- 这里可加表格、表单，实现分类的增删改查 -->
  </div>
</template>

<script>
export default {
  name: 'ProductCategory',
  // 可通过 axios 调用后端接口
  created() {
    // 示例：获取分类列表
    this.$http.get('/product/getProductList').then(res => {
      console.log(res.data);
    });
  }
}
</script>

<style scoped>
.product-category {
  padding: 20px;
}
</style>