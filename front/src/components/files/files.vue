<template>
  <div>
    <Row style="background:#eee;padding:20px">
      <i-col span="24" :style="{paddingBottom: '1rem',width: '62rem'}" v-for="item in yearPostList" :key="item.Year">
        <Card>
          <p slot="title">{{item.Year}}年</p>
          <p v-for="item2 in item.PostInfo" :key="item2.ID">
            <span @click="postDetail(item2.ID)" :style="{cursor: 'pointer'}" class="articleHandle">
              {{item2.CreatedAt}}---{{item2.Title}}
            </span>
          </p>
        </Card>
      </i-col>
    </Row>
    <div :style="{minHeight: '20rem'}"></div>
  </div>
</template>

<script>
export default {
  name: 'files',
  data () {
    return {
      yearPostList: []
    }
  },
  created () {
    this.yearfiles()
  },
  methods: {
    postDetail (postsId) {
      this.$router.push({path: `/article/${postsId}`})
    },
    async yearfiles () {
      const { data: { code, msg, result } } = await this.$http.get('/posts/files/list')
      if (code === 200) {
        this.$Message.success(msg)
        this.yearPostList = result.data
      }
    }
  }
}
</script>

<style scoped>
.articleHandle {
  color: #333333;
}
.articleHandle:hover{
  color: #c0c3c8;
}
</style>
