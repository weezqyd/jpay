<template>
  <BaseHeader/>
  <div style="display: flex">
    <BaseSide/>
    <el-container style="margin: 1.2em">
      <el-col :span="24">
        <el-row>
          <el-container style="width: 100%">
            <el-table :data="tableData" style="width: 100%; margin: 1.2em" border>
              <el-table-column prop="id" label="#" width="50"/>
              <el-table-column prop="name" label="Name"/>
              <el-table-column prop="phone" label="Phone Number"/>
              <el-table-column prop="country">
                <template #header>
                  <el-select v-model="search" size="small" placeholder="Select country to filter" @change="getCustomers(search)">
                    <el-option
                        v-for="(country, idx) in filters"
                        :key="idx"
                        :label="country.toUpperCase()"
                        :value="country === 'All' ? '' : country"
                    />
                  </el-select>
                </template>
              </el-table-column>
            </el-table>
          </el-container>
        </el-row>
      </el-col>
    </el-container>
  </div>
</template>

<style>
</style>
<script lang="ts" setup>
import BaseHeader from "./components/layouts/BaseHeader.vue"
import BaseSide from "./components/layouts/BaseSide.vue"
import {ref} from 'vue'

const search = ref("")
const filters = ["All", "Unknown", "Morocco", "Uganda", "Cameroon", "Mozambique", "Ethiopia"]

let tableData = ref([])

async function getCustomers(country: string = "") {
  fetch(`http://localhost:9000/api/v1/customers?country=${country}`).then(res => res.json()).then(data => {
    tableData.value = data.data
  })
}

getCustomers()
</script>
