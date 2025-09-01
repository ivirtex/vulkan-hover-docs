# vkGetCudaModuleCacheNV(3) Manual Page

## Name

vkGetCudaModuleCacheNV - Get CUDA module cache



## [](#_c_specification)C Specification

To get the CUDA module cache call:

```c++
// Provided by VK_NV_cuda_kernel_launch
VkResult vkGetCudaModuleCacheNV(
    VkDevice                                    device,
    VkCudaModuleNV                              module,
    size_t*                                     pCacheSize,
    void*                                       pCacheData);
```

## [](#_parameters)Parameters

- `device` is the logical device that destroys the Function.
- `module` is the CUDA module.
- `pCacheSize` is a pointer containing the amount of bytes to be copied in `pCacheData`
- `pCacheData` is a pointer to a buffer in which to copy the binary cache

## [](#_description)Description

If `pCacheData` is `NULL`, then the size of the binary cache, in bytes, is returned in `pCacheSize`. Otherwise, `pCacheSize` **must** point to a variable set by the application to the size of the buffer, in bytes, pointed to by `pCacheData`, and on return the variable is overwritten with the amount of data actually written to `pCacheData`. If `pCacheSize` is less than the size of the binary shader code, nothing is written to `pCacheData`, and `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`.

The returned cache **may** then be used later for further initialization of the CUDA module, by sending this cache *instead* of the PTX code when using [vkCreateCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateCudaModuleNV.html).

Note

Using the binary cache instead of the original PTX code **should** significantly speed up initialization of the CUDA module, given that the whole compilation and validation will not be necessary.

As with [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html), the binary cache depends on the specific implementation. The application **must** assume the cache upload might fail in many circumstances and thus **may** have to get ready for falling back to the original PTX code if necessary. Most often, the cache **may** succeed if the same device driver and architecture is used between the cache generation from PTX and the use of this cache. In the event of a new driver version or if using a different device architecture, this cache **may** become invalid.

Note

This query does not behave consistently with the behavior described in [Opaque Binary Data Results](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fundamentals-binaryresults), for historical reasons.

If the amount of data available is larger than the passed `pDataSize`, the query returns a `VK_INCOMPLETE` success status instead of a `VK_ERROR_NOT_ENOUGH_SPACE_KHR` error status.

Valid Usage (Implicit)

- [](#VUID-vkGetCudaModuleCacheNV-device-parameter)VUID-vkGetCudaModuleCacheNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetCudaModuleCacheNV-module-parameter)VUID-vkGetCudaModuleCacheNV-module-parameter  
  `module` **must** be a valid [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaModuleNV.html) handle
- [](#VUID-vkGetCudaModuleCacheNV-pCacheSize-parameter)VUID-vkGetCudaModuleCacheNV-pCacheSize-parameter  
  `pCacheSize` **must** be a valid pointer to a `size_t` value
- [](#VUID-vkGetCudaModuleCacheNV-pCacheData-parameter)VUID-vkGetCudaModuleCacheNV-pCacheData-parameter  
  If the value referenced by `pCacheSize` is not `0`, and `pCacheData` is not `NULL`, `pCacheData` **must** be a valid pointer to an array of `pCacheSize` bytes
- [](#VUID-vkGetCudaModuleCacheNV-module-parent)VUID-vkGetCudaModuleCacheNV-module-parent  
  `module` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_cuda\_kernel\_launch](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cuda_kernel_launch.html), [VkCudaModuleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaModuleNV.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetCudaModuleCacheNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0