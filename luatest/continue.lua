print("continue语句")
print("continue语句在for循环中")
for i = 1, 10 do
    if i % 2 == 0 then
      continue  -- 跳过偶数
    end
    print(i)
  end

print("continue语句在for ... in ...循环中")

local tbl = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
for i, v in ipairs(tbl) do
    if v % 2 == 0 then
      continue  -- 跳过偶数
    end
    print(v)
end

print("continue语句在while循环中")
local i = 1
while i <= 10 do
    i = i + 1
    if i % 2 == 0 then
      continue  -- 跳过偶数
    end
    print(i)
end

print("continue语句在repeat ... until循环中")
i = 1
repeat
    i = i + 1
    if i % 2 == 0 then
      continue  -- 跳过偶数
    end
    print(i)
until i > 10

print("continue语句与to-be-closed变量的交互")
-- 测试continue与to-be-closed变量的交互
local function test_to_be_closed()
  for i = 1, 5 do
    local file <close> = io.open("test.txt", "w")
    file:write("Line " .. i)
    
    if i == 3 then
      print("Before continue at i=" .. i)
      continue  -- 这里file应该被关闭
      print("This should not be printed")
    end
    
    print("Processing i=" .. i)
  end
  print("Loop ended")
  
  -- 验证文件是否正确关闭的方法
  local status, err = pcall(function()
    local check = io.open("test.txt", "r")
    print("File content: " .. check:read("*a"))
    check:close()
  end)
  print("File check status:", status)
end

test_to_be_closed()


print("continue语句的栈展开处理")
-- 测试嵌套函数调用中的continue栈展开
local function inner(x)
  print("Inner start with x=" .. x)
  if x == 3 then
    print("About to continue from inner function")
    return "continue" -- 模拟从内部函数触发continue
  end
  print("Inner end")
  return "normal"
end

local function test_stack_unwinding()
  for i = 1, 5 do
    print("Iteration " .. i .. " start")
    
    local result = inner(i)
    if result == "continue" then
      print("Processing continue from inner")
      continue
    end
    
    print("Iteration " .. i .. " end")
  end
  print("Loop completed")
end

test_stack_unwinding()

-- 带异常处理的栈展开测试
local function test_stack_with_pcall()
  for i = 1, 5 do
    print("Outer iteration " .. i)
    
    pcall(function()
      for j = 1, 3 do
        print("  Inner iteration " .. j)
        if i == 2 and j == 2 then
          continue  -- 这应该只影响内部循环
        end
        if i == 4 and j == 2 then
          error("Test error")  -- 测试错误与continue的交互
        end
        print("  Completed inner " .. j)
      end
    end)
    
    print("Outer iteration " .. i .. " end")
  end
end

test_stack_with_pcall()

print("continue语句调用钩子函数")
-- 测试continue与调试钩子的交互
local function test_debug_hooks()
  local hook_counts = {line = 0, call = 0, ret = 0}
  
  local function hook(event)
    hook_counts[event] = hook_counts[event] + 1
    print("Hook: " .. event .. " (total: " .. hook_counts[event] .. ")")
  end
  
  debug.sethook(hook, "crl")
  
  for i = 1, 3 do
    print("Iteration: " .. i)
    if i == 2 then
      continue
    end
    print("End of iteration " .. i)
  end
  
  debug.sethook()  -- 清除钩子
  print("Final hook counts:", 
        "line="..hook_counts.line, 
        "call="..hook_counts.call, 
        "return="..hook_counts.ret)
end

test_debug_hooks()