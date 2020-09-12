<template>
  <Content :style="{ background: '#fff', minHeight: '70rem', padding: '1rem'}">
    <div style="padding: 2rem">
      <List item-layout="vertical" v-if="this.article_list">
        <ListItem v-for="item in article_list" :key="item.title">
          <div style="position: flex;margin-left: 20rem">
            <Tag color="purple">
              <Icon type="ios-book"></Icon>
              <span>{{item.Category.name}}</span>
            </Tag>
          </div>
          <ListItemMeta :avatar="item.avatar" :title="item.title" :description="item.preface" />
          创建时间:{{ item.created_at }}
          <template slot="action">
            <li :style="{marginTop: '3rem'}" :class="item.iszan ? 'zanHandle': 'nozanHandle'">
              <Icon type="ios-thumbs-up-outline" @click="zanHandle(item)"/> {{item.zan_count}}
            </li>
            <li>
              <Icon type="ios-eye" /> {{item.views}}
            </li>
            <li>
              <Button :style="{left:'2rem', height:'22px', lineHeight:'22px', fontSize:'10px', color: '#818694'}" type="warning" icon="ios-search" shape="circle" @click="postDetail(item.id)">立刻阅读</Button>
            </li>
          </template>
          <template slot="extra" :style="{marginTop: '-10px'}">
            <img :src="item.head_img" style="width: 280px">
          </template>
        </ListItem>
      </List>
      <List v-else>
        <ListItem>
            <div>文章消失在遥远的🌕</div>
        </ListItem>
      </List>
    </div>
    <!--          -->
    <div class="pageHandle">
      <Page :total="total" show-elevator :page-size="pageSize" @on-change="changepage" @on-page-size-change="changepagesize"></Page>
    </div>

  </Content>
</template>

<script>
export default {
  name: 'posts',
  props: {
    tagId: {
      type: Number
    },
    yearMonth: {
      type: String
    }
  },
  data () {
    return {
      total: 1,
      pageSize: 10,
      article_list: '',
      zanData: {
        status: '',
        id: ''
      }
    }
  },
  created () {
    this.httpPostList(1)
    this.tagChange()
  },
  watch: {
    tagId (newVal) {
      this.httpPostList(1)
    },
    yearMonth (newVal) {
      this.httpPostList(1)
    }
  },
  methods: {
    async httpZanPost (status, id) {
      this.zanData.status = status
      this.zanData.id = id
      const { data: { code } } = await this.$http.post('/api/zan', this.zanData)
      if (code === 200) {
      }
    },
    zanHandle (item) {
      var status = 'add'
      if (item.iszan) {
        item.zan_count -= 1
        status = 'sub'
      } else {
        item.zan_count += 1
      }
      item.iszan = !item.iszan
      this.httpZanPost(status, item.id)
    },
    tagChange () {
      console.log('tag为:', this.tagId)
    },
    postDetail (postsId) {
      this.$router.push({path: `article/${postsId}`})
    },
    //
    async httpPostList (pageNum) {
      const { data: { code, msg, result } } = await this.$http.get('/posts/page/list', {params: {'pageNum': pageNum, 'pageSize': this.pageSize, 'tagId': this.tagId, 'yearMonth': this.yearMonth}})
      if (code === 200) {
        this.total = result.total
        this.article_list = result.data
        // 文章刷新
        this.$Message.success(msg)
      } else {
        this.$Message.error(msg)
      }
    },
    changepage (pageNum) {
      this.httpPostList(pageNum)
    },
    changepagesize (pageSize) {
      this.pageSize = pageSize
    }
  }
}
</script>

<style scoped>
  .pageHandle {
    width: 100%;
    min-height: 8rem;
    display: flex;
    justify-content: center;
  }
  .ivu-page-options-elevator >>> input {
    width: 2rem;
  }
  .zanHandle {
    color: orange;
  }
  .nozanHandle {
    color: #797979;
  }
</style>
