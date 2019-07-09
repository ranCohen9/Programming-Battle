#include <stdio.h>
#include <math.h>
#include <windows.h> // QueryPerformanceFrequency, QueryPerformanceCounter

enum
{
    MAX_TEST_ITERATIONS = 10,
    MAX_LEIBNIZ_ITERATIONS = 300000
};

#define MEASURE_START() \
    static LARGE_INTEGER nFrequency, nStartingTime, nEndingTime, nElapsedMicroseconds; \
    QueryPerformanceFrequency(&nFrequency); \
    QueryPerformanceCounter(&nStartingTime);

#define MEASURE_STOP() \
    QueryPerformanceCounter(&nEndingTime); \
    nElapsedMicroseconds.QuadPart = (nEndingTime.QuadPart - nStartingTime.QuadPart) * 1000000 / nFrequency.QuadPart;

#define MEASURE_GET_MSEC() \
    (nElapsedMicroseconds.QuadPart / 1000.0)

// Note: Not using pow() due to performance issue
#define LEIBNIZ_ONCE(n) \
    (n % 2 == 0 ? 1.0f : -1.0f) / (2 * n + 1)

double SeriesCalculation()
{
    const double PI = 3.141592653589;
    double dSeriesSum = 0.0;
    double dGoal = PI / 4.0;
    const double dGoalResolution = 0.000001;
    printf("This is the goal %f (resolution: %f)\n", dGoal, dGoalResolution);
    
    MEASURE_START();
    int nIteration = 0;
    for (; (fabs(dSeriesSum - dGoal) > dGoalResolution) && (nIteration < MAX_LEIBNIZ_ITERATIONS); ++nIteration)
    {
        dSeriesSum += LEIBNIZ_ONCE(nIteration);
    }
    MEASURE_STOP();

    const double dElapsedMsec = MEASURE_GET_MSEC();
    printf("elapsed: %f ms | score: %f | iterations: %d\n", dElapsedMsec, dSeriesSum, nIteration);
    return dElapsedMsec;
}

int main()
{
    double dIterationsSum = 0.0;
    for (int nIteration = 0; nIteration < MAX_TEST_ITERATIONS; ++nIteration)
    {
        dIterationsSum += SeriesCalculation();
    }
    printf("\nTotal Avegrage of %d iterations: %f ms", MAX_TEST_ITERATIONS, dIterationsSum / MAX_TEST_ITERATIONS);
    return 0;
}
