from collections import Counter
import sys

def analyze_data_and_output_string(data):
    # 假设 data 是一个列表的列表，其中不包含列名
    data_values = data
    
    result_str = ""
    
    for i, column_data in enumerate(data[0]):
        column_name = column_data
        column_index = i
        
        frequency = Counter([row[column_index] for row in data_values[1:]])
        total_count = sum(frequency.values())
        
        stats_data = []
        cumulative_percentage = 0
        for value, count in frequency.items():
            percentage = (count / total_count) * 100
            cumulative_percentage += percentage
            stats_data.append((value, count, percentage, cumulative_percentage))
        
        column_result = f"属性: {column_name}\n"
        column_result += "  值\t频数\t百分比\t累计百分比\n"
        for value, count, percentage, cumulative_percentage in stats_data:
            column_result += f"  {value}\t{count}\t{percentage:.2f}%\t{cumulative_percentage:.2f}%\n"
        
        result_str += column_result + "\n"
    
    return result_str.strip()

# 示例使用
if __name__ == "__main__":
   first_arg = sys.argv[1]
   print(analyze_data_and_output_string(first_arg))