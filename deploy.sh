#!/bin/bash

# Step 1. 检查 
#   - 删除错误或者不使用的modules
go mod tidy

# Step 2. 测试
#   - 包括运行直接和间接依赖项的测试, 验证当前选定的软件包版本是否兼容
go test all

# Setp3. 确保 go.sum + go.mod 一起提交到版本库

# ... 待补充


# Step 4. 更新到服务器

# scp app root@192.168.1.120:/app/xxx
