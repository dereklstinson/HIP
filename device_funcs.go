package hip

//#include <hip/hip_runtime_api.h>
import "C"
import "unsafe"

//Device is a device
type Device C.hipDevice_t

func (d Device) c() C.hipDevice_t      { return (C.hipDevice_t)(d) }
func (d *Device) cptr() *C.hipDevice_t { return (*C.hipDevice_t)(d) }

func DeviceSynchronize() error {
	return status(C.hipDeviceSynchronize()).error("hipDeviceSynchronize")
}
func DeviceReset() error { return status(C.hipDeviceReset()).error("hipDeviceReset") }

func pSetDevice(deviceID int32) error {
	return status(C.hipSetDevice((C.int)(deviceID))).error("hipSetDevice")
}
func GetDevice() (deviceID int32, err error) {
	err = status(C.hipGetDevice((*C.int)(&deviceID))).error("hipGetDevice")
	return deviceID, err
}
func GetDeviceCount() (count int32, err error) {
	err = status(C.hipGetDeviceCount((*C.int)(&count))).error("hipGetDeviceCount")
	return count, err
}

func DeviceGetAttribute(attr DeviceAttribute, deviceID int32) (val int32, err error) {
	err = status(C.hipDeviceGetAttribute((*C.int)(&val), attr.c(), (C.int)(deviceID))).error("hipDeviceGetAttribute")
	return val, err
}

func GetDeviceProperties(deviceID int32) (prop *DeviceProp, err error) {
	prop = new(DeviceProp)
	err = status(C.hipGetDeviceProperties(prop.cptr(), (C.int)(deviceID))).error("hipGetDeviceProperties")
	return prop, err
}

func DeviceSetCacheConfig(cacheConfig FuncCache) error {
	return status(C.hipDeviceSetCacheConfig(cacheConfig.c())).error("hipDeviceSetCacheConfig")
}

func DeviceGetCacheConfig() (cacheConfig FuncCache, err error) {
	err = status(C.hipDeviceGetCacheConfig(cacheConfig.cptr())).error("hipDeviceGetCacheConfig")
	return cacheConfig, err
}

func DeviceGetLimit(limit Limit) (pValue uint, err error) {
	var sizet C.size_t
	err = status(C.hipDeviceGetLimit(&sizet, limit.c())).error("hipDeviceGetLimit")
	pValue = (uint)(sizet)
	return pValue, err
}

//for hipgo this has to be used with .so functions
func FuncSetCacheConfig(sofunc unsafe.Pointer, config FuncCache) error {
	return status(C.hipFuncSetCacheConfig(sofunc, config.c())).error("hipFuncSetCacheConfig")
}
func DeviceGetSharedMemConfig() (pConfig SharedMemConfig, err error) {
	err = status(C.hipDeviceGetSharedMemConfig(pConfig.cptr())).error("hipDeviceGetSharedMemConfig")
	return pConfig, err
}
func DeviceSetSharedMemConfig(pConfig SharedMemConfig) error {
	return status(C.hipDeviceSetSharedMemConfig(pConfig.c())).error("hipDeviceSetSharedMemConfig")
}
func SetDeviceFlags(dflag DeviceFlag) error {
	return status(C.hipSetDeviceFlags(dflag.c())).error("hipSetDeviceFlags")
}

func ChooseDevice(prop *DeviceProp) (device int32, err error) {
	err = status(C.hipChooseDevice((*C.int)(&device), prop.cptr())).error("hipChooseDevice")
	return device, err
}

func DeviceGetPCIBusId(device int32) (pciBusId string, err error) {
	length := C.int(255)
	bid := make([]C.char, length)
	err = status(C.hipDeviceGetPCIBusId(&bid[0], length, (C.int)(device))).error("hipDeviceGetPCIBusId")
	pciBusId = C.GoString(&bid[0])
	return pciBusId, err

}

func DeviceGetByPCIBusId(pciBusId string) (device int32, err error) {
	bid := []byte(pciBusId)
	cbid := make([]C.char, len(bid))
	for i := range bid {
		cbid[i] = (C.char)(bid[i])
	}
	err = status(C.hipDeviceGetByPCIBusId((*C.int)(&device), &cbid[0])).error("hipDeviceGetByPCIBusId")
	return device, err
}

func DeviceGet(ordinal int32) (d Device, err error) {
	err = status(C.hipDeviceGet(d.cptr(), (C.int)(ordinal))).error("hipDeviceGet")
	return d, err
}
func (d Device) ComputeCapability() (major, minor int32, err error) {
	err = status(C.hipDeviceComputeCapability((*C.int)(&major), (*C.int)(&minor), d.c())).error("hipDeviceComputeCapability")
	return major, minor, err
}
func (d Device) GetName() (name string, err error) {
	length := C.int(255)
	cname := make([]C.char, length)
	err = status(C.hipDeviceGetName(&cname[0], length, d.c())).error("hipDeviceGetName")
	name = C.GoString(&cname[0])
	return name, err
}
func (d Device) GetTotalMem() (totalSIB uint, err error) {
	var sib C.size_t
	err = status(C.hipDeviceTotalMem(&sib, d.c())).error("GetTotalMem")
	totalSIB = (uint)(sib)
	return totalSIB, err
}

//func  hipDevicePrimaryCtxGetState(hipDevice_t dev, unsigned int* flags, int* active)error{return status(C.hipDevicePrimaryCtxGetState(hipDevice_t dev, unsigned int* flags, int* active)).error("hipDevicePrimaryCtxGetState")}
//func  hipDevicePrimaryCtxRelease(hipDevice_t dev)error{return status(C.hipDevicePrimaryCtxRelease(hipDevice_t dev)).error("hipDevicePrimaryCtxRelease")}
//func  hipDevicePrimaryCtxRetain(hipCtx_t* pctx, hipDevice_t dev)error{return status(C.hipDevicePrimaryCtxRetain(hipCtx_t* pctx, hipDevice_t dev)).error("hipDevicePrimaryCtxRetain")}
//func  hipDevicePrimaryCtxReset(hipDevice_t dev)error{return status(C.hipDevicePrimaryCtxReset(hipDevice_t dev)).error("hipDevicePrimaryCtxReset")}
//func  hipDevicePrimaryCtxSetFlags(hipDevice_t dev, unsigned int flags)error{return status(C.hipDevicePrimaryCtxSetFlags(hipDevice_t dev, unsigned int flags)).error("hipDevicePrimaryCtxSetFlags")}