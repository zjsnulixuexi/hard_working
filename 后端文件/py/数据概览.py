import numpy as np
from scipy import stats
import sys

def analyze_data_as_string(data_str):
    # 将字符串转换为列表
    data_list = eval(data_str)
    
    # 移除列名，只保留数据部分
    data_list = data_list[1:]
    
    # 将列表转换为numpy数组
    data_array = np.array(data_list)
    
    # 获取列名
    column_names = data_list[0]
    
    # 初始化存储统计数据的字典
    stats_data = {column: {} for column in column_names}
    
    # 对每个属性进行数据分析
    for column_name, column_data in zip(column_names, data_array.T):
        stats_data[column_name]['最大值'] = np.max(column_data)
        stats_data[column_name]['最小值'] = np.min(column_data)
        stats_data[column_name]['平均值'] = np.mean(column_data)
        stats_data[column_name]['方差'] = np.var(column_data, ddof=0)  # 总体方差
        stats_data[column_name]['标准差'] = np.std(column_data, ddof=0)  # 总体标准差
        stats_data[column_name]['中位数'] = np.median(column_data)
        
        # 计算众数
        try:
            mode_result = stats.mode(column_data)
            stats_data[column_name]['众数'] = mode_result.mode[0]
        except Exception as e:
            # 如果有多个众数，我们只取第一个
            mode_value = np.bincount(column_data).argmax()
            stats_data[column_name]['众数'] = mode_value
    
    # 将统计数据转换为字符串
    result_str_list = ["属性统计数据:"]
    for column, values in stats_data.items():
        result_str_list.append(f"属性: {column}")
        for stat_name, stat_value in values.items():
            result_str_list.append(f"  {stat_name}: {stat_value:.2f}")  # 格式化为两位小数
    
    # 将列表转换为字符串并返回
    return "\n".join(result_str_list)

# 示例使用
if __name__ == "__main__":
   first_arg = sys.argv[1]
   print(analyze_data_as_string(first_arg))