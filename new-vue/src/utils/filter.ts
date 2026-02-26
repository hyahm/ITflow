export interface TimestampFormatOptions {
  // 年月日时分秒的占位符
  y?: string; // 年
  m?: string; // 月
  d?: string; // 日
  h?: string; // 时
  i?: string; // 分
  s?: string; // 秒
  a?: string; // 星期
  q?: string; // 季度
}

/**
 * 预设格式模板
 */
export const TimePatterns = {
  // 完整时间
  FULL: '{y}-{m}-{d} {h}:{i}:{s}',
  // 日期时间
  DATETIME: '{y}-{m}-{d} {h}:{i}',
  // 日期
  DATE: '{y}-{m}-{d}',
  // 时间
  TIME: '{h}:{i}:{s}',
  // 简短时间
  SHORT_TIME: '{h}:{i}',
  // 月日
  MONTH_DAY: '{m}-{d}',
  // 中文格式
  ZH_FULL: '{y}年{m}月{d}日 {h}时{i}分{s}秒',
  ZH_DATETIME: '{y}年{m}月{d}日 {h}:{i}',
  ZH_DATE: '{y}年{m}月{d}日',
  ZH_TIME: '{h}时{i}分{s}秒',
  // 带星期
  WITH_WEEK: '{y}-{m}-{d} 周{a}',
  ZH_WITH_WEEK: '{y}年{m}月{d}日 星期{a}',
} as const;

/**
 * 格式化时间戳
 * @param timestamp - 时间戳（支持秒/毫秒）
 * @param pattern - 格式模板，支持预设key或自定义字符串
 * @returns 格式化后的时间字符串
 */
export function parseTimestamp(timestamp: number): string {
  // 处理空值
  if (timestamp === null || timestamp === undefined) {
    return '';
  }

  // 转换时间戳为数字

  // 判断时间戳是秒还是毫秒
  const isSeconds = timestamp.toString().length <= 10;
  const date = isSeconds ? new Date(timestamp * 1000) : new Date(timestamp);

  // 验证日期是否有效
  if (isNaN(date.getTime())) {
    console.warn(`Invalid timestamp: ${timestamp}`);
    return String(timestamp);
  }

  // 获取预设模板或使用自定义模板

  // 获取时间各部分
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate();
  const hour = date.getHours();
  const minute = date.getMinutes();
  const second = date.getSeconds();
  //   const week = date.getDay();
  //   const quarter = Math.floor((date.getMonth() + 3) / 3);
  return `${year}-${month}-${day} ${hour}:${minute}:${second}`;
  // 星期映射
  //   const weekMap = ['日', '一', '二', '三', '四', '五', '六'];

  // 替换占位符
  //   const result = formatPattern.replace(/{([ymdhisaq])+}/g, (match: string, key: string): string => {
  //     switch (key) {
  //       case 'y':
  //         return year.toString();
  //       case 'm':
  //         return month.toString().padStart(2, '0');
  //       case 'd':
  //         return day.toString().padStart(2, '0');
  //       case 'h':
  //         return hour.toString().padStart(2, '0');
  //       case 'i':
  //         return minute.toString().padStart(2, '0');
  //       case 's':
  //         return second.toString().padStart(2, '0');
  //       case 'a':
  //         return weekMap[week] as string;
  //       case 'q':
  //         return quarter.toString();
  //       default:
  //         return match;
  //     }
  //   });
}

/**
 * 相对时间格式化（如：刚刚，3分钟前）
 */
export function formatRelativeTime(
  timestamp: number | string,
  baseTime: number = Date.now(),
): string {
  const ts = typeof timestamp === 'string' ? parseFloat(timestamp) : timestamp;
  const isSeconds = ts.toString().length <= 10;
  const time = isSeconds ? ts * 1000 : ts;
  const diff = baseTime - time;

  const minute = 60 * 1000;
  const hour = 60 * minute;
  const day = 24 * hour;
  const month = 30 * day;
  const year = 365 * day;

  if (diff < minute) {
    return '刚刚';
  } else if (diff < hour) {
    const minutes = Math.floor(diff / minute);
    return `${minutes}分钟前`;
  } else if (diff < day) {
    const hours = Math.floor(diff / hour);
    return `${hours}小时前`;
  } else if (diff < month) {
    const days = Math.floor(diff / day);
    return `${days}天前`;
  } else if (diff < year) {
    const months = Math.floor(diff / month);
    return `${months}个月前`;
  } else {
    const years = Math.floor(diff / year);
    return `${years}年前`;
  }
}

/**
 * 获取星期几（中文）
 */
export function getChineseWeekday(timestamp: number | string): string {
  const ts = typeof timestamp === 'string' ? parseFloat(timestamp) : timestamp;
  const isSeconds = ts.toString().length <= 10;
  const date = isSeconds ? new Date(ts * 1000) : new Date(ts);
  const week = date.getDay();
  const weekdays = ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六'];
  return weekdays[week] as string;
}

/**
 * 计算年龄（根据生日时间戳）
 */
export function calculateAge(birthTimestamp: number): number {
  const birthDate = new Date(
    birthTimestamp.toString().length <= 10 ? birthTimestamp * 1000 : birthTimestamp,
  );
  const today = new Date();
  let age = today.getFullYear() - birthDate.getFullYear();
  const monthDiff = today.getMonth() - birthDate.getMonth();

  if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birthDate.getDate())) {
    age--;
  }

  return age;
}

/**
 * 获取时间戳的日期部分（去除时分秒）
 */
export function getDatePartTimestamp(timestamp: number): number {
  const date = new Date(timestamp.toString().length <= 10 ? timestamp * 1000 : timestamp);
  date.setHours(0, 0, 0, 0);
  return Math.floor(date.getTime() / 1000); // 返回秒级时间戳
}

export default parseTimestamp;
