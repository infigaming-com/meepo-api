#!/usr/bin/env python3
import os
import re
import json
import subprocess

# 获取脚本所在目录
script_dir = os.path.dirname(os.path.abspath(__file__))
# 项目根目录（脚本目录的上一级）
root_dir = os.path.dirname(script_dir)

# 查找所有error_reason.proto文件
def find_error_reason_files():
    result = subprocess.run(
        ["find", root_dir, "-name", "error_reason.proto"],
        capture_output=True,
        text=True
    )
    return result.stdout.strip().split("\n")

# 从proto文件中提取错误码
def extract_error_codes(file_path):
    error_codes = {}
    module_name = os.path.dirname(file_path).split('/')[-3]  # 获取模块名称
    
    with open(file_path, 'r') as f:
        content = f.read()
        
    # 使用正则表达式提取枚举值
    pattern = r'\s+(\w+)\s+=\s+(\d+)'
    matches = re.findall(pattern, content)
    
    for name, code in matches:
        if name != "UNSPECIFIED" and int(code) > 0:  # 忽略UNSPECIFIED和值为0的项
            error_codes[int(code)] = f"{module_name.upper()}_{name}"
    
    return error_codes

# 主函数
def main():
    all_error_codes = {}
    
    # 获取所有error_reason.proto文件
    error_files = find_error_reason_files()
    
    # 从每个文件中提取错误码
    for file_path in error_files:
        error_codes = extract_error_codes(file_path)
        all_error_codes.update(error_codes)
    
    # 按错误码排序
    sorted_error_codes = {str(k): v for k, v in sorted(all_error_codes.items())}
    
    # 输出为JSON格式
    print(json.dumps(sorted_error_codes, indent=2))
    
    # 保存到文件
    output_path = os.path.join(root_dir, 'error_codes.json')
    with open(output_path, 'w') as f:
        json.dump(sorted_error_codes, f, indent=2)
    
    print(f"已提取 {len(sorted_error_codes)} 个错误码并保存到 error_codes.json")

if __name__ == "__main__":
    main() 