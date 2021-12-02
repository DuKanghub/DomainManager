<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="IP:">
        <el-input v-model="formData.ip" clearable placeholder="请输入" />
      </el-form-item>

      <el-form-item label="主机名:">
        <el-input v-model="formData.hostname" clearable placeholder="请输入" />
      </el-form-item>

      <el-form-item label="端口:">
        <el-input v-model.number="formData.port" clearable placeholder="请输入" />
      </el-form-item>

      <el-form-item label="启用:">
        <el-switch v-model="formData.active" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
      </el-form-item>

      <el-form-item label="用户ID:">
        <el-input v-model.number="formData.user_id" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item>
        <el-button size="mini" type="primary" @click="save">保存</el-button>
        <el-button size="mini" type="primary" @click="back">返回</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import {
  createHostInfo,
  updateHostInfo,
  findHostInfo
} from '@/api/hostInfo' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'HostInfo',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
        ip: '',
        hostname: '',
        port: 0,
        active: false,
        user_id: 0

      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findHostInfo({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.rehost
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createHostInfo(this.formData)
          break
        case 'update':
          res = await updateHostInfo(this.formData)
          break
        default:
          res = await createHostInfo(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>
