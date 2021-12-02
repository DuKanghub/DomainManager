<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="名称:">
        <el-input v-model="formData.name" clearable placeholder="请输入" />
      </el-form-item>

      <el-form-item label="用户名:">
        <el-input v-model="formData.username" clearable placeholder="请输入" />
      </el-form-item>

      <el-form-item label="密码:">
        <el-input v-model="formData.password" clearable placeholder="请输入" />
      </el-form-item>

      <el-form-item label="私钥:">
        <el-input v-model="formData.privateKey" clearable placeholder="请输入" />
      </el-form-item>

      <el-form-item label="启用sudo:">
        <el-switch v-model="formData.become" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
      </el-form-item>

      <el-form-item label="秘钥认证:">
        <el-switch v-model="formData.key_auth" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
      </el-form-item>

      <el-form-item label="备注:">
        <el-input v-model="formData.comment" clearable placeholder="请输入" />
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
  createSSHUser,
  updateSSHUser,
  findSSHUser
} from '@/api/sshUser' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'SSHUser',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
        name: '',
        username: '',
        password: '',
        privateKey: '',
        become: false,
        key_auth: false,
        comment: ''

      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findSSHUser({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.resshUser
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
          res = await createSSHUser(this.formData)
          break
        case 'update':
          res = await updateSSHUser(this.formData)
          break
        default:
          res = await createSSHUser(this.formData)
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
