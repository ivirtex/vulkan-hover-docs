# vkGetExecutionGraphPipelineScratchSizeAMDX(3) Manual Page

## Name

vkGetExecutionGraphPipelineScratchSizeAMDX - Query scratch space
required to dispatch an execution graph



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the scratch space required to dispatch an execution graph,
call:

``` c
// Provided by VK_AMDX_shader_enqueue
VkResult vkGetExecutionGraphPipelineScratchSizeAMDX(
    VkDevice                                    device,
    VkPipeline                                  executionGraph,
    VkExecutionGraphPipelineScratchSizeAMDX*    pSizeInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the that `executionGraph` was created on.

- `executionGraph` is the execution graph pipeline to query the scratch
  space for.

- `pSizeInfo` is a pointer to a
  [VkExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExecutionGraphPipelineScratchSizeAMDX.html)
  structure that will contain the required scratch size.

## <a href="#_description" class="anchor"></a>Description

After this function returns, information about the scratch space
required will be returned in `pSizeInfo`.

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-device-parameter"
  id="VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-device-parameter"></a>
  VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-executionGraph-parameter"
  id="VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-executionGraph-parameter"></a>
  VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-executionGraph-parameter  
  `executionGraph` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)
  handle

- <a
  href="#VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-pSizeInfo-parameter"
  id="VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-pSizeInfo-parameter"></a>
  VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-pSizeInfo-parameter  
  `pSizeInfo` **must** be a valid pointer to a
  [VkExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExecutionGraphPipelineScratchSizeAMDX.html)
  structure

- <a
  href="#VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-executionGraph-parent"
  id="VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-executionGraph-parent"></a>
  VUID-vkGetExecutionGraphPipelineScratchSizeAMDX-executionGraph-parent  
  `executionGraph` **must** have been created, allocated, or retrieved
  from `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMDX_shader_enqueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMDX_shader_enqueue.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExecutionGraphPipelineScratchSizeAMDX.html),
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetExecutionGraphPipelineScratchSizeAMDX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
