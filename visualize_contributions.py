import pandas as pd
import matplotlib.pyplot as plt

df = pd.read_csv('contributions.csv', parse_dates=['Date'])
df.set_index('Date', inplace=True)

all_days = pd.date_range(start=df.index.min(), end=df.index.max(), freq='D')
df = df.reindex(all_days, fill_value=0)

df['CumulativeContributions'] = df['Contributions'].cumsum()
df['YearMonth'] = df.index.to_period('M')
monthly_contributions = df.groupby('YearMonth')['Contributions'].sum()

day_of_week_contributions = df.groupby(df.index.dayofweek)['Contributions'].sum()
day_of_week_labels = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday']

fig, ax = plt.subplots(3, 1, figsize=(15, 18))

ax[0].plot(df.index, df['CumulativeContributions'], color='green', linewidth=2)
ax[0].set_xlabel('Date', fontsize=12)
ax[0].set_ylabel('Cumulative Contributions', fontsize=12)
ax[0].set_title('Cumulative GitHub Contributions Over Time', fontsize=16)

ax[1].bar(monthly_contributions.index.astype(str), monthly_contributions.values, color='blue')
ax[1].set_xlabel('Year-Month', fontsize=12)
ax[1].set_ylabel('Contributions', fontsize=12)
ax[1].set_title('Monthly GitHub Contributions', fontsize=16)
ax[1].tick_params(axis='x', rotation=90)

ax[2].pie(day_of_week_contributions, labels=day_of_week_labels, autopct='%1.1f%%', startangle=140, colors=plt.cm.Paired.colors)
ax[2].set_title('GitHub Contributions by Day of the Week', fontsize=16)

plt.tight_layout()
plt.show()

cumulative_contributions_table = df[['CumulativeContributions']].reset_index()
cumulative_contributions_table.columns = ['Date', 'Cumulative Contributions']

print(cumulative_contributions_table.head(10))

cumulative_contributions_table.to_csv('cumulative_contributions_table.csv', index=False)

