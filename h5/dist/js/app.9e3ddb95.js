(function(e){function t(t){for(var o,l,i=t[0],s=t[1],c=t[2],f=0,p=[];f<i.length;f++)l=i[f],Object.prototype.hasOwnProperty.call(a,l)&&a[l]&&p.push(a[l][0]),a[l]=0;for(o in s)Object.prototype.hasOwnProperty.call(s,o)&&(e[o]=s[o]);u&&u(t);while(p.length)p.shift()();return r.push.apply(r,c||[]),n()}function n(){for(var e,t=0;t<r.length;t++){for(var n=r[t],o=!0,i=1;i<n.length;i++){var s=n[i];0!==a[s]&&(o=!1)}o&&(r.splice(t--,1),e=l(l.s=n[0]))}return e}var o={},a={app:0},r=[];function l(t){if(o[t])return o[t].exports;var n=o[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,l),n.l=!0,n.exports}l.m=e,l.c=o,l.d=function(e,t,n){l.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},l.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},l.t=function(e,t){if(1&t&&(e=l(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(l.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var o in e)l.d(n,o,function(t){return e[t]}.bind(null,o));return n},l.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return l.d(t,"a",t),t},l.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},l.p="/h5/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],s=i.push.bind(i);i.push=t,i=i.slice();for(var c=0;c<i.length;c++)t(i[c]);var u=s;r.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},"034f":function(e,t,n){"use strict";var o=n("85ec"),a=n.n(o);a.a},"56d7":function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d");var o=n("2b0e"),a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:"app"}},[n("h1",{staticStyle:{margin:"auto"}},[e._v("数据挖宝")]),n("el-container",[n("el-header",[n("p",[e._v("2021程序员节")])]),n("keep-alive",[n(e.currentTabComponent,{tag:"component"})],1)],1)],1)},r=[],l=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-dialog",{attrs:{title:"登录",visible:e.dialogFormVisible,"before-close":e.handleClose},on:{"update:visible":function(t){e.dialogFormVisible=t}}},[n("el-form",{attrs:{model:e.form}},[n("el-form-item",{attrs:{label:"输入体验码:"}},[n("el-input",{attrs:{autocomplete:"off"},model:{value:e.form.name,callback:function(t){e.$set(e.form,"name",t)},expression:"form.name"}})],1)],1),n("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{on:{click:e.cancelClick}},[e._v("取 消")]),n("el-button",{attrs:{type:"primary"},on:{click:e.confirmClick}},[e._v("确 定")])],1)],1)},i=[],s=(n("b0c0"),n("d3b7"),n("bc3a")),c=n.n(s);function u(e){o["default"].bus.emit("needlogin",e)}var f={post:function(e,t,n,o,a){c.a.post(e,t).then((function(e){console.log(e),2===e.data.err?(u((function(){console.log("login ok.")})),o(e.data)):n(e)})).catch(o).finally(a)},get:function(e,t,n,o,a){c.a.get(e,{params:t}).then((function(e){2===e.data.err?(u((function(){console.log("login ok.")})),o(e.data)):n(e)})).catch(o).finally(a)},save:function(e,t){if("number"===typeof t||"string"===typeof t||"boolean"===typeof t)localStorage.setItem(e,t);else{var n=JSON.stringify(t);localStorage.setItem(e+"@json",n)}},load:function(e){var t=localStorage.getItem(e+"@json");if(!t)return localStorage.getItem(e);try{return JSON.parse(t)}catch(n){return null}}},p=f,m={name:"Login",props:{},created:function(){this.$bus.on("needlogin",this.needLoginEV),console.log("login created")},beforeDestroy:function(){this.$bus.off("needlogin",this.needLoginEV),console.log("login will destroy.")},methods:{handleClose:function(e){console.log("login close."),e()},confirmClick:function(){var e=this;p.post("/login",{vcode:e.form.name},(function(t){console.log(t),t.data.err?e.$message({message:"登录失败:"+t.data.err,type:"error"}):(e.$message({message:"登录成功!",type:"success"}),p.save("user",t.data))}),(function(t){console.log("error=====",t),e.$message({message:"登录失败"+t,type:"error"})}),(function(){console.log("finally"),e.dialogFormVisible=!1}))},cancelClick:function(e){console.log(e),this.dialogFormVisible=!1},needLoginEV:function(e){console.log("needLoginEV call...",e),this.callback=e,this.dialogFormVisible=!0}},data:function(){return{dialogFormVisible:!1,form:{name:""}}}},d=m,g=n("2877"),b=Object(g["a"])(d,l,i,!1,null,"658a5cff",null),h=b.exports,v=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-container",[n("el-aside",{attrs:{width:"50px"}}),n("el-main",{},[n("el-row",{attrs:{gutter:20}},e._l(e.teamscore,(function(t){return n("el-col",{key:t.name,attrs:{span:8}},[e._v(e._s(t.name)+" : "+e._s(t.score))])})),1),n("el-table",{staticStyle:{width:"100%",margin:"auto","margin-top":"50px"},attrs:{"max-height":"450",data:e.tableData,"header-cell-style":e.hcs,"cell-style":e.hcs}},[n("el-table-column",{attrs:{prop:"name",label:"队伍名"}}),n("el-table-column",{attrs:{prop:"score",label:"得分"}}),n("el-table-column",{attrs:{prop:"type",label:"得分项"}})],1)],1),n("el-aside",{attrs:{width:"120px"}},[n("h4",[e._v("剩余时间 "+e._s(e.lefttime)+" 秒")]),n("el-input",{attrs:{size:"mini",placeholder:"密令"},model:{value:e.input,callback:function(t){e.input=t},expression:"input"}}),n("el-button",{staticStyle:{"margin-top":"20px"},attrs:{size:"mini",type:"warning"},on:{click:e.reset}},[e._v("重置")])],1)],1)},y=[],_={name:"Welcome",props:{},data:function(){return{tableData:[],progress:0,input:"",lefttime:180,teamscore:[]}},created:function(){var e=this;this.getList(),setInterval((function(){e.getList()}),1e3)},methods:{hcs:function(){return"text-align : center;"},reset:function(){var e=this;p.get("reset",{code:this.input},(function(t){console.log(t),0==t.data.errorno&&e.$message.success("重置成功")}),(function(e){console.log(e)}))},getList:function(){var e=this;p.get("info",{},(function(t){var n=t.data.data;e.lefttime=n.lefttime;var o={},a=[];for(var r in n.magics){var l=n.magics[r],i={};i.name=l,i.score=1,i.type=r,a.push(i),o[l]||(o[l]=0),o[l]+=1}for(var s in n.formulas){var c=n.formulas[s],u=c[0],f={};f.name=u;var p=JSON.parse(s).length;f.score=p*p,f.type=c[1],a.push(f),o[u]||(o[u]=0),o[u]+=f.score}var m=[];for(var d in o)m.push({name:d,score:o[d]});e.teamscore=m,e.tableData=a}),(function(t){console.log(t),e.$message.warning("比赛暂未开始")}))}}},w=_,k=Object(g["a"])(w,v,y,!1,null,"ca0b28aa",null),x=k.exports,O={name:"app",components:{Login:h,Welcome:x},methods:{},data:function(){return{currentTabComponent:"Welcome"}}},S=O,j=(n("034f"),Object(g["a"])(S,a,r,!1,null,null,null)),$=j.exports,C=n("5c96"),L=n.n(C);n("c69f");o["default"].use(L.a);var V=n("b828");o["default"].use(V["a"]),console.log("====="),console.log("production"),c.a.defaults.baseURL="http://47.104.220.230",c.a.defaults.withCredentials=!0,o["default"].config.productionTip=!1,new o["default"]({render:function(e){return e($)}}).$mount("#app")},"85ec":function(e,t,n){},c69f:function(e,t,n){}});
//# sourceMappingURL=app.9e3ddb95.js.map