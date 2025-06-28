# VkCheckpointDataNV(3) Manual Page

## Name

VkCheckpointDataNV - Return structure for command buffer checkpoint data



## [](#_c_specification)C Specification

The [VkCheckpointDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCheckpointDataNV.html) structure is defined as:

```c++
// Provided by VK_NV_device_diagnostic_checkpoints
typedef struct VkCheckpointDataNV {
    VkStructureType            sType;
    void*                      pNext;
    VkPipelineStageFlagBits    stage;
    void*                      pCheckpointMarker;
} VkCheckpointDataNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `stage` is a `VkPipelineStageFlagBits` value specifying which pipeline stage the checkpoint marker data refers to.
- `pCheckpointMarker` contains the value of the last checkpoint marker executed in the stage that `stage` refers to.

## [](#_description)Description

The stages at which a checkpoint marker **can** be executed are implementation-defined and **can** be queried by calling [vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html).

Valid Usage (Implicit)

- [](#VUID-VkCheckpointDataNV-sType-sType)VUID-VkCheckpointDataNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CHECKPOINT_DATA_NV`
- [](#VUID-VkCheckpointDataNV-pNext-pNext)VUID-VkCheckpointDataNV-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_NV\_device\_diagnostic\_checkpoints](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_diagnostic_checkpoints.html), [VkPipelineStageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetQueueCheckpointDataNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetQueueCheckpointDataNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCheckpointDataNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0