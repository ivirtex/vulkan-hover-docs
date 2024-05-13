# VkDeviceOrHostAddressKHR(3) Manual Page

## Name

VkDeviceOrHostAddressKHR - Union specifying a device or host address



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDeviceOrHostAddressKHR` union is defined as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef union VkDeviceOrHostAddressKHR {
    VkDeviceAddress    deviceAddress;
    void*              hostAddress;
} VkDeviceOrHostAddressKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `deviceAddress` is a buffer device address as returned by the
  [vkGetBufferDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferDeviceAddressKHR.html)
  command.

- `hostAddress` is a host memory address.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html),
[VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html),
[VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapToMemoryInfoEXT.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html),
[VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapBuildInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceOrHostAddressKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
