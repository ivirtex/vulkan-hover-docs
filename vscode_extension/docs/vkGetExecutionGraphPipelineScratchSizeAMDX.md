# vkGetExecutionGraphPipelineScratchSizeAMDX(3) Manual Page

## Name

vkGetExecutionGraphPipelineScratchSizeAMDX - Query scratch space required to dispatch an execution graph



## [](#_c_specification)C Specification

To query the scratch space required to dispatch an execution graph, call:

```c++
// Provided by VK_AMDX_shader_enqueue
VkResult vkGetExecutionGraphPipelineScratchSizeAMDX(
    VkDevice                                    device,
    VkPipeline                                  executionGraph,
    VkExecutionGraphPipelineScratchSizeAMDX*    pSizeInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device that `executionGraph` was created on.
- `executionGraph` is the execution graph pipeline to query the scratch space for.
- `pSizeInfo` is a pointer to a [VkExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineScratchSizeAMDX.html) structure that will contain the required scratch size.

## [](#_description)Description

After this function returns, information about the scratch space required will be returned in `pSizeInfo`.

Valid Usage (Implicit)

- [](#VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-device-parameter)VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-executionGraph-parameter)VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-executionGraph-parameter  
  `executionGraph` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-pSizeInfo-parameter)VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-pSizeInfo-parameter  
  `pSizeInfo` **must** be a valid pointer to a [VkExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineScratchSizeAMDX.html) structure
- [](#VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-executionGraph-parent)VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-executionGraph-parent  
  `executionGraph` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## [](#_see_also)See Also

[VK\_AMDX\_shader\_enqueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_shader_enqueue.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineScratchSizeAMDX.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetExecutionGraphPipelineScratchSizeAMDX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0