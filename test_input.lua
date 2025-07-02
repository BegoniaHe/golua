-- Test file for type checking transformation

function test_function(param1:string, param2:number, param3)
    print("Hello from test_function")
    print("param1:", param1)
    print("param2:", param2)
    print("param3:", param3)
end

local function local_test(x:number, y:string)
    return x + #y
end

-- Function without type annotations should remain unchanged
function no_types(a, b, c)
    return a + b + c
end

-- Function with mixed typed and untyped parameters
function mixed_function(typed:boolean, untyped, another:string)
    if typed then
        print(untyped .. another)
    end
end
