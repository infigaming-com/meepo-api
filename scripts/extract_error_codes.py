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

# 确保目录存在
def ensure_dir_exists(directory):
    if not os.path.exists(directory):
        os.makedirs(directory)

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
    
    # 定义local目录和错误码文件路径
    local_dir = os.path.join(root_dir, "locales", "en")
    error_codes_file = os.path.join(local_dir, "error_codes.json")
    
    # 确保local目录存在
    ensure_dir_exists(local_dir)
    
    # 读取已存在的错误码文件（如果存在）
    existing_error_codes = {}
    if os.path.exists(error_codes_file):
        try:
            with open(error_codes_file, 'r') as f:
                existing_error_codes = json.load(f)
            print(f"已读取现有错误码文件: {error_codes_file}")
        except Exception as e:
            print(f"读取现有错误码文件失败: {str(e)}")
    
    # 合并错误码，优先保留已存在的数据
    merged_error_codes = sorted_error_codes.copy()
    for code, message in existing_error_codes.items():
        merged_error_codes[code] = message  # 已存在的错误码优先
    
    # 保存到文件
    with open(error_codes_file, 'w', encoding='utf-8') as f:
        json.dump(merged_error_codes, f, indent=2, ensure_ascii=False)
    
    print(f"已生成错误码文件: {error_codes_file}")

if __name__ == "__main__":
    main() 