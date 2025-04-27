// 'use client'
import { useState } from 'react';
import { createConnectTransport } from '@connectrpc/connect-web';
import { createPromiseClient } from '@connectrpc/connect';

import { CalculatorService } from '../gen/proto/calculator_connect';

export default function Calculator() {
  const [firstNumber, setFirstNumber] = useState('');
  const [secondNumber, setSecondNumber] = useState('');
  const [operator, setOperator] = useState('+');
  const [result, setResult] = useState<number | null>(null);
  const [error, setError] = useState('');

  // Create transport
  // const transport = createConnectTransport({
  //   baseUrl: 'http://localhost:8080',
  // });

  // // Create client
  // const client = createPromiseClient(CalculatorService, transport);

  // // 计算方法
  const calculate = async () => {
    try {
      setError('');
      
      // 验证输入
      const num1 = parseFloat(firstNumber);
      const num2 = parseFloat(secondNumber);
      
      if (isNaN(num1) || isNaN(num2)) {
        setError('请输入有效的数字');
        return;
      }
      
      // // 调用RPC方法
      // const response = await client.calculate({
      //   firstNumber: num1,
      //   secondNumber: num2,
      //   operator: operator
      // });
      
      // 设置结果
      // setResult(response.result);
      
    } catch (err: any) {
      console.error('计算错误:', err);
      // 处理错误信息
      setError(err.message || '计算过程中发生错误');
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-100 via-white to-purple-100 flex items-center justify-center p-4 font-sans">
      <div className="bg-white/80 backdrop-blur-md rounded-3xl shadow-2xl p-10 w-full max-w-md border border-gray-200">
        <h1 className="text-4xl font-extrabold text-center mb-10 text-blue-700 drop-shadow-lg tracking-wide">ConnectRPC 计算器</h1>
        <div className="space-y-7">
          <div className="relative">
            <input
              type="number"
              value={firstNumber}
              onChange={(e) => setFirstNumber(e.target.value)}
              className="w-full px-6 py-4 rounded-2xl border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400 focus:border-transparent transition shadow-md text-black bg-white/70 placeholder-gray-400 text-lg"
              placeholder="第一个数"
            />
          </div>
          <div className="flex justify-center">
            <select
              value={operator}
              onChange={(e) => setOperator(e.target.value)}
              className="w-40 px-5 py-3 rounded-2xl border border-gray-300 focus:outline-none focus:ring-2 focus:ring-purple-400 focus:border-transparent transition bg-white/80 shadow-md appearance-none cursor-pointer text-center font-bold text-xl text-blue-700 hover:bg-blue-50"
            >
              <option value="+">加法 (+)</option>
              <option value="-">减法 (-)</option>
              <option value="*">乘法 (×)</option>
              <option value="/">除法 (÷)</option>
            </select>
          </div>
          <div className="relative">
            <input
              type="number"
              value={secondNumber}
              onChange={(e) => setSecondNumber(e.target.value)}
              className="w-full px-6 py-4 rounded-2xl border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-400 focus:border-transparent transition shadow-md text-black bg-white/70 placeholder-gray-400 text-lg"
              placeholder="第二个数"
            />
          </div>
          <button
            onClick={calculate}
            className="w-full bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500 text-white font-bold py-4 rounded-2xl transition-all duration-300 transform hover:scale-105 active:scale-95 focus:outline-none focus:ring-2 focus:ring-purple-400 focus:ring-offset-2 shadow-xl hover:shadow-2xl text-xl tracking-wide"
          >
            计算
          </button>
          {result !== null && (
            <div className="p-6 bg-gradient-to-r from-blue-100 to-purple-100 border border-blue-200 rounded-2xl text-center shadow-inner mt-2 animate-fade-in">
              <div className="text-purple-700 text-2xl font-extrabold mb-1 drop-shadow">计算结果</div>
              <div className="text-blue-600 text-3xl font-mono tracking-widest select-all">
                {result}
              </div>
            </div>
          )}
          {error && (
            <div className="mt-6 p-4 bg-red-100 border border-red-300 rounded-2xl text-red-700 flex items-center shadow-md animate-shake">
              <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6 mr-3 flex-shrink-0" viewBox="0 0 20 20" fill="currentColor">
                <path fillRule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clipRule="evenodd" />
              </svg>
              <span className="font-semibold text-lg">{error}</span>
            </div>
          )}
        </div>
        <div className="mt-10 text-center text-gray-400 text-xs select-none">
          © {new Date().getFullYear()} ConnectRPC 计算器
        </div>
      </div>
    </div>
  );
}