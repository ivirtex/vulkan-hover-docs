# VkAccelerationStructureBuildTypeKHR(3) Manual Page

## Name

VkAccelerationStructureBuildTypeKHR - Acceleration structure build type



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of `buildType` in
[vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureBuildSizesKHR.html)
are:

``` c
// Provided by VK_KHR_acceleration_structure
typedef enum VkAccelerationStructureBuildTypeKHR {
    VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_KHR = 0,
    VK_ACCELERATION_STRUCTURE_BUILD_TYPE_DEVICE_KHR = 1,
    VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_OR_DEVICE_KHR = 2,
} VkAccelerationStructureBuildTypeKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_KHR` requests the memory
  requirement for operations performed by the host.

- `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_DEVICE_KHR` requests the memory
  requirement for operations performed by the device.

- `VK_ACCELERATION_STRUCTURE_BUILD_TYPE_HOST_OR_DEVICE_KHR` requests the
  memory requirement for operations performed by either the host, or the
  device.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureBuildSizesKHR.html),
[vkGetMicromapBuildSizesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMicromapBuildSizesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureBuildTypeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
