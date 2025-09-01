# VkExecutionGraphPipelineScratchSizeAMDX(3) Manual Page

## Name

VkExecutionGraphPipelineScratchSizeAMDX - Structure describing the scratch space required to dispatch an execution graph



## [](#_c_specification)C Specification

The `VkExecutionGraphPipelineScratchSizeAMDX` structure is defined as:

```c++
// Provided by VK_AMDX_shader_enqueue
typedef struct VkExecutionGraphPipelineScratchSizeAMDX {
    VkStructureType    sType;
    void*              pNext;
    VkDeviceSize       minSize;
    VkDeviceSize       maxSize;
    VkDeviceSize       sizeGranularity;
} VkExecutionGraphPipelineScratchSizeAMDX;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `minSize` indicates the minimum scratch space required for dispatching the queried execution graph.
- `maxSize` indicates the maximum scratch space that can be used for dispatching the queried execution graph.
- `sizeGranularity` indicates the granularity at which the scratch space can be increased from `minSize`.

## [](#_description)Description

Applications **can** use any amount of scratch memory greater than `minSize` for dispatching a graph, however only the values equal to `minSize` + an integer multiple of `sizeGranularity` will be used. Greater values **may** result in higher performance, up to `maxSize` which indicates the most memory that an implementation can use effectively.

Valid Usage (Implicit)

- [](#VUID-VkExecutionGraphPipelineScratchSizeAMDX-sType-sType)VUID-VkExecutionGraphPipelineScratchSizeAMDX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXECUTION_GRAPH_PIPELINE_SCRATCH_SIZE_AMDX`

## [](#_see_also)See Also

[VK\_AMDX\_shader\_enqueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_shader_enqueue.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetExecutionGraphPipelineScratchSizeAMDX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExecutionGraphPipelineScratchSizeAMDX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0