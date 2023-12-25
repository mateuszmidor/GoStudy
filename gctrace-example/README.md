# gctrace

Print garbage collector run statistics. This demo will allocate 1MB of memory every second.

## RUN

```sh
./run_all.sh
```

, output:
```text
gc 1 @4.002s 0%: 0.050+0.17+0.003 ms clock, 0.15+0/0.19/0.078+0.009 ms cpu, 4->4->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 3 P
gc 2 @8.007s 0%: 0.028+0.098+0.003 ms clock, 0.086+0/0.063/0.030+0.009 ms cpu, 4->4->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 3 P
gc 3 @12.009s 0%: 0.043+0.13+0.003 ms clock, 0.13+0/0.11/0.015+0.010 ms cpu, 4->4->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 3 P
gc 4 @16.011s 0%: 0.051+0.13+0.002 ms clock, 0.15+0/0.18/0.011+0.007 ms cpu, 4->4->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 3 P
gc 5 @20.013s 0%: 0.035+0.099+0.003 ms clock, 0.10+0/0.095/0.040+0.010 ms cpu, 4->4->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 3 P
gc 6 @24.016s 0%: 0.055+0.19+0.004 ms clock, 0.16+0/0.13/0.020+0.014 ms cpu, 4->4->0 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 3 P
```