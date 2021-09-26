<template>
  <el-container>
    <el-aside width="50px">

    </el-aside>
    <el-main style="">
    

    <!-- <el-progress :text-inside="true" :stroke-width="24" :percentage="progress" status="success"></el-progress> -->
    

    <el-row :gutter="20">
      <el-col :span="8" v-for="(item) in teamscore" v-bind:key="item.name">{{item.name}} : {{item.score}}</el-col>
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

    <el-aside width="120px">
        <h4>剩余时间 {{lefttime}} 秒</h4>
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
        // console.log(res.data.data)

        let data = res.data.data;
        out_this.lefttime = data.lefttime;

        let teamscoremap = {};
        let tbd = [];
        for(let k in data.magics){
          let v = data.magics[k]
          let item = {};
          item.name = v;
          item.score = 1;
          item.type = k;
          tbd.push(item);

          if(!teamscoremap[v]){
            teamscoremap[v] = 0;
          }

          teamscoremap[v] += 1;
        }

        for(let k in data.formulas){
          let v = data.formulas[k]

          let tname = v[0];

          let item = {};
          item.name = tname;
          let idcount = JSON.parse(k).length;
          item.score = idcount*idcount;
          item.type = v[1];
          tbd.push(item);

          if(!teamscoremap[tname]){
            teamscoremap[tname] = 0;
          }

          teamscoremap[tname] += item.score;
        }

        let tmp = [];
        for(let k in teamscoremap){
          tmp.push({name:k,score:teamscoremap[k]})
        }
        out_this.teamscore = tmp;
        
        out_this.tableData = tbd;

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
