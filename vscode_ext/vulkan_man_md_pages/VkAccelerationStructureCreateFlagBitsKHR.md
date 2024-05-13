# VkAccelerationStructureCreateFlagBitsKHR(3) Manual Page

## Name

VkAccelerationStructureCreateFlagBitsKHR - Bitmask specifying additional
creation parameters for acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html)::`createFlags`,
specifying additional creation parameters for acceleration structures,
are:

``` c
// Provided by VK_KHR_acceleration_structure
typedef enum VkAccelerationStructureCreateFlagBitsKHR {
    VK_ACCELERATION_STRUCTURE_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR = 0x00000001,
  // Provided by VK_EXT_descriptor_buffer
    VK_ACCELERATION_STRUCTURE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT = 0x00000008,
  // Provided by VK_NV_ray_tracing_motion_blur
    VK_ACCELERATION_STRUCTURE_CREATE_MOTION_BIT_NV = 0x00000004,
} VkAccelerationStructureCreateFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_ACCELERATION_STRUCTURE_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR`
  specifies that the acceleration structure’s address **can** be saved
  and reused on a subsequent run.

- `VK_ACCELERATION_STRUCTURE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT`
  specifies that the acceleration structure **can** be used with
  descriptor buffers when capturing and replaying (e.g. for trace
  capture and replay), see
  [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html)
  for more detail.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureCreateFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
