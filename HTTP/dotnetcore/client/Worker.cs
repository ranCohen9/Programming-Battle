using System;
using System.Net;
using System.Net.Http;
using System.Threading.Tasks;

public class Worker
{
    public async Task<double> GetOnce(int id)
    {
        try
        {
            using (HttpClient httpClient = new HttpClient())
            {
                DateTime start=DateTime.Now;
                var response = await httpClient.GetAsync($"http://localhost:5000/api/values/{id}");
                var resString = await response.Content.ReadAsStringAsync();
                DateTime stop=DateTime.Now;
                //System.Console.WriteLine($"elapsed get time {(stop-start).TotalMilliseconds}");
                return (stop-start).TotalMilliseconds;
            }
        }
        catch (System.Exception ex)
        {
            System.Console.WriteLine(ex.Message);
            return double.MaxValue;
        }
    }

    public async Task GetMany(){
        double sum=0;
        int i=0;
        for (; i < 100; i++)
        {
            double result=await this.GetOnce(i);
            sum+=result;
        }
        System.Console.WriteLine($"finished GET loop avg {sum/100} | {i}");
    }

    public async Task<double> PostOnce(int id){
        try
        {
            using(HttpClient httpClient=new HttpClient()){
                DateTime start=DateTime.Now;
                HttpContent body= new StringContent(id.ToString());
                var response=await httpClient.PostAsync("http://localhost:5000/api/values",body);
                var reslut=await response.Content.ReadAsStringAsync();
                DateTime stop=DateTime.Now;
                return (stop-start).TotalMilliseconds;
            }
        }
        catch (System.Exception ex)
        {
            System.Console.WriteLine(ex.Message);
            return double.MaxValue;
        }
    }

    public async Task PostMany(){
        double sum=0;
        int i=0;
        for (; i < 100; i++)
        {
            double result=await this.PostOnce(i);
            sum+=result;
        }
        System.Console.WriteLine($"finished POST loop avg {sum/100} | {i}");
    }


}