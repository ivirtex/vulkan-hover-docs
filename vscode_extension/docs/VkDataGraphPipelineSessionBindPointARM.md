# VkDataGraphPipelineSessionBindPointARM(3) Manual Page

## Name

VkDataGraphPipelineSessionBindPointARM - Enumeration describing the bind points of a data graph pipeline session



## [](#_c_specification)C Specification

Possible values of `VkDataGraphPipelineSessionBindPointARM`, specifying the bind point of a data graph pipeline session, are:

```c++
// Provided by VK_ARM_data_graph
typedef enum VkDataGraphPipelineSessionBindPointARM {
    VK_DATA_GRAPH_PIPELINE_SESSION_BIND_POINT_TRANSIENT_ARM = 0,
} VkDataGraphPipelineSessionBindPointARM;
```

## [](#_description)Description

- `VK_DATA_GRAPH_PIPELINE_SESSION_BIND_POINT_TRANSIENT_ARM` corresponds to the transient data produced and consumed during one dispatch of a data graph pipeline in a data graph pipeline session. This transient data is never reused by subsequent dispatches and can safely be clobbered once a [vkCmdDispatchDataGraphARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchDataGraphARM.html) command completes execution.

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkBindDataGraphPipelineSessionMemoryInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindDataGraphPipelineSessionMemoryInfoARM.html), [VkDataGraphPipelineSessionBindPointRequirementARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionBindPointRequirementARM.html), [VkDataGraphPipelineSessionMemoryRequirementsInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineSessionMemoryRequirementsInfoARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineSessionBindPointARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0