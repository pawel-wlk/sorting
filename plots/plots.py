import pandas as pd
import matplotlib
import matplotlib.pyplot as plt

data = pd.read_csv(
    '../stats10.csv', names=['type', 'n', 'comparisons', 'swaps', 'time'])

data['c/n'] = data.apply(lambda row: row['comparisons']/row['n'], axis=1)
data['s/n'] = data.apply(lambda row: row['swaps']/row['n'], axis=1)

# for t in data['type'].unique():
    # means = data[data['type'] == t].groupby('n').mean()
    # means.plot()


axes = []
for i, key in enumerate(data.keys()[2:]):
    types =  data['type'].unique()
    axes.append(None)
    for t in types:
        df = data[data['type'] == t][['n', key]]
        axes[i] = df.groupby('n').mean().plot(ax=axes[i], title=key)
    
    axes[i].legend(labels=list(types))
plt.show()
