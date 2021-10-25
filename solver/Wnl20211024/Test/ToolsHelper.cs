using NeinMath;
using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using System.Numerics;
using System.Text;
using System.Threading.Tasks;

namespace Test
{
    public static class ToolsHelper
    {
        public static Integer Num1024 = Integer.Parse("1024");
        /// 随机排列数组元素
        /// </summary>
        /// <param name="myList"></param>
        /// <returns></returns>
        public static List<T> RandomSort<T>(this List<T> list)
        {
            var random = new Random();
            var newList = new List<T>();
            foreach (var item in list)
            {
                newList.Insert(random.Next(newList.Count), item);
            }
            return newList;
        }

        /// <summary>
        /// 验证数字
        /// </summary>
        /// <param name="Num"></param>
        /// <returns></returns>
        public static bool VerifyDigital(Integer LocationidNum, Integer Magic)
        {
            if (LocationidNum.CompareTo(Magic) == 1)
            {
                if (
                (LocationidNum.Subtract(Num1024)).CompareTo(Magic) == 0
                || (LocationidNum.Remainder(Num1024)).CompareTo(Magic) == 0
                )
                {
                    return true;
                }
            }
            else
            {
                if (
                (LocationidNum.Add(Num1024)).CompareTo(Magic) == 0
                || (LocationidNum.Multiply(Num1024)).CompareTo(Magic) == 0
                )
                {
                    return true;
                }
            }

            return false;
        }

        /// <summary>
        /// 验证数字
        /// </summary>
        /// <param name="Num"></param>
        /// <returns></returns>
        public static bool VerifyDigital_long(long LocationidNum, long Magic)
        {
            if (LocationidNum >= Magic)
            {
                if (
                (LocationidNum - 1024) == Magic
                || (LocationidNum % 1024) == (Magic)
                )
                {
                    return true;
                }
            }
            else
            {
                if (
                 (LocationidNum + 1024) == Magic
                || (LocationidNum * 1024) == (Magic)
                )
                {
                    return true;
                }
            }

            return false;
        }


        /// <summary>
        /// 提取字符串的数字(快)
        /// </summary>
        /// <param name="Str"></param>
        /// <returns></returns>
        public static string ExtractionDigital(string Str)
        {
            //string result = System.Text.RegularExpressions.Regex.Replace(s, @"[^0-9]+", "");
            //return int.Parse(result);
            StringBuilder sb = new StringBuilder();
            foreach (char c in Str)
            {
                if (Convert.ToInt32(c) >= 48 && Convert.ToInt32(c) <= 57)
                {
                    sb.Append(c);
                }
            }
            return sb.ToString();
        }
        /// <summary>
        /// 提取字符串的数字（满）
        /// </summary>
        /// <param name="Str"></param>
        /// <returns></returns>
        public static string ExtractionDigital1(string Str)
        {
            string result = System.Text.RegularExpressions.Regex.Replace(Str, @"[^0-9]+", "");
            return result;
        }

        public static List<T> PageBy<T>(this List<T> query, int pageNum, int pageSize)
        {
            if (query == null)
            {
                throw new ArgumentNullException("query");
            }

            return query.Skip((pageNum - 1) * pageSize).Take(pageSize).ToList();
        }





    }
}
