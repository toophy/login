#include <stdio.h>
#include <stdlib.h>

#if !defined(_WIN32)
#include <sys/prctl.h>
void ChangePrcName(int sid)
{
        char szTempName[32];
        sprintf(szTempName, "scs_%04d", sid);
        prctl(PR_SET_NAME, szTempName, 0, 0, 0);
        //printf(szTempName);
}
#else
void ChangePrcName(int sid)
{
	printf("scs_%04d\n", sid);
}
#endif // !defined(_WIN32)

