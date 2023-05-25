import service from '@/utils/request'
export const deleteImage = (data) => {
  return service({
    url: 'private/deleteImage',
    method: 'delete',
    data
  })
}
