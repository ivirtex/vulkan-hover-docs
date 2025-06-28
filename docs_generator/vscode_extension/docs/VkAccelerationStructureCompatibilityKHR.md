# VkAccelerationStructureCompatibilityKHR(3) Manual Page

## Name

VkAccelerationStructureCompatibilityKHR - Acceleration structure compatibility



## [](#_c_specification)C Specification

Possible values of `pCompatibility` returned by [vkGetDeviceAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceAccelerationStructureCompatibilityKHR.html) are:

```c++
// Provided by VK_KHR_acceleration_structure
typedef enum VkAccelerationStructureCompatibilityKHR {
    VK_ACCELERATION_STRUCTURE_COMPATIBILITY_COMPATIBLE_KHR = 0,
    VK_ACCELERATION_STRUCTURE_COMPATIBILITY_INCOMPATIBLE_KHR = 1,
} VkAccelerationStructureCompatibilityKHR;
```

## [](#_description)Description

- `VK_ACCELERATION_STRUCTURE_COMPATIBILITY_COMPATIBLE_KHR` if the `pVersionData` version acceleration structure is compatible with `device`.
- `VK_ACCELERATION_STRUCTURE_COMPATIBILITY_INCOMPATIBLE_KHR` if the `pVersionData` version acceleration structure is not compatible with `device`.

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [vkGetDeviceAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceAccelerationStructureCompatibilityKHR.html), [vkGetDeviceMicromapCompatibilityEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceMicromapCompatibilityEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureCompatibilityKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0