# SmartLogger2-UserLicenseStat
SmartLogger2 - User license statistics


Parsing the log file SmartLogger2 (DataService)  
Search for minimum and maximum values are admin and lite license by each hour.  
Display search progress. Output in tabular form to the console.  
c: / programdata / speech technology center / smartlogger2 / dataservice / log / dataservice.log *

File Format:  
ASCII text, with CRLF line terminators

String:  
06/07/2016 08: 31: 08,813 [SessionsManager.Cleaner] DEBUG - Protection.PrintInfo (): Free Client licenses: 6 admins, 24 lites.

Options start:  
1) Search the current date  
./NezUserLicenseStat  
2) Search the specific date  
./NezUserLicenseStat -d=2016-06-14  


Парсинг лог-файла SmartLogger2  
Поиск минимального и максимального значения admin и lite лицензий в разрезе каждого часа.  
Отображение прогресса поиска. Вывод в табличном виде в консоль.  
c:/programdata/speech technology center/smartlogger2/dataservice/log/dataservice.log*  

Формат файлов:  
ASCII text, with CRLF line terminators  

Парсим строку вида:  
2016-06-07 08:31:08,813 [SessionsManager.Cleaner] DEBUG - Protection.PrintInfo(): Free Client licenses: 6 admins, 24 lites.

Варианты запуска:  
1) Поиск по текущей дате  
./NezUserLicenseStat  
2) Поиск по заданной дате  
./NezUserLicenseStat -d=2016-06-14  


EXAMPLE:  
[somalipirate@localhost sszdr-slfs1]$ ./NezUserLicenseStat -d=2016-06-14  
DataService.log 16-Jun-2016  
Loading...  
522483/522483 100% 11.79sec  
DataService.log.1 16-Jun-2016  
Loading...  
641669/641669 100% 14.47sec  
DataService.log.2 15-Jun-2016  
Loading...  
641618/641618 100% 14.71sec  
DataService.log.3 15-Jun-2016  
Loading...  
641229/641229 100% 14.46sec  
DataService.log.4 15-Jun-2016  
Loading...  
641187/641187 100% 14.34sec  
DataService.log.5 14-Jun-2016  
Loading...  
641242/641242 100% 14.35sec  
DataService.log.6 14-Jun-2016  
Loading...  
641125/641125 100% 14.38sec  
DataService.log.7 13-Jun-2016  
Loading...  
640886/640886 100% 14.48sec  
       Date Hour AdminMax AdminMin LiteMax LiteMin  
14-Jun-2016    0        8        8      36      36  
14-Jun-2016    1        8        8      36      36  
14-Jun-2016    2        8        8      36      36  
14-Jun-2016    3        8        8      36      36  
14-Jun-2016    4        8        8      36      36  
14-Jun-2016    5        8        8      36      36  
14-Jun-2016    6        8        8      36      36  
14-Jun-2016    7        8        8      36      36  
14-Jun-2016    8        8        7      36      21  
14-Jun-2016    9        7        6      22      17  
14-Jun-2016   10        6        6      18      11  
14-Jun-2016   11        6        6      12       9  
14-Jun-2016   12        6        5       9       8  
14-Jun-2016   13        6        5       8       6  
14-Jun-2016   14        6        5       6       5  
14-Jun-2016   15        5        5       5       2  
14-Jun-2016   16        6        5       3       2  
14-Jun-2016   17        7        6      19       3  
14-Jun-2016   18        7        7      23      18  
14-Jun-2016   19        8        7      23      23  
14-Jun-2016   20        8        8      24      22  
14-Jun-2016   21        8        8      27      24  
14-Jun-2016   22        8        8      27      27  
14-Jun-2016   23        8        8      27      27  

