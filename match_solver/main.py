"""Main Module"""

import urllib.request
import threading
from multiprocessing import Process
import time
import json
import random


class LocMagic(object):
    def __init__(self, locId, magicId, locNum, magicNum, f_no):
        self.locId = locId
        self.magicId = magicId
        self.locNum = magicNum
        self.numLen = len(str(locNum))
        self.magicNum = magicNum
        self.f_no = f_no

    def __eq__(self, other):
        return self.locId == other.locId

    def __hash__(self) -> int:
        return hash(self.locId)


found_loc_magics = set()
reported_loc_magics = list()
formulated_loc_magics = list()

found_lock = threading.Lock()
report_lock = threading.Lock()

skip_nos = set()

class Factor(object):
    def __init__(self, result, expr, fac_loc_magics):
        self.result = result
        self.expr = expr
        self.fac_loc_magics = fac_loc_magics

    def add(self, factor):
        self.result += factor.result
        self.expr = "%s+%s" % (self.expr, factor.expr)
        self.fac_loc_magics += factor.fac_loc_magics

    def substract(self, factor):
        self.result -= factor.result
        self.expr = "%s-%s" % (self.expr, factor.expr)
        self.fac_loc_magics += factor.fac_loc_magics

    def multiply(self, factor):
        self.result *= factor.result
        self.expr = "(%s)*(%s)" % (self.expr, factor.expr)
        self.fac_loc_magics += factor.fac_loc_magics

    def divide(self, factor):
        self.result = self.result // factor.result
        self.expr = "(%s)/(%s)" % (self.expr, factor.expr)
        self.fac_loc_magics += factor.fac_loc_magics


def find_bao(file_path, f_number):
    """find bao
    file_path -- Bao file path
    """
    print("start find file %s", file_path)
    bao_count = 0
    line_count = 0
    with open(file_path, encoding='ascii') as bao_file:
        for line in bao_file:
            if f_number in skip_nos:
                # print("skippppppppppppp %d", f_number)
                # print(skip_nos)
                return

            loc_magic = line.split(',')
            loc_id = loc_magic[0][15:79]
            magic_id = loc_magic[1][9:-3]
            loc_num = int("".join(list(filter(str.isdigit, loc_id))))
            magic_num = int(magic_id)
            if magic_num in (loc_num + 1024, loc_num - 1024,
                             loc_num * 1024, loc_num % 1024):
                bao_count += 1
                found_lock.acquire()
                found_loc_magics.add(
                    LocMagic(loc_id, magic_id, loc_num, magic_num, f_number))
                found_lock.release()
            line_count += 1
    # print(line_count)
    # print(bao_count)


def test_line():
    """test unit"""
    line = '{"locationid":"zdn0lox372z2u4rpdp84rld1dn5e95gwbp7eeenldoupug5ohkg90svy0ije2jrs","magic":"3598006090807"}\n'
    loc_magic = line.split(',')
    loc_id = loc_magic[0][15:79]
    magic_id = loc_magic[1][9:-3]
    print(line)
    print(loc_id)
    print(magic_id)


url_token = "01aef5b0d5d5d8aedb710459fcfac94c"


def report_bao_run():
    """
    {
        "locationid":"2ixewekdyvgbfmfzli9iifm1w9hnd2ij5kr1avy1zw3c7rl", //宝藏地点id
        "token":"ooaksuquwqiw=928182ijasj" //请求token
    }
    """
    while True:
        if found_loc_magics:
            found_lock.acquire()
            any_found = found_loc_magics.pop()
            found_lock.release()
            req_json = {'locationid': any_found.locId, 'token': url_token}
            req_json_data = bytes(json.dumps(req_json), 'utf8')
            req = urllib.request.Request(
                "http://47.104.220.230/dig", data=req_json_data)
            try:
                resp = urllib.request.urlopen(req, timeout=3)
            except:
                print("exception dig")
                time.sleep(0.02)
                continue

            resp = resp.read().decode('utf-8')
            resp_json = json.loads(resp)
            resp_err_no = resp_json['errorno']
            if resp_err_no in (0, 2):
                if resp_err_no == 2 and len(skip_nos) < 20:
                    skip_nos.add(any_found.f_no)
                    # print("skip noooooooo %d, ---%d", any_found.f_no, len(skip_nos))
                    # print(skip_nos)
                    continue
                report_lock.acquire()
                reported_loc_magics.append(any_found)
                report_lock.release()
                # print("didi" + ' ' + str(resp_json['errorno']) + ' ' + str(
                    # any_found.locNum) + ' ' + str(len(found_loc_magics)))
        else:
            time.sleep(0.05)


target_max = 40


def report_formula_run():
    """report formula run"""
    while True:
        if reported_loc_magics:
            report_lock.acquire()
            reported_count = len(reported_loc_magics)

            if reported_count < 20:
                report_lock.release()
                time.sleep(0.02)
                continue
            
            cal_locs = reported_loc_magics[0:reported_count]
            report_lock.release()

            unorder_locs = cal_locs.copy()
            random.shuffle(unorder_locs)
            # print("impossible len = %d , %d" % (len(unorder_locs), reported_count))
            factor = find_formular(unorder_locs)
            if factor == None:
                time.sleep(0.02)
                continue
            req_json = {'formula': factor.expr, 'token': url_token}
            req_json_data = bytes(json.dumps(req_json), 'utf8')
            req = urllib.request.Request(
                "http://47.104.220.230/formula", data=req_json_data)
            try:
                resp = urllib.request.urlopen(req, timeout=3)
            except TimeoutError as timeout_error:
                print("exception formula")
                print(timeout_error)
                time.sleep(0.02)
                continue
            except:
                print("exception formula unknown")
                time.sleep(0.02)
                continue

            resp = resp.read().decode('utf-8')
            resp_json = json.loads(resp)
            error_no = resp_json['errorno']
            if error_no == 0:
                report_lock.acquire()
                for cal_factor in factor.fac_loc_magics:
                    reported_loc_magics.remove(cal_factor)
                report_lock.release()
                # print("formula found %s, %d" % (factor.expr, factor.result))
            elif error_no == 3:
                wrong_loc_ids = resp_json["data"]
                report_lock.acquire()
                for cal_factor in factor.fac_loc_magics:
                    if cal_factor.locId in wrong_loc_ids:
                        reported_loc_magics.remove(cal_factor)
                report_lock.release()
            else:
                # exit(1000)
                # print("errorrrrrrr %s, %d" % (factor.expr, factor.result))
                time.sleep(0.02)
        else:
            time.sleep(0.05)


def find_formular(cal_locs, factor_ori=None, num_target=1024):
    """find formular"""
    if len(cal_locs) < 2:
        return None
    else:
        diff = 0
        add_or_sub = True
        if factor_ori:
            diff = num_target - factor_ori.result
            add_or_sub = diff >= 0
            diff = abs(diff)
        else:
            diff = num_target

        diff = min(diff, 300)
        two_locs = find_two_locs(cal_locs, diff)
        if two_locs:
            first_loc = two_locs[0]
            second_loc = two_locs[1]
            find_factor = Factor(first_loc.locNum // second_loc.locNum, "%s/%s" % (
                first_loc.locId, second_loc.locId), [first_loc, second_loc])
            if factor_ori:
                if add_or_sub:
                    factor_ori.add(find_factor)
                    if factor_ori.result == 1024:
                        return factor_ori
                else:
                    factor_ori.substract(find_factor)
                    if factor_ori.result == 1024:
                        return factor_ori
            else:
                factor_ori = find_factor
            next_locs = cal_locs.copy()
            next_locs.remove(first_loc)
            next_locs.remove(second_loc)
            return find_formular(next_locs, factor_ori)
        else:
            # print("find formula none, %d" % len(cal_locs))
            return None


def find_two_locs(cal_locs, diff):
    """find two locs"""
    find_locs = cal_locs.copy()
    if len(find_locs) < 2:
        return None
    try_count = max(5, len(cal_locs) // 3)
    for i in range(5):
        first_loc = random.sample(find_locs, 1)[0]
        for find_loc in find_locs:
            if find_loc.locId == first_loc.locId:
                continue
            if abs(find_loc.numLen - first_loc.numLen) > 3:
                continue
            if find_loc.locNum < first_loc.locNum:
                loc_value = first_loc.locNum // find_loc.locNum
                if loc_value <= diff:
                    return (first_loc, find_loc)
            else:
                loc_value = find_loc.locNum // first_loc.locNum
                if loc_value <= diff:
                    return (find_loc, first_loc)
    # print("find_two_locs len=%d, %d" % (len(cal_locs), diff))
    return None


def find_bao_run(file_list):
    for file_no in file_list:
        find_bao("data/Treasure_%d.data" % file_no)


def main2(bao_file_list):
    """main func"""
    # test_line()
    t_report_formula = threading.Thread(target=report_formula_run)
    t_report_bao1 = threading.Thread(target=report_bao_run, name='report1')
    t_report_bao2 = threading.Thread(target=report_bao_run, name='report2')
    t_report_bao3 = threading.Thread(target=report_bao_run, name='report3')

    t_report_formula.start()
    t_report_bao1.start()
    t_report_bao2.start()
    t_report_bao3.start()

    # bao_file_list = []
    # for i in range(40):
    #     bao_file_list.append(i)
    # random.shuffle(bao_file_list)

    for file_no in bao_file_list:
        find_bao("data/Treasure_%d.data" % file_no, file_no)

    # t_find_bao1 = threading.Thread(target=find_bao_run, name='find1', args=(bao_file_list[0:20],))
    # t_find_bao2 = threading.Thread(target=find_bao_run, name='find2', args=(bao_file_list[20:40],))
    # t_find_bao1.start()
    # t_find_bao2.start()

    print("Done.")


def main():
    bao_file_list = []
    for i in range(128):
        bao_file_list.append(i)
    random.shuffle(bao_file_list)

    p1 = Process(target=main2, args=(bao_file_list[0:64],))
    p2 = Process(target=main2, args=(bao_file_list[64:128],))
    p1.start()
    p2.start()
    p1.join()
    p2.join()


if __name__ == '__main__':
    main()
