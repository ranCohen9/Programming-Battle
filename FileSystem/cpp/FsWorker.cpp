#include <stdio.h>
#include <cassert>
#include "Measure.h"

static const char DIRECTORY_NAME[] = "files";

enum
{
    TEST_ITERATIONS = 1000,
    BUFFER_LENGTH = 256,
};

void CreateDirectory()
{
    char szCommand[BUFFER_LENGTH];
    sprintf(szCommand, "rd %s /S /Q && md %s", DIRECTORY_NAME, DIRECTORY_NAME);
    system(szCommand);
}

void CreateFile(const int nIndex)
{
    char szFileName[BUFFER_LENGTH];
    sprintf(szFileName, "%s/file-%04d.txt", DIRECTORY_NAME, nIndex);
    FILE *pFile = fopen(szFileName, "w");
    assert(NULL != pFile);
    fclose(pFile);
}

void CreateManyFiles()
{
    CreateDirectory();

    MEASURE_START();
    for (int nIndex = 0; nIndex < TEST_ITERATIONS; nIndex++)
    {
        CreateFile(nIndex);
    }
    MEASURE_STOP();

    printf("Create files elapsed: %f msec\n", MEASURE_GET_MSEC());
}

void WriteToFile()
{
    char szFilePath[BUFFER_LENGTH];
    sprintf(szFilePath, "%s/contentFile.txt", DIRECTORY_NAME);

    //CreateDirectory(); // Note: Disabled to keep the previous test files

    MEASURE_START();
    FILE  *pFile = fopen(szFilePath, "w");
    assert(NULL != pFile);
    for (int nIndex = 0; nIndex < TEST_ITERATIONS; nIndex++)
    {
        fprintf(pFile, "index is %d\n", nIndex);
    }
    fclose(pFile);
    MEASURE_STOP();

    printf("Write to file elapsed: %f msec\n", MEASURE_GET_MSEC());
}

int main()
{
    CreateManyFiles();
    WriteToFile();
    return 0;
}
