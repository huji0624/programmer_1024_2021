(function(e){function t(t){for(var n,l,i=t[0],s=t[1],c=t[2],f=0,p=[];f<i.length;f++)l=i[f],Object.prototype.hasOwnProperty.call(r,l)&&r[l]&&p.push(r[l][0]),r[l]=0;for(n in s)Object.prototype.hasOwnProperty.call(s,n)&&(e[n]=s[n]);u&&u(t);while(p.length)p.shift()();return a.push.apply(a,c||[]),o()}function o(){for(var e,t=0;t<a.length;t++){for(var o=a[t],n=!0,i=1;i<o.length;i++){var s=o[i];0!==r[s]&&(n=!1)}n&&(a.splice(t--,1),e=l(l.s=o[0]))}return e}var n={},r={app:0},a=[];function l(t){if(n[t])return n[t].exports;var o=n[t]={i:t,l:!1,exports:{}};return e[t].call(o.exports,o,o.exports,l),o.l=!0,o.exports}l.m=e,l.c=n,l.d=function(e,t,o){l.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:o})},l.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},l.t=function(e,t){if(1&t&&(e=l(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var o=Object.create(null);if(l.r(o),Object.defineProperty(o,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var n in e)l.d(o,n,function(t){return e[t]}.bind(null,n));return o},l.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return l.d(t,"a",t),t},l.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},l.p="/h5/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],s=i.push.bind(i);i.push=t,i=i.slice();for(var c=0;c<i.length;c++)t(i[c]);var u=s;a.push([0,"chunk-vendors"]),o()})({0:function(e,t,o){e.exports=o("56d7")},"034f":function(e,t,o){"use strict";var n=o("85ec"),r=o.n(n);r.a},"56d7":function(e,t,o){"use strict";o.r(t);o("e260"),o("e6cf"),o("cca6"),o("a79d");var n=o("2b0e"),r=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("div",{attrs:{id:"app"}},[o("h1",{staticStyle:{margin:"auto"}},[e._v("数据挖宝")]),o("el-container",[o("el-header",[o("p",[e._v("2021程序员节")])]),o("keep-alive",[o(e.currentTabComponent,{tag:"component"})],1)],1)],1)},a=[],l=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("el-dialog",{attrs:{title:"登录",visible:e.dialogFormVisible,"before-close":e.handleClose},on:{"update:visible":function(t){e.dialogFormVisible=t}}},[o("el-form",{attrs:{model:e.form}},[o("el-form-item",{attrs:{label:"输入体验码:"}},[o("el-input",{attrs:{autocomplete:"off"},model:{value:e.form.name,callback:function(t){e.$set(e.form,"name",t)},expression:"form.name"}})],1)],1),o("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[o("el-button",{on:{click:e.cancelClick}},[e._v("取 消")]),o("el-button",{attrs:{type:"primary"},on:{click:e.confirmClick}},[e._v("确 定")])],1)],1)},i=[],s=(o("b0c0"),o("d3b7"),o("bc3a")),c=o.n(s);function u(e){n["default"].bus.emit("needlogin",e)}var f={post:function(e,t,o,n,r){c.a.post(e,t).then((function(e){console.log(e),2===e.data.err?(u((function(){console.log("login ok.")})),n(e.data)):o(e)})).catch(n).finally(r)},get:function(e,t,o,n,r){c.a.get(e,{params:t}).then((function(e){2===e.data.err?(u((function(){console.log("login ok.")})),n(e.data)):o(e)})).catch(n).finally(r)},save:function(e,t){if("number"===typeof t||"string"===typeof t||"boolean"===typeof t)localStorage.setItem(e,t);else{var o=JSON.stringify(t);localStorage.setItem(e+"@json",o)}},load:function(e){var t=localStorage.getItem(e+"@json");if(!t)return localStorage.getItem(e);try{return JSON.parse(t)}catch(o){return null}}},p=f,d={name:"Login",props:{},created:function(){this.$bus.on("needlogin",this.needLoginEV),console.log("login created")},beforeDestroy:function(){this.$bus.off("needlogin",this.needLoginEV),console.log("login will destroy.")},methods:{handleClose:function(e){console.log("login close."),e()},confirmClick:function(){var e=this;p.post("/login",{vcode:e.form.name},(function(t){console.log(t),t.data.err?e.$message({message:"登录失败:"+t.data.err,type:"error"}):(e.$message({message:"登录成功!",type:"success"}),p.save("user",t.data))}),(function(t){console.log("error=====",t),e.$message({message:"登录失败"+t,type:"error"})}),(function(){console.log("finally"),e.dialogFormVisible=!1}))},cancelClick:function(e){console.log(e),this.dialogFormVisible=!1},needLoginEV:function(e){console.log("needLoginEV call...",e),this.callback=e,this.dialogFormVisible=!0}},data:function(){return{dialogFormVisible:!1,form:{name:""}}}},m=d,g=o("2877"),b=Object(g["a"])(m,l,i,!1,null,"658a5cff",null),h=b.exports,v=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("el-container",[o("el-aside",{attrs:{width:"50px"}}),o("el-main",{},[o("el-row",{attrs:{gutter:20}},e._l(e.teamscore,(function(t){return o("el-col",{key:t.name,staticStyle:{border:"solid 1px blue"},attrs:{span:8}},[e._v(e._s(t.name)+" : "+e._s(t.score))])})),1),o("el-table",{staticStyle:{width:"100%",margin:"auto","margin-top":"50px"},attrs:{"max-height":"450",data:e.tableData,"header-cell-style":e.hcs,"cell-style":e.hcs}},[o("el-table-column",{attrs:{prop:"name",label:"队伍名"}}),o("el-table-column",{attrs:{prop:"score",label:"得分"}}),o("el-table-column",{attrs:{prop:"type",label:"得分项"}})],1)],1),o("el-aside",{attrs:{width:"120px"}},[o("h4",[e._v("剩余时间 "+e._s(e.lefttime)+" 秒")]),o("el-input",{attrs:{size:"mini",placeholder:"密令"},model:{value:e.input,callback:function(t){e.input=t},expression:"input"}}),o("el-button",{staticStyle:{"margin-top":"20px"},attrs:{size:"mini",type:"warning"},on:{click:e.reset}},[e._v("重置")])],1)],1)},y=[],_=(o("26e9"),{name:"Welcome",props:{},data:function(){return{tableData:[],progress:0,input:"",lefttime:180,teamscore:[]}},created:function(){var e=this;this.getList(),setInterval((function(){e.getList()}),1e3)},methods:{hcs:function(){return"text-align : center;"},reset:function(){var e=this;p.get("reset",{code:this.input},(function(t){console.log(t),0==t.data.errorno&&e.$message.success("重置成功")}),(function(e){console.log(e)}))},getList:function(){var e=this;p.get("info",{},(function(t){console.log(t.data.data);var o=t.data.data;e.lefttime=o.lefttime;var n={},r=[];for(var a in o.records){var l=o.records[a];r.push({name:l.team,score:l.score,type:l.record}),void 0==n[l.team]&&(n[l.team]={name:l.team,score:0}),n[l.team].score+=l.score}var i=[];for(var s in n)i.push(n[s]);e.teamscore=i,r.reverse(),e.tableData=r}),(function(t){console.log(t),e.$message.warning("比赛暂未开始")}))}}}),w=_,x=Object(g["a"])(w,v,y,!1,null,"0a2467be",null),k=x.exports,O={name:"app",components:{Login:h,Welcome:k},methods:{},data:function(){return{currentTabComponent:"Welcome"}}},S=O,j=(o("034f"),Object(g["a"])(S,r,a,!1,null,null,null)),$=j.exports,C=o("5c96"),L=o.n(C);o("c69f");n["default"].use(L.a);var V=o("b828");n["default"].use(V["a"]),console.log("====="),console.log("production"),c.a.defaults.baseURL="http://47.104.220.230",c.a.defaults.withCredentials=!0,n["default"].config.productionTip=!1,new n["default"]({render:function(e){return e($)}}).$mount("#app")},"85ec":function(e,t,o){},c69f:function(e,t,o){}});
//# sourceMappingURL=app.d3b3870b.js.map