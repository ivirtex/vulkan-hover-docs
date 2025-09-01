# VkCheckpointData2NV(3) Manual Page

## Name

VkCheckpointData2NV - Return structure for command buffer checkpoint data



## [](#_c_specification)C Specification

The [VkCheckpointData2NV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCheckpointData2NV.html) structure is defined as:

```c++
// Provided by VK_NV_device_diagnostic_checkpoints with VK_VERSION_1_3 or VK_KHR_synchronization2
typedef struct VkCheckpointData2NV {
    VkStructureType          sType;
    void*                    pNext;
    VkPipelineStageFlags2    stage;
    void*                    pCheckpointMarker;
} VkCheckpointData2NV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stage` indicates a single pipeline stage which the checkpoint marker data refers to.
- `pCheckpointMarker` contains the value of the last checkpoint marker executed in the stage that `stage` refers to.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkCheckpointData2NV-sType-sType)VUID-VkCheckpointData2NV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CHECKPOINT_DATA_2_NV`
- [](#VUID-VkCheckpointData2NV-pNext-pNext)VUID-VkCheckpointData2NV-pNext-pNext  
  `pNext` **must** be `NULL`

The stages at which a checkpoint marker **can** be executed are implementation-defined and **can** be queried by calling [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html).

## [](#_see_also)See Also

[VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html), [VK\_NV\_device\_diagnostic\_checkpoints](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_diagnostic_checkpoints.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetQueueCheckpointData2NV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetQueueCheckpointData2NV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCheckpointData2NV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0