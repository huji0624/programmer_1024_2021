<template>
  <div id="app">
    <!-- <Login/> -->
    <h1 style="margin:auto;">数据挖宝</h1>
    <el-container>
      <el-header><p>2021程序员节</p></el-header>
      <keep-alive>
        <component v-bind:is="currentTabComponent"></component>
      </keep-alive>
    </el-container>
  </div>
</template>

<script>
import Login from './components/Login.vue'
import Welcome from './components/Welcome.vue'

export default {
  name: 'app',
  components: {
    Login,
    Welcome,
  },
  methods:{
    
  },
  data:function(){
    return {
      currentTabComponent:"Welcome",
      };
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
