<template>
  <div class="login-section">
    
    <el-form label-position="top" label-width="100px" class="demo-ruleForm" :rules="rules" :model="rulesForm" status-icon ref="ruleForm">
      <h2>注册新用户</h2>
      <el-form-item label="用户名" prop="name">
        <el-input type="text" v-model="rulesForm.name"></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input type="password" v-model="rulesForm.password"></el-input>
      </el-form-item>
      <el-form-item class="register-buttons">
        <el-button type="primary" @click="submitForm('ruleForm')">注册</el-button>
        <el-button @click="closeRegisterDialog">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { mapMutations } from 'vuex'
export default {
  data() {
    return {
      //存储数据的对象
      rulesForm: {
        name: '',
        password: '',
      },
      rules: {
        name: [
          { required: true, message: '名字', trigger: 'blur' },
          { min: 3, max: 9, message: '长度在3到9', trigger: 'blur' },
        ],
        password: [
          { required: true, message: '密码', trigger: 'blur' },
          { min: 3, max: 16, message: '长度在3到16', trigger: 'blur' },
        ],
      },
    }
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          //如果校检通过，在这里向后端发送用户名和密码
          this.$store.dispatch('registerOptions/register', this.rulesForm)
        }
      })
    },
    ...mapMutations('registerOptions', {closeRegisterDialog: 'CLOSE_REGISTER_DIALOG'})
  },
}
</script>
<style scoped>
h2{
  text-align: center;
  font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
}
.login-section {
  /* 用来 mask */
  position: fixed; /*固定定位*/
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: hsla(120, 0%, 30%, 0.3);
  z-index: 98;
}
.el-form {
  position: fixed; /*固定定位*/
  top: 50%; /*距顶部50%*/
  left: 50%;
  /* margin: -100px 0 0 -150px; 设定这个div的margin-top的负值为自身的高度的一半,margin-left的值也是自身的宽度的一半的负值.(感觉在绕口令) */
  transform: translateX(-50%) translateY(-50%); /* 用平移实现居中更好，自适应 */
  width: 300px; /*宽为400,那么margin-top为-200px*/
  padding: 20px;
  border-radius: 8px;
  background-color: #f3f5f6;
  z-index: 99; /*浮动在最上层 */
}
.register-buttons {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
