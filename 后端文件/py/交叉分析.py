import pandas as pd
import sys


def perform_chi_square_analysis(data_str):
    # 安全地将输入的字符串转换为列表
    data_list = eval(data_str)  # 使用 eval 来解析列表，注意：eval 可能带来安全风险

    # 检查数据是否有效
    if not data_list or not all(isinstance(row, list) for row in data_list):
        return "输入数据无效。"

    # 将列表转换为 pandas DataFrame
    df = pd.DataFrame(data_list)

    # 选择分组项和分析项
    # 假设第一列是分组项，其余列是分析项
    grouping_column = df.columns[0]
    analysis_columns = df.columns[1:]

    # 进行列联交叉分析
    # 这里我们只分析分组项与第一分析项的关系
    cross_tab = pd.crosstab(df[grouping_column], df[analysis_columns[0]])

    # 初始化输出字符串
    output_str = "列联交叉分析结果（频数）:\n" + str(cross_tab) + "\n"

    # 计算百分比
    percentage_cross_tab = cross_tab.div(cross_tab.sum(axis=1), axis=0) * 100
    output_str += "列联交叉分析结果（百分比）:\n" + str(percentage_cross_tab) + "\n"

    # 返回输出字符串
    return output_str


# 示例使用

if __name__ == "__main__":
    first_arg = sys.argv[1]
    print(perform_chi_square_analysis(first_arg))
