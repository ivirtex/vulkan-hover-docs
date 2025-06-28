# VkPipelineBinaryKHR(3) Manual Page

## Name

VkPipelineBinaryKHR - Opaque handle to a pipeline binary object



## [](#_c_specification)C Specification

Pipeline binary objects allow the result of pipeline construction to be reused between pipelines and between runs of an application. Reuse is achieved by extracting pipeline binaries from a [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) object, associating them with a corresponding [VkPipelineBinaryKeyKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeyKHR.html) and then adding a [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html) to the `pNext` chain of any `Vk*PipelineCreateInfo` when creating a pipeline. Pipeline binaries can be reused between runs by extracting `VkPipelineBinaryDataKHR` from `VkPipelineBinaryKHR` objects, saving the contents, and then using them to create a `VkPipelineBinaryKHR` object on subsequent runs.

When creating a pipeline that includes [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html) in the `pNext` chain, or has the `VK_PIPELINE_CREATE_2_CAPTURE_DATA_BIT_KHR` flag set, the use of [VkPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCache.html) objects is not allowed.

Pipeline binary objects are represented by `VkPipelineBinaryKHR` handles:

```c++
// Provided by VK_KHR_pipeline_binary
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkPipelineBinaryKHR)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_KHR\_pipeline\_binary](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_pipeline_binary.html), [VkPipelineBinaryDataInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryDataInfoKHR.html), [VkPipelineBinaryHandlesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryHandlesInfoKHR.html), [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html), [vkDestroyPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyPipelineBinaryKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineBinaryKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0