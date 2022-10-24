<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="listQuery.info" placeholder="服务名称/服务描述" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">
        搜索
      </el-button>
      <router-link :to="'/service/service_create_http'">
        <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">
          添加HTTP服务
        </el-button>
      </router-link>
      <router-link :to="'/service/service_create_tcp'">
        <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">
          添加TCP服务
        </el-button>
      </router-link>
      <router-link :to="'/service/service_create_grpc'">
        <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleCreate">
          添加GRPC服务
        </el-button>
      </router-link>
    </div>

    <el-table
      :key="tableKey"
      v-loading="listLoading"
      :data="list"
      element-loading-text="拼命加载中"
      element-loading-spinner="el-icon-loading"
      height="500"
      :border="false"
      stripe
      fit
      highlight-current-row
      style="width: 100%;"
    >
      <el-table-column fixed sortable label="ID" prop="id" align="center" width="60">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="服务名称" align="center" min-width="130px">
        <template slot-scope="{row}">
          <span>{{ row.service_name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="服务描述" align="center" min-width="150px">
        <template slot-scope="{row}">
          <span>{{ row.service_desc }}</span>
        </template>
      </el-table-column>
      <el-table-column label="服务类型" align="center" min-width="100px">
        <template slot-scope="{row}">
          <span>{{ row.load_type | typeFilter }}</span>
        </template>
      </el-table-column>
      <el-table-column label="服务地址" align="center" min-width="220px">
        <template slot-scope="{row}">
          <span>{{ row.service_addr }}</span>
        </template>
      </el-table-column>
      <el-table-column label="QPS" align="center" min-width="80px">
        <template slot-scope="{row}">
          <span>{{ row.qps }}</span>
        </template>
      </el-table-column>
      <el-table-column label="日请求量" align="center" min-width="90px">
        <template slot-scope="{row}">
          <span>{{ row.qpd }}</span>
        </template>
      </el-table-column>
      <el-table-column label="节点数" align="center" min-width="80px">
        <template slot-scope="{row}">
          <span>{{ row.total_node }}</span>
        </template>
      </el-table-column>

      <el-table-column fixed="right" label="操作" align="center" width="250" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <router-link :to="'/service/service_stat/'+row.id">
            <el-button type="primary" size="mini" style="margin-right: 10px;">
              统计
            </el-button>
          </router-link>
          <router-link v-if="row.load_type===0" :to="'/service/service_edit_http/'+row.id">
            <el-button type="primary" size="mini" style="margin-right: 10px;">
              修改
            </el-button>
          </router-link>
          <router-link v-if="row.load_type===1" :to="'/service/service_edit_tcp/'+row.id">
            <el-button type="primary" size="mini" style="margin-right: 10px;">
              修改
            </el-button>
          </router-link>
          <router-link v-if="row.load_type===2" :to="'/service/service_edit_grpc/'+row.id">
            <el-button type="primary" size="mini" style="margin-right: 10px;">
              修改
            </el-button>
          </router-link>
          <el-button size="mini" type="danger" @click="handleDelete(row,$index)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />
  </div>
</template>

<script>
import { serviceList, serviceDelete } from '@/api/service'
import waves from '@/directive/waves' // waves directive
import Pagination from '@/components/Pagination' // secondary package based on el-pagination

const serviceTypeOptions = [
  { key: '0', display_name: 'HTTP' },
  { key: '1', display_name: 'TCP' },
  { key: '2', display_name: 'GRPC' }
]

const serviceTypeKeyValue = serviceTypeOptions.reduce((acc, cur) => {
  acc[cur.key] = cur.display_name
  return acc
}, {})

export default {
  name: 'ServiceList',
  components: { Pagination },
  directives: { waves },
  filters: {
    typeFilter(type) {
      return serviceTypeKeyValue[type]
    }
  },
  data() {
    return {
      tableKey: 0,
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10,
        info: ''
      },
      temp: {
        id: undefined
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      serviceList(this.listQuery).then(response => {
        // console.log(response.data.total, this.total)
        this.list = response.data.list
        if (response.data.total !== 0) {
          this.total = response.data.total
        }
        // Just to simulate the time of the request
        setTimeout(() => {
          this.listLoading = false
        }, 0.5 * 1000)
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        importance: 1,
        remark: '',
        timestamp: new Date(),
        title: '',
        status: 'published',
        type: ''
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handleDelete(row, index) {
      this.$confirm('此操作将删除该记录, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        const deleteQuery = {
          'id': row.id
        }
        this.listLoading = true
        serviceDelete(deleteQuery).then(response => {
          this.$notify({
            title: 'Success',
            message: '删除成功',
            type: 'success',
            duration: 2000
          })
          this.getList()
        })
      }).catch(() => {
        this.$notify({
          title: '信息',
          message: '已取消删除',
          type: 'info',
          duration: 2000
        })
      })
    }
  }
}
</script>
