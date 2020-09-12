<template>
  <div class="mavonEditor">
    <Form ref="formInline" :model="formInline" :rules="ruleInline" inline>
      <Row>
        <Col span="24">
          <FormItem prop="title">
            <Input type="text" v-model="formInline.title" placeholder="请输入标题名"  :style="{width:'26rem'}">
              <Icon type="md-bookmarks"  slot="prepend"/>
            </Input>
          </FormItem>
        </Col>
      </Row>
      <Row>
        <Col span="24">
          <FormItem prop="preface">
            <Input type="textarea" v-model="formInline.preface" :style="{width:'26rem'}" :rows="8" placeholder="请输入文章前言"></Input>
          </FormItem>
        </Col>
      </Row>
      <Row>
        <Col span="24">
          <FormItem prop="category_id">
            <Icon type="ios-keypad" slot="prepend"/>
            <i-select v-model="formInline.category_id" :style="{width:'26rem'}">
              <i-option v-for="item in category_list" :value="item.id" :key="item.id">{{ item.name }}</i-option>
            </i-select>
          </FormItem>
        </Col>
      </Row>
      <Row>
        <Col span="24">
          <Form-item>
            <div class="demo-upload-list" v-for="item in uploadList" :key="item.uid">
              <template v-if="item.status === 'finished'">
                <img :src="item.url">
                <div class="demo-upload-list-cover">
                  <Icon type="ios-eye-outline" @click.native="handleView(item.name)"></Icon>
                  <Icon type="ios-trash-outline" @click.native="handleRemove(item)"></Icon>
                </div>
              </template>
              <template v-else>
                <Progress v-if="item.showProgress" :percent="item.percentage" hide-info></Progress>
              </template>
            </div>
            <Upload
              ref="upload"
              :show-upload-list="false"
              :default-file-list="defaultList"
              :on-success="handleSuccess"
              :format="['jpg','jpeg','png']"
              :max-size="2048"
              :on-format-error="handleFormatError"
              :on-exceeded-size="handleMaxSize"
              :before-upload="handleBeforeUpload"
              type="drag"
              action="http://127.0.0.1:8888/upload/title/picture"
              style="display: inline-block;width:58px;">
              <div style="width: 58px;height:58px;line-height: 58px;">
                <Icon type="ios-camera" size="20"></Icon>
              </div>
            </Upload>
            <Modal title="图片展示" v-model="visible">
              <img :src="'http://127.0.0.1:8888/export/title/picture/?filename=' + imgName" v-if="visible" style="width: 100%">
            </Modal>
          </Form-item>
        </Col>
      </Row>
      <Row>
        <Col span="24">
          <FormItem>
            <mavon-editor :style="{ zIndex: '0',width:'54rem'}" ref="md" :toolbars="markdownOption" v-model="formInline.content" @imgAdd="imgAdd" @imgDel="imgDel"/>
          </FormItem>
        </Col>
      </Row>
      <FormItem>
        <Button type="primary" @click="handleSubmit('formInline')">提交文章</Button>
      </FormItem>
    </Form>
  </div>
</template>
<script>
import axios from 'axios'
export default {
  props: {
    category_list: {
      type: Array
    }
  },
  data () {
    return {
      markdownOption: {
        bold: true, // 粗体
        italic: true, // 斜体
        header: true, // 标题
        underline: true, // 下划线
        mark: true, // 标记
        superscript: true, // 上角标
        quote: true, // 引用
        ol: true, // 有序列表
        link: true, // 链接
        imagelink: true, // 图片链接
        help: true, // 帮助
        code: true, // code
        subfield: true, // 是否需要分栏
        fullscreen: true, // 全屏编辑
        readmodel: true, // 沉浸式阅读
        /* 1.3.5 */
        undo: true, // 上一步
        trash: true, // 清空
        save: true, // 保存（触发events中的save事件）
        /* 1.4.2 */
        navigation: true // 导航目录
      },
      // category_list: [],
      formInline: {
        // 图片head
        head_img: '',
        // 标题
        title: '',
        // 类别
        category_id: '',
        // 文章
        content: '',
        // 前言
        preface: ''
      },
      ruleInline: {
        title: [
          { required: true, message: '文章标题不能为空', trigger: 'blur' }
        ],
        category_id: [
          { required: true, message: '文章类型不能为空', trigger: 'blur' }
        ],
        content: [
          { required: true, message: '文章内容不能为空', trigger: 'blur' }
        ],
        preface: [
          { required: true, message: '文章前言不能为空', trigger: 'blur' }
        ]
      },
      img_file: {},
      // 文章图片变量
      defaultList: [],
      imgName: '',
      visible: false,
      uploadList: []
    }
  },
  mounted () {
    this.uploadList = this.$refs.upload.fileList
  },
  methods: {
    // 文章头图片上传
    handleView (name) {
      this.imgName = name
      this.visible = true
    },
    async httpRemoveTitlePic (removeFileName) {
      const { data: { code, msg } } = await this.$http.delete('/remove/title/picture/' + removeFileName)
      if (code === 200) {
        this.formInline.head_img = ''
        this.$Message.success(msg)
      } else {
        this.$Message.error(msg)
      }
    },
    handleRemove (file) {
      const fileList = this.$refs.upload.fileList
      var removeFileName = this.$refs.upload.fileList.splice(fileList.indexOf(file), 1)[0].name
      if (!removeFileName) {
        this.$Message.error('要删除图片不存在')
        return
      }
      this.httpRemoveTitlePic(removeFileName)
    },
    handleSuccess (res, file) {
      this.formInline.head_img = res.result.urlpath
      file.url = res.result.urlpath
      file.name = res.result.filename
    },
    handleFormatError (file) {
      this.$Notice.warning({
        title: '标题图片格式不正确。',
        desc: '上传图片名 ' + file.name + ' 图片格式不正确请选jpg 或 png.'
      })
    },
    handleMaxSize (file) {
      this.$Notice.warning({
        title: '超出单张图片限制',
        desc: '图片  ' + file.name + ' 超过了2M，请传更小的图片.'
      })
    },
    handleBeforeUpload () {
      const check = this.uploadList.length < 1
      if (!check) {
        this.$Notice.warning({
          title: '标题图片最多上传一个'
        })
      }
      return check
    },
    imgAdd (pos, $file) {
      // 第一步.将图片上传到服务器.
      var formdata = new FormData()
      formdata.append('image', $file)
      this.img_file[pos] = $file
      axios({
        // 请求地址
        url: this.$http.defaults.baseURL + '/upload/content/picture',
        method: 'post',
        data: formdata,
        headers: { 'Content-Type': 'multipart/form-data' }
      }).then((data) => {
        // 第二步.将返回的url替换到文本原位置![...](./0) -> ![...](url)
        /**
         * $vm 指为mavonEditor实例，可以通过如下两种方式获取
         * 1.  通过引入对象获取: `import {mavonEditor} from ...` 等方式引入后，`$vm`为`mavonEditor`
         * 2. 通过$refs获取: html声明ref : `<mavon-editor ref=md ></mavon-editor>，`$vm`为 `this.$refs.md`
         * 3. 由于vue运行访问的路径只能在static下，so，我就把图片保存到它这里了
         */
        this.$refs.md.$img2Url(pos, data.data.result.urlpath)
      })
    },
    imgDel (pos) {
      delete this.img_file[pos]
    },
    async httpPosts () {
      const { data: { code, msg } } = await this.$http.post('/posts', this.formInline)
      if (code === 200) {
        this.$Message.success(msg)
        this.$router.push({ path: 'posts' })
      } else {
        this.$Message.error(msg)
      }
    },
    handleSubmit (name) {
      var token = window.sessionStorage.getItem('token')
      if (!token) {
        this.$Message.error('请先登陆,才可以上传文章')
        return
      }
      if (!this.formInline.content) {
        this.$Message.error('您没有输入文章内容')
        return
      }
      if (!this.formInline.head_img) {
        this.$Message.error('您没有上传文章图片')
        return
      }
      this.$refs[name].validate((valid) => {
        if (!valid) {
          this.httpPosts()
        } else {
          this.$Message.error('上传文章不合法!')
        }
      })
    }
  }
}
</script>
<style scoped>
.mavonEditor {
  padding: 2rem 4rem;
  background: #fff;
  height: 100%;
  min-height: 45rem;
}
.demo-upload-list{
  display: inline-block;
  width: 60px;
  height: 60px;
  text-align: center;
  line-height: 60px;
  border: 1px solid transparent;
  border-radius: 4px;
  overflow: hidden;
  background: #fff;
  position: relative;
  box-shadow: 0 1px 1px rgba(0,0,0,.2);
  margin-right: 4px;
}
.demo-upload-list img{
  width: 100%;
  height: 100%;
}
.demo-upload-list-cover{
  display: none;
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  background: rgba(0,0,0,.6);
}
.demo-upload-list:hover .demo-upload-list-cover{
  display: block;
}
.demo-upload-list-cover i{
  color: #fff;
  font-size: 20px;
  cursor: pointer;
  margin: 0 2px;
}
</style>
