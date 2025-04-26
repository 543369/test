<template>
  <div class="min-h-screen bg-white flex items-center justify-center p-4 font-sans">
    <div class="bg-white rounded-3xl shadow-[0_10px_30px_rgba(0,_0,_0,_0.1)] p-8 w-full max-w-md border border-gray-100">
      <h1 class="text-3xl font-bold text-center mb-8 text-gray-800">ConnectRPC 计算器</h1>
      
      <div class="space-y-6">
        <div class="relative">
          <input
            type="number"
            v-model="firstNumber"
            class="w-full px-5 py-4 rounded-xl border border-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition shadow-sm text-black"
            placeholder="第一个数字"
          />
        </div>
        
        <select
          v-model="operator"
          class="w-full px-5 py-4 rounded-xl border border-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition bg-white shadow-sm appearance-none cursor-pointer text-center font-bold text-lg text-black"
        >
          <option value="+">加法 (+)</option>
          <option value="-">减法 (-)</option>
          <option value="*">乘法 (×)</option>
          <option value="/">除法 (÷)</option>
        </select>
        
        <div class="relative">
          <input
            type="number"
            v-model="secondNumber"
            class="w-full px-5 py-4 rounded-xl border border-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition shadow-sm text-black"
            placeholder="第二个数字"
          />
        </div>
        
        <button
          @click="calculate"
          class="w-full bg-gradient-to-r from-blue-500 to-blue-700 text-white font-semibold py-4 rounded-xl transition-all duration-300 transform hover:scale-[1.02] active:scale-[0.98] focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 shadow-lg hover:shadow-xl"
        >
          计算
        </button>
        
        <div v-if="result !== null" class="p-5 bg-blue-50 border border-blue-100 rounded-xl text-center">
          <div class="text-blue-700 text-xl font-bold mb-1">计算结果</div>
          <div class="text-blue-600 text-2xl font-mono tracking-wider">
            {{ result }}
          </div>
        </div>
        
        <div v-if="error" class="mt-6 p-4 bg-red-50 border border-red-200 rounded-xl text-red-600 flex items-center">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2 flex-shrink-0" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
          </svg>
          <span>{{ error }}</span>
        </div>
      </div>
      
      <div class="mt-8 text-center text-gray-400 text-xs">
        © {{ new Date().getFullYear() }} ConnectRPC 计算器
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
// import { createConnectTransport } from '@bufbuild/connect-web';
// import { CalculatorServiceClient } from '../gen/calculator_connect';

// 创建Connect客户端
// const transport = createConnectTransport({
//   baseUrl: 'http://localhost:8080',
// });
// const client = new CalculatorServiceClient(transport);

// 状态
const firstNumber = ref('');
const secondNumber = ref('');
const operator = ref('+');
const result = ref(null);
const error = ref('');

// 计算方法
const calculate = async () => {
  try {
    error.value = '';
    
    // 验证输入
    const num1 = parseFloat(firstNumber.value);
    const num2 = parseFloat(secondNumber.value);
    
    if (isNaN(num1) || isNaN(num2)) {
      error.value = '请输入有效的数字';
      result.value = null;
      return;
    }
    
    // 调用后端服务
    const response = await client.calculate({
      firstNumber: num1,
      secondNumber: num2,
      operator: operator.value
    });
    
    result.value = response.result;
  } catch (err) {
    error.value = err.message || '计算出错';
    result.value = null;
  }
};
</script>