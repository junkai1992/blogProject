<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col-lg-12" style="padding: 0">
        <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
          <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarTogglerDemo01" aria-controls="navbarTogglerDemo01" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse" id="navbarTogglerDemo01">
            <a class="navbar-brand" href="#">Xjk blog</a>
            <ul class="navbar-nav mr-auto mt-2 mt-lg-0">
              <!--              <li class="nav-item active">-->
              <!--                <a class="nav-link" href="#">Home <span class="sr-only">(current)</span></a>-->
              <!--              </li>-->
              <li class="nav-item">
                <div class="title-item-class"  @click="postsHandle()">
                  <Icon type="ios-navigate"></Icon>文章
                </div>
              </li>
              <li class="nav-item">
                <div class="title-item-class" @click="postCreateHandle()">
                  <Icon type="ios-keypad"></Icon>
                  新建
                </div>
              </li>
              <li class="nav-item">
                <div v-if="judgeLogin"  @click="showlogin()" class="title-item-class">
                  <Icon type="ios-analytics"></Icon>
                  用户登陆
                </div>
              </li>
              <!--              <li class="nav-item">-->
              <!--                <a class="nav-link disabled" href="#" tabindex="-1" aria-disabled="true">Disabled</a>-->
              <!--              </li>-->
            </ul>
            <form class="form-inline my-2 my-lg-0">
              <input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search">
              <button class="btn btn-outline-secondary my-2 my-sm-0" type="submit">Search</button>
            </form>
          </div>
        </nav>
      </div>
    </div>
    <div class="row" style="height: auto;min-height: 1000px;">
      <div class="col-lg-2" :style="{background:'#c0c3c8'}">
        <div :style="{display: 'flex', justifyContent: 'center', alignItems:'center', height: '15rem'}">
          <Avatar src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1599907997423&di=5081e04ad52f0d6d05d8cc51310c7133&imgtype=0&src=http%3A%2F%2Fcdn.duitang.com%2Fuploads%2Fitem%2F201309%2F26%2F20130926095128_SiPMh.jpeg" :style="{width: '9rem', height: '9rem'}"/>
        </div>
        <div :style="{background: '#808695', padding:'40px 9px', color: '#ffffff'}">
          <div :style="{color: '#FFFFFF', fontSize: '14px'}">好好学习，天天向上</div>
          <p>python,lua,go,javaScript....</p>
        </div>
        <div :style="{lineHeight: '4rem', padding: '2rem 0.1rem 3rem', fontSize: '20px'}">
          <p><Icon type="ios-eye" />访问人数:
            <animate-number from="1" :to="vistor" duration="3000" easing="easeOutQuad" fromColor="#44CC00" to-color="#e27559" v-if="showvistor"></animate-number>
          </p>
          <p>
            <span>关于作者</span>
            <router-link to="/files">
              <span>存档</span>
            </router-link>
            <span>私信</span>
          </p>
        </div>
        <div>
          <div :style="{color: '#40485a', fontSize: '18px'}">标签:</div>
        </div>
        <div :style="{padding:'10px',display: 'flex',justifyContent: 'center',justifyItems: 'center', lineHeight: '2rem'}">
          <p>
            <Tag color="green" :style="{cursor: 'pointer'}" @click.native="httpCategoryPostList(0)">全部</Tag>
            <Tag v-for="item in category_list" :key="item.id" color="green" :style="{cursor: 'pointer'}" @click.native="httpCategoryPostList(item.id)">{{item.name}}</Tag>
          </p>
        </div>
        <div :style="{minHeight: '20rem'}"></div>
      </div>
      <div class="col-lg-7" style="padding: 0;">
        <router-view v-if="isRouterAlive" :category_list="category_list" :tagId="tagId" :yearMonth="yearMonth" :currentReadPostId="currentReadPostId">
        </router-view>
        <div v-if="!isRouterAlive">hello world</div>
      </div>
      <div class="col-lg-3" :style="{background:'#c0c3c8'}">
        <!--日历-->
        <div class="dateHandle">
          <vc-calendar :attributes='calendarDate'></vc-calendar>
        </div>
        <div class="dateHandle">
          <i-col span="20">
            <Card>
              <p slot="title">
                <Icon type="md-bookmarks"></Icon>
                最新更新
              </p>
              <a href="#" slot="extra" @click.prevent="changeLimit">
                <Icon type="ios-loop-strong"></Icon>
                换一批
              </a>
              <ul style="list-style-type:none">
                <li v-for="item in randomPostList" :key="item.id" class="posts-archives" @click="postDetial(item.id)">
                    <span>
                    {{ item.title }}<Icon type="ios-calendar-outline" />{{ item.created_at }}
                    </span>
                </li>
              </ul>
            </Card>
          </i-col>
        </div>
        <div class="dateHandle">
          <i-col span="20">
            <Card>
              <p slot="title">
                <Icon type="md-bookmarks"></Icon>
                文章档案
              </p>
              <ul style="list-style-type:none">
                <li v-for="item in MonthPostList" :key="item.ArchiveDate" @click="httpzYMPostsList(item.Year,item.Month)" class="posts-archives">
                    <span>
                    {{ item.Year }}年{{item.Month}}月
                    </span>
                  <span>共计:{{ item.Total }}篇</span>
                </li>
              </ul>
            </Card>
          </i-col>
        </div>
      </div>
    </div>
    <div class="row footer">
      <div class="col-lg-12">2020-2099 &copy; xjk Blog</div>
    </div>
    <Modal
      v-model="modal1"
      :loading="loading"
      ok-text = "登陆"
      title="用户登陆"
      @on-ok="ok"
      @on-cancel="cancel"
    >
      <div>
        <div class="login">
          <div class="from-wrap">
            <Form ref="loginDataId" :model="loginData" :rules="ruleValidate" :label-width="80">
              <FormItem label="邮箱" prop="email">
                <Input type="text" v-model="loginData.email" placeholder="请输入邮箱" clearable :style="{width: '21rem'}">
                  <Icon type="ios-person-outline" slot="prepend"></Icon>
                </Input>
              </FormItem>
              <FormItem label="密码" prop="password">
                <Input type="password" v-model="loginData.password" placeholder="请输入密码" :style="{width: '21rem'}" clearable>
                  <Icon type="ios-lock-outline" slot="prepend"></Icon>
                </Input>
              </FormItem>
            </Form>
          </div>
        </div>
      </div>
    </Modal>
    <template>
      <Back-top :height="100" :bottom="100">
        <div class="top">
          <div>
            <Icon type="ios-arrow-dropup-circle" />
          </div>
          <div>
            返回顶端
          </div>
        </div>
      </Back-top>
    </template>
  </div>
</template>

<script>
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap.min'
export default {
  name: 'home',
  data () {
    return {
      MonthPostList: [],
      tagId: 0,
      vistor: 0,
      showvistor: false,
      judgeLogin: true,
      // 表单验证
      loading: true,
      modal1: false,
      // 当前读取文章id
      currentReadPostId: '',
      isRouterAlive: true,
      limitNum: 5,
      limitFrom: 0,
      newtitleList: [],
      randomPostList: [],
      postList: [],
      category_list: [],
      // 年月
      yearMonth: '',
      // 日历
      calendarDate: [
        {
          key: 'today',
          // 括号内传递日期可点亮指定日期，如new Date(2019, 6, 1)，也可传递多个日期：如dates: [ new Date(2018, 0, 1), new Date(2018, 0, 15) ]
          dates: new Date(),
          highlight: true,
          // popover 点亮的日期上出现提示内容
          popover: {
            label: '美好的一天！要开心呦！'
          }
        },
        {
          key: 'yestoday',
          // 括号内传递日期可点亮指定日期，如new Date(2019, 6, 1)，也可传递多个日期：如dates: [ new Date(2018, 0, 1), new Date(2018, 0, 15) ]
          dates: new Date(2020, 8, 20),
          highlight: true,
          // popover 点亮的日期上出现提示内容
          popover: {
            label: '明天又事情！'
          }
        }
      ],
      // 每页数据
      count: 10,
      user: '未登陆',
      loginData: {
        email: '',
        password: ''
      },
      ruleValidate: {
        email: [
          { required: true, message: 'email账号不能为空', trigger: 'blur' },
          { type: 'string', min: 6, message: '账号长度至少6个字符', trigger: 'blur' },
          { type: 'email', message: '邮箱格式不正确', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '密码不能为空', trigger: 'blur' },
          { type: 'string', min: 6, message: '密码长度至少6个字符', trigger: 'blur' }
        ]
      }
    }
  },
  mounted () {
    this.httpPutVistor()
    this.httpGetVistor()
  },
  created () {
    this.changeLimit()
    // 获取最新上传文章信息
    this.httpNewPost()
    // 获取标签
    this.httpCategory()
    // 月分组文章
    this.httpMonthPostList()
  },
  methods: {
    // 最新更新：文章详情
    postDetial (postsId) {
      this.currentReadPostId = postsId
      this.$router.push({path: `/article/${postsId}`})
    },
    postCreateHandle () {
      if (this.$route.path !== '/relative') {
        this.$router.push({ path: 'relative' })
      }
    },
    postsHandle () {
      this.yearMonth = ''
      if (this.$route.path !== '/posts') {
        this.$router.push({ path: '/posts' })
      }
    },
    httpzYMPostsList (year, month) {
      this.yearMonth = year + '-' + month
      if (this.$route.path !== '/posts') {
        this.$router.push({ path: '/posts' })
      }
    },
    async httpMonthPostList () {
      const { data: { code, result } } = await this.$http.get('/posts/month/list')
      if (code === 200) {
        this.MonthPostList = result.data
      }
    },
    async httpCategoryPostList (categoryId) {
      if (this.$route.path !== '/posts') {
        this.$router.push({ path: '/posts' })
      }
      this.tagId = categoryId
    },
    // 获取标签
    async httpCategory () {
      const { data: { code, msg, result } } = await this.$http.get('/categories/lists')
      if (code === 200) {
        this.category_list = result.categorys
        this.$Message.success(msg)
      }
    },
    // 退出登陆
    loginOut () {
      this.judgeLogin = true
      window.sessionStorage.clear()
      this.user = '未登陆'
    },
    // 不让关闭弹窗
    messageWarningFn (text) {
      this.$Message.warning(text)
      setTimeout(() => {
        this.loading = false
        this.$nextTick(() => {
          this.loading = true
        })
      }, 500)
    },
    changeLimit () {
      this.randomPostList = this.editArrayItems(this.postList, 5)
    },
    showlogin () {
      this.modal1 = true
    },
    handleSubmit (name) {
      this.$refs[name].validate((valid) => {
        if (valid) {
          this.$Message.success('提交成功!')
        } else {
          this.$Message.error('表单验证失败!')
        }
      })
    },
    handleReset (name) {
      this.$refs[name].resetFields()
    },
    editArrayItems (arr, num) {
      const tempArray = []
      for (let index in arr) {
        tempArray.push(arr[index])
      }
      const returnArray = []
      for (let i = 0; i < num; i++) {
        if (tempArray.length > 0) {
          const arrIndex = Math.floor(Math.random() * tempArray.length)
          returnArray[i] = tempArray[arrIndex]
          tempArray.splice(arrIndex, 1)
        } else {
          break
        }
      }
      return returnArray
    },
    async httpNewPost () {
      const { data: { code, result } } = await this.$http.get('/posts/new/list')
      if (code === 200) {
        this.postList = result.data
        this.randomPostList = this.editArrayItems(this.postList, 5)
        // this.editArrayItems(this.postList, 5)
      }
    },
    async httpPutVistor () {
      await this.$http.put('/visit/putvistor')
    },
    async httpGetVistor () {
      const { data: { code, msg, result } } = await this.$http.get('/visit/getvistor')
      if (code === 200) {
        this.$Message.success(msg)
        this.vistor = result.visitcount.visitcount
        this.showvistor = true
      }
    },
    async httpLogin () {
      const { data: { code, msg, result } } = await this.$http.post('/api/auth/login', this.loginData)
      if (code === 200) {
        this.$Message.success(msg)
        this.user = result.user
        window.sessionStorage.setItem('token', 'Bearer ' + result.token)
        window.sessionStorage.setItem('user', result.user)
        window.sessionStorage.setItem('user_id', result.user_id)
        this.judgeLogin = false
        this.modal1 = false
      } else {
        this.$Message.error(msg)
      }
    },
    ok: function () {
      this.$refs['loginDataId'].validate((valid) => {
        if (!valid) {
          this.messageWarningFn('用户验证失败!')
          return
        }
        // 发送axios请求登陆
        this.httpLogin()
      })
    },
    cancel () {
      this.$Message.info('取消登陆')
    }
  }
}
</script>

<style scoped>
.top{
  width: 70px;
  font-size: 10px;
  padding: 10px;
  background: #40485a;
  color: #fff;
  text-align: center;
  border-radius: 2px;
}
.layout{
  border: 1px solid #d7dde4;
  background: #f5f7f9;
  position: relative;
  border-radius: 4px;
  overflow: hidden;
  min-height: 49rem;
}
.layout-logo{
  width: 100px;
  height: 30px;
  background: #5b6270;
  border-radius: 3px;
  float: left;
  position: relative;
  top: 15px;
  left: 20px;
}
.layout-nav{
  width: 420px;
  margin: 0 17.6rem;
  left: 4%;
  /*margin-right: 20px;*/
}
.layout-footer-center{
  text-align: center;
}
.demo-avatar-badge >>> .ivu-badge-count{
  top: 11px;
  min-width: 10px;
  height: 13px;
  line-height: 10px;
  right: 4px;
}
.demo-avatar-badge >>> .bgEdt {
  left: 95%;
  position: absolute;
  padding: 3px;
}
.demo-avatar-badge >>> .bgEdt2 {
  left: 98%;
  position: absolute;
  padding: 3px;
}
.col-6-style {
  background-color: white;
  height: 50rem;
  top: 64px;
}
.ivu-layout >>> .ivu-layout-sider {
  width: 21rem;
}

.dateHandle {
  display: flex;
  justify-content: center;
  padding: 2rem 0;
}
.btn span {
  border-right: #dad9d5 1px solid;
  padding: 0 10px;
}
.title-item-class{
  float: left;
  margin-left: 3rem;
  color: #c0c3c8;
  margin-top: 0.5rem;
}
.title-item-class:hover{
  color: white;
  cursor: pointer;
}
.posts-archives {
  cursor: pointer;
}
.posts-archives:hover{
  color: #c0c3c8;
  display: -webkit-box;
  /* -webkit-box-orient: vertical; */
  /*! autoprefixer: off */
  -webkit-box-orient: vertical;
  /* autoprefixer: on */
  -webkit-line-clamp: 2;
  overflow: hidden;
}
.breadHandle:hover {
  color: white;
}
.footer{
  min-height: 4rem;
  background: #272c31;
  color: #b3b6bb;
  text-align: center;
  line-height: 4rem;
  font-size: 18px;
}
</style>
