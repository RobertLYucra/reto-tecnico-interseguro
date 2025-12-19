class ResponseUtil {
  static success(message, data = null) {
    return {
      success: true,
      message,
      data,
    };
  }

  static error(message, errorDetail = null) {
    const errors = errorDetail ? [errorDetail] : [message];
    return {
      success: false,
      message,
      errors,
    };
  }
}

module.exports = ResponseUtil;
