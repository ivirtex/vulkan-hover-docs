# VkDeviceOrHostAddressKHR(3) Manual Page

## Name

VkDeviceOrHostAddressKHR - Union specifying a device or host address



## [](#_c_specification)C Specification

The `VkDeviceOrHostAddressKHR` union is defined as:

```c++
// Provided by VK_KHR_acceleration_structure, VK_NV_cooperative_vector
typedef union VkDeviceOrHostAddressKHR {
    VkDeviceAddress    deviceAddress;
    void*              hostAddress;
} VkDeviceOrHostAddressKHR;
```

## [](#_members)Members

- `deviceAddress` is a buffer device address as returned by the [vkGetBufferDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddressKHR.html) command.
- `hostAddress` is a host memory address.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VK\_NV\_cooperative\_vector](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cooperative_vector.html), [VkAccelerationStructureBuildGeometryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureBuildGeometryInfoKHR.html), [VkConvertCooperativeVectorMatrixInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConvertCooperativeVectorMatrixInfoNV.html), [VkCopyAccelerationStructureToMemoryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureToMemoryInfoKHR.html), [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapToMemoryInfoEXT.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceOrHostAddressKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0