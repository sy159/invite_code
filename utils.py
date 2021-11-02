# python3，Python2把//改为/

CHARS = ['5', 'd', 'f', 'u', 'c', 'y', 'a', 'r', '1', 'j',
         '2', 't', 'x', 'w', 'g', 's', '8', 'm', 'v', 'p',
         '4', 'q', 'h', 'b', '3', 'n', '6', 'k', '7', 'e',
         'z', '9']

CHAR_LEN = len(CHARS)
DIVIDER = 'i'  # 分割标识(区分补位，应该是chars里面没出现的字符)


def id_to_code(from_id: int, min_length: int = 6) -> str:
    code = ""
    try:
        from_id = int(from_id)
        while from_id // CHAR_LEN > 0:
            code += CHARS[from_id % CHAR_LEN]  # 通过余数获取索引位置生成码
            from_id = from_id // CHAR_LEN
        code += CHARS[from_id % CHAR_LEN]  # 处理未除尽的余数
        code = code[::-1]
        fix_len = min_length - len(code)  # 需要补码的长度
        if fix_len > 0:
            code += DIVIDER
            for i in range(fix_len - 1):
                code += CHARS[i]
    except Exception as e:
        pass
    return code


def code_to_id(code: str) -> int:
    if not code:
        return 0
    code = str(code)
    res_id = 0
    try:
        for i in range(len(code)):
            if code[i] == DIVIDER:
                break
            try:
                char_index = CHARS.index(code[i])
            except ValueError as e:
                char_index = 0
            if i > 0:
                res_id = res_id * CHAR_LEN + char_index
            else:
                res_id = char_index
    except Exception as e:
        pass
    return res_id

