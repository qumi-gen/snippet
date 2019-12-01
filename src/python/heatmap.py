# 航空機のデータを読み込み
flights = sns.load_dataset("flights")
 
# ピボットを生成
flights = flights.pivot("month", "year", "passengers")
 
# X軸に年、Y軸に月、色の濃さで便数を表す
sns.heatmap(flights, annot=True, fmt="d")
