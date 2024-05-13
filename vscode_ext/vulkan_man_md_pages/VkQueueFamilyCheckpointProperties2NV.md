# VkQueueFamilyCheckpointProperties2NV(3) Manual Page

## Name

VkQueueFamilyCheckpointProperties2NV - Return structure for queue family
checkpoint information query



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkQueueFamilyCheckpointProperties2NV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyCheckpointProperties2NV.html)
structure is defined as:

``` c
// Provided by VK_KHR_synchronization2 with VK_NV_device_diagnostic_checkpoints
typedef struct VkQueueFamilyCheckpointProperties2NV {
    VkStructureType          sType;
    void*                    pNext;
    VkPipelineStageFlags2    checkpointExecutionStageMask;
} VkQueueFamilyCheckpointProperties2NV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `checkpointExecutionStageMask` is a mask indicating which pipeline
  stages the implementation can execute checkpoint markers in.

## <a href="#_description" class="anchor"></a>Description

Additional queue family information can be queried by setting
[VkQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueueFamilyProperties2.html)::`pNext` to
point to a `VkQueueFamilyCheckpointProperties2NV` structure.

Valid Usage (Implicit)

- <a href="#VUID-VkQueueFamilyCheckpointProperties2NV-sType-sType"
  id="VUID-VkQueueFamilyCheckpointProperties2NV-sType-sType"></a>
  VUID-VkQueueFamilyCheckpointProperties2NV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_QUEUE_FAMILY_CHECKPOINT_PROPERTIES_2_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_NV_device_diagnostic_checkpoints](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_diagnostic_checkpoints.html),
[VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkQueueFamilyCheckpointProperties2NV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
