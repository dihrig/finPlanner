<template>
  <div>
    <h1>finPlanner Frontend</h1>
    <button @click="checkDb">Check DB Health</button>
    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script lang="ts">
import { ref } from "vue";
import axios from "axios";

export default {
  setup() {
    const message = ref<string | null>(null);

    const checkDb = async () => {
      try {
        const res = await axios.get("/api/db");
        message.value = JSON.stringify(res.data);
      } catch (err) {
        message.value = "DB check failed";
      }
    };

    return { message, checkDb };
  },
};
</script>

<style>
body {
  font-family: sans-serif;
  padding: 2rem;
}
</style>
