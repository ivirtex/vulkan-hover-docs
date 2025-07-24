# vkCreatePipelineBinariesKHR(3) Manual Page

## Name

vkCreatePipelineBinariesKHR - Create pipeline binaries from a pipeline or previously retrieved data



## [](#_c_specification)C Specification

To create pipeline binary objects, call:

```c++
// Provided by VK_KHR_pipeline_binary
VkResult vkCreatePipelineBinariesKHR(
    VkDevice                                    device,
    const VkPipelineBinaryCreateInfoKHR*        pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkPipelineBinaryHandlesInfoKHR*             pBinaries);
```

## [](#_parameters)Parameters

- `device` is the logical device that creates the pipeline binary objects.
- `pCreateInfo` is a pointer to a [VkPipelineBinaryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryCreateInfoKHR.html) structure that contains the data to create the pipeline binaries from.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pBinaries` is a pointer to a [VkPipelineBinaryHandlesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryHandlesInfoKHR.html) structure in which the resulting pipeline binaries are returned.

## [](#_description)Description

The implementation will attempt to create all pipeline binaries. If creation fails for any pipeline binary, then:

- The corresponding entry in the `pPipelineBinaries` output array will be filled with [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html).
- The `VkResult` returned by [vkCreatePipelineBinariesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePipelineBinariesKHR.html) will contain the error value for the first entry in the output array in `pBinaries` containing [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html).

Valid Usage (Implicit)

- [](#VUID-vkCreatePipelineBinariesKHR-device-parameter)VUID-vkCreatePipelineBinariesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkCreatePipelineBinariesKHR-pCreateInfo-parameter)VUID-vkCreatePipelineBinariesKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkPipelineBinaryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryCreateInfoKHR.html) structure
- [](#VUID-vkCreatePipelineBinariesKHR-pAllocator-parameter)VUID-vkCreatePipelineBinariesKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreatePipelineBinariesKHR-pBinaries-parameter)VUID-vkCreatePipelineBinariesKHR-pBinaries-parameter  
  `pBinaries` **must** be a valid pointer to a [VkPipelineBinaryHandlesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryHandlesInfoKHR.html) structure

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_PIPELINE_BINARY_MISSING_KHR`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_INITIALIZATION_FAILED`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipelineBinaryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryCreateInfoKHR.html), [VkPipelineBinaryHandlesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryHandlesInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreatePipelineBinariesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0