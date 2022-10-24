<template>
  <div class="mixin-components-container">
    <el-row>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span v-if="isEdit===false">添加租户</span>
          <span v-if="isEdit===true">修改租户</span>
        </div>
        <div style="margin-bottom:50px;">
          <el-form ref="form" :model="form" label-width="130px">
            <el-form-item label="app_id" class="is-required">
              <el-input v-model="form.app_id" :disabled="isEdit" placeholder="6-128英文数字下划线" />
            </el-form-item>
            <el-form-item label="租户名称" class="is-required">
              <el-input v-model="form.name" placeholder="最多255字符，必填" />
            </el-form-item>
            <el-form-item label="32位密钥">
              <el-input v-model="form.secret" placeholder="32位字符串，非必填，自动生成" />
            </el-form-item>
            <el-form-item label="QPS限流">
              <el-input v-model="form.qps" placeholder="0表示无限制" />
            </el-form-item>
            <el-form-item label="日请求量限流">
              <el-input v-model="form.qpd" placeholder="0表示无限制" />
            </el-form-item>
            <el-form-item>
              <el-button v-if="isEdit === false" type="primary" :disabled="submitDisable" @click="handleCreate">立即创建</el-button>
              <el-button v-if="isEdit === true" type="primary" :disabled="submitDisable" @click="handleCreate">立即修改</el-button>
              <el-button @click="$store.dispatch('tagsView/delView', $route);$router.push('/app/list')">取消</el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-card>
    </el-row>
  </div>
</template>

<script>
import { appCreate, appDetail, appUpdate } from '@/api/app'

export default {
  name: 'AppCreate',
  data() {
    return {
      isEdit: false,
      submitDisable: false,
      form: {
        app_id: '',
        name: '',
        qpd: '',
        qps: '',
        secret: '',
        white_ips: '',
        id: ''
      }
    }
  },
  created() {
    const id = this.$route.params && this.$route.params.id
    if (id > 0) {
      this.isEdit = true
      this.fetchData(id)
    }
  },
  methods: {
    fetchData(id) {
      const query = { 'id': id }
      appDetail(query).then(response => {
        this.form.app_id = response.data.app_id
        this.form.name = response.data.name
        this.form.qpd = response.data.qpd
        this.form.qps = response.data.qps
        this.form.qps = response.data.qps
        this.form.secret = response.data.secret
        this.form.white_ips = response.data.white_ips
        this.form.id = response.data.id
      }).catch(() => {
      })
    },

    handleCreate(row, index) {
      this.submitDisable = true
      const query = Object.assign({}, this.form)
      if (this.isEdit) {
        appUpdate(query).then(response => {
          this.submitDisable = false
          this.$notify({
            title: 'Success',
            message: '修改成功',
            type: 'success',
            duration: 2000
          })
          this.$store.dispatch('tagsView/delView', this.$route)
          this.$router.push('/app/list')
        }).catch(() => {
          this.submitDisable = false
        })
      } else {
        appCreate(query).then(response => {
          this.submitDisable = false
          this.$notify({
            title: 'Success',
            message: '添加成功',
            type: 'success',
            duration: 2000
          })
          this.$store.dispatch('tagsView/delView', this.$route)
          this.$router.push('/app/list')
        }).catch(() => {
          this.submitDisable = false
        })
      }
    }
  }
}
</script>

<style scoped>
.mixin-components-container {
  background-color: #f0f2f5;
  padding: 30px;
  min-height: calc(100vh - 84px);
}

.component-item {
  min-height: 100px;
}

.el-select .el-input {
  width: 130px;
}

.input-with-select .el-input-group__prepend {
  background-color: #fff;
}
</style>
