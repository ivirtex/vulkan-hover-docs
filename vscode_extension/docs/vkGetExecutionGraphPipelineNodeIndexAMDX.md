# vkGetExecutionGraphPipelineNodeIndexAMDX(3) Manual Page

## Name

vkGetExecutionGraphPipelineNodeIndexAMDX - Query internal id of a node in an execution graph



## [](#_c_specification)C Specification

To query the internal node index for a particular node in an execution graph, call:

```c++
// Provided by VK_AMDX_shader_enqueue
VkResult vkGetExecutionGraphPipelineNodeIndexAMDX(
    VkDevice                                    device,
    VkPipeline                                  executionGraph,
    const VkPipelineShaderStageNodeCreateInfoAMDX* pNodeInfo,
    uint32_t*                                   pNodeIndex);
```

## [](#_parameters)Parameters

- `device` is the logical device that `executionGraph` was created on.
- `executionGraph` is the execution graph pipeline to query the internal node index for.
- `pNodeInfo` is a pointer to a [VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html) structure identifying the name and index of the node to query.
- `pNodeIndex` is the returned internal node index of the identified node.

## [](#_description)Description

Once this function returns, the contents of `pNodeIndex` contain the internal node index of the identified node.

Valid Usage

- [](#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-09140)VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-09140  
  `pNodeInfo->pName` **must** not be `NULL`
- [](#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-09141)VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-09141  
  `pNodeInfo->index` **must** not be `VK_SHADER_INDEX_UNUSED_AMDX`
- [](#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-09142)VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-09142  
  There **must** be a node in `executionGraph` with a shader name and index equal to `pNodeInfo->pName` and `pNodeInfo->index`

Valid Usage (Implicit)

- [](#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-device-parameter)VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-parameter)VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-parameter  
  `executionGraph` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-parameter)VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-parameter  
  `pNodeInfo` **must** be a valid pointer to a valid [VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html) structure
- [](#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeIndex-parameter)VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeIndex-parameter  
  `pNodeIndex` **must** be a valid pointer to a `uint32_t` value
- [](#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-parent)VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-parent  
  `executionGraph` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_AMDX\_shader\_enqueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_shader_enqueue.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetExecutionGraphPipelineNodeIndexAMDX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0