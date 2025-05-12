<template>
  <div class="bt"
       :style="`--clr:#0f0`"
       @mousemove="handleMouseMove"
  >
  </div>
</template>

<script setup lang="ts">


// 鼠标移动事件处理（包含 TypeScript 类型标注）
const handleMouseMove = (event: MouseEvent) => {
  const element = event.currentTarget as HTMLElement
  const rect = element.getBoundingClientRect()

  // 计算相对位置（包含非空断言）
  const x = event.clientX - rect.left
  const y = event.clientY - rect.top

  // 设置 CSS 自定义属性
  element.style.setProperty('--x', `${x}px`)
  element.style.setProperty('--y', `${y}px`)
}
</script>

<style scoped lang="scss">

.bt {
  height: 80px;
  width: 200px;

  background-color: rgba(45, 45, 45, 1);

  transition: 0.5s; // 设定动画延迟，对伪类，css动态变动，窗口大小变动，生效。

  position: relative; //将这个元素作为边界，防止内部的绝对定位，向外寻找祖先元素。
  overflow: hidden; // 把发光效果，外部屏蔽
}

// 效果
.bt::before {
  content: '';
  position: absolute;

  top: var(--y);
  left: var(--x); // 让发光遮罩随x,y移动

  width: 400px;
  height: 200px; // 发光效果的大小

  transform: translate(-50%, -50%); // 移动，把发光效果移动到左上，占1/4
  transition: 0.1s; // 动画延迟0.1s，指 transform

  opacity: 0; // 发光效果，默认隐藏，透明度0


  background: radial-gradient(var(--clr), transparent, transparent); // 发光渐变，从颜色1/3快速透明

}

// 效果，范围控制
// 把发光内部，显示屏蔽，只显示1px
.bt::after {
  content: ''; // 必须声明
  position: absolute; // 绝对定位，常用定位类型
  background-color: rgba(45, 45, 45, 0.8); // 半透明深灰色背景,常用于创建遮罩层、按钮悬停效果或装饰性覆盖层。
  // 如果本体不设置背景色，这里因产生背景色的效果。因为这遮掩层就是用背景色做的。
  inset: 1px; // 使伪元素在父元素内部各边缩进1px，形成内嵌效果
}

// 效果，开关
.bt:hover::before {
  opacity: 1; // 当鼠标移入时，发光效果，设置透明度1，完全显示
}
</style>
