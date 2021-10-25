using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using Newtonsoft.Json.Serialization;
using Swifter.Json;
using System;
using System.Collections.Generic;
using System.Data;
using System.Text;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.Text.Unicode;

namespace Test
{
    /// <summary>
    /// json 序列化扩展
    /// </summary>
    /// <summary>
    /// json 序列化扩展
    /// </summary>
    public static class JsonExtensionHelper
    {
        /// <summary>
        /// object转换为Json
        /// </summary>
        /// <param name="obj">object</param>
        /// <returns></returns>
        public static string ToJson(this object obj)
        {
            JsonSerializerOptions options = new JsonSerializerOptions();
            //日期格式化
            options.Encoder = System.Text.Encodings.Web.JavaScriptEncoder.Create(UnicodeRanges.All);
            options.Converters.Add(new SystemTextJsonConvert.DateTimeConverter());

            //return System.Text.Json.JsonSerializer.Serialize(obj, options);
            var timeConverter = new IsoDateTimeConverter { DateTimeFormat = "yyyy-MM-dd HH:mm:ss" };
            return JsonConvert.SerializeObject(obj, timeConverter);
        }
        /// <summary>
        /// object转换为Json(自定义日期格式)
        /// </summary>
        /// <param name="obj">object</param>
        /// <param name="datetimeformats">日期格式</param>
        /// <returns></returns>
        public static string ToJson(this object obj, string datetimeformats)
        {
            var timeConverter = new IsoDateTimeConverter { DateTimeFormat = datetimeformats };
            return JsonConvert.SerializeObject(obj, timeConverter);
        }
        /// <summary>
        /// Json转换为object
        /// </summary>
        /// <typeparam name="T">类型</typeparam>
        /// <param name="Json">Json</param>
        /// <returns></returns>
        public static T ToObject<T>(this string Json)
        {
            //return Json == null ? default(T) : System.Text.Json.JsonSerializer.Deserialize<T>(Json);
            return Json == null ? default(T) : JsonFormatter.DeserializeObject<T>(Json);
        }
        /// <summary>
        /// Json转换为object
        /// </summary>
        /// <typeparam name="T">类型</typeparam>
        /// <param name="Json">Json</param>
        /// <returns></returns>
        public static T ToObject1<T>(this string Json)
        {
            return Json == null ? default(T) : System.Text.Json.JsonSerializer.Deserialize<T>(Json);
            //return Json == null ? default(T) : JsonFormatter.DeserializeObject<T>(Json);
        }
        /// <summary>
        /// Json转换为List
        /// </summary>
        /// <typeparam name="T">类型</typeparam>
        /// <param name="Json">Json</param>
        /// <returns></returns>
        public static List<T> ToList<T>(this string Json)
        {
            JsonSerializerOptions options = new JsonSerializerOptions();
            //日期格式化
            options.Converters.Add(new SystemTextJsonConvert.DateTimeConverter());
            try
            {
                return Json == null ? null : System.Text.Json.JsonSerializer.Deserialize<List<T>>(Json, options);
            }
            catch (Exception)
            {
                return Json == null ? null : JsonConvert.DeserializeObject<List<T>>(Json);
            }

        }
       
        /// <summary>
        /// Json转换为JObject
        /// </summary>
        /// <param name="Json">Json</param>
        /// <returns></returns>
        public static JObject ToJObject(this string Json)
        {
            return Json == null ? JObject.Parse("{}") : JObject.Parse(Json.Replace("&nbsp;", ""));
        }


    }

    public class SystemTextJsonConvert
    {
        public class DateTimeConverter : System.Text.Json.Serialization.JsonConverter<DateTime>
        {
            public override DateTime Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
            {
                return DateTime.Parse(reader.GetString());
            }
            public override void Write(Utf8JsonWriter writer, DateTime value, JsonSerializerOptions options)
            {
                writer.WriteStringValue(value.ToString("yyyy-MM-dd HH:mm:ss"));
            }
        }

        public class DateTimeNullableConverter : System.Text.Json.Serialization.JsonConverter<DateTime?>
        {
            public override DateTime? Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
            {
                return string.IsNullOrEmpty(reader.GetString()) ? default(DateTime?) : DateTime.Parse(reader.GetString());
            }
            public override void Write(Utf8JsonWriter writer, DateTime? value, JsonSerializerOptions options)
            {
                writer.WriteStringValue(value?.ToString("yyyy-MM-dd HH:mm:ss"));
            }
        }
    }
}
