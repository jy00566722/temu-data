<template>
    <div class="nav">
    <n-flex vertical>
        <router-link to="/">主页</router-link>
        <router-link to="/collect">数据收集</router-link>
      </n-flex>
    </div>
</template>

<style scoped>
.nav a{
    text-decoration: none; /* 去除下划线 */
    color: inherit; /* 保持颜色不变 */
}
.nav a.router-link-active{
    color:blueviolet;
}
</style>