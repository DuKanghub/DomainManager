<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="所属账号:">
          <el-select v-model="searchInfo.account" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in DnsRamUserOptions" :key="key" :label="item.label" :value="item.value"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="dns提供商:">
          <el-select v-model="searchInfo.dnsProvider" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in CloudPlatFormOptions" :key="key" :label="item.label" :value="item.value"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="公司:">
          <el-select v-model="searchInfo.company" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in companyOptions" :key="key" :label="item.label" :value="item.value"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="域名:">
          <el-input placeholder="dukanghub.com" v-model="searchInfo.domain" @keyup.enter.native="onSubmit" />
        </el-form-item>
        <el-form-item label="解析数">
          <el-input placeholder="大于多少" v-model="searchInfo.recordTotal" @keyup.enter.native="onSubmit" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
          <el-button icon="el-icon-edit" size="mini" type="primary" @click="updateDialog">批量更新</el-button>
          <el-button icon="el-icon-edit" size="mini" type="primary" @click="flushdomains">刷新域名缓存</el-button>
          <el-popover v-model="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <el-button slot="reference" icon="el-icon-delete" size="mini" type="danger" style="margin-left: 10px;">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      ref="multipleTable"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      :data="tableData"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column label="更新日期" width="180">
        <template slot-scope="scope">{{ scope.row.UpdatedAt|formatDate }}</template>
      </el-table-column>
      <el-table-column label="域名" prop="domain" width="120" />
      <el-table-column label="公司" prop="company" width="120">
        <template slot-scope="scope">
          {{filterDict(scope.row.company,"company")}}
        </template>
      </el-table-column>
      <el-table-column label="解析数" prop="recordTotal" width="120" /> 
      <el-table-column label="dns提供商" prop="dnsProvider" width="120">
        <template slot-scope="scope">
          {{filterDict(scope.row.dnsProvider,"CloudPlatForm")}}
        </template>
      </el-table-column>
      <el-table-column label="所属账号" prop="account" width="120">
        <template slot-scope="scope">
          {{filterDict(scope.row.account,"DnsRamUser")}}
        </template>
      </el-table-column>
      <el-table-column label="注册日期" prop="registerTime" width="180">
        <template slot-scope="scope">
          {{scope.row.registerTime|formatDate}}
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button size="mini" type="success" @click="toDetail(scope.row)">详情</el-button>
          <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="updateDomain(scope.row)">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 20, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
    <!-- 变更修改的弹窗 -->
    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="域名:" v-if='isShow'>
          <el-input v-model="formData.domain" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="公司:">
          <el-select v-model="formData.company" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in companyOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="解析数:" v-if="isShow">
          <el-input v-model.number="formData.recordTotal" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="dns提供商:">
          <el-select v-model="formData.dnsProvider" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in CloudPlatFormOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="所属账号:">
          <el-select v-model="formData.account" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in DnsRamUserOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="注册日期:" v-if='isShow'>
          <el-date-picker type="datetime" placeholder="选择日期" v-model="formData.registerTime" clearable />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createDomain,
  deleteDomain,
  deleteDomainByIds,
  updateDomain,
  updateDomains,
  findDomain,
  flushDomains,
  getDomainList
} from '@/api/domainList' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'
export default {
  name: 'DomainList',
  mixins: [infoList],
  data() {
    return {
      listApi: getDomainList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      updateVisible: false,
      multipleSelection: [],
      companyOptions: [],
      CloudPlatFormOptions: [],
      DnsRamUserOptions: [],
      formData: {
        domain: '',
        company: 0,
        recordTotal: 0,
        dnsProvider: 0,
        account: 0,
        registerTime: new Date()
      },
      isShow: true
    }
  },
  filters: {
    formatDate: function(time) {
      if (time !== null && time !== '') {
        var date = new Date(time)
        return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
      } else {
        return ''
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? '是' : '否'
      } else {
        return ''
      }
    }
  },
  async created() {
    await this.getTableData()
    await this.getDict('company')
    await this.getDict('CloudPlatForm')
    await this.getDict('DnsRamUser')
  },
  methods: {
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    toDetail(row) {
      this.$router.push({
        name: 'recordList',
        params: {
          domain: row.domain
        }
      })
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteDomain(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteDomainByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async flushdomains() {
      const res = await flushDomains()
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '更新成功'
        })
        this.getTableData()
      }
    },
    // 更新单条数据
    async updateDomain(row) {
      this.isShow = true
      const res = await findDomain({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.reDomain
        this.dialogFormVisible = true
      }
    },
    // 批量更新
    async onUpdateMulti(formData) {
      const domains = this.$refs.multipleTable.selection
      for (let i = 0; i < domains.length; i++) {
        domains[i].company = formData.company || domains[i].company
        domains[i].dnsProvider = formData.dnsProvider || domains[i].dnsProvider
        domains[i].account = formData.account || domains[i].account
      }
      const res = await updateDomains({ domains })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '更新成功'
        })
        if (this.tableData.length === domains.length && this.page > 1) {
          this.page--
        }
        this.updateVisible = false
        this.getTableData()
        this.closeDialog()
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        domain: '',
        company: 0,
        recordTotal: 0,
        dnsProvider: 0,
        account: 0,
        registerTime: new Date()
      }
    },
    async deleteDomain(row) {
      const res = await deleteDomain({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case 'create':
          res = await createDomain(this.formData)
          break
        case 'update':
          res = await updateDomain(this.formData)
          break
        case 'updateMulti':
          this.onUpdateMulti(this.formData)
          break
        default:
          res = await createDomain(this.formData)
          break
      }
      if (this.type !== 'updateMulti') {
        alert(this.type)
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '创建/更改成功'
          })
          this.closeDialog()
          this.getTableData()
        }
      }
    },
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
      this.isShow = true
    },
    updateDialog() {
      if (this.$refs.multipleTable.selection.length <= 0) {
        this.$message.error('请选择需要批量更新的内容')
      } else {
        this.formData = {}
        this.type = 'updateMulti'
        this.dialogFormVisible = true
        this.isShow = false
      }
    }
  }
}
</script>

<style>
</style>
