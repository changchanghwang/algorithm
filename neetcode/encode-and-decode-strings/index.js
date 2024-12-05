class Solution {
  separator = ';';
  /**
   * @param {string[]} strs
   * @returns {string}
   */
  encode(strs) {
    if (strs.length === 0) {
      return null;
    }
    return strs
      .map((str) =>
        str.includes(this.separator) ? str.replaceAll(this.separator, `${this.separator}${this.separator}`) : str,
      )
      .join(this.separator);
  }

  /**
   * @param {string} str
   * @returns {string[]}
   */
  decode(str) {
    if (str === null) {
      return [];
    }
    return str.split(this.separator);
  }
}

const input = ['1,23', '45,6', '7,8,9'];
const solution = new Solution();

console.log(solution.decode(solution.encode(input)));
