# VkBuildAccelerationStructureModeKHR(3) Manual Page

## Name

VkBuildAccelerationStructureModeKHR - Enum specifying the type of build
operation to perform



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBuildAccelerationStructureModeKHR` enumeration is defined as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef enum VkBuildAccelerationStructureModeKHR {
    VK_BUILD_ACCELERATION_STRUCTURE_MODE_BUILD_KHR = 0,
    VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR = 1,
} VkBuildAccelerationStructureModeKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_BUILD_ACCELERATION_STRUCTURE_MODE_BUILD_KHR` specifies that the
  destination acceleration structure will be built using the specified
  geometries.

- `VK_BUILD_ACCELERATION_STRUCTURE_MODE_UPDATE_KHR` specifies that the
  destination acceleration structure will be built using data in a
  source acceleration structure, updated by the specified geometries.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBuildAccelerationStructureModeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
