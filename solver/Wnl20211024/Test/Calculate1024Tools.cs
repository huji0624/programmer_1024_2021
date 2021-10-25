using System;
using System.Collections.Generic;
using System.Linq;
using System.Numerics;
using System.Text;
using System.Threading.Tasks;

namespace Test
{
    public class Calculate1024Tools
    {

        public static List<string> GetData(List<long> nums)
        {
            List<express> list = calc(nums);
            list = list.Where((x, i) => list.FindIndex(z => z.exp == x.exp) == i).ToList();
            List<string> answer = list.Where(t => t.val == 1024).Select(t => t.exp).ToList();
            

            return answer;
        }

        private static List<express> calc(List<long> nums)
        {
            List<express> list = new List<express>();
            if (nums.Count == 1)
            {
                express p = new express() { oper = 0, val = nums[0], exp = nums[0].ToString() };
                list.Add(p);
                return list;
            }
            foreach (var item in nums)
            {
                List<long> t = new List<long>(nums);
                t.Remove(item);
                List<express> temp = calc(t);
                foreach (var tt in temp)
                {
                    String str1 = getstr(item, tt, 1);
                    String str2 = getstr(item, tt, 2);
                    String str3 = getstr(item, tt, 3);
                    String str4 = getstr(item, tt, 4);
                    String str5 = getstr(item, tt, 5);
                    String str6 = getstr(item, tt, 6);
                    list.Add(new express() { oper = 1, val = item + tt.val, exp = str1 });
                    list.Add(new express() { oper = 2, val = item - tt.val, exp = str2 });
                    list.Add(new express() { oper = 3, val = item * tt.val, exp = str3 });

                    if (tt.val != 0) list.Add(new express() { oper = 4, val = item / tt.val, exp = str4 });
                    list.Add(new express() { oper = 5, val = tt.val - item, exp = str5 });
                    if (item != 0) list.Add(new express() { oper = 6, val = tt.val / item, exp = str6 });
                }

            }
            return list;
        }

        private static String getstr(long num, express p, int oper)
        {
            switch (oper)
            {
                case 1:
                    return num + "+" + p.exp;
                case 2:
                    return num + "-" + (p.oper == 1 || p.oper == 2 || p.oper == 5 ? "(" + p.exp + ")" : p.exp);
                case 3:
                    return num + "*" + (p.oper == 1 || p.oper == 2 || p.oper == 5 ? "(" + p.exp + ")" : p.exp);
                case 4:
                    return num + "/" + (p.oper == 0 ? p.exp : "(" + p.exp + ")");
                case 5:
                    return p.exp + "-" + num;
                case 6:
                    return (p.oper == 1 || p.oper == 2 || p.oper == 5 ? "(" + p.exp + ")" : p.exp) + "/" + num;
                default:
                    return "";
            }
        }

        class express
        {
            public long val { get; set; }
            public string exp { get; set; }
            public int oper { get; set; }
        }
    }
}
