<script setup lang="ts">
import WelcomeItem from './WelcomeItem.vue'
import DocumentationIcon from './icons/IconDocumentation.vue'
import ToolingIcon from './icons/IconTooling.vue'
import EcosystemIcon from './icons/IconEcosystem.vue'
import CommunityIcon from './icons/IconCommunity.vue'
import SupportIcon from './icons/IconSupport.vue'
import { ref, reactive, type Ref } from 'vue'
import axios from 'axios'


// give each packSize a unique id
let id = 0

let newPackSize = ref(0)
let packSizes = ref([])
let items: Ref<number> = ref(0)

let packsResult = reactive([{ count: 0, size: 0 }])
let summaryResult = {
  extraItems: 0,
  count: 0,
  total: 0
}

function addPackSize() {
  packSizes.value.push({ id: id++, size: newPackSize.value })
  newPackSize.value = 0
}

function removePackSize(packSize: any) {
  packSizes.value = packSizes.value.filter((t) => t !== packSize)
}

async function calculate() {
  let sizes: any = []
  packSizes.value.forEach(function(s) {
    sizes.push(parseInt(s.size))
  })

  console.log(sizes)

  const { data, status } = await axios.post(
      'http://localhost:8090/packing/calculate',
      {
        'items': parseInt(String(items.value)),
        'packs': sizes
      },
      {
        headers: {
          'Content-Type': 'application/json',
          'Access-Control-Allow-Origin': '*',
          withCredentials: true,
          mode: 'no-cors',
          Accept: 'application/json'
        }
      }
  )

  const packsRes = { sizes: Object.keys(data.Packs), counts: Object.values(data.Packs) }

  packsResult.splice(0)
  for (let i = 0; i < packsRes.sizes.length; i++) {
    packsResult.push({ size: parseInt(packsRes.sizes[i]), count: parseInt(<string>packsRes.counts[i]) })
  }

  summaryResult.total = data.Total
  summaryResult.count = data.Count
  summaryResult.extraItems = data.ExtraItems
}

</script>

<template>

  <WelcomeItem>
    <template #icon>
      <CommunityIcon />
    </template>
    <template #heading>Set Pack Sizes</template>
    <form @submit.prevent="addPackSize">
      <input v-model="newPackSize">
      <button>Add Pack Size</button>
    </form>
    <ul>
      <li v-for="packSize in packSizes" :key="id">
        {{ packSize.size }}
        <button @click="removePackSize(packSize)">X</button>
      </li>
    </ul>
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <CommunityIcon />
    </template>
    <template #heading>Set Items Count</template>
    <form @submit.prevent="">
      <input v-model="items">
      <br>
      <button @click="calculate()">Calculate</button>
    </form>

  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <CommunityIcon />
    </template>

    <h2>Packs Mix:</h2>
    <ul>
      <li v-for="p in packsResult" :key="p.size">
        {{ p.count }} packs of size {{ p.size }}
      </li>
    </ul>
    <h2>Summary</h2>
    <ul>
      <li>
        Items Count: {{ summaryResult.total }}
      </li>
      <li>
        Packs Count: {{ summaryResult.count }}
      </li>
      <li>
        Extra Items Count: {{ summaryResult.extraItems }}
      </li>
    </ul>
  </WelcomeItem>
</template>

