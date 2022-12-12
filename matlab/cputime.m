format long
x = 1:10;
y1 = mean(transpose(dlmread ('cputimepointers.csv', ',', 0,0)));
y2 = mean(transpose(dlmread ('cputimeints.csv', ',', 0,0)))/100;
b1 = y1/x;
b2 = y2/x;
yCalc1 = b1*x;
yCalc2 = b2*x;
hold on
plot(x, yCalc1)
plot(x, yCalc2)
scatter(x, y1, '.')
scatter(x, y2, '.')
%title('CPU time for pointers and integers', 'FontSize', 14)
ylabel('Time (ms)')
xlabel('Number of pointers/integers (10^7)')
xlim([1 10])
ylim([0 50])
xticks(1:1:10)
yticks(0:5:100)
legend('Pointers','Integers')
set(gca, 'FontName', 'IBM Plex Sans')

a = yCalc1(:) - y1(:)
figure
scatter(x,a)

b = yCalc2(:) - y2(:)
figure
scatter(x,b)

hold off
