# vkGetPipelineBinaryDataKHR(3) Manual Page

## Name

vkGetPipelineBinaryDataKHR - Get the data store from a pipeline binary



## [](#_c_specification)C Specification

Data **can** be retrieved from a pipeline binary object using the command:

```c++
// Provided by VK_KHR_pipeline_binary
VkResult vkGetPipelineBinaryDataKHR(
    VkDevice                                    device,
    const VkPipelineBinaryDataInfoKHR*          pInfo,
    VkPipelineBinaryKeyKHR*                     pPipelineBinaryKey,
    size_t*                                     pPipelineBinaryDataSize,
    void*                                       pPipelineBinaryData);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the pipeline binary.
- `pInfo` is a pointer to a [VkPipelineBinaryDataInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryDataInfoKHR.html) structure which describes the pipeline binary to get data from.
- `pPipelineBinaryKey` is a pointer to a [VkPipelineBinaryKeyKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeyKHR.html) structure where the key for this binary will be written.
- `pPipelineBinaryDataSize` is a pointer to a `size_t` value related to the amount of data in the pipeline binary, as described below.
- `pPipelineBinaryData` is either `NULL` or a pointer to a buffer.

## [](#_description)Description

If `pPipelineBinaryData` is `NULL`, then the size of the data, in bytes, that is required to store the binary is returned in `pPipelineBinaryDataSize`. Otherwise, `pPipelineBinaryDataSize` **must** contain the size of the buffer, in bytes, pointed to by `pPipelineBinaryData`, and on return `pPipelineBinaryDataSize` is overwritten with the size of the data, in bytes, that is required to store the binary. If `pPipelineBinaryDataSize` is less than the size that is required to store the binary, nothing is written to `pPipelineBinaryData` and `VK_ERROR_NOT_ENOUGH_SPACE_KHR` will be returned, instead of `VK_SUCCESS`.

If the call returns one of the success return codes, the pipeline binary key is written to `pPipelineBinaryKey`, regardless of whether `pPipelineBinaryData` is `NULL` or not.

If [pipelineBinaryCompressedData](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-pipelineBinaryCompressedData) is `VK_FALSE`, implementations **should** not return compressed pipeline binary data to the application.

Valid Usage (Implicit)

- [](#VUID-vkGetPipelineBinaryDataKHR-device-parameter)VUID-vkGetPipelineBinaryDataKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetPipelineBinaryDataKHR-pInfo-parameter)VUID-vkGetPipelineBinaryDataKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkPipelineBinaryDataInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryDataInfoKHR.html) structure
- [](#VUID-vkGetPipelineBinaryDataKHR-pPipelineBinaryKey-parameter)VUID-vkGetPipelineBinaryDataKHR-pPipelineBinaryKey-parameter  
  `pPipelineBinaryKey` **must** be a valid pointer to a [VkPipelineBinaryKeyKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeyKHR.html) structure
- [](#VUID-vkGetPipelineBinaryDataKHR-pPipelineBinaryDataSize-parameter)VUID-vkGetPipelineBinaryDataKHR-pPipelineBinaryDataSize-parameter  
  `pPipelineBinaryDataSize` **must** be a valid pointer to a `size_t` value
- [](#VUID-vkGetPipelineBinaryDataKHR-pPipelineBinaryData-parameter)VUID-vkGetPipelineBinaryDataKHR-pPipelineBinaryData-parameter  
  If the value referenced by `pPipelineBinaryDataSize` is not `0`, and `pPipelineBinaryData` is not `NULL`, `pPipelineBinaryData` **must** be a valid pointer to an array of `pPipelineBinaryDataSize` bytes

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_NOT_ENOUGH_SPACE_KHR`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipelineBinaryDataInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryDataInfoKHR.html), [VkPipelineBinaryKeyKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeyKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPipelineBinaryDataKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0