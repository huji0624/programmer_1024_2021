<template>
  <el-container>
    <el-aside width="320px">
    </el-aside>
    <el-main style="">
    
    <el-row style="margin-top:10px;text-align:center;" v-for="(item,index) in teamscore" v-bind:key="item.name">
      <el-button style="width:512px;" :type="typeFromIndex(index)">{{item.name}} : {{item.score}}</el-button>
    </el-row>

    <el-table
      max-height="450"
      :data="tableData"
      style="width: 100%;margin:auto;margin-top:50px;"
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
      <el-table-column
        prop="type"
        label="得分项">
      </el-table-column>
    </el-table>
      
    </el-main>

    <el-aside width="320px">
        <p>剩余时间</p>
        <el-progress :percentage="timeleft" :format="format"></el-progress>
        <p> </p>
        <p> </p>
        <p> </p>
        <el-input size="mini" v-model="input" placeholder="密令"></el-input>
        <el-button size="mini" @click="reset" type="warning" style="margin-top:20px;">重置</el-button>
    </el-aside>
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
      progress:0,
      input:"",
      timeleft:100,
      lefttime:180,
      teamscore:[]
    }
  },
  created:function(){
    this.getList()

    setInterval(() => {
      this.getList()
    }, 1000);
  },
  methods:{
    typeFromIndex(index){
      if (index==0){
        return "warning"
      }else if(index==1){
        return "success"
      }

      return "danger"
    },
    format() {
        return `${this.lefttime}`;
    },
    hcs:function(){
      return "text-align : center;"
    },
    reset:function(){
      let out_this = this;
      common.get("reset",{code:this.input},function(res){
        console.log(res)
        if(res.data.errorno==0){
          out_this.$message.success("重置成功")
        }

      },function(err){
        console.log(err)
      })
    },
    getList:function(){
      let out_this = this;
      common.get("info",{},function(res){
        console.log(res.data.data)

        let data = res.data.data;
        out_this.timeleft = data.lefttime/180;
        out_this.lefttime = data.lefttime;

        let records = []
        for(let k in data.records){
          let v = data.records[k]
          records.push({name:v.team,score:v.score,type:v.record})
        }

        data.teams.sort(function(a,b){return b.score-a.score})
        out_this.teamscore = data.teams;

        records.reverse()
        out_this.tableData = records

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
