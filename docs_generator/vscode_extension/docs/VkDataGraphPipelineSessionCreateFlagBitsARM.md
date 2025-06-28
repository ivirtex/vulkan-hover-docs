# VkDataGraphPipelineSessionCreateFlagBitsARM(3) Manual Page

## Name

VkDataGraphPipelineSessionCreateFlagBitsARM - Bitmask specifying additional parameters of a data graph pipeline session



## [](#_c_specification)C Specification

Bits which **can** be set in [VkDataGraphPipelineSessionCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionCreateInfoARM.html)::`flags`, specifying additional parameters of a data graph pipeline session, are:

```c++
// Provided by VK_ARM_data_graph
// Flag bits for VkDataGraphPipelineSessionCreateFlagBitsARM
typedef VkFlags64 VkDataGraphPipelineSessionCreateFlagBitsARM;
static const VkDataGraphPipelineSessionCreateFlagBitsARM VK_DATA_GRAPH_PIPELINE_SESSION_CREATE_PROTECTED_BIT_ARM = 0x00000001ULL;
```

## [](#_description)Description

- `VK_DATA_GRAPH_PIPELINE_SESSION_CREATE_PROTECTED_BIT_ARM` specifies that the data graph pipeline session is backed by protected memory.

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelineSessionCreateFlagsARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionCreateFlagsARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineSessionCreateFlagBitsARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0