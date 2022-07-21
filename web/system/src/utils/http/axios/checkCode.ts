import { useI18n } from '/@/hooks/web/useI18n';
import { useUserStoreWithOut } from '/@/store/modules/user';
import projectSetting from '/@/settings/projectSetting';
import { SessionTimeoutProcessingEnum } from '/@/enums/appEnum';
import { useMessage } from '/@/hooks/web/useMessage';
import httpStatus from 'http-status';

const { createErrorModal } = useMessage();
const stp = projectSetting.sessionTimeoutProcessing;

export function checkCode(status: number, msg: string): boolean {
  if (status === 200) {
    return false;
  }
  const { t } = useI18n();
  const userStore = useUserStoreWithOut();
  let errMessage = '';

  switch (status) {
    case 400:
      errMessage = t(msg);
      break;
    // 401: Not logged in
    // Jump to the login page if not logged in, and carry the path of the current page
    // Return to the current page after successful login. This step needs to be operated on the login page.
    case httpStatus.UNAUTHORIZED:
      userStore.setToken(undefined);
      errMessage = msg || t('sys.api.errMsg401');
      if (stp === SessionTimeoutProcessingEnum.PAGE_COVERAGE) {
        userStore.setSessionTimeout(true);
      } else {
        userStore.logout(true);
      }
      break;
    case httpStatus.FORBIDDEN:
      errMessage = t('sys.api.errMsg403');
      break;
    // 404请求不存在
    case httpStatus.NOT_FOUND:
      errMessage = t('sys.api.errMsg404');
      break;
    case httpStatus.METHOD_NOT_ALLOWED:
      errMessage = t('sys.api.errMsg405');
      break;
    case httpStatus.REQUEST_TIMEOUT:
      errMessage = t('sys.api.errMsg408');
      break;
    case httpStatus.INTERNAL_SERVER_ERROR:
      errMessage = t('sys.api.errMsg500');
      break;
    case httpStatus.NOT_IMPLEMENTED:
      errMessage = t('sys.api.errMsg501');
      break;
    case httpStatus.BAD_GATEWAY:
      errMessage = t('sys.api.errMsg502');
      break;
    case httpStatus.SERVICE_UNAVAILABLE:
      errMessage = t('sys.api.errMsg503');
      break;
    case httpStatus.GATEWAY_TIMEOUT:
      errMessage = t('sys.api.errMsg504');
      break;
    case httpStatus.HTTP_VERSION_NOT_SUPPORTED:
      errMessage = t('sys.api.errMsg505');
      break;
    default:
  }
  createErrorModal({ title: t('sys.api.errorTip'), content: errMessage });
  return true;
}
