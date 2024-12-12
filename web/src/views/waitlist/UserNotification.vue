<template>
  <div class="user-notification">
    <div class="message">
      Your account is currently on the waiting list. <br />Please contact the
      admin for further assistance.
    </div>
  </div>
</template>
<script setup lang="ts">
defineOptions({
  name: 'WaitlistNotification',
})
</script>
<style lang="scss" scoped>
.user-notification {
  display: flex;
  height: 100vh;
  justify-content: center;
  align-items: center;
  > .message {
    padding: 16px;
    background: #fff;
    border-radius: 6px;
    font-size: 18px;
    width: 500px;
  }
}
</style>
