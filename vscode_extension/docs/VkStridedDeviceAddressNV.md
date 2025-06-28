# VkStridedDeviceAddressNV(3) Manual Page

## Name

VkStridedDeviceAddressNV - Structure specifying a device addresses with a stride



## [](#_c_specification)C Specification

The [VkStridedDeviceAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressNV.html) structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkStridedDeviceAddressNV {
    VkDeviceAddress    startAddress;
    VkDeviceSize       strideInBytes;
} VkStridedDeviceAddressNV;
```

## [](#_members)Members

- `startAddress` is the device address (as returned by the [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddress.html) command) at which the region starts, or zero if the region is unused.
- `strideInBytes` is the byte stride between consecutive elements. Only the bottom 32 bits are used. The field is 64 bits to ensure consistent alignment across all containing structures.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkBuildPartitionedAccelerationStructureIndirectCommandNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildPartitionedAccelerationStructureIndirectCommandNV.html), [VkClusterAccelerationStructureInstantiateClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInstantiateClusterInfoNV.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkStridedDeviceAddressNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0