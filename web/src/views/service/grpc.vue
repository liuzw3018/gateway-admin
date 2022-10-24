<template>
  <div class="mixin-components-container">
    <el-row>
      <el-card class="box-card">
        <div slot="header" class="clearfix">
          <span v-if="isEdit===false">创建GRPC服务</span>
          <span v-if="isEdit===true">修改GRPC服务</span>
        </div>
        <div style="margin-bottom:50px;">
          <el-form ref="form" :model="form" label-width="130px">
            <el-form-item label="服务名称" class="is-required">
              <el-input v-model="form.service_name" :disabled="isEdit" placeholder="6-128英文数字下划线" />
            </el-form-item>
            <el-form-item label="服务描述" class="is-required">
              <el-input v-model="form.service_desc" placeholder="最多255字符，必填" />
            </el-form-item>
            <el-form-item label="服务端口" class="is-required">
              <el-input v-model="form.port" :disabled="isEdit" placeholder="端口范围8001-8999，必填" />
            </el-form-item>

            <el-form-item label="metadata转换">
              <el-input
                v-model="form.header_transfor"
                type="textarea"
                autosize
                placeholder="metadata转换支持 增加(add)，删除(del)，修改(edit)；格式：add head value，多条换行"
              />
            </el-form-item>

            <el-form-item label="开启验证">
              <el-switch v-model="form.open_auth" :active-value="1" :inactive-value="0" />
            </el-form-item>

            <el-form-item label="IP白名单">
              <el-input
                v-model="form.white_list"
                type="textarea"
                autosize
                placeholder="格式：127.0.0.1 多条换行，白名单优先于黑名单"
              />
            </el-form-item>

            <el-form-item label="IP黑名单">
              <el-input
                v-model="form.black_list"
                type="textarea"
                autosize
                placeholder="格式：127.0.0.1 多条换行，白名单优先于黑名单"
              />
            </el-form-item>

            <el-form-item label="客户端限流">
              <el-input v-model="form.clientip_flow_limit" placeholder="0表示无限制" />
            </el-form-item>
            <el-form-item label="服务端限流">
              <el-input v-model="form.service_flow_limit" placeholder="0表示无限制" />
            </el-form-item>

            <el-form-item label="轮询方式">
              <el-radio-group v-model="form.round_type">
                <el-radio :label="0">random</el-radio>
                <el-radio :label="1">round-robin</el-radio>
                <el-radio :label="2">weight_round-robin</el-radio>
                <el-radio :label="3">ip_hash</el-radio>
              </el-radio-group>
            </el-form-item>

            <el-form-item label="IP列表" class="is-required">
              <el-input
                v-model="form.ip_list"
                type="textarea"
                autosize
                placeholder="格式：127.0.0.1:80 多条换行"
              />
            </el-form-item>
            <el-form-item label="权重列表" class="is-required">
              <el-input v-model="form.weight_list" type="textarea" autosize placeholder="格式：50 多条换行" />
            </el-form-item>

            <el-form-item>
              <el-button v-if="isEdit===false" type="primary" :disabled="submitDisable" @click="handleCreate">立即创建</el-button>
              <el-button v-if="isEdit===true" type="primary" :disabled="submitDisable" @click="handleCreate">立即修改</el-button>
              <el-button @click="$store.dispatch('tagsView/delView', $route);$router.push('/service/list')">取消</el-button>
            </el-form-item>
          </el-form>
        </div>
      </el-card>
    </el-row>
  </div>
</template>

<script>
import { addGRPCService, updateGRPCService, getServiceDetail } from '@/api/service'

export default {
  name: 'ServiceCreateGRPC',
  data() {
    return {
      isEdit: false,
      submitDisable: false,
      form: {
        service_name: '',
        service_desc: '',
        port: '',
        open_auth: '',
        white_list: '',
        black_list: '',
        white_host_name: '',
        header_transfor: '',
        clientip_flow_limit: '',
        service_flow_limit: '',
        round_type: 2,
        forbid_list: '',
        ip_list: '',
        weight_list: ''
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
      getServiceDetail(query).then(response => {
        this.form.id = response.data.info.id
        this.form.service_name = response.data.info.service_name
        this.form.service_desc = response.data.info.service_desc

        this.form.port = response.data.grpc.port
        this.form.header_transfor = response.data.grpc.header_transfor

        this.form.black_list = response.data.access_control.black_list.replace(/,/g, '\n')
        this.form.clientip_flow_limit = response.data.access_control.clientip_flow_limit
        this.form.open_auth = response.data.access_control.open_auth
        this.form.service_flow_limit = response.data.access_control.service_flow_limit
        this.form.white_host_name = response.data.access_control.white_host_name.replace(/,/g, '\n')
        this.form.white_list = response.data.access_control.white_list.replace(/,/g, '\n')

        this.form.forbid_list = response.data.load_balance.forbid_list.replace(/,/g, '\n')
        this.form.ip_list = response.data.load_balance.ip_list.replace(/,/g, '\n')
        this.form.round_type = response.data.load_balance.round_type
        this.form.upstream_connect_timeout = response.data.load_balance.upstream_connect_timeout
        this.form.upstream_header_timeout = response.data.load_balance.upstream_header_timeout
        this.form.upstream_idle_timeout = response.data.load_balance.upstream_idle_timeout
        this.form.upstream_max_idle = response.data.load_balance.upstream_max_idle
        this.form.weight_list = response.data.load_balance.weight_list.replace(/,/g, '\n')
      }).catch(() => {

      })
    },
    handleCreate(row, index) {
      this.submitDisable = true
      const createQuery = Object.assign({}, this.form)
      console.log(createQuery)
      createQuery.white_list = createQuery.white_list.replace(/\n/g, ',')
      createQuery.weight_list = createQuery.weight_list.replace(/\n/g, ',')
      createQuery.black_list = createQuery.black_list.replace(/\n/g, ',')
      createQuery.ip_list = createQuery.ip_list.replace(/\n/g, ',')

      createQuery.port = Number(createQuery.port)
      createQuery.clientip_flow_limit = Number(createQuery.clientip_flow_limit)
      createQuery.service_flow_limit = Number(createQuery.service_flow_limit)
      if (this.isEdit === true) {
        updateGRPCService(createQuery).then(response => {
          this.submitDisable = false
          this.$notify({
            title: 'Success',
            message: '修改成功',
            type: 'success',
            duration: 2000
          })
          this.$store.dispatch('tagsView/delView', this.$route)
          this.$router.push('/service/list')
        }).catch(() => {
          this.submitDisable = false
        })
      } else {
        addGRPCService(createQuery).then(response => {
          this.submitDisable = false
          this.$notify({
            title: 'Success',
            message: '添加成功',
            type: 'success',
            duration: 2000
          })
          this.$store.dispatch('tagsView/delView', this.$route)
          this.$router.push('/service/list')
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
