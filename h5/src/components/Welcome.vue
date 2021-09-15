<template>
  <el-container>
    <!-- <el-aside width="200px">
    </el-aside> -->
    <el-main style="">
    

    <el-progress :text-inside="true" :stroke-width="24" :percentage="progress" status="success"></el-progress>
    <p>被挖出宝藏占比</p>

    <el-table
      :data="tableData"
      style="width: 70%;margin:auto;margin-top:50px;"
      :header-cell-style="hcs"
      :cell-style="hcs">
      <el-table-column
        prop="name"
        label="队伍名">
      </el-table-column>
      <el-table-column
        prop="score"
        label="得分">
      </el-table-column>
    </el-table>
      
    </el-main>
  </el-container>
</template>

<script>
import common from '../plugins/common';

export default {
  name: 'Welcome',
  props: {
  },
  data:function(){
    return {
      tableData:[],
      progress:0
    }
  },
  created:function(){
    this.getList()

    setInterval(() => {
      this.getList()
    }, 1000);
  },
  methods:{
    hcs:function(){
      return "text-align : center;"
    },
    getList:function(){
      let out_this = this;
      common.get("info",{},function(res){
        // console.log(res)
        // out_this.$message.success("OK.")
        let tbs = [];
        let ret = res.data.data.result;
        let tt = 0;
        for (let key in ret) {
          let v = ret[key]
          tt += v
          tbs.push({name:key,score:v})
        }
        tbs = tbs.sort(function(a,b){
          return b["score"]-a["score"]
        })
        out_this.tableData = tbs;

        if(res.data.data.total){
          out_this.progress = Math.floor(tt/res.data.data.total*10000)/100
          // console.log(out_this.progress)
        }

      },function(err){
        console.log(err)
        out_this.$message.warning("比赛暂未开始")
      })
    },
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
