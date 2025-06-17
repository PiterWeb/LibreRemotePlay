export default async function wrapAsync<T>(promise: Promise<T>) {
  return promise
    .then(data => [data, null])
    .catch((err: string) => [null, err])
}