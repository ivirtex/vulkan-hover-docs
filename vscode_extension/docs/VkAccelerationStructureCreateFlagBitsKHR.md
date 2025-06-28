# VkAccelerationStructureCreateFlagBitsKHR(3) Manual Page

## Name

VkAccelerationStructureCreateFlagBitsKHR - Bitmask specifying additional creation parameters for acceleration structure



## [](#_c_specification)C Specification

Bits which **can** be set in [VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateInfoKHR.html)::`createFlags`, specifying additional creation parameters for acceleration structures, are:

```c++
// Provided by VK_KHR_acceleration_structure
typedef enum VkAccelerationStructureCreateFlagBitsKHR {
    VK_ACCELERATION_STRUCTURE_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR = 0x00000001,
  // Provided by VK_EXT_descriptor_buffer
    VK_ACCELERATION_STRUCTURE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT = 0x00000008,
  // Provided by VK_NV_ray_tracing_motion_blur
    VK_ACCELERATION_STRUCTURE_CREATE_MOTION_BIT_NV = 0x00000004,
} VkAccelerationStructureCreateFlagBitsKHR;
```

## [](#_description)Description

- `VK_ACCELERATION_STRUCTURE_CREATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT_KHR` specifies that the acceleration structure’s address **can** be saved and reused on a subsequent run.
- `VK_ACCELERATION_STRUCTURE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` specifies that the acceleration structure **can** be used with descriptor buffers when capturing and replaying (e.g. for trace capture and replay), see [VkOpaqueCaptureDescriptorDataCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOpaqueCaptureDescriptorDataCreateInfoEXT.html) for more detail.
- `VK_ACCELERATION_STRUCTURE_CREATE_MOTION_BIT_NV` specifies that the acceleration structure will be used with motion information, see [VkAccelerationStructureMotionInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMotionInfoNV.html) for more detail.

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureCreateFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureCreateFlagBitsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0