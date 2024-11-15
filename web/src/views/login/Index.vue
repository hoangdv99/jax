<template>
  <div class="login-page">
    <div class="wrapper">
      <div class="form-container">
        <div class="logo">
          <img src="/images/logoG.png" alt="" class="image" />
        </div>
        <label for="email" class="label">Email</label>
        <InputText
          v-model="email"
          id="email"
          type="text"
          placeholder="Email address"
          class="input"
        />

        <label for="password" class="label">Password</label>
        <InputText
          v-model="password"
          id="password"
          type="password"
          placeholder="Password"
          class="input"
        />

        <div class="forgot">
          <span class="button">Forgot password?</span>
        </div>
        <Button label="Sign In" @click="login" />
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { useAccountStore } from '@/stores/account'
import { InputText, Button } from 'primevue'
import { ref } from 'vue'

const accountStore = useAccountStore()

defineOptions({
  name: 'LoginPage',
})

const email = ref('')
const password = ref('')

function login() {
  accountStore.login({ email: email.value, password: password.value })
}
</script>
<style lang="scss" scoped>
.login-page {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  > .wrapper {
    border-radius: 56px;
    padding: 0.3rem;
    background: linear-gradient(
      180deg,
      var(--primary-color) 10%,
      rgba(33, 150, 243, 0) 30%
    );
  }
}

.form-container {
  display: flex;
  flex-direction: column;
  padding: 5rem;
  border-radius: 53px;
  background-color: #fff;
  > .logo {
    display: flex;
    justify-content: center;
  }
  > .logo > .image {
    width: 4rem;
  }
  > .label {
    margin-top: 2rem;
    margin-bottom: 0.2rem;
    font-weight: 500;
    font-size: 1.25rem;
    color: #0f172a;
  }
  > .input {
    width: 30rem;
  }
  > .forgot {
    display: flex;
    justify-content: flex-end;
    margin-top: 1rem;
    margin-bottom: 2rem;
  }
  > .forgot > .button {
    color: #10b981;
    cursor: pointer;
  }
}
</style>
