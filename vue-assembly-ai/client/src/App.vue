<script setup lang='ts'>
async function getText() {}
</script>

<template>
  <form @submit.prevent='getText'>
    select Your mp3 Clip
    <br />
    <input type='file' ref="file" />
    <br />
    <button>click to get result</button>
  </form>
</template>

<style scoped></style>
