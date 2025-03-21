import dayjs from "dayjs";

export const formatTime = (time: string) => {
  return dayjs(time).format("YYYY-MM-DD");
};

// 计算传入时间与当前时间的时间差
export const timeAgo = (time: string) => {
  const now = dayjs();
  const minuteDiff = now.diff(dayjs(time), "minute");
  const hourDiff = now.diff(dayjs(time), "hour");
  const dayDiff = now.diff(dayjs(time), "day");
  const monthDiff = now.diff(dayjs(time), "month");
  const yearDiff = now.diff(dayjs(time), "year");
  if (minuteDiff < 60) {
    return `${minuteDiff}分钟前`;
  } else if (hourDiff < 24) {
    return `${hourDiff}小时前`;
  } else if (dayDiff < 30) {
    return `${dayDiff}天前`;
  } else if (monthDiff < 12) {
    return `${monthDiff}个月前`;
  } else {
    return `${yearDiff}年前`;
  }
};

// 计算文章字数和阅读时间
export const readTime = (content: string) => {
  const wordsPerMinute = 500;
  const words = content.length;
  const minutes = Math.ceil(words / wordsPerMinute);
  return { words, minutes };
};
