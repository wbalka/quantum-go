## Hydro Projekt

This project focuses on numerical simulations of quantum-mechanical problems.
I'm not using any AI code generation tools, because it would not allow me to learn and would be much less fun.
I use AI however to refresh my memory when I encounter issues of mathematical nature. 

### Hydrogen
Current model has dedicated mechanisms for energy calculations of wave functions of spherical and axial symmetries,\
as well as an universal calculator for any 3D continuous and differentiable function in spherical coordinates\
Current calculation results for hydrogen:\
Ground state energy:\
-0.49999999999995487 a.u.\
compared to -0.5 exact value\
First excited state (2s orbital):\
-0.12500000000000228 a.u.\
compared to -0.125 exact value\
2p orbital (cylindrical symmetry):\
-0.12500000000050221 a.u.\
3d orbital (n=3, l=2, m=0):\
-0.055555555556499145 (vs -0.055(5) a.u.)\
4d orbital (n=4, l=2, m=0):\
-0.031249999997674131 (vs -0.03125 a.u.)
5f orbital (n=5, l=3, m=-1)
-0.019999999998824177 (vs -0.02)

### Program output
> Calculation parameters\
> Orbital 1s\
> interval size: 2.000000, number of intervals: 100, G-L >points range: 14-80\
>Most accurate energy:  -0.49999999999999223 - calculated with >32 G-L points\
>Average energy: -0.49999999999091027\
>Computation time 0.001597273 seconds\
>Calculation parameters\
> Orbital 2s\
> interval size: 2.000000, number of intervals: 100, G-L >points range: 14-80\
>Most accurate energy:  -0.12500000000000208 - calculated with >23 G-L points\
>Average energy: -0.12499999997649475\
>Computation time 0.00058263 seconds\
>Calculation parameters\
> Orbital 2s\
> interval size: 1.000000, number of intervals: 200, G-L >points range: 14-80\
>Most accurate energy:  -0.12500000000009889 - calculated with >15 G-L points\
>Average energy: -0.12499999960014936\
>Computation time 0.001320209 seconds\
>Calculation parameters\
> Orbital n=2 l=1\
> Area radial size: 2.000000, number of areas: 100, radial G-L >points range: 24-40\
> angular G-L points range 24-40\
>Most accurate energy:  -0.12500000000048977 - calculated with >33 radial and 35 angular G-L points\
>Average energy: -0.12499999998956519\
>Computation time 0.060611044 seconds\
>Calculation parameters\
> Orbital n=3 l=1\
> Area radial size: 2.000000, number of areas: 100, radial G-L >points range: 24-40\
> angular G-L points range 24-40\
>Most accurate energy:  -0.055555555556979178 - calculated >with 31 radial and 29 angular G-L points\
>Average energy: -0.055555555517606124\
>Computation time 0.047706554 seconds\
>Calculation parameters\
> Orbital n=3 l=2\
> Area radial size: 2.000000, number of areas: 100, radial G-L >points range: 24-40\
> angular G-L points range 24-40\
>Most accurate energy:  -0.05555555555650446 - calculated with >26 radial and 40 angular G-L points\
>Average energy: -0.055555555524384639\
>Computation time 0.0535891 seconds\
>Calculation parameters\
> Orbital n=4 l=1\
> Area radial size: 2.000000, number of areas: 100, radial G-L >points range: 24-40\
> angular G-L points range 24-40\
>Most accurate energy:  -0.031250000002788866 - calculated >with 25 radial and 39 angular G-L points\
>Average energy: -0.0312499999542056\
>Computation time 0.046089561 seconds\
>Calculation parameters\
> Orbital n=4 l=2\
> Area radial size: 2.000000, number of areas: 100, radial G-L >points range: 24-40\
> angular G-L points range 24-40\
>Most accurate energy:  -0.031249999997673084 - calculated >with 31 radial and 30 angular G-L points\
>Average energy: -0.031249999957004452\
>Computation time 0.050126469 seconds