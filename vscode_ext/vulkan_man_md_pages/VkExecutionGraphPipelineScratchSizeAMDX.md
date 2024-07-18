# VkExecutionGraphPipelineScratchSizeAMDX(3) Manual Page

## Name

VkExecutionGraphPipelineScratchSizeAMDX - Structure describing the
scratch space required to dispatch an execution graph



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkExecutionGraphPipelineScratchSizeAMDX` structure is defined as:

``` c
// Provided by VK_AMDX_shader_enqueue
typedef struct VkExecutionGraphPipelineScratchSizeAMDX {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       size;
} VkExecutionGraphPipelineScratchSizeAMDX;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `size` indicates the scratch space required for dispatch the queried
  execution graph.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkExecutionGraphPipelineScratchSizeAMDX-sType-sType"
  id="VUID-VkExecutionGraphPipelineScratchSizeAMDX-sType-sType"></a>
  VUID-VkExecutionGraphPipelineScratchSizeAMDX-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_EXECUTION_GRAPH_PIPELINE_SCRATCH_SIZE_AMDX`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMDX_shader_enqueue](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMDX_shader_enqueue.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetExecutionGraphPipelineScratchSizeAMDX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExecutionGraphPipelineScratchSizeAMDX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
