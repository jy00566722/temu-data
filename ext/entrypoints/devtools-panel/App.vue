<template>

  <n-config-provider >
    <div class="title">TEMU后台数据分析插件</div>
    <n-divider />
    
    <n-space vertical size="large">

      <n-layout has-sider>
        <n-layout-sider
          collapse-mode="transform"
          :collapsed-width="25"
          :width="150"
          show-trigger="bar"
          content-style="padding: 24px;"
          bordered
        >
       
          <Nav></Nav>
        </n-layout-sider>
        <n-layout-content content-style="padding: 24px;">
          <router-view></router-view>
        </n-layout-content>
      </n-layout>
    </n-space>
  </n-config-provider>

</template>

<script setup>
  import { defineComponent, onBeforeMount} from 'vue'
  import { darkTheme, NButton,NDivider} from 'naive-ui'
  import Nav from './views/Nav.vue'
  // import {db} from './db.ts';
  import {getOrderList} from "./comm/orderList"


  // db.friends.add({name:'mihu1',age:32})
  // db.friends.add({name:'mihu2',age:32})
  onBeforeMount(() => {
    collect_data();
  } )
  const collect_data = () => {
    chrome.devtools.network.onRequestFinished.addListener(request => { 
    request.getContent((body) => {
       if(request.request && request.request.url == "https://seller.kuajingmaihuo.com/bgSongbird-api/supplier/deliverGoods/management/pageQueryDeliveryOrders") { //发货单列表

        if (request.request.url.includes('.')) { 
        //   chrome.runtime.sendMessage({ 
        //       response: body 
        //   }); 
          console.log('body:');
          getOrderList(body)

        } 
      } 
    }); 
  });
  }
</script>



<style scoped>
  .title {
    text-align: center;
    margin-top: 5px;
    font-size: 30px;
    background: linear-gradient(to right, #d92121, #87c890);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }
</style>
