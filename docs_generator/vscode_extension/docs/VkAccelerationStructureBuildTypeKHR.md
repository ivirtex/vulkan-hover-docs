# VkAccelerationStructureBuildTypeKHR(3) Manual Page

## Name

VkAccelerationStructureBuildTypeKHR - Acceleration structure build type



## [](#_c_specification)C Specification

Possible values of `buildType` in [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureBuildSizesKHR.html) are:

```c++
// Provided by VK_KHR_acceleration_structure
typedef enum VkAccelerationStructureBuildTypeKHR {
    VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_KHR = 0,
    VK_ACCELERATION_STRUCTURE_BUILD_TYPE_DEVICE_KHR = 1,
    VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_OR_DEVICE_KHR = 2,
} VkAccelerationStructureBuildTypeKHR;
```

## [](#_description)Description

- `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_KHR` requests the memory requirement for operations performed by the host.
- `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_DEVICE_KHR` requests the memory requirement for operations performed by the device.
- `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_OR_DEVICE_KHR` requests the memory requirement for operations performed by either the host, or the device.

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureBuildSizesKHR.html), [vkGetMicromapBuildSizesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMicromapBuildSizesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureBuildTypeKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0