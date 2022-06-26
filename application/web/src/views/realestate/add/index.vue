<template>
  <div class="app-container">
    <el-form ref="ruleForm" v-loading="loading" :model="ruleForm" :rules="rules" label-width="100px">

      <el-form-item label="客户" prop="proprietor">
        <el-select v-model="ruleForm.proprietor" placeholder="请选择客户" @change="selectGet">
          <el-option
            v-for="item in accountList"
            :key="item.accountId"
            :label="item.userName"
            :value="item.accountId"
          >
            <span style="float: left">{{ item.userName }}</span>
            <span style="float: right; color: #8492a6; font-size: 13px">{{ item.accountId }}</span>
          </el-option>
        </el-select>
      </el-form-item>


      <!-- <el-form-item label="总空间 ㎡" prop="totalArea">
        <el-input-number v-model="ruleForm.totalArea" :precision="2" :step="0.1" :min="0" />
      </el-form-item> -->
      <el-form-item label="作品名称" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.carmodel" placeholder="请输入内容"></el-input>
      </el-form-item>
      <el-form-item label="原作者" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.car_scz" placeholder="请输入内容"></el-input>
      </el-form-item>
      <el-form-item label="创作时间" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.car_scd" placeholder="请输入内容"></el-input>
      </el-form-item>
      <el-form-item label="作品简介" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.car_scsj" placeholder="请输入内容"></el-input>
      </el-form-item>
      <el-form-item label="作品分类" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.car_lbjph" placeholder="请输入内容"></el-input>
      </el-form-item>
      <el-form-item label="受众人群" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.car_lbjscz" placeholder="请输入内容"></el-input>
      </el-form-item>
      <!-- <el-form-item label="商品零部件生产地" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.car_lbjscd" placeholder="请输入内容"></el-input>
      </el-form-item>
      <el-form-item label="商品零部件生产时间" prop="totalArea">
        <el-input size="medium"  v-model="ruleForm.car_lbjscsj" placeholder="请输入内容"></el-input>
      </el-form-item> -->

      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAccountList } from '@/api/account'
import { createRealEstate } from '@/api/realEstate'

export default {
  name: 'AddRealeState',
  data() {
    var checkArea = (rule, value, callback) => {
      if (value <= 0) {
        callback(new Error('必须大于0'))
      } else {
        callback()
      }
    }
    return {
      ruleForm: {
        proprietor: '',
        // totalArea: 0,
        // livingSpace: 0
        carmodel: '',
        car_scz: '',
        car_scd: '',
        car_scsj: '',
        car_lbjph: '',
        car_lbjscz: '',
        car_lbjscd: '',
        car_lbjscsj: '',

      },
      accountList: [],
      rules: {
        proprietor: [
          { required: true, message: '请选择客户', trigger: 'change' }
        ],
        // totalArea: [
        //   { validator: checkArea, trigger: 'blur' }
        // ],
        // livingSpace: [
        //   { validator: checkArea, trigger: 'blur' }
        // ]
      },
      loading: false
    }
  },
  computed: {
    ...mapGetters([
      'accountId'
    ])
  },
  created() {
    queryAccountList().then(response => {
      if (response !== null) {
        // 过滤掉管理员
        this.accountList = response.filter(item =>
          item.userName !== '管理员'
        )
      }
    })
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('是否立即创建?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }).then(() => {
            this.loading = true
            createRealEstate({
              accountId: this.accountId,
              proprietor: this.ruleForm.proprietor,
              // totalArea: this.ruleForm.totalArea,
              // livingSpace: this.ruleForm.livingSpace
              carmodel: this.ruleForm.carmodel,
              car_scz: this.ruleForm.car_scz,
              car_scd: this.ruleForm.car_scd,
              car_scsj: this.ruleForm.car_scsj,
              car_lbjph: this.ruleForm.car_lbjph,
              car_lbjscz: this.ruleForm.car_lbjscz,
              car_lbjscd: this.ruleForm.car_lbjscd,
              car_lbjscsj: this.ruleForm.car_lbjscsj,

            }).then(response => {
              this.loading = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: '创建成功!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: '创建失败!'
                })
              }
            }).catch(_ => {
              this.loading = false
            })
          }).catch(() => {
            this.loading = false
            this.$message({
              type: 'info',
              message: '已取消创建'
            })
          })
        } else {
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    selectGet(accountId) {
      this.ruleForm.proprietor = accountId
    }
  }
}
</script>

<style scoped>
</style>
