# vkGetExecutionGraphPipelineNodeIndexAMDX(3) Manual Page

## Name

vkGetExecutionGraphPipelineNodeIndexAMDX - Query internal id of a node
in an execution graph



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the internal node index for a particular node in an execution
graph, call:

``` c
// Provided by VK_AMDX_shader_enqueue
VkResult vkGetExecutionGraphPipelineNodeIndexAMDX(
    VkDevice                                    device,
    VkPipeline                                  executionGraph,
    const VkPipelineShaderStageNodeCreateInfoAMDX* pNodeInfo,
    uint32_t*                                   pNodeIndex);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the that `executionGraph` was created on.

- `executionGraph` is the execution graph pipeline to query the internal
  node index for.

- `pNodeInfo` is a pointer to a
  [VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html)
  structure identifying the name and index of the node to query.

- `pNodeIndex` is the returned internal node index of the identified
  node.

## <a href="#_description" class="anchor"></a>Description

Once this function returns, the contents of `pNodeIndex` contain the
internal node index of the identified node.

Valid Usage

- <a href="#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-09140"
  id="VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-09140"></a>
  VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-09140  
  `pNodeInfo->pName` **must** not be `NULL`

- <a href="#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-09141"
  id="VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-09141"></a>
  VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-09141  
  `pNodeInfo->index` **must** not be `VK_SHADER_INDEX_UNUSED_AMDX`

- <a
  href="#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-09142"
  id="VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-09142"></a>
  VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-09142  
  There **must** be a node in `executionGraph` with a shader name and
  index equal to `pNodeInfo->pName` and `pNodeInfo->index`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-device-parameter"
  id="VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-device-parameter"></a>
  VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-parameter"
  id="VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-parameter"></a>
  VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-parameter  
  `executionGraph` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)
  handle

- <a
  href="#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-parameter"
  id="VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-parameter"></a>
  VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeInfo-parameter  
  `pNodeInfo` **must** be a valid pointer to a valid
  [VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html)
  structure

- <a
  href="#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeIndex-parameter"
  id="VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeIndex-parameter"></a>
  VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-pNodeIndex-parameter  
  `pNodeIndex` **must** be a valid pointer to a `uint32_t` value

- <a
  href="#VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-parent"
  id="VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-parent"></a>
  VUID-vkGetExecutionGraphPipelineNodeIndexAMDX-executionGraph-parent  
  `executionGraph` **must** have been created, allocated, or retrieved
  from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMDX_shader_enqueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMDX_shader_enqueue.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html),
[VkPipelineShaderStageNodeCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageNodeCreateInfoAMDX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetExecutionGraphPipelineNodeIndexAMDX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
