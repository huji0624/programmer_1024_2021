using Flurl.Http;
using NeinMath;
using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Numerics;
using System.Runtime.InteropServices;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace Test
{
    class Program
    {
        public static Stopwatch watch = new Stopwatch();
        public static int MaxNum = 5000000;
        public static Integer Num1000 = Integer.Parse("1000");
        public static Integer Num0 = Integer.Parse("0");
        public static Integer Num1 = Integer.Parse("1");
        public static string[] strs = new string[] { "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K" };
        /// <summary>
        /// 从文本读取到的数据
        /// </summary>
        public static BlockingCollection<string> _readDataModel = new BlockingCollection<string>(MaxNum);
        /// <summary>
        /// 判断为正确的数据
        /// </summary>
        public static BlockingCollection<ReadDataModel> _justData = new BlockingCollection<ReadDataModel>(MaxNum);

        /// <summary>
        /// 准备提交的算式
        /// </summary>
        public static BlockingCollection<Submit1024DataModel> _submit1024Data = new BlockingCollection<Submit1024DataModel>(MaxNum);

        /// <summary>
        /// 储存用于计算的数据
        /// </summary>
        public static List<NumDicModel> _numDic = new List<NumDicModel>();
        public static List<WaitNumDicModel> _waitNumDic = new List<WaitNumDicModel>();
        /// <summary>
        /// 已读取完毕的文件
        /// </summary>
        public static List<long> _readFinishData = new List<long>();
        /// <summary>
        /// 已经被他人扫过的文件
        /// </summary>
        public static List<long> _conflictData = new List<long>();

        /// <summary>
        /// 成功提交的算式
        /// </summary>
        public static List<string> _Data1024 = new List<string>();
        public const string _digTreasureUrl = "http://47.104.220.230/dig";
        public const string _submit1024Url = "http://47.104.220.230/formula";
        public const string _token = "ad517f45ce05f7f9b8e9805dfec1cf60";
        public static long _readNum = 0;
        public static long _justNum = 0;
        static async Task Main(string[] args)
        {
            var program = new Program();
            ThreadPool.SetMinThreads(30000, 1000);//设置最小线程并发数
            ThreadPool.SetMaxThreads(30001, 1000);//设置最大线程并发数

            Task.Run(async () => { await program.DoNumDicModel(); });
            for (int i = 0; i < 50; i++)
            {
                //Task.Run(async () => { await program.ProcessDataMain(); });
            }
            for (int i = 0; i < 5; i++)
            {
                Task.Run(async () => { await program.DigTreasureMain(); });
            }
            for (int i = 0; i < 5; i++)
            {
                Task.Run(async () => { await program.Submit1024Main(); });
            }
            Task.Run(async () => { await program.Calculate1024(4, 0); });
            Task.Run(async () => { await program.Calculate1024(4, 0); });
            Task.Run(async () => { await program.Calculate1024(4, 0); });
            Task.Run(async () => { await program.Calculate1024(4, 0); });
            Task.Run(async () => { await program.Calculate1024(4, 0); });
            Task.Run(async () => { await program.Calculate1024(4, 1); });
            Task.Run(async () => { await program.Calculate1024(4, -1); });
            Task.Run(async () => { await program.Calculate1024(5, 0); });
            Task.Run(async () => { await program.Calculate1024(8, 0); });
            Task.Run(async () => { await program.Calculate1024(10, 0); });
            Task.Run(async () => { await program.Calculate1024(20, 0); });
            Task.Run(async () => { await program.Calculate1024(30, 0); });

            Console.WriteLine("开始!");
            watch.Start();
            Console.WriteLine("读取文件模块启动成功");
            string dirPath = AppContext.BaseDirectory + @"data/";

            DirectoryInfo dirInfo = new DirectoryInfo(dirPath);
            var files = Directory.GetFiles(dirPath, "*.data").ToList();
            files = files.RandomSort();

            List<List<string>> ff = new List<List<string>>();
            var fNum = 4;
            //var fNum = 16;
            for (int i = 0; i < 128 / fNum; i++)
            {
                ff.Add(files.PageBy(i + 1, fNum));
            }

            foreach (var f in ff)
            {
                Task.Run(async () => { await program.ReadDataMain(f); });
                Task.Delay(500).Wait();
            }

            await Task.Run(() => { program.Test(); });

            Console.Read();
            Console.WriteLine("结束!");
        }


        /// <summary>
        /// 读取数据
        /// </summary>
        /// <param name="files"></param>
        /// <returns></returns>
        private async Task ReadDataMain(List<string> files)
        {
            foreach (var fileName in files)
            {
                await ReadData(fileName);
            }
        }

        /// <summary>
        /// <summary>
        /// 加载数据
        /// </summary>
        /// <param name="fileName"></param>
        /// <returns></returns>
        public async Task ReadData(string fileName)
        {

            long.TryParse(ToolsHelper.ExtractionDigital(fileName), out long filesNum);
            try
            {
                Console.WriteLine($"开始读取{fileName}");

                using (StreamReader streamReader = new StreamReader(fileName, Encoding.UTF8))
                {
                    string s = String.Empty;
                    List<string> strs = new List<string>();
                    while ((s = streamReader.ReadLine()) != null)
                    {
                        _readNum++;
                        strs.Add(s);
                        if (strs.Count == 500000)
                        {
                            await ProcessData(strs);
                            strs = null;
                            System.GC.Collect();
                            strs = new List<string>();
                        }
                    }
                    if (strs.Count > 0)
                    {
                        await ProcessData(strs);
                        strs = null;
                        System.GC.Collect();
                        strs = new List<string>();
                    }
                    _readFinishData.Add(filesNum);
                }
            }

            catch (Exception)
            {
                return;
            }

        }

        /// <summary>
        /// 处理数据
        /// </summary>
        /// <returns></returns>
        public async Task ProcessDataMain()
        {
            Console.WriteLine("处理数据模块启动成功");
            await Task.Run(() => { });
            while (_readDataModel != null)
            {
                var DataModel = _readDataModel.Take();
            }

        }

        /// <summary>
        /// 处理数据
        /// </summary>
        /// <param name="DataModel"></param>
        private async Task ProcessData(List<string> DataModels)
        {
            await Task.Run(() => { });
            ReadDataModel Data = new ReadDataModel();
            string locationid = "";
            Integer LocationidNum = Num0;
            Integer magicNum = Num0;
            foreach (var DataModel in DataModels)
            {
                try
                {
                    if (string.IsNullOrWhiteSpace(DataModel)) continue;
                    Data = DataModel.ToObject<ReadDataModel>();
                    locationid = ToolsHelper.ExtractionDigital(Data.locationid);
                    LocationidNum = Integer.Parse(locationid);
                    magicNum = Integer.Parse(Data.magic);
                    if (!ToolsHelper.VerifyDigital(LocationidNum, magicNum)) continue;
                    _justData.Add(Data);
                    _justNum++;
                    _numDic.Add(new NumDicModel()
                    {
                        magicNum = magicNum,
                        locationid = Data.locationid
                    });
                }
                catch (Exception)
                {

                }

            }
        }

        /// <summary>
        /// 清洗数据
        /// </summary>
        /// <returns></returns>
        public async Task DoNumDicModel()
        {
            await Task.Run(() => { });
            Console.WriteLine("清洗数据模块启动成功");
            while (true)
            {
                try
                {
                    if (_numDic.Count < 2)
                    {
                        Task.Delay(500).Wait();
                        continue;
                    }

                    var _smallNum = _numDic.RandomSort().ToList().Where(t => t.magicNum.CompareTo(Num1000) != 1).ToList();
                    foreach (var _s in _smallNum)
                    {
                        _waitNumDic.Add(new WaitNumDicModel()
                        {
                            magicNum = long.Parse(_s.magicNum.ToString()),
                            locationid = $"{_s.locationid}",
                            locationidList = new List<string>() { _s.locationid }
                        });
                        _numDic.RemoveAll(t => t.locationid == _s.locationid);
                    }

                    var __numDic = _numDic.RandomSort().ToList();
                    for (int i = 0; i < __numDic.Count - 1; i++)
                    {

                        if (__numDic[i].magicNum.CompareTo(Num1000) != 1)
                        {
                            _waitNumDic.Add(new WaitNumDicModel()
                            {
                                magicNum = long.Parse(__numDic[i].magicNum.ToString()),
                                locationid = $"{__numDic[i].locationid}",
                                locationidList = new List<string>() { __numDic[i].locationid }
                            });
                            _numDic.RemoveAll(t => t.locationid == __numDic[i].locationid);
                            continue;
                        }
                        var _num1 = new NumDicModel();
                        var _num2 = new NumDicModel();
                        if (__numDic[i].magicNum.CompareTo(__numDic[i + 1].magicNum) == 1)
                        {
                            _num1.magicNum = __numDic[i].magicNum;
                            _num1.locationid = __numDic[i].locationid;
                            _num2.magicNum = __numDic[i + 1].magicNum;
                            _num2.locationid = __numDic[i + 1].locationid;
                        }
                        else
                        {
                            _num1.magicNum = __numDic[i + 1].magicNum;
                            _num1.locationid = __numDic[i + 1].locationid;
                            _num2.magicNum = __numDic[i].magicNum;
                            _num2.locationid = __numDic[i].locationid;
                        }


                        var m1 = _num1.magicNum.Subtract(_num2.magicNum);
                        if (m1.CompareTo(Num1000) != 1)
                        {
                            _waitNumDic.Add(new WaitNumDicModel()
                            {
                                magicNum = long.Parse(m1.ToString()),
                                locationid = $"({_num1.locationid}-{_num2.locationid})",
                                locationidList = new List<string>() { _num1.locationid, _num2.locationid }
                            });
                            _numDic.RemoveAll(t => t.locationid == _num1.locationid);
                            _numDic.RemoveAll(t => t.locationid == _num2.locationid);
                            continue;
                        }
                        var m2 = _num1.magicNum.Divide(_num2.magicNum);
                        if (m2.CompareTo(Num1000) != 1 && m2.CompareTo(Num1) == 1)
                        {
                            _waitNumDic.Add(new WaitNumDicModel()
                            {
                                magicNum = long.Parse(m2.ToString()),
                                locationid = $"({_num1.locationid}/{_num2.locationid})",
                                locationidList = new List<string>() { _num1.locationid, _num2.locationid }
                            });
                            _numDic.RemoveAll(t => t.locationid == _num1.locationid);
                            _numDic.RemoveAll(t => t.locationid == _num2.locationid);
                            continue;
                        }

                    }
                    Task.Delay(500).Wait();
                }
                catch (Exception)
                {

                }

            }

        }


        /// <summary>
        /// 提交挖宝数据
        /// </summary>
        /// <returns></returns>
        public async Task DigTreasureMain()
        {
            Console.WriteLine("提交挖宝数据模块启动成功");
            //从队列中取元素。

            ReadDataModel Data = new ReadDataModel();
            while (_justData != null)
            {
                try
                {
                    Data = _justData.Take();
                    await DigTreasure(Data);
                }
                catch (Exception)
                {

                }

            }

        }
        /// <summary>
        /// 提交挖宝数据
        /// </summary>
        /// <returns></returns>
        private async Task DigTreasure(ReadDataModel Data)
        {
            try
            {
                var responseData = await _digTreasureUrl.WithTimeout(3).PostJsonAsync(new { locationid = Data.locationid, token = _token }).ReceiveJson<ResponseModel>();
                //请求失败，重新放回队列
                if (responseData.errorno == -1 && Data.errorNum < 10)
                {
                    Data.errorNum++;
                    _justData.Add(Data);
                }
            }
            catch (Exception ex)
            {
                _justData.Add(Data);
            }
        }

        /// <summary>
        /// 提交1024数据
        /// </summary>
        /// <returns></returns>
        public async Task Submit1024Main()
        {
            Console.WriteLine("提交1024数据模块启动成功");
            //从队列中取元素。

            while (_submit1024Data != null)
            {
                try
                {
                    var Data = _submit1024Data.Take();
                    await Submit1024(Data);
                }
                catch (Exception)
                {

                }

            }

        }
        /// <summary>
        /// 提交1024数据
        /// </summary>
        /// <returns></returns>
        private async Task Submit1024(Submit1024DataModel Data)
        {
            try
            {
                var responseData = await _submit1024Url.WithTimeout(3).PostJsonAsync(new { formula = Data.formula, token = _token }).ReceiveJson<ResponseModel>();
                Console.WriteLine("===========>responseData：" + responseData.errorno);
                if (responseData.errorno == 0)
                {
                    foreach (var l in Data.locationid)
                    {
                        _waitNumDic.RemoveAll(t => t.locationidList.Contains(l));
                    }
                }
                if (responseData.errorno == 3)
                {
                    foreach (var d in responseData.data)
                    {
                        _waitNumDic.RemoveAll(t => t.locationidList.Contains(d));
                    }
                    if (Data.locationid.Intersect(responseData.data).Count() > 0) return;
                    await Submit1024(Data);
                }
                if (responseData.errorno == 1 && Data.errorNum < 10)
                {
                    Data.errorNum++;
                    _submit1024Data.Add(Data);
                }
                //请求失败，重新放回队列
                if (responseData.errorno == -1 && Data.errorNum < 10)
                {
                    Data.errorNum++;
                    _submit1024Data.Add(Data);
                }
            }
            catch (Exception ex)
            {
                Console.WriteLine("Submit1024 ex：" + ex.Message);
                _submit1024Data.Add(Data);
            }
        }



        /// <summary>
        /// 计算1024
        /// </summary>
        /// <returns></returns>
        public async Task Calculate1024(int Num, int type)
        {
            await Task.Run(() => { });
            Console.WriteLine("计算1024模块启动成功");
            while (true)
            {
                try
                {
                    var NewNumDic = _waitNumDic.RandomSort();
                    if (NewNumDic.Count < Num)
                    {
                        Task.Delay(3000).Wait();
                        continue;
                    }
                    switch (type)
                    {
                        case -1:
                            NewNumDic = NewNumDic.OrderBy(t => t.magicNum).ToList();
                            break;
                        case 1:
                            NewNumDic = NewNumDic.OrderByDescending(t => t.magicNum).ToList();
                            break;
                        default:
                            break;
                    }
                    if (Num == 30) Num = NewNumDic.Count;
                    if (Num == 20) Num = NewNumDic.Count / 2;
                    Console.WriteLine($"开始计算1024,{NewNumDic.Count}个，Num:{Num}，type:{type}");
                    for (int i = 0; i < NewNumDic.Count; i++)
                    {
                        var _k = NewNumDic.Skip(i).Take(Num).ToList();
                        //如果大于4，则把多余的数字合并为4个
                        if (_k.Count > 4)
                        {
                            for (int j = 4; j < _k.Count; j++)
                            {
                                if (j % 2 == 0)
                                {
                                    _k[3].magicNum = _k[3].magicNum - _k[j].magicNum;
                                    _k[3].locationid = _k[3].locationid + "-" + _k[j].locationid;
                                }
                                else
                                {
                                    _k[3].magicNum = _k[3].magicNum + _k[j].magicNum;
                                    _k[3].locationid = _k[3].locationid + "+" + _k[j].locationid;
                                }
                            }
                        }
                        await Js(_k);
                    }
                    Task.Delay(500).Wait();
                }
                catch (Exception ex)
                {
                    Console.WriteLine("=====>Calculate1024 ex:" + ex.Message);
                }
            }

        }

        private async Task Js(List<WaitNumDicModel> _k)
        {
            await Task.Run(() => { });
            var answer = Calculate1024Tools.GetData(_k.Select(t => t.magicNum).ToList());
            foreach (var a in answer)
            {
                Console.WriteLine("提交算式：" + a);
                var formula = a;
                Dictionary<string, string> d = new Dictionary<string, string>();
                int re = 0;
                List<string> locationidList = new List<string>();
                foreach (var _newK in _k.OrderByDescending(t => t.magicNum))
                {
                    locationidList.AddRange(_newK.locationidList);
                    d.Add($"xin{strs[re]}", _newK.locationid);
                    formula = formula.Replace(_newK.magicNum.ToString(), $"xin{strs[re]}");
                    re++;
                }
                foreach (var key in d.Keys)
                {
                    formula = formula.Replace(key, d[key]);
                }
                if (_Data1024.Contains(formula)) continue;
                _Data1024.Add(formula);

                _submit1024Data.Add(new Submit1024DataModel()
                {
                    formula = formula,
                    locationid = locationidList
                });
                return;
            }
        }

        public void Test()
        {
            int MaxWorkerThreads, miot, AvailableWorkerThreads, aiot;

            while (true)
            {
                //获得最大的线程数量
                ThreadPool.GetMaxThreads(out MaxWorkerThreads, out miot);

                AvailableWorkerThreads = aiot = 0;

                //获得可用的线程数量
                ThreadPool.GetAvailableThreads(out AvailableWorkerThreads, out aiot);

                //返回线程池中活动的线程数
                var aaa = MaxWorkerThreads - AvailableWorkerThreads;
                var s = (int)watch.Elapsed.TotalSeconds;
                Console.WriteLine($"{aaa}活动线程；{s}秒；读取数据共{_readNum}；平均每秒{(_readNum / (s == 0 ? 1 : s)) / 10000}万个；正确：{_justNum}个；_readDataModel{_readDataModel.Count}条数据；_waitNumDic{_waitNumDic.Count}条数据；_justData有{_justData.Count}条数据；_readFinishData有{_readFinishData.Count}条数据；_conflictData有{_conflictData.Count}条数据；");
                Task.Delay(1000).Wait();
            }



            var strs = "{\"locationid\":\"6q63m72jy3loe25y250fddddrq5knqciq9gcnvh1bsux4ilpky8bsqne6dv3df8o4ilpky8bsqne6dv3df8o\",\"magic\":\"9939024584589050276010107099339\"}";
            for (int i = 0; i < 10; i++)
            {
                Console.WriteLine("方法1开始!");
                watch = new Stopwatch();
                watch.Start();
                for (int n = 0; n < 1000000; n++)
                {
                    //var sfasfs = strs.ToObject<ReadDataModel>();
                    var Data = strs.ToObject<ReadDataModel>();
                    var locationid = ToolsHelper.ExtractionDigital(Data.locationid);
                    // byte[] byteArray = BitConverter.GetBytes(locationid);
                    var LocationidNum = Integer.Parse(locationid); //new Integer(locationid, 10); 
                    //byte[] byteArray1 = BitConverter.GetBytes(Data.magic);
                    var magicNum = Integer.Parse(Data.magic);  //new Integer(Data.magic, 10); 
                    //Data.magicNum = magicNum;
                    var aaaa1 = LocationidNum.Add(magicNum);
                    var aaaa2 = LocationidNum.Subtract(magicNum);
                    var aaaa3 = LocationidNum.Multiply(magicNum);
                    var aaaa4 = LocationidNum.Divide(magicNum);
                    var aaaa5 = LocationidNum.Remainder(magicNum);
                    var aaaa6 = LocationidNum.CompareTo(magicNum);
                }
                watch.Stop();
                Console.WriteLine($"方法1结束!耗时：{watch.Elapsed.TotalMilliseconds}");

                Console.WriteLine("方法2开始!");
                watch = new Stopwatch();
                watch.Start();
                for (int m = 0; m < 1000000; m++)
                {
                    //var sfasfs = strs.ToObject1<ReadDataModel>();
                    var Data = strs.ToObject<ReadDataModel>();
                    var locationid = ToolsHelper.ExtractionDigital(Data.locationid);
                    //var LocationidNum = System.Numerics.BigInteger.Parse(locationid);
                    //var magicNum = System.Numerics.BigInteger.Parse(Data.magic);
                    var LocationidNum = new BigInteger.NetCore.BigInteger(locationid); //new Integer(locationid, 10); 
                    //byte[] byteArray1 = BitConverter.GetBytes(Data.magic);
                    var magicNum = new BigInteger.NetCore.BigInteger(Data.magic);  //new Integer(Data.magic, 10); 
                    //Data.magicNum = magicNum;
                    var aaaa1 = LocationidNum.Add(magicNum);
                    var aaaa2 = LocationidNum.Subtract(magicNum);
                    var aaaa3 = LocationidNum.Multiply(magicNum);
                    var aaaa4 = LocationidNum.Divide(magicNum);
                    var aaaa5 = LocationidNum.Remainder(magicNum);
                    var aaaa6 = LocationidNum.CompareTo(magicNum);
                }
                watch.Stop();
                Console.WriteLine($"方法2结束!耗时：{watch.Elapsed.TotalMilliseconds}");
                Console.WriteLine($"");
                Console.WriteLine($"====================");
                Console.WriteLine($"");
            }

            Thread.Sleep(1000000);
        }

    }
}




