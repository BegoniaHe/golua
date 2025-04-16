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
  