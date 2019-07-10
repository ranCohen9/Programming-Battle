#ifndef MEASURE__H
#define MEASURE__H

#define MEASURE_START() \
    static LARGE_INTEGER nFrequency, nStartingTime, nEndingTime, nElapsedMicroseconds; \
    QueryPerformanceFrequency(&nFrequency); \
    QueryPerformanceCounter(&nStartingTime);

#define MEASURE_STOP() \
    QueryPerformanceCounter(&nEndingTime); \
    nElapsedMicroseconds.QuadPart = (nEndingTime.QuadPart - nStartingTime.QuadPart) * 1000000 / nFrequency.QuadPart;

#define MEASURE_GET_MSEC() \
    (nElapsedMicroseconds.QuadPart / 1000.0)

#endif // MEASURE__H
