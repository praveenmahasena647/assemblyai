<script setup lang="ts">
import { io } from "socket.io-client";
import { ref } from "vue";
let file = ref(null);

async function getText() {
}
</script>

<template>
  <form @submit.prevent="getText">
    select Your mp3 Clip
    <br />
    <input type="file" ref="file" />
    <br />
    <button>click to get result</button>
  </form>
</template>

<style scoped></style>
