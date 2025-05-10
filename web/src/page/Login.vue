<script lang="ts" setup>
import {reactive, ref} from 'vue'
import type {FormInstance, FormRules} from 'element-plus'

interface LoginForm {
  account: string
  email: string
  password: string
}

const loginFormRef = ref<FormInstance>()
const loginForm = reactive<LoginForm>({
  account: '',
  email: '',
  password: '',
})


const rules = reactive<FormRules<LoginForm>>({
  account: [
    {required: true, message: '请输入账号', trigger: 'blur'},
    {min: 3, max: 20, message: '长度3-20', trigger: 'blur'},
  ],
  password: [
    {
      required: true,
      message: '请输入密码',
      trigger: 'blur',
    },
  ],
})

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid, fields) => {
    if (valid) {
      console.log('submit!')
    } else {
      console.log('error submit!', fields)
    }
  })
}

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return
  formEl.resetFields()
}

</script>

<template>
  <el-form
      ref="loginFormRef"
      style="max-width: 600px"
      :model="loginForm"
      :rules="rules"
      label-width="auto"
  >
    <el-form-item label="账号" prop="account">
      <el-input v-model="loginForm.account"/>
    </el-form-item>
    <el-form-item label="密码" prop="password">
      <el-input v-model="loginForm.password"/>
    </el-form-item>

    <el-form-item>
      <el-button @click="resetForm(loginFormRef)">重置</el-button>
      <el-button type="primary" @click="submitForm(loginFormRef)">
        登陆
      </el-button>
    </el-form-item>
  </el-form>
</template>
