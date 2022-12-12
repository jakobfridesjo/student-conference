format long
x = transpose(10:10:500);
xq = 1:1:500;
y = dlmread ('memoryoverhead.csv', ',', 0,0);

hold on
f = fit(x,y,'exp2');
plot(f,x,y)

ylabel('GC Cycles')
xlabel('GC Percentage')
xlim([0 500])
ylim([0 y(1,1)])
xticks(0:50:500)
yticks(0:50:y(1,1))
set(gca, 'FontName', 'IBM Plex Sans')
hold off