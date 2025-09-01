# vkReleaseCapturedPipelineDataKHR(3) Manual Page

## Name

vkReleaseCapturedPipelineDataKHR - Release captured pipeline binary data



## [](#_c_specification)C Specification

To release pipeline resources captured with `VK_PIPELINE_CREATE_2_CAPTURE_DATA_BIT_KHR`, call:

```c++
// Provided by VK_KHR_pipeline_binary
VkResult vkReleaseCapturedPipelineDataKHR(
    VkDevice                                    device,
    const VkReleaseCapturedPipelineDataInfoKHR* pInfo,
    const VkAllocationCallbacks*                pAllocator);
```

## [](#_parameters)Parameters

- `device` is the logical device that created the pipeline object.
- `pInfo` is a pointer to a [VkReleaseCapturedPipelineDataInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkReleaseCapturedPipelineDataInfoKHR.html) structure which describes the pipeline to release the data from.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.

## [](#_description)Description

The implementation **may** free any resources captured as a result of creating the pipeline with `VK_PIPELINE_CREATE_2_CAPTURE_DATA_BIT_KHR` and put the pipeline into a state as if `VK_PIPELINE_CREATE_2_CAPTURE_DATA_BIT_KHR` had not been provided at pipeline creation time.

Note

Any resources captured as a result of creating the pipeline with `VK_PIPELINE_CREATE_2_CAPTURE_DATA_BIT_KHR` are implicitly freed by [vkDestroyPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyPipeline.html).

Valid Usage

- [](#VUID-vkReleaseCapturedPipelineDataKHR-pipeline-09611)VUID-vkReleaseCapturedPipelineDataKHR-pipeline-09611  
  If `VkAllocationCallbacks` were provided when `pipeline` was created, a compatible set of callbacks **must** be provided in `pAllocator`
- [](#VUID-vkReleaseCapturedPipelineDataKHR-pipeline-09612)VUID-vkReleaseCapturedPipelineDataKHR-pipeline-09612  
  If no [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) were provided when `pipeline` was created, `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- [](#VUID-vkReleaseCapturedPipelineDataKHR-device-parameter)VUID-vkReleaseCapturedPipelineDataKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkReleaseCapturedPipelineDataKHR-pInfo-parameter)VUID-vkReleaseCapturedPipelineDataKHR-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkReleaseCapturedPipelineDataInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkReleaseCapturedPipelineDataInfoKHR.html) structure
- [](#VUID-vkReleaseCapturedPipelineDataKHR-pAllocator-parameter)VUID-vkReleaseCapturedPipelineDataKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkReleaseCapturedPipelineDataInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkReleaseCapturedPipelineDataInfoKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkReleaseCapturedPipelineDataKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0