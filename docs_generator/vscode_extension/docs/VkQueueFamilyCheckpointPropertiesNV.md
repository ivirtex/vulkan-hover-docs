# VkQueueFamilyCheckpointPropertiesNV(3) Manual Page

## Name

VkQueueFamilyCheckpointPropertiesNV - Return structure for queue family checkpoint information query



## [](#_c_specification)C Specification

The [VkQueueFamilyCheckpointPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyCheckpointPropertiesNV.html) structure is defined as:

```c++
// Provided by VK_NV_device_diagnostic_checkpoints
typedef struct VkQueueFamilyCheckpointPropertiesNV {
    VkStructureType         sType;
    void*                   pNext;
    VkPipelineStageFlags    checkpointExecutionStageMask;
} VkQueueFamilyCheckpointPropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `checkpointExecutionStageMask` is a mask indicating which pipeline stages the implementation can execute checkpoint markers in.

## [](#_description)Description

Additional queue family information can be queried by setting [VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueueFamilyProperties2.html)::`pNext` to point to a `VkQueueFamilyCheckpointPropertiesNV` structure.

Valid Usage (Implicit)

- [](#VUID-VkQueueFamilyCheckpointPropertiesNV-sType-sType)VUID-VkQueueFamilyCheckpointPropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_device\_diagnostic\_checkpoints](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_diagnostic_checkpoints.html), [VkPipelineStageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkQueueFamilyCheckpointPropertiesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0