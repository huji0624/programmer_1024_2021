using NeinMath;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Numerics;
using System.Text;
using System.Threading.Tasks;

namespace Test
{
    public class ReadDataModel
    {
        public string locationid { get; set; }
        public string magic { get; set; }
        //public Integer magicNum { get; set; }
        public int errorNum { get; set; }

    }
    public class ResponseModel
    {
        public int errorno { get; set; }
        public List<string> data { get; set; }
    }

    public class Submit1024DataModel
    {
        public string formula { get; set; }
        public int errorNum { get; set; }
        public List<string> locationid { get; set; }
    }
    public class NumDicModel
    {
        public Integer magicNum { get; set; }
        public string locationid { get; set; }

    }
    public class WaitNumDicModel
    {
        public long magicNum { get; set; }
        public string locationid { get; set; }
        public List<string> locationidList { get; set; }

    }
    public class ZeroModel
    {
        public string str { get; set; }
        public List<string> locationidList { get; set; } = new List<string>();

    }

}
