# VkCheckpointData2NV(3) Manual Page

## Name

VkCheckpointData2NV - Return structure for command buffer checkpoint
data



## <a href="#_c_specification" class="anchor"></a>C Specification

The [VkCheckpointData2NV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCheckpointData2NV.html) structure is defined
as:

``` c
// Provided by VK_KHR_synchronization2 with VK_NV_device_diagnostic_checkpoints
typedef struct VkCheckpointData2NV {
    VkStructureType          sType;
    void*                    pNext;
    VkPipelineStageFlags2    stage;
    void*                    pCheckpointMarker;
} VkCheckpointData2NV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `stage` indicates a single pipeline stage which the checkpoint marker
  data refers to.

- `pCheckpointMarker` contains the value of the last checkpoint marker
  executed in the stage that `stage` refers to.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkCheckpointData2NV-sType-sType"
  id="VUID-VkCheckpointData2NV-sType-sType"></a>
  VUID-VkCheckpointData2NV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_CHECKPOINT_DATA_2_NV`

- <a href="#VUID-VkCheckpointData2NV-pNext-pNext"
  id="VUID-VkCheckpointData2NV-pNext-pNext"></a>
  VUID-VkCheckpointData2NV-pNext-pNext  
  `pNext` **must** be `NULL`

The stages at which a checkpoint marker **can** be executed are
implementation-defined and **can** be queried by calling
[vkGetPhysicalDeviceQueueFamilyProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceQueueFamilyProperties2.html).

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html),
[VK_NV_device_diagnostic_checkpoints](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_diagnostic_checkpoints.html),
[VkPipelineStageFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlags2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetQueueCheckpointData2NV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetQueueCheckpointData2NV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCheckpointData2NV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
